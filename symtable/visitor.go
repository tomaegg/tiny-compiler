package symtable

import (
	"reflect"
	"strconv"
	"tj-compiler/g4/parser"
	"tj-compiler/utils"

	"github.com/antlr4-go/antlr/v4"
)

type ExprAttribute struct {
	Type   SymType
	Value  any
	Symbol Symbol
}

var _ parser.RustLikeParserVisitor = (*Visitor)(nil)

type Visitor struct {
	parser.BaseRustLikeParserVisitor

	globalScope  Scope
	currentScope Scope
	symTable     *SymTable

	errorCount int

	loopDepth int // loop 相关的层数
}

func NewVisitor() *Visitor {
	return &Visitor{}
}

func (v Visitor) TotalError() int {
	return v.errorCount
}

func (v *Visitor) LogDefine(s Symbol) {
	token := s.Token()
	utils.PosLogger(token).Infof("define: %s", s.String())
}

func (v *Visitor) LogResolve(res Symbol, atToken antlr.Token) {
	if res == nil {
		utils.PosLogger(atToken).Fatalf("unresolved symbol at token: <%s>", atToken.GetText())
	}
	utils.PosLogger(atToken).Infof("resolve: %s", res.String())
}

func (v *Visitor) LogError(err SemanticErr, atToken antlr.Token) {
	v.errorCount++
	utils.PosLogger(atToken).Errorf("%s", err)
}

func (v *Visitor) LogInfer(s Symbol, at antlr.Token) {
	utils.PosLogger(at).Infof("infer: %s", s)
}

func (v *Visitor) checkFuncReturn(scope FuncScope) {
	funcSymbol := scope.GetSymbol()
	returned := scope.HasReturn()
	funcHasReturn := !funcSymbol.RetType().SameWith(Void)
	if funcHasReturn && !returned {
		err := NewSematicErr(RetErr).
			Message("FuncScope <%s> should have return statement, but not", funcSymbol.Name())
		v.LogError(err, funcSymbol.Token())
	}
}

func (v *Visitor) checkUninferredSymbol(scope Scope) {
	for _, sym := range scope.Symbols() {
		if sym.Type().SameWith(ToInfer) {
			err := NewSematicErr(TypeErr).
				Message("Symbol <%s> is uninferred", sym.Name())
			v.LogError(err, sym.Token())
		}
	}
}

func (v *Visitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *Visitor) VisitProg(ctx *parser.ProgContext) any {
	v.globalScope = NewGlobalScope(nil)
	v.symTable = NewSymTable(v.globalScope) // symtable从此开始定义

	v.currentScope = v.globalScope
	ret := v.Visit(ctx.Declaration())
	v.currentScope = v.currentScope.Enclosed()
	return ret
}

func (v *Visitor) VisitBlock(ctx *parser.BlockContext) any {
	localScope := NewLocalScope(v.currentScope)
	v.symTable.AddEdge(localScope, v.currentScope)
	v.currentScope = localScope
	for _, stat := range ctx.AllStat() {
		v.Visit(stat)
	}
	// 退出块作用域前检查未推断类型
	v.checkUninferredSymbol(v.currentScope)
	v.currentScope = v.currentScope.Enclosed()
	return nil
}

func (v *Visitor) VisitFuncBlock(ctx *parser.FuncBlockContext) any {
	for _, stat := range ctx.AllStat() {
		v.Visit(stat)
	}
	return nil
}

func (v *Visitor) VisitDeclaration(ctx *parser.DeclarationContext) any {
	for _, fd := range ctx.AllFuncDeclaration() {
		v.Visit(fd)
	}
	return nil
}

func (v *Visitor) VisitFuncParam(ctx *parser.FuncParamContext) any {
	fp := ctx
	tokenID := fp.ID().GetSymbol()
	varType := v.Visit(ctx.Rtype()).(SymType)
	varName := tokenID.GetText()
	mutable := fp.MUT() != nil
	varSymbol := NewBaseSymbol(varName, varType, mutable, tokenID)
	return varSymbol
}

func (v *Visitor) VisitFuncParamsList(ctx *parser.FuncParamsListContext) any {
	fps := ctx.FuncParams().AllFuncParam()
	params := make([]BaseSymbol, 0, len(fps))
	for _, fp := range fps {
		params = append(params, v.Visit(fp).(BaseSymbol))
	}
	return params
}

func (v *Visitor) VisitRtype(ctx *parser.RtypeContext) any {
	// NOTE: 目前有两种类型 1.i32 ; 2.array
	switch {
	case ctx.INT32() != nil:
		return Int32
	case ctx.ArrayType() != nil:
		etype := v.Visit(ctx.ArrayType().Rtype()).(SymType)
		number := v.Visit(ctx.ArrayType().ExprNumber()).(ExprAttribute)
		l := int32(number.Value.(int64))
		return NewSymArray(etype, l)
	default:
		panic("rtype cannot be nil type")
	}
}

func (v *Visitor) VisitFuncSignature(ctx *parser.FuncSignatureContext) any {
	funcToken := ctx.ID().GetSymbol()
	funcName := funcToken.GetText()

	params := v.Visit(ctx.FuncParamsList()).([]BaseSymbol)
	retType := Void
	if funcRet := ctx.FuncDeclarationReturn(); funcRet != nil {
		retType = v.Visit(funcRet.Rtype()).(SymType)
	}

	funcScope := NewFuncScope(v.currentScope, funcName)
	funcSymbol := NewFuncSymbol(
		funcName,
		funcScope,
		params,
		retType,
		funcToken,
	)

	funcScope.SetSymbol(funcSymbol)
	funcScope.Define(funcSymbol)
	v.LogDefine(funcSymbol)
	for _, s := range params {
		funcScope.Define(s)
		v.LogDefine(s)
	}

	return funcScope
}

func (v *Visitor) VisitFuncDeclaration(ctx *parser.FuncDeclarationContext) any {
	funcScope := v.Visit(ctx.FuncSignature()).(FuncScope)
	v.symTable.AddEdge(funcScope, v.currentScope)
	v.currentScope.Define(funcScope.GetSymbol()) // func scope定义在global之中
	v.currentScope = funcScope

	ret := v.Visit(ctx.FuncBlock())

	// check uninferred
	v.checkUninferredSymbol(v.currentScope)
	// check func return
	v.checkFuncReturn(funcScope)

	v.currentScope = v.currentScope.Enclosed()
	return ret
}

func (v *Visitor) VisitStatFuncReturn(ctx *parser.StatFuncReturnContext) any {
	// NOTE: 获取func scope, 通过scope链条向上查找
	funcScope := GetFuncScope(v.currentScope)
	if funcScope == nil {
		err := NewSematicErr(RetErr).Message("return statement not in function scope")
		v.LogError(err, ctx.RETURN().GetSymbol())
		return nil
	}

	funcSymbol := funcScope.GetSymbol()
	wantType := funcSymbol.RetType()
	getType := Void
	if ctx.Expr() != nil {
		attr := v.Visit(ctx.Expr()).(ExprAttribute)
		getType = attr.Type
	}

	if !wantType.SameWith(getType) {
		err := NewSematicErr(TypeErr).Message(
			"function return type mismatch: want <%s>, get <%s>",
			wantType, getType,
		)
		tokenRet := ctx.RETURN().GetSymbol()
		v.LogError(err, tokenRet)
	}

	// func scope returned
	funcScope.SetReturn()
	return nil
}

func (v *Visitor) VisitStatVarDeclare(ctx *parser.StatVarDeclareContext) any {
	tokenID := ctx.ID().GetSymbol()
	// NOTE: 声明时可以没有类型, 需要定义为to infer
	var lhsType SymType
	if ctx.VarType() == nil {
		lhsType = ToInfer // NOTE: 此时标注为待推断
	} else {
		lhsType = v.Visit(ctx.VarType().Rtype()).(SymType)
	}

	mutable := ctx.MUT() != nil
	lhsSymbol := NewBaseSymbol(tokenID.GetText(), lhsType, mutable, tokenID)

	v.LogDefine(lhsSymbol)
	if v.currentScope.Exist(lhsSymbol) {
		v.currentScope.Shallow(lhsSymbol)
		utils.PosLogger(tokenID).Infof("shallow: %s", lhsSymbol)
	} else {
		v.currentScope.Define(lhsSymbol)
	}

	// varinit exist
	if ctx.VarInit() != nil {
		// 访问varinit得到属性
		rhsAttr := v.Visit(ctx.VarInit().Expr()).(ExprAttribute)
		utils.PosLogger(tokenID).Debugf("init with assignment: rhs type: %v", rhsAttr.Type)
		// cannot be void
		if rhsAttr.Type.SameWith(Void) {
			err := NewSematicErr(TypeErr).Message("rhs cannot be void type")
			v.LogError(err, tokenID)
		}

		switch lhsSymbol.Type().(type) {
		case SymToInfer:
			// 当前为待推断
			lhsSymbol.Infer(rhsAttr.Type)
			v.LogInfer(lhsSymbol, tokenID)

		default:
			if !lhsType.SameWith(rhsAttr.Type) {
				err := NewSematicErr(TypeErr).
					Message("mismatch type assignment: <%s> != <%v>", lhsSymbol.Type(), rhsAttr.Type)
				v.LogError(err, tokenID)
			}
		}
	}

	return nil
}

// different expression

func (v *Visitor) VisitExprFuncCall(ctx *parser.ExprFuncCallContext) any {
	funcName := ctx.ID().GetSymbol().GetText()
	funcSymbol := v.currentScope.Resolve(funcName).(FuncSymbol)
	paramsRequired := funcSymbol.Params()

	// 检查参数数量
	funcParamCtx := ctx.FuncCallList().FuncCallParam()
	if len(paramsRequired) != len(funcParamCtx.AllExpr()) {
		err := NewSematicErr(ArgsErr).Message("function <%s> requires %d arguments, but %d are provided",
			funcName, len(paramsRequired), len(funcParamCtx.AllExpr()),
		)
		v.LogError(err, funcSymbol.Token())
	}

	// 检查参数类型
	for i := range min(len(funcParamCtx.AllExpr()), len(paramsRequired)) {
		fp := funcParamCtx.AllExpr()[i]
		attr := v.Visit(fp).(ExprAttribute)
		if !attr.Type.SameWith(paramsRequired[i].Type()) {
			err := NewSematicErr(ArgsErr).Message("function <%s> args at %d requires %s, but get type <%s>",
				funcName, i, paramsRequired[i], attr.Type,
			)
			v.LogError(err, funcSymbol.Token())
		}
	}

	v.LogResolve(funcSymbol, ctx.ID().GetSymbol())
	v.Visit(ctx.FuncCallList())
	return ExprAttribute{Type: funcSymbol.RetType()}
}

func (v *Visitor) VisitExprCmp(ctx *parser.ExprCmpContext) any {
	attrLhs := v.Visit(ctx.GetLhs()).(ExprAttribute)
	attrRhs := v.Visit(ctx.GetRhs()).(ExprAttribute)
	op := ctx.GetOp()
	if !IsNumberTypes(attrLhs.Type, attrRhs.Type) {
		err := NewSematicErr(TypeErr).Message("binary op only for numbers: lhs: %v, rhs: %v", attrLhs.Type, attrRhs.Type)
		v.LogError(err, op)
		return ExprAttribute{Type: Error}
	}

	if !attrLhs.Type.SameWith(attrRhs.Type) {
		err := NewSematicErr(TypeErr).Message("type mismatch with: <%s> %s <%s>",
			attrLhs.Type, op.GetText(), attrRhs.Type,
		)
		v.LogError(err, op)
		return ExprAttribute{Type: Error}
	}
	// NOTE: 由于暂时不支持bool类型，因此返回Int32
	return ExprAttribute{Type: Int32}
}

func (v *Visitor) VisitExprMulDiv(ctx *parser.ExprMulDivContext) any {
	attrLhs := v.Visit(ctx.GetLhs()).(ExprAttribute)
	attrRhs := v.Visit(ctx.GetRhs()).(ExprAttribute)
	op := ctx.GetOp()
	if !IsNumberTypes(attrLhs.Type, attrRhs.Type) {
		err := NewSematicErr(TypeErr).Message("binary op only for numbers: lhs: %v, rhs: %v", attrLhs.Type, attrRhs.Type)
		v.LogError(err, op)
		return ExprAttribute{Type: Error}
	}

	if !attrLhs.Type.SameWith(attrRhs.Type) {
		err := NewSematicErr(TypeErr).Message("type mismatch: <%s> %s <%s>",
			attrLhs.Type, op.GetText(), attrRhs.Type,
		)
		v.LogError(err, op)
		// 如果类型不一致返回 error 类型
		return ExprAttribute{Type: Error}
	}
	// 类型一致返回 int32 类型
	return ExprAttribute{Type: Int32}
}

func (v *Visitor) VisitExprAddSub(ctx *parser.ExprAddSubContext) any {
	attrLhs := v.Visit(ctx.GetLhs()).(ExprAttribute)
	attrRhs := v.Visit(ctx.GetRhs()).(ExprAttribute)
	op := ctx.GetOp()
	if !IsNumberTypes(attrLhs.Type, attrRhs.Type) {
		err := NewSematicErr(TypeErr).Message("binary op only for numbers: lhs: %v, rhs: %v", attrLhs.Type, attrRhs.Type)
		v.LogError(err, op)
		return ExprAttribute{Type: Error}
	}
	if !attrLhs.Type.SameWith(attrRhs.Type) {
		err := NewSematicErr(TypeErr).Message("type mismatch: <%s> %s <%s>",
			attrLhs.Type, op.GetText(), attrRhs.Type,
		)
		v.LogError(err, op)
		// 如果类型不一致返回 undefined 类型
		return ExprAttribute{Type: Error}
	}
	// 类型一致返回 int32 类型
	return ExprAttribute{Type: Int32}
}

func (v *Visitor) VisitExprParen(ctx *parser.ExprParenContext) any {
	return v.Visit(ctx.Expr())
}

func (v *Visitor) VisitExprNumber(ctx *parser.ExprNumberContext) any {
	numToken := ctx.NUMBER().GetSymbol()
	val, err := strconv.ParseInt(numToken.GetText(), 10, 32)
	if err != nil {
		err := NewSematicErr(IntOverflowErr).Message("overflow int32: %s", numToken.GetText())
		v.LogError(err, numToken)
		val = -1 // invalid
	}
	return ExprAttribute{Type: Int32, Value: val}
}

func (v *Visitor) VisitExprNum(ctx *parser.ExprNumContext) any {
	return v.Visit(ctx.ExprNumber())
}

func (v *Visitor) VisitExprArray(ctx *parser.ExprArrayContext) any {
	var (
		exprs    = ctx.ArrayElems().AllExpr()
		valid    = true
		flagType SymType
	)
	for _, e := range exprs {
		elem := v.Visit(e).(ExprAttribute)
		if flagType == nil {
			flagType = elem.Type
		}

		if !elem.Type.SameWith(flagType) {
			// elems mismatch
			err := NewSematicErr(TypeErr).Message("array elements type mismatch, first: %s, current: %v", flagType, elem.Type)
			v.LogError(err, e.GetStart())
			valid = false
		}
	}
	if !valid {
		return ExprAttribute{Type: Error, Value: nil} // semantic阶段ignore value
	}
	arrayType := NewSymArray(flagType, int32(len(exprs)))
	return ExprAttribute{Type: arrayType, Value: nil} // semantic阶段ignore value
}

func (v *Visitor) VisitExprLValue(ctx *parser.ExprLValueContext) any {
	return v.Visit(ctx.LValue())
}

func (v *Visitor) VisitLValueID(ctx *parser.LValueIDContext) any {
	token := ctx.ID().GetSymbol()
	s := v.currentScope.Resolve(token.GetText())
	var t SymType
	switch s := s.(type) {
	case BaseSymbol:
		t = s.Type()
	case FuncSymbol:
		t = Func
	default:
		utils.PosLogger(token).Fatalf("use unknown symbol: <%v>", token.GetText())
	}
	v.LogResolve(s, token)

	return ExprAttribute{Type: t, Symbol: s}
}

func (v *Visitor) checkArrayIdx(ctx parser.IExprContext) {
	ret := v.Visit(ctx).(ExprAttribute)
	switch ret.Type.(type) {
	case SymInt32:
		return

	default:
		err := NewSematicErr(TypeErr).Message("index must be integer, but %v", ret.Type)
		v.LogError(err, ctx.GetStart())
	}
}

func (v *Visitor) VisitLValueArrayAccess(ctx *parser.LValueArrayAccessContext) any {
	exprIdx := ctx.Expr()
	lvalue := ctx.LValue()
	token := lvalue.GetStart()
	arraySymbol := v.currentScope.Resolve(token.GetText())

	v.LogResolve(arraySymbol, token)

	v.checkArrayIdx(exprIdx)
	array := v.unwrapArrayVisit(lvalue)
	switch unwrapped := array.(type) {
	case SymError:
		return ExprAttribute{Type: Error}

	case SymArray:
		utils.PosLogger(lvalue.GetStart()).Debugf("unwrap array access: elems: %v", unwrapped.ElemType)
		return ExprAttribute{Type: unwrapped.ElemType, Symbol: arraySymbol}

	default: // error
		err := NewSematicErr(TypeErr).Message("operator [] should only use for array type, actual type: %v", unwrapped)
		v.LogError(err, exprIdx.GetStart())
		return ExprAttribute{Type: Error, Value: nil}
	}
}

func (v *Visitor) unwrapArrayVisit(ctx parser.ILValueContext) SymType {
	switch ctx := ctx.(type) {
	case *parser.LValueIDContext:
		// 叶子节点，访问达到ID
		token := ctx.ID().GetSymbol()
		symbol := v.currentScope.Resolve(token.GetText())
		// 必须是array type
		if s, ok := symbol.Type().(SymArray); !ok {
			err := NewSematicErr(TypeErr).Message("operator [] should only use for array type, actual type: %v", symbol.Type())
			v.LogError(err, token)
			return Error
		} else {
			utils.PosLogger(token).Debugf("unwrap array access: elems: %v", s)
		}
		return symbol.Type()

	case *parser.LValueArrayAccessContext:
		idxExpr := ctx.Expr()
		v.checkArrayIdx(idxExpr)
		// 递归访问
		lvalue := ctx.LValue()
		ret := v.unwrapArrayVisit(ctx.LValue())
		if s, ok := ret.(SymArray); ok {
			utils.PosLogger(lvalue.GetStart()).Debugf("unwrap array access: elems: %v", s.ElemType)
			return s.ElemType
		}
		return ret

	default:
		// 出现错误
		err := NewSematicErr(TypeErr).Message("operator [] should only use for array type, actual node type: %v", reflect.TypeOf(ctx))
		v.LogError(err, ctx.GetStart())
		return Error
	}
}

func (v *Visitor) VisitStatExpr(ctx *parser.StatExprContext) any {
	return v.Visit(ctx.Expr())
}

func (v *Visitor) VisitStatBreak(ctx *parser.StatBreakContext) any {
	if v.loopDepth <= 0 {
		err := NewSematicErr(BreakErr).Message("break outside loop")
		v.LogError(err, ctx.BREAK().GetSymbol())
	}
	return nil
}

func (v *Visitor) VisitStatContinue(ctx *parser.StatContinueContext) any {
	if v.loopDepth <= 0 {
		err := NewSematicErr(ContiErr).Message("continue outside loop")
		v.LogError(err, ctx.CONTINUE().GetSymbol())
	}
	return nil
}

func (v *Visitor) VisitStatWhile(ctx *parser.StatWhileContext) any {
	v.Visit(ctx.Expr())
	v.loopDepth++
	v.Visit(ctx.Block())
	v.loopDepth--
	return nil
}

func (v *Visitor) VisitStatLoop(ctx *parser.StatLoopContext) any {
	v.loopDepth++
	v.Visit(ctx.Block())
	v.loopDepth--
	return nil
}

func (v *Visitor) VisitStatIfElse(ctx *parser.StatIfElseContext) any {
	// if branch
	v.Visit(ctx.IfBranch().Expr())
	v.Visit(ctx.IfBranch().Block())

	// else if branch
	for _, b := range ctx.AllElifBranch() {
		v.Visit(b.Expr())
		v.Visit(b.Block())
	}

	// else branch
	if ctx.ElseBranch() != nil {
		v.Visit(ctx.ElseBranch().Block())
	}

	return nil
}

func (v *Visitor) VisitStatVarAssign(ctx *parser.StatVarAssignContext) any {
	lv := ctx.LValue()
	var atToken antlr.Token

	switch lv := lv.(type) {
	case *parser.LValueIDContext:
		atToken = lv.ID().GetSymbol()

	case *parser.LValueArrayAccessContext:
		// check lvalue for array access
		lhs := v.Visit(lv).(ExprAttribute)
		if lhs.Type.SameWith(Error) {
			return nil
		}
		atToken = lv.GetStart()

	default:
		panic("should not reach here")
	}

	lhs := v.Visit(lv).(ExprAttribute)
	rhs := v.Visit(ctx.Expr()).(ExprAttribute)

	// check type match
	if lhs.Type.SameWith(ToInfer) {
		lhs.Symbol.Infer(rhs.Type)
		v.LogInfer(lhs.Symbol, atToken)
	} else if !lhs.Type.SameWith(rhs.Type) {
		// 如果已经推断类型且不匹配，则报错
		err := NewSematicErr(TypeErr).
			Message("mismatch type assignment: <%s> != <%s>", lhs.Type, rhs.Type)
		v.LogError(err, atToken)
	} else if !lhs.Symbol.(BaseSymbol).Mutable() {
		// 如果lhs是常量
		err := NewSematicErr(AttrErr).Message("modify const var: %s", lhs.Symbol)
		v.LogError(err, atToken)
	}

	return nil
}

func (v *Visitor) VisitStatBlock(ctx *parser.StatBlockContext) any {
	return v.Visit(ctx.Block())
}
