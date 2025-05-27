// Code generated from RustLikeParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // RustLikeParser

import "github.com/antlr4-go/antlr/v4"

type BaseRustLikeParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseRustLikeParserVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRustLikeParserVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRustLikeParserVisitor) VisitExprAddSub(ctx *ExprAddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRustLikeParserVisitor) VisitExprParen(ctx *ExprParenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRustLikeParserVisitor) VisitExprNum(ctx *ExprNumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRustLikeParserVisitor) VisitExprMulDiv(ctx *ExprMulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRustLikeParserVisitor) VisitExprCmp(ctx *ExprCmpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRustLikeParserVisitor) VisitExprFuncCall(ctx *ExprFuncCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRustLikeParserVisitor) VisitExprID(ctx *ExprIDContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRustLikeParserVisitor) VisitFuncCallList(ctx *FuncCallListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRustLikeParserVisitor) VisitFuncCallParam(ctx *FuncCallParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRustLikeParserVisitor) VisitFuncDeclaration(ctx *FuncDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRustLikeParserVisitor) VisitFuncSignature(ctx *FuncSignatureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRustLikeParserVisitor) VisitFuncDeclarationReturn(ctx *FuncDeclarationReturnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRustLikeParserVisitor) VisitFuncParamsList(ctx *FuncParamsListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRustLikeParserVisitor) VisitFuncParams(ctx *FuncParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRustLikeParserVisitor) VisitFuncParam(ctx *FuncParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRustLikeParserVisitor) VisitFuncBlock(ctx *FuncBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRustLikeParserVisitor) VisitRtype(ctx *RtypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRustLikeParserVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRustLikeParserVisitor) VisitStatBlock(ctx *StatBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRustLikeParserVisitor) VisitStatFuncReturn(ctx *StatFuncReturnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRustLikeParserVisitor) VisitStatVarDeclare(ctx *StatVarDeclareContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRustLikeParserVisitor) VisitStatVarAssign(ctx *StatVarAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRustLikeParserVisitor) VisitStatExpr(ctx *StatExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRustLikeParserVisitor) VisitStatIfElse(ctx *StatIfElseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRustLikeParserVisitor) VisitStatWhile(ctx *StatWhileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRustLikeParserVisitor) VisitStatLoop(ctx *StatLoopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRustLikeParserVisitor) VisitStatEmpty(ctx *StatEmptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRustLikeParserVisitor) VisitVarType(ctx *VarTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRustLikeParserVisitor) VisitVarInit(ctx *VarInitContext) interface{} {
	return v.VisitChildren(ctx)
}
