package ir

import (
	"tj-compiler/symtable"

	"github.com/antlr4-go/antlr/v4"
	log "github.com/sirupsen/logrus"
	"tinygo.org/x/go-llvm"
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
		log.Debug("skip for generated done")
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

func (ig *IRGenerator) Module() llvm.Module {
	ig.generate()
	return ig.llvmMod
}
