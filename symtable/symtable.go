package symtable

import "github.com/antlr4-go/antlr/v4"

type SymTable struct{}

func NewSymTable(tree antlr.Tree) *SymTable {
	walker := antlr.NewParseTreeWalker()
	listner := NewListener()
	walker.Walk(listner, tree)
	// TODO:
	return nil
}
