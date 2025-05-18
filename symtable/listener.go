package symtable

import "tj-compiler/g4/parser"

type Listener struct {
	parser.BaseRustLikeParserListener
	globalScope  Scope
	currentScope Scope
}

func (l *Listener) EnterProg(ctx *parser.ProgContext) {
}

func (l *Listener) EnterBlock(ctx *parser.BlockContext) {
}

func (l *Listener) EnterFuncDeclaration(ctx *parser.FuncDeclarationContext) {
}
