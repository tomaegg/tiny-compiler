package ir

import (
	"tj-compiler/g4/parser"
	"tj-compiler/symtable"

	"github.com/antlr4-go/antlr/v4"
	log "github.com/sirupsen/logrus"

	"tinygo.org/x/go-llvm"
)

func (v *Visitor) LLVMType(t symtable.SymType) llvm.Type {
	ctx := v.ctx
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
	ctx     llvm.Context
	mod     llvm.Module
	builder llvm.Builder

	// symbol table prepared
	symTable     *symtable.SymTable
	currentScope symtable.Scope
}

func NewVisitor(symTable *symtable.SymTable) (v *Visitor, cancel func()) {
	ctx := llvm.NewContext()
	mod := ctx.NewModule("example") // TODO: module name
	builder := ctx.NewBuilder()

	v = &Visitor{
		symTable: symTable,
		ctx:      ctx,
		mod:      mod,
		builder:  builder,
	}

	cancel = func() {
		v.ctx.Dispose()
		v.builder.Dispose()
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
	ret := v.Visit(ctx.FuncSignature())
	// TODO:
	return ret
}

func (v *Visitor) VisitFuncSignature(ctx *parser.FuncSignatureContext) any {
	funcToken := ctx.ID().GetSymbol()
	funcName := funcToken.GetText()
	funcScope := v.symTable.Scope(funcName).(symtable.FuncScope)
	funcSymbol := funcScope.GetSymbol()
	log.Debugf("FuncScope: %s", funcScope.Name())

	// TODO: you know do what
	for range funcSymbol.Params() {
	}

	// NOTE: 只是demo, 用于适配环境，后续需要修改
	retType := v.LLVMType(funcSymbol.RetType())
	fnType := llvm.FunctionType(retType, nil, false)
	fn := llvm.AddFunction(v.mod, funcName, fnType)
	block := v.ctx.AddBasicBlock(fn, "entry")

	v.builder.SetInsertPointAtEnd(block)
	v.builder.CreateRet(llvm.ConstInt(retType, 42, false))

	return funcScope
}
