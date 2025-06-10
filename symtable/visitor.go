package symtable

import (
	"strconv"
	"tj-compiler/g4/parser"

	"github.com/antlr4-go/antlr/v4"
	log "github.com/sirupsen/logrus"
)

var typeMap = map[string]SymType{
	"i32": SymInt32,
}

type ExprAttribute struct {
	Type  SymType
	Value any
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
	line, col := token.GetLine(), token.GetColumn()
	log.Infof("[%02d:%02d] define: %s", line, col, s.String())
}

func (v *Visitor) LogResolve(res Symbol, atToken antlr.Token) {
	line, col := atToken.GetLine(), atToken.GetColumn()
	if res == nil {
		log.Fatalf("unresolved symbol at token[%d:%d]: <%s>", line, col, atToken.GetText())
	}
	log.Infof("[%02d:%02d] resolve: %s", line, col, res.String())
}

func (v *Visitor) LogError(err SemanticErr, atToken antlr.Token) {
	v.errorCount++
	line, col := atToken.GetLine(), atToken.GetColumn()
	log.Errorf("[%02d:%02d] %s", line, col, err)
}

func (v *Visitor) LogInfer(s Symbol, at antlr.Token) {
	log.Infof("[%02d:%02d] infer: %s",
		at.GetLine(), at.GetColumn(), s)
}

func (v *Visitor) checkFuncReturn(scope FuncScope) {
	funcSymbol := scope.GetSymbol()
	returned := scope.HasReturn()
	funcHasReturn := funcSymbol.RetType() != SymVoid
	if funcHasReturn && !returned {
		err := NewSematicErr(RetErr).
			Message("FuncScope <%s> should have return statement, but not", funcSymbol.Name())
		v.LogError(err, funcSymbol.Token())
	}
}

func (v *Visitor) checkUninferredSymbol(scope Scope) {
	for _, sym := range scope.Symbols() {
		if sym.Type() == SymToInfer {
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
	retType := typeMap[ctx.INT32().GetSymbol().GetText()]
	return retType
}

func (v *Visitor) VisitFuncSignature(ctx *parser.FuncSignatureContext) any {
	funcToken := ctx.ID().GetSymbol()
	funcName := funcToken.GetText()

	params := v.Visit(ctx.FuncParamsList()).([]BaseSymbol)
	retType := SymVoid
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
	getType := SymVoid
	if ctx.Expr() != nil {
		attr := v.Visit(ctx.Expr()).(ExprAttribute)
		getType = attr.Type
	}

	if wantType != getType {
		err := NewSematicErr(TypeErr).Message(
			"function return type dismatch: want <%s>, get <%s>",
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
	var varType SymType
	if ctx.VarType() == nil {
		varType = SymToInfer // NOTE: 此时标注为待推断
	} else {
		varType = v.Visit(ctx.VarType().Rtype()).(SymType)
	}
	mutable := ctx.MUT() != nil
	varSymbol := NewBaseSymbol(tokenID.GetText(), varType, mutable, tokenID)

	if v.currentScope.Exist(varSymbol) {
		v.currentScope.Shallow(varSymbol)
	} else {
		v.currentScope.Define(varSymbol)
	}

	v.LogDefine(varSymbol)

	// varinit exist
	if ctx.VarInit() != nil {
		// 访问varinit得到属性
		attr := v.Visit(ctx.VarInit().Expr()).(ExprAttribute)
		switch varSymbol.Type() {
		case SymToInfer:
			// 当前为待推断
			varSymbol.Infer(attr.Type)
			v.LogInfer(varSymbol, tokenID)
		default:
			err := NewSematicErr(TypeErr).
				Message("dismatch type assignment: <%s> != <%s>", varSymbol.Type(), attr.Type)
			v.LogError(err, tokenID)
		}
	}
	return nil
}

// different expression

func (v *Visitor) VisitExprID(ctx *parser.ExprIDContext) any {
	s := v.currentScope.Resolve(ctx.ID().GetSymbol().GetText())
	var t SymType
	switch s := s.(type) {
	case BaseSymbol:
		t = s.Type()
	case FuncSymbol:
		t = SymFunc
	default:
		log.Fatalf("use unknown symbol: <%v>", ctx.ID().GetText())
	}
	return ExprAttribute{Type: t}
}

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
		if attr.Type != paramsRequired[i].Type() {
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
	if attrLhs.Type != attrRhs.Type {
		err := NewSematicErr(TypeErr).Message("type mismatch with: <%s> %s <%s>",
			attrLhs.Type, op.GetText(), attrRhs.Type,
		)
		v.LogError(err, op)
		// NOTE: 如果类型不符合，默认按照int32处理
		return ExprAttribute{Type: SymError}
	}
	// NOTE: 由于暂时不支持bool类型，因此返回Int32
	return ExprAttribute{Type: SymInt32}
}

func (v *Visitor) VisitExprMulDiv(ctx *parser.ExprMulDivContext) any {
	attrLhs := v.Visit(ctx.GetLhs()).(ExprAttribute)
	attrRhs := v.Visit(ctx.GetRhs()).(ExprAttribute)
	op := ctx.GetOp()
	if attrLhs.Type != attrRhs.Type {
		err := NewSematicErr(TypeErr).Message("type mismatch: <%s> %s <%s>",
			attrLhs.Type, op.GetText(), attrRhs.Type,
		)
		v.LogError(err, op)
		// 如果类型不一致返回 error 类型
		return ExprAttribute{Type: SymError}
	}
	// 类型一致返回 int32 类型
	return ExprAttribute{Type: SymInt32}
}

func (v *Visitor) VisitExprAddSub(ctx *parser.ExprAddSubContext) any {
	attrLhs := v.Visit(ctx.GetLhs()).(ExprAttribute)
	attrRhs := v.Visit(ctx.GetRhs()).(ExprAttribute)
	op := ctx.GetOp()
	if attrLhs.Type != attrRhs.Type {
		err := NewSematicErr(TypeErr).Message("type mismatch: <%s> %s <%s>",
			attrLhs.Type, op.GetText(), attrRhs.Type,
		)
		v.LogError(err, op)
		// 如果类型不一致返回 undefined 类型
		return ExprAttribute{Type: SymError}
	}
	// 类型一致返回 int32 类型
	return ExprAttribute{Type: SymInt32}
}

func (v *Visitor) VisitExprParen(ctx *parser.ExprParenContext) any {
	return v.Visit(ctx.Expr())
}

func (v *Visitor) VisitExprNum(ctx *parser.ExprNumContext) any {
	// TODO: check int
	numToken := ctx.NUMBER().GetSymbol()
	val, err := strconv.ParseInt(numToken.GetText(), 10, 32)
	if err != nil {
		err := NewSematicErr(IntOverflowErr).Message("overflow int32: %s", numToken.GetText())
		v.LogError(err, numToken)
	}
	return ExprAttribute{Type: SymInt32, Value: val}
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
	tokenID := ctx.ID().GetSymbol()
	res := v.currentScope.Resolve(tokenID.GetText())
	v.LogResolve(res, ctx.ID().GetSymbol())

	attr := v.Visit(ctx.Expr()).(ExprAttribute)
	if res.Type() == SymToInfer {
		res.Infer(attr.Type)
		v.LogInfer(res, tokenID)
	} else if res.Type() != attr.Type {
		// 如果已经推断类型且不匹配，则报错
		err := NewSematicErr(TypeErr).
			Message("dismatch type assignment: <%s> != <%s>", res.Type(), attr.Type)
		v.LogError(err, tokenID)
	}
	return nil
}

func (v *Visitor) VisitStatBlock(ctx *parser.StatBlockContext) any {
	return v.Visit(ctx.Block())
}
