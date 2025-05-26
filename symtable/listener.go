package symtable

import (
	"tj-compiler/g4/parser"

	"github.com/antlr4-go/antlr/v4"
	log "github.com/sirupsen/logrus"
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

func (l *Listener) Define(s Symbol, line, col int) {
	l.currentScope.Define(s)
	log.Infof("[%02d:%02d] define: %s", line, col, s.String())
}

func (l *Listener) Resolve(s antlr.Token) {
	varName := s.GetText()
	line, col := s.GetLine(), s.GetColumn()
	symbol := l.currentScope.Resolve(varName)
	log.Infof("[%02d:%02d] resolve: %s", line, col, symbol.String())
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
	funcToken := ctx.FuncDeclarationHeader().ID().GetSymbol()
	funcName := funcToken.GetText()
	funcScope := NewFuncScope(l.currentScope, funcName)
	l.graph.AddEdge(funcScope.Name(), l.currentScope.Name())
	l.currentScope = funcScope

	fps := ctx.FuncDeclarationHeader().FuncParamsList().FuncParams().AllFuncParam()
	params := make([]BaseSymbol, 0, len(fps))
	for _, fp := range fps {
		tokenID := fp.ID().GetSymbol()
		tokenType := fp.Rtype().INT32().GetSymbol()
		varType := typeMap[tokenType.GetText()]
		varName := tokenID.GetText()
		mutable := fp.MUT() != nil
		varSymbol := NewBaseSymbol(varName, varType, mutable)
		l.Define(varSymbol, tokenID.GetLine(), tokenID.GetColumn())
		params = append(params, varSymbol)
	}
	funcSymbol := NewFuncSymbol(funcName, l.currentScope, params)

	l.Define(funcSymbol, funcToken.GetLine(), funcToken.GetColumn())
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
	l.Define(varSymbol, tokenID.GetLine(), tokenID.GetColumn())
}

// (3) 使用符号
func (l *Listener) EnterExprID(ctx *parser.ExprIDContext) {
	l.Resolve(ctx.ID().GetSymbol())
}

func (l *Listener) ExitFuncDeclaration(ctx *parser.FuncDeclarationContext) {
	l.graph.AddNode(string(l.graph.ToScopeDot(l.currentScope)))
	l.currentScope = l.currentScope.Enclosed()
}
