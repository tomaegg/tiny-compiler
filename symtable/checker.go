package symtable

import (
	"github.com/antlr4-go/antlr/v4"
)

type SemanticChecker struct {
	v    *Visitor
	root antlr.ParseTree
}

func NewSemanticChecker(root antlr.ParseTree) *SemanticChecker {
	visitor := NewVisitor()
	return &SemanticChecker{v: visitor, root: root}
}

func (s *SemanticChecker) Check() (totalError int) {
	s.v.Visit(s.root)
	return s.v.TotalError()
}

func (s *SemanticChecker) SymbolTable() *SymTable {
	return s.v.symTable
}
