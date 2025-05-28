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
	graph        *SymTableGraph
}

func NewVisitor() *Visitor {
	return &Visitor{
		graph: NewSymTableGraph(),
	}
}

func (v *Visitor) GetDotGraph() []byte {
	return v.graph.ToDot()
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
	line, col := atToken.GetLine(), atToken.GetColumn()
	log.Errorf("[%02d:%02d] %s", line, col, err)
}

func (v *Visitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *Visitor) VisitProg(ctx *parser.ProgContext) any {
	v.globalScope = NewGlobalScope(nil)
	v.currentScope = v.globalScope
	ret := v.Visit(ctx.Declaration())
	v.graph.AddNode(string(v.graph.ToScopeDot(v.currentScope)))
	v.currentScope = v.currentScope.Enclosed()
	return ret
}

func (v *Visitor) VisitBlock(ctx *parser.BlockContext) any {
	localScope := NewLocalScope(v.currentScope)
	v.graph.AddEdge(localScope.Name(), v.currentScope.Name())
	v.currentScope = localScope
	for _, stat := range ctx.AllStat() {
		v.Visit(stat)
	}
	v.graph.AddNode(string(v.graph.ToScopeDot(v.currentScope)))
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
	v.graph.AddEdge(funcScope.Name(), v.currentScope.Name())
	v.currentScope.Define(funcScope.GetSymbol()) // func scope定义在global之中
	v.currentScope = funcScope

	ret := v.Visit(ctx.FuncBlock())

	v.graph.AddNode(string(v.graph.ToScopeDot(v.currentScope)))
	v.currentScope = v.currentScope.Enclosed()

	return ret
}

func (v *Visitor) VisitStatFuncReturn(ctx *parser.StatFuncReturnContext) any {
	// NOTE: 获取func scope, 通过scope链条向上查找
	var funcScope FuncScope
	for s := range Iter(v.currentScope) {
		if _, ok := s.(FuncScope); ok {
			funcScope = s.(FuncScope)
			break
		}
	}
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
	v.currentScope.Define(varSymbol)
	v.LogDefine(varSymbol)
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
		log.Panicf("unknown symbol: %v", s)
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
		// NOTE: 如果类型不符合，默认按照int32处理
	}
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
		// NOTE: 如果类型不符合，默认按照int32处理
	}
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

func (v *Visitor) VisitStatWhile(ctx *parser.StatWhileContext) any {
	v.Visit(ctx.Expr())
	v.Visit(ctx.Block())
	return nil
}

func (v *Visitor) VisitStatLoop(ctx *parser.StatLoopContext) any {
	v.Visit(ctx.Block())
	return nil
}

func (v *Visitor) VisitStatIfElse(ctx *parser.StatIfElseContext) any {
	for i, expr := range ctx.AllExpr() {
		v.Visit(expr)
		v.Visit(ctx.Block(i))
	}

	// expr 会比 block 少一个
	blocks := ctx.AllBlock()
	v.Visit(blocks[len(blocks)-1])
	return nil
}

func (v *Visitor) VisitStatVarAssign(ctx *parser.StatVarAssignContext) any {
	tokenID := ctx.ID().GetSymbol()
	res := v.currentScope.Resolve(tokenID.GetText())
	v.LogResolve(res, ctx.ID().GetSymbol())

	attr := v.Visit(ctx.Expr()).(ExprAttribute)
	if attr.Type != res.Type() {
		err := NewSematicErr(TypeErr).Message("dismatch type assignment: <%s> != <%s>", attr.Type, res.Type())
		v.LogError(err, tokenID)
	}
	return nil
}

func (v *Visitor) VisitStatReturn(ctx *parser.StatFuncReturnContext) any {
	if ctx.Expr() != nil {
		v.Visit(ctx.Expr())
	}
	return nil
}

func (v *Visitor) VisitStatBlock(ctx *parser.StatBlockContext) any {
	return v.Visit(ctx.Block())
}
