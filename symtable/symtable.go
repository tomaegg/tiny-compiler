package symtable

import "github.com/antlr4-go/antlr/v4"

type SymTable struct {
	v *Visitor
}

func NewSymTable(root antlr.ParseTree) *SymTable {
	visitor := NewVisitor()
	visitor.Visit(root)
	return &SymTable{v: visitor}
}

func (s *SymTable) DotGraph() []byte {
	return s.v.GetDotGraph()
}
