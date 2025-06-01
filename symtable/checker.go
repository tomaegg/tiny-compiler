package symtable

import (
	"github.com/antlr4-go/antlr/v4"
)

type SemanticChecker struct {
	*Visitor
}

func NewSemanticChecker(root antlr.ParseTree) *SemanticChecker {
	visitor := NewVisitor()
	visitor.Visit(root)
	visitor.symTable.SetTree(root)
	return &SemanticChecker{Visitor: visitor}
}

func (s *SemanticChecker) SymbolTable() *SymTable {
	return s.symTable
}
