package symtable

import (
	"github.com/antlr4-go/antlr/v4"
)

type SemanticChecker struct {
	v *Visitor
}

func NewSemanticChecker(root antlr.ParseTree) *SemanticChecker {
	visitor := NewVisitor()
	visitor.Visit(root)
	visitor.symTable.SetTree(root)
	return &SemanticChecker{v: visitor}
}

func (s *SemanticChecker) SymbolTable() *SymTable {
	return s.v.symTable
}

func (s *SemanticChecker) TotalErrors() int {
	return s.v.TotalError()
}
