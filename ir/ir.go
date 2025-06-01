package ir

import "tj-compiler/symtable"

type IRGenerator struct {
	*Visitor
	generated bool
}

func NewIRGenerator(symTable *symtable.SymTable) (*IRGenerator, func()) {
	v, f := NewVisitor(symTable)
	return &IRGenerator{Visitor: v}, f
}

func (ig *IRGenerator) generate() {
	if ig.generated {
		return
	}
	ig.generated = true
	tree := ig.symTable.Tree()
	ig.Visit(tree)
}

func (ig *IRGenerator) IR() []byte {
	ig.generate()

	// Print the module to human-readable IR (stdout)
	return []byte(ig.mod.String())
}
