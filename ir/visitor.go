package ir

import (
	"tj-compiler/g4/parser"
	"tj-compiler/symtable"

	"github.com/antlr4-go/antlr/v4"
	log "github.com/sirupsen/logrus"

	"tinygo.org/x/go-llvm"
)

func (v *Visitor) LLVMType(t symtable.SymType) llvm.Type {
	ctx := v.llvmCtx
	switch t {
	case symtable.SymInt32:
		return ctx.Int32Type()
	case symtable.SymVoid:
		return ctx.VoidType()
	default:
		log.Panicf("unsupported type: %s", t)
	}
	return ctx.VoidType()
}

type Visitor struct {
	parser.BaseRustLikeParserVisitor

	// llvm context
	llvmCtx     llvm.Context
	llvmMod     llvm.Module
	llvmBuilder llvm.Builder

	// current llvm block
	currentBB llvm.BasicBlock

	// symbol table prepared
	symTable     *symtable.SymTable
	currentScope symtable.Scope
}

func NewVisitor(module string, symTable *symtable.SymTable) (v *Visitor, cancel func()) {
	ctx := llvm.NewContext()
	mod := ctx.NewModule(module)
	builder := ctx.NewBuilder()

	v = &Visitor{
		symTable:    symTable,
		llvmCtx:     ctx,
		llvmMod:     mod,
		llvmBuilder: builder,
	}

	cancel = func() {
		v.llvmCtx.Dispose()
		v.llvmBuilder.Dispose()
	}

	return
}

func (v *Visitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *Visitor) VisitProg(ctx *parser.ProgContext) any {
	v.currentScope = v.symTable.GlobalScope()
	ret := v.Visit(ctx.Declaration())
	v.currentScope = v.currentScope.Enclosed()
	return ret
}

func (v *Visitor) VisitDeclaration(ctx *parser.DeclarationContext) any {
	for _, fd := range ctx.AllFuncDeclaration() {
		v.Visit(fd)
	}
	return nil
}

func (v *Visitor) VisitFuncDeclaration(ctx *parser.FuncDeclarationContext) any {
	v.Visit(ctx.FuncSignature())
	v.Visit(ctx.FuncBlock())
	return nil
}

func (v *Visitor) VisitFuncSignature(ctx *parser.FuncSignatureContext) any {
	funcToken := ctx.ID().GetSymbol()
	funcName := funcToken.GetText()
	funcScope := v.symTable.Scope(funcName).(symtable.FuncScope)
	funcSymbol := funcScope.GetSymbol()
	log.Debugf("FuncScope: %s", funcScope.Name())

	// define function
	paramTypes := make([]llvm.Type, len(funcSymbol.Params()))
	for i, p := range funcSymbol.Params() {
		paramTypes[i] = v.LLVMType(p.Type())
	}

	retType := v.LLVMType(funcSymbol.RetType())
	fnType := llvm.FunctionType(retType, paramTypes, false)
	fn := llvm.AddFunction(v.llvmMod, funcName, fnType)

	// define function basic entry block
	fnBlock := v.llvmCtx.AddBasicBlock(fn, "entry")
	v.currentBB = fnBlock
	return nil
}

func (v *Visitor) VisitFuncBlock(ctx *parser.FuncBlockContext) any {
	for _, stat := range ctx.AllStat() {
		v.Visit(stat)
	}
	return nil
}
