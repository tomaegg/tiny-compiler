package symtable

import (
	"tj-compiler/g4/parser"
)

var typeMap = map[string]SymType{
	"i32": SymInt32,
}

type Listener struct {
	parser.BaseRustLikeParserListener
	globalScope  Scope
	currentScope Scope
	graph        *SymTableGraph
}

var _ parser.RustLikeParserListener = (*Listener)(nil)

func NewListener() *Listener {
	return &Listener{
		graph: NewSymTableGraph(),
	}
}

func (l *Listener) GetDotGraph() []byte {
	return l.graph.ToDot()
}

// (1) 创建新的scope

func (l *Listener) EnterProg(ctx *parser.ProgContext) {
	l.globalScope = NewGlobalScope(nil)
	l.currentScope = l.globalScope
}

func (l *Listener) ExitProg(ctx *parser.ProgContext) {
	l.graph.AddNode(string(l.graph.ToScopeDot(l.currentScope)))
	l.currentScope = l.currentScope.Enclosed()
}

func (l *Listener) EnterBlock(ctx *parser.BlockContext) {
	// 需要判定是否继承父节点的func scope
	_, isFuncScope := l.currentScope.(FuncScope)
	if isFuncScope {
		return
	}
	localScope := NewLocalScope(l.currentScope)
	l.graph.AddEdge(localScope.Name(), l.currentScope.Name())
	l.currentScope = localScope
}

func (l *Listener) ExitBlock(ctx *parser.BlockContext) {
	_, isFuncScope := l.currentScope.(FuncScope)
	if isFuncScope {
		return
	}
	l.graph.AddNode(string(l.graph.ToScopeDot(l.currentScope)))
	l.currentScope = l.currentScope.Enclosed()
}

// func declare的scope应该和func param的scope是同一个, 因此进入func declare的时候
// 就应当建立scope,
func (l *Listener) EnterFuncDeclaration(ctx *parser.FuncDeclarationContext) {
	funcName := ctx.FuncDeclarationHeader().ID().GetText()
	funcScope := NewFuncScope(l.currentScope, funcName)
	l.graph.AddEdge(funcScope.Name(), l.currentScope.Name())
	l.currentScope = funcScope

	fps := ctx.FuncDeclarationHeader().FuncParamsList().FuncParams().AllFuncParam()
	params := make([]BaseSymbol, 0, len(fps))
	for _, fp := range fps {
		tokenType := fp.Rtype().INT32().GetSymbol()
		varType := typeMap[tokenType.GetText()]
		varName := fp.ID().GetText()
		mutable := fp.MUT() != nil
		varSymbol := NewBaseSymbol(varName, varType, mutable)
		l.currentScope.Define(varSymbol)
		params = append(params, varSymbol)
	}
	funcSymbol := NewFuncSymbol(funcName, l.currentScope, params)

	l.currentScope.Define(funcSymbol)
}

// (2) 创建新的Symbol

func (l *Listener) EnterStatVarDeclare(ctx *parser.StatVarDeclareContext) {
	tokenID := ctx.ID().GetSymbol()
	// NOTE: 声明时可以没有类型, 需要定义为to infer
	var varType SymType
	if ctx.VarType() == nil {
		varType = SymToInfer // NOTE: 此时标注为待推断
	} else {
		tokenType := ctx.VarType().Rtype().INT32().GetSymbol()
		varType = typeMap[tokenType.GetText()]
	}
	mutable := ctx.MUT() != nil
	varSymbol := NewBaseSymbol(tokenID.GetText(), varType, mutable)
	l.currentScope.Define(varSymbol)
}

// (3) 使用符号
func (l *Listener) EnterExprID(ctx *parser.ExprIDContext) {
	varName := ctx.ID().GetText()
	l.currentScope.Resolve(varName)
}

func (l *Listener) ExitFuncDeclaration(ctx *parser.FuncDeclarationContext) {
	l.graph.AddNode(string(l.graph.ToScopeDot(l.currentScope)))
	l.currentScope = l.currentScope.Enclosed()
}
