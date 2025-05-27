package symtable

import (
	"tj-compiler/g4/parser"

	"github.com/antlr4-go/antlr/v4"
	log "github.com/sirupsen/logrus"
)

var typeMap = map[string]SymType{
	"i32": SymInt32,
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

func (v *Visitor) LogResolve(res Symbol) {
	s := res.Token()
	line, col := s.GetLine(), s.GetColumn()
	if res == nil {
		log.Fatalf("unresolved symbol at token[%d:%d]: <%s>", line, col, s.GetText())
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
	funcScope := v.Visit(ctx.FuncSignature()).(*FuncScope)
	v.graph.AddEdge(funcScope.Name(), v.currentScope.Name())
	v.currentScope.Define(funcScope.GetSymbol()) // func scope定义在global之中
	v.currentScope = funcScope

	ret := v.Visit(ctx.FuncBlock())

	v.graph.AddNode(string(v.graph.ToScopeDot(v.currentScope)))
	v.currentScope = v.currentScope.Enclosed()

	return ret
}

func (v *Visitor) VisitStatFuncReturn(ctx *parser.StatFuncReturnContext) any {
	tokenRet := ctx.RETURN().GetSymbol()
	getType := SymVoid
	if ctx.Expr() != nil {
		// TODO: type checking
		v.Visit(ctx)
	}

	funcScope := v.currentScope.(*FuncScope)
	funcSymbol := funcScope.GetSymbol()
	wantType := funcSymbol.RetType()

	if wantType != getType {
		err := NewSematicErr(TypeErr).Message(
			"function return type dismatch: want %s, get %s",
			wantType, getType,
		)
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
	v.currentScope.Resolve(ctx.ID().GetSymbol().GetText())
	return nil
}

func (v *Visitor) VisitExprMulDiv(ctx *parser.ExprMulDivContext) any {
	v.Visit(ctx.GetLhs())
	v.Visit(ctx.GetRhs())
	return nil
}

func (v *Visitor) VisitExprAddSub(ctx *parser.ExprAddSubContext) any {
	v.Visit(ctx.GetLhs())
	v.Visit(ctx.GetRhs())
	return nil
}

func (v *Visitor) VisitExprParen(ctx *parser.ExprParenContext) any {
	return v.Visit(ctx.Expr())
}

func (v *Visitor) VisitExprNum(ctx *parser.ExprNumContext) any {
	// TODO: check int
	return nil
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
	v.LogResolve(res)

	v.Visit(ctx.Expr())
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
