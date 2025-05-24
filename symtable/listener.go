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
}

var _ parser.RustLikeParserListener = (*Listener)(nil)

func NewListener() *Listener {
	return &Listener{}
}

// (1) 创建新的scope

func (l *Listener) EnterProg(ctx *parser.ProgContext) {
	l.globalScope = NewGlobalScope(nil)
	l.currentScope = l.globalScope
}

func (l *Listener) EnterBlock(ctx *parser.BlockContext) {
	// 需要判定是否继承父节点的func scope
	_, isFuncScope := l.currentScope.(FuncScope)
	if isFuncScope {
		return
	}
	localScope := NewLocalScope(l.currentScope)
	l.currentScope = localScope
}

// func declare的scope应该和func param的scope是同一个, 因此进入func declare的时候
// 就应当建立scope,
func (l *Listener) EnterFuncDeclaration(ctx *parser.FuncDeclarationContext) {
	funcName := ctx.FuncDeclarationHeader().ID().GetText()
	funcScope := NewFuncScope(l.currentScope, funcName)
	l.currentScope = funcScope

	funcSymbol := NewFuncSymbol(funcName, l.currentScope)
	l.currentScope.Define(funcSymbol)
}

// (2) 创建新的Symbol

func (l *Listener) EnterStatVarDeclare(ctx *parser.StatVarDeclareContext) {
	varName := ctx.ID().GetText()
	// NOTE: 声明时可以没有类型, 需要定义为to infer
	var varType SymType
	if ctx.VarType().Rtype().IsEmpty() {
		varType = SymToInfer // NOTE: 此时标注为待推断
	} else {
		varType = typeMap[ctx.VarType().Rtype().GetText()]
	}
	mutable := ctx.MUT() != nil
	varSymbol := NewBaseSymbol(varName, varType, mutable)
	l.currentScope.Define(varSymbol)
}

func (l *Listener) EnterFuncParam(ctx *parser.FuncParamContext) {
	varName := ctx.ID().GetText()
	varType := typeMap[ctx.Rtype().GetText()]
	mutable := ctx.MUT() != nil
	varSymbol := NewBaseSymbol(varName, varType, mutable)
	l.currentScope.Define(varSymbol)
}

// (3) 使用符号
func (l *Listener) EnterExprID(ctx *parser.ExprIDContext) {
	varName := ctx.ID().GetText()
	l.currentScope.Resolve(varName)
}

// (4) 回到father
func (l *Listener) ExitProg(ctx *parser.ProgContext) {
	l.currentScope = l.currentScope.Enclosed()
}

func (l *Listener) ExitBlock(ctx *parser.BlockContext) {
	_, isFuncScope := l.currentScope.(FuncScope)
	if isFuncScope {
		return
	}
	l.currentScope = l.currentScope.Enclosed()
}

func (l *Listener) ExitFuncDeclaration(ctx *parser.FuncDeclarationContext) {
	l.currentScope = l.currentScope.Enclosed()
}
