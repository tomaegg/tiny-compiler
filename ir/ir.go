package ir

import (
	"tj-compiler/symtable"

	"github.com/antlr4-go/antlr/v4"
)

type IRGenerator struct {
	*Visitor
	generated bool
	tree      antlr.ParseTree
}

func NewIRGenerator(module string, symTable *symtable.SymTable, tree antlr.ParseTree) (*IRGenerator, func()) {
	v, f := NewVisitor(module, symTable)
	return &IRGenerator{Visitor: v, tree: tree}, f
}

func (ig *IRGenerator) generate() {
	if ig.generated {
		return
	}
	ig.generated = true
	ig.Visit(ig.tree)
}

func (ig *IRGenerator) IR() []byte {
	ig.generate()

	// Print the module to human-readable IR (stdout)
	return []byte(ig.llvmMod.String())
}
