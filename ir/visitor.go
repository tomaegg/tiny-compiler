package ir

import (
	"fmt"
	"tj-compiler/g4/parser"
	"tj-compiler/symtable"
	"tj-compiler/utils/dsa"

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
	currentFn llvm.Value

	// symbol table prepared
	symTable     *symtable.SymTable
	currentScope symtable.Scope

	// if count
	counter map[string]int
	// error count
	errorCount int
	// loop Basic Block
	loopBB dsa.Stack[[2]llvm.BasicBlock] // (exit, header)
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
		counter:     make(map[string]int),
	}

	cancel = func() {
		v.llvmCtx.Dispose()
		v.llvmBuilder.Dispose()
	}

	return
}

func (v *Visitor) LogError(err symtable.SemanticErr, atToken antlr.Token) {
	v.errorCount++
	line, col := atToken.GetLine(), atToken.GetColumn()
	log.Errorf("[%02d:%02d] %s", line, col, err)
}

func (v *Visitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *Visitor) VisitProg(ctx *parser.ProgContext) any {
	v.currentScope = v.symTable.GlobalScope()
	ret := v.Visit(ctx.Declaration())
	v.currentScope = v.currentScope.Enclosed()

	// check module
	llvm.VerifyModule(v.llvmMod, llvm.AbortProcessAction)
	return ret
}

func (v *Visitor) VisitDeclaration(ctx *parser.DeclarationContext) any {
	for _, fd := range ctx.AllFuncDeclaration() {
		v.Visit(fd)
	}
	return nil
}

func (v *Visitor) VisitFuncDeclaration(ctx *parser.FuncDeclarationContext) any {
	funcName := ctx.FuncSignature().ID().GetText()
	funcScope := v.symTable.Scope(funcName).(symtable.FuncScope)
	funcSymbol := funcScope.GetSymbol()
	v.currentScope = funcScope
	v.Visit(ctx.FuncSignature())
	v.Visit(ctx.FuncBlock())
	v.currentScope = v.currentScope.Enclosed()

	// 对于return void 为每一个basic block添加return, 针对void函数
	isVoid := GetType(funcSymbol).ReturnType() == v.llvmCtx.VoidType()
	for _, bb := range v.currentFn.BasicBlocks() {
		inst := bb.LastInstruction()
		if isVoid && !IsATerminator(inst) { // 如果是void, 那么优先加return而不是unreachable
			// 修复return
			v.llvmBuilder.SetInsertPointAtEnd(bb)
			v.llvmBuilder.CreateRetVoid()
			continue
		}

		if inst.IsNil() {
			// 意味着空语句，create unreachable
			v.llvmBuilder.SetInsertPointAtEnd(bb)
			v.llvmBuilder.CreateUnreachable()
			log.Debug("found nil inst, created unreachable tag")
			continue
		}
		if IsATerminator(inst) {
			continue
		}
		err := symtable.NewSematicErr(symtable.RetErr).
			Message("missing return in function, basic block: %s", bb.AsValue().String())
		v.LogError(err, ctx.FuncSignature().ID().GetSymbol())
		log.Fatal("unrecoverable error detected, aborted")
	}

	// check function
	fn := v.llvmMod.NamedFunction(funcName)
	llvm.VerifyFunction(fn, llvm.PrintMessageAction)
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
	// bind function type with symbol
	SetType(funcSymbol, fnType)

	// define function basic entry block
	fnBlock := v.llvmCtx.AddBasicBlock(fn, "entry")
	v.llvmBuilder.SetInsertPointAtEnd(fnBlock)
	v.currentFn = fn
	return fnType
}

func (v *Visitor) VisitFuncBlock(ctx *parser.FuncBlockContext) any {
	return v.flattenBlock(ctx.AllStat())
}

func (v *Visitor) VisitBlock(ctx *parser.BlockContext) any {
	return v.flattenBlock(ctx.AllStat())
}

// true: 发现function return
// false: 没有发现function return
func (v *Visitor) flattenBlock(stats []parser.IStatContext) bool {
	// 去除嵌套作用域, 也就是StatBlock得到的序列
	for _, stat := range stats {
		switch stat := stat.(type) {
		case *parser.StatFuncReturnContext:
			v.Visit(stat)
			token := stat.GetStart()
			log.Debugf("[%d:%d] first return found in block near token <%s>, block ends",
				token.GetLine(), token.GetColumn(), token.GetText(),
			)
			return true // 代表找到了了function return
		case *parser.StatBreakContext:
			v.Visit(stat)
			token := stat.GetStart()
			log.Debugf("[%d:%d] first break found in block near token <%s>, block ends",
				token.GetLine(), token.GetColumn(), token.GetText(),
			)
			return true // 代表找到了了break
		case *parser.StatContinueContext:
			v.Visit(stat)
			token := stat.GetStart()
			log.Debugf("[%d:%d] first continue found in block near token <%s>, block ends",
				token.GetLine(), token.GetColumn(), token.GetText(),
			)
			return true // 代表找到了了continue
		case *parser.StatBlockContext: // block则进行flatten
			if v.Visit(stat.Block()).(bool) {
				return true
			}
		default:
			v.Visit(stat)
		}
	}
	return false
}

func (v *Visitor) VisitStatVarDeclare(ctx *parser.StatVarDeclareContext) any {
	varSymbol := v.currentScope.Resolve(ctx.ID().GetText())
	val := v.llvmBuilder.CreateAlloca(v.LLVMType(varSymbol.Type()), "var_"+varSymbol.Name())
	SetValue(varSymbol, val) // 在declare时分配空间, 加入到符号表中
	if ctx.VarInit() != nil {
		// TODO:
		rhs := v.Visit(ctx.VarInit().Expr()).(llvm.Value)
		v.llvmBuilder.CreateStore(rhs, val)
	}
	return nil
}

func (v *Visitor) getCondi(value llvm.Value) llvm.Value {
	width := value.Type().IntTypeWidth()
	if width > 1 {
		// 整数 -> i1：通过比较非零实现
		zero := llvm.ConstInt(value.Type(), 0, false)
		return v.llvmBuilder.CreateICmp(llvm.IntNE, value, zero, "cmp")
	}
	// i1 无需转换
	return value
}

// TODO: else if block
func (v *Visitor) VisitStatIfElse(ctx *parser.StatIfElseContext) any {
	label := func(l string) string {
		ret := fmt.Sprintf("if.%s.%d", l, v.counter[l])
		v.counter[l]++
		return ret
	}

	// if condition
	ifCondiValue := v.Visit(ctx.IfBranch().Expr()).(llvm.Value)
	ifCondi := v.getCondi(ifCondiValue)

	// merge block
	mergeBlock := v.llvmCtx.AddBasicBlock(v.currentFn, label("merge"))
	// then block
	thenBlock := v.llvmCtx.AddBasicBlock(v.currentFn, label("then"))
	elseBlock := mergeBlock
	if ctx.ElseBranch() != nil {
		elseBlock = v.llvmCtx.AddBasicBlock(v.currentFn, label("else"))
	}

	v.llvmBuilder.CreateCondBr(ifCondi, thenBlock, elseBlock)

	// 填充块
	// then block
	v.llvmBuilder.SetInsertPointAtEnd(thenBlock)
	if !v.Visit(ctx.IfBranch().Block()).(bool) {
		v.llvmBuilder.CreateBr(mergeBlock)
	}

	// else block
	if ctx.ElseBranch() != nil {
		v.llvmBuilder.SetInsertPointAtEnd(elseBlock)
		if !v.Visit(ctx.ElseBranch().Block()).(bool) {
			v.llvmBuilder.CreateBr(mergeBlock)
		}
	}

	// go to merge block
	v.llvmBuilder.SetInsertPointAtEnd(mergeBlock)
	return nil
}

func (v *Visitor) VisitStatBlock(ctx *parser.StatBlockContext) any {
	return v.Visit(ctx.Block())
}

func (v *Visitor) VisitStatExpr(ctx *parser.StatExprContext) any {
	return v.Visit(ctx.Expr())
}

func (v *Visitor) VisitStatVarAssign(ctx *parser.StatVarAssignContext) any {
	lhsVal := GetValue(v.currentScope.Resolve(ctx.ID().GetText()))
	rhsVal := v.Visit(ctx.Expr()).(llvm.Value)
	v.llvmBuilder.CreateStore(rhsVal, lhsVal) // val -> value, p -> pointer
	return nil
}

func (v *Visitor) VisitStatFuncReturn(ctx *parser.StatFuncReturnContext) any {
	if ctx.Expr() == nil {
		return v.llvmBuilder.CreateRetVoid()
	}
	val := v.Visit(ctx.Expr()).(llvm.Value)
	v.llvmBuilder.CreateRet(val)
	return nil
}

func (v *Visitor) VisitStatLoop(ctx *parser.StatLoopContext) any {
	// create a basic block
	loopBody := v.llvmCtx.AddBasicBlock(v.currentFn, "loop_body")
	loopHeader := loopBody
	// else block, 无论是否存在，创建空的即可
	exitBlock := v.llvmCtx.AddBasicBlock(v.currentFn, ("loop_exit"))
	v.llvmBuilder.CreateBr(loopBody) // branch to loop basic block
	v.llvmBuilder.SetInsertPointAtEnd(loopBody)

	// push block
	v.loopBB.Push([2]llvm.BasicBlock{exitBlock, loopHeader}) // exit, header

	if !v.Visit(ctx.Block()).(bool) {
		// 加入回边
		v.llvmBuilder.CreateBr(loopBody) // branch to itself
	}

	v.loopBB.Pop()

	// 生成后续
	v.llvmBuilder.SetInsertPointAtEnd(exitBlock)

	return nil
}

func (v *Visitor) VisitStatWhile(ctx *parser.StatWhileContext) any {
	whileHeader := v.llvmCtx.AddBasicBlock(v.currentFn, "while.header")
	whileBody := v.llvmCtx.AddBasicBlock(v.currentFn, "while.body")
	whileExit := v.llvmCtx.AddBasicBlock(v.currentFn, "while.exit")

	v.llvmBuilder.CreateBr(whileHeader)

	v.llvmBuilder.SetInsertPointAtEnd(whileHeader)
	condi := v.getCondi(v.Visit(ctx.Expr()).(llvm.Value))
	v.llvmBuilder.CreateCondBr(condi, whileBody, whileExit)

	// 进入while body
	v.llvmBuilder.SetInsertPointAtEnd(whileBody)

	// push block
	v.loopBB.Push([2]llvm.BasicBlock{whileExit, whileHeader})

	if !v.Visit(ctx.Block()).(bool) {
		// 回边, 每一次都需要计算
		v.llvmBuilder.CreateBr(whileHeader)
	}

	// pop block
	v.loopBB.Pop()

	// 生成后续
	v.llvmBuilder.SetInsertPointAtEnd(whileExit)

	return nil
}

func (v *Visitor) VisitStatContinue(ctx *parser.StatContinueContext) any {
	bodyBlock := v.loopBB.Top()[1]
	v.llvmBuilder.CreateBr(bodyBlock)
	return nil
}

func (v *Visitor) VisitStatBreak(ctx *parser.StatBreakContext) any {
	// peek but not pop
	exitBlock := v.loopBB.Top()[0]
	v.llvmBuilder.CreateBr(exitBlock)
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
		ret = v.llvmBuilder.CreateAdd(lhs, rhs, "tadd")
	case parser.RustLikeLexerMINUS:
		ret = v.llvmBuilder.CreateSub(lhs, rhs, "tsub")
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
		ret = v.llvmBuilder.CreateMul(lhs, rhs, "tMult")
	case parser.RustLikeLexerDIV:
		ret = v.llvmBuilder.CreateSDiv(lhs, rhs, "tDiv")
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
	// get function type
	funcSymbol := v.currentScope.Resolve(ctx.ID().GetText())
	fnType := GetType(funcSymbol)
	if fnType.IsNil() {
		log.Panicf("should not get nil function type for: %s", ctx.ID().GetText())
	}
	fn := v.llvmMod.NamedFunction(ctx.ID().GetText())
	if fnType.IsNil() {
		log.Panicf("should not get nil function value for: %s", ctx.ID().GetText())
	}
	log.Debugf("LLVM call function: %s", fnType.String())
	args := v.Visit(ctx.FuncCallList()).([]llvm.Value)
	return v.llvmBuilder.CreateCall(fnType, fn, args, "retCall"+funcSymbol.Name())
}

func (v *Visitor) VisitExprParen(ctx *parser.ExprParenContext) any {
	return v.Visit(ctx.Expr())
}

func (v *Visitor) VisitExprID(ctx *parser.ExprIDContext) any {
	varSymbol := v.currentScope.Resolve(ctx.ID().GetText())
	llvmType := v.LLVMType(varSymbol.Type())
	llvmVal := GetValue(varSymbol)
	ret := v.llvmBuilder.CreateLoad(llvmType, llvmVal, "tload")
	return ret
}
