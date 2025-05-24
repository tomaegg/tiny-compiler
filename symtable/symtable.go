package symtable

import "github.com/antlr4-go/antlr/v4"

type SymTable struct {
	l *Listener
}

func NewSymTable(root antlr.Tree) *SymTable {
	walker := antlr.NewParseTreeWalker()
	listener := NewListener()
	walker.Walk(listener, root)
	return &SymTable{l: listener}
}

func (s *SymTable) DotGraph() []byte {
	return s.l.GetDotGraph()
}
