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
	v.currentScope = v.symTable.Scope(ctx.FuncSignature().ID().GetText())
	v.Visit(ctx.FuncSignature())
	v.Visit(ctx.FuncBlock())
	v.currentScope = v.currentScope.Enclosed()
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

func (v *Visitor) VisitStatVarDeclare(ctx *parser.StatVarDeclareContext) any {
	varSymbol := v.currentScope.Resolve(ctx.ID().GetText())
	val := v.llvmBuilder.CreateAlloca(v.LLVMType(varSymbol.Type()), "var_declare_tmp") // TODO: naming
	varSymbol.SetLLVMValue(val)                                                        // 在declare时分配空间, 加入到符号表中
	return nil
}

func (v *Visitor) VisitStatVarAssign(ctx *parser.StatVarAssignContext) any {
	lhsVal := v.currentScope.Resolve(ctx.ID().GetText()).LLVMValue()
	rhsVal := v.Visit(ctx.Expr()).(llvm.Value)
	v.llvmBuilder.CreateStore(rhsVal, lhsVal) // val -> value, p -> pointer
	return nil
}

func (v *Visitor) VisitExprNum(ctx *parser.ExprNumContext) any {
	val := llvm.ConstIntFromString(v.llvmCtx.Int32Type(), ctx.GetText(), 10)
	return val
}

func (v *Visitor) VisitExprAddSub(ctx *parser.ExprAddSubContext) any {
	lhs := v.Visit(ctx.GetLhs()).(llvm.Value)
	rhs := v.Visit(ctx.GetRhs()).(llvm.Value)
	var ret llvm.Value
	// NOTE: 目前值只有INT32
	switch ctx.GetOp().GetTokenType() {
	case parser.RustLikeLexerPLUS:
		ret = v.llvmBuilder.CreateAdd(lhs, rhs, "add_tmp") // TODO: naming
	case parser.RustLikeLexerMINUS:
		ret = v.llvmBuilder.CreateSub(lhs, rhs, "sub_tmp") // TODO: naming
	default:
		panic("should not get token other than +,-")
	}
	return ret
}

func (v *Visitor) VisitExprCmp(ctx *parser.ExprCmpContext) any {
	lhs := v.Visit(ctx.GetLhs()).(llvm.Value)
	rhs := v.Visit(ctx.GetRhs()).(llvm.Value)
	var name string
	var op llvm.IntPredicate
	// NOTE: 目前值只有INT32
	switch ctx.GetOp().GetTokenType() {
	case parser.RustLikeLexerLE:
		name = "le_tmp"
		op = llvm.IntSLE
	case parser.RustLikeLexerLT:
		name = "lt_tmp"
		op = llvm.IntSLT
	case parser.RustLikeLexerGE:
		name = "ge_tmp"
		op = llvm.IntSGE
	case parser.RustLikeLexerGT:
		name = "gt_tmp"
		op = llvm.IntSGT
	case parser.RustLikeLexerEQ:
		name = "eq_tmp"
		op = llvm.IntEQ
	case parser.RustLikeLexerNE:
		name = "ne_tmp"
		op = llvm.IntNE
	default:
		panic("should not get token other than <=,<,>=,>,==,!=")
	}
	ret := v.llvmBuilder.CreateICmp(op, lhs, rhs, name)
	return ret
}

func (v *Visitor) VisitExprMulDiv(ctx *parser.ExprMulDivContext) any {
	lhs := v.Visit(ctx.GetLhs()).(llvm.Value)
	rhs := v.Visit(ctx.GetRhs()).(llvm.Value)
	var ret llvm.Value
	// NOTE: 目前值只有INT32
	switch ctx.GetOp().GetTokenType() {
	case parser.RustLikeLexerMULT:
		ret = v.llvmBuilder.CreateMul(lhs, rhs, "mult_mp") // TODO: naming
	case parser.RustLikeLexerDIV:
		ret = v.llvmBuilder.CreateSDiv(lhs, rhs, "div_tmp") // TODO: naming
	default:
		panic("should not get token other than *,/")
	}
	return ret
}

func (v *Visitor) VisitFuncCallList(ctx *parser.FuncCallListContext) any {
	exprs := ctx.FuncCallParam().AllExpr()
	args := make([]llvm.Value, len(exprs))
	for i, e := range exprs {
		args[i] = v.Visit(e).(llvm.Value)
	}
	return args
}

func (v *Visitor) VisitExprFuncCall(ctx *parser.ExprFuncCallContext) any {
	fn := v.llvmMod.NamedFunction(ctx.ID().GetText())
	retType := fn.CalledFunctionType().ReturnType()
	args := v.Visit(ctx.FuncCallList()).([]llvm.Value)
	return v.llvmBuilder.CreateCall(retType, fn, args, "func_call_tmp") // TODO: naming
}

func (v *Visitor) VisitExprParen(ctx *parser.ExprParenContext) any {
	return v.Visit(ctx.Expr())
}

func (v *Visitor) VisitExprID(ctx *parser.ExprIDContext) any {
	varSymbol := v.currentScope.Resolve(ctx.ID().GetText())
	llvmType := v.LLVMType(varSymbol.Type())
	llvmVal := varSymbol.LLVMValue()
	v.llvmBuilder.CreateLoad(llvmType, llvmVal, "load_tmp") // TODO: naming
	return nil
}
