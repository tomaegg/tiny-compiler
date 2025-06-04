// Code generated from RustLikeParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // RustLikeParser

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by RustLikeParser.
type RustLikeParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by RustLikeParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by RustLikeParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by RustLikeParser#ExprAddSub.
	VisitExprAddSub(ctx *ExprAddSubContext) interface{}

	// Visit a parse tree produced by RustLikeParser#ExprParen.
	VisitExprParen(ctx *ExprParenContext) interface{}

	// Visit a parse tree produced by RustLikeParser#ExprNum.
	VisitExprNum(ctx *ExprNumContext) interface{}

	// Visit a parse tree produced by RustLikeParser#ExprMulDiv.
	VisitExprMulDiv(ctx *ExprMulDivContext) interface{}

	// Visit a parse tree produced by RustLikeParser#ExprCmp.
	VisitExprCmp(ctx *ExprCmpContext) interface{}

	// Visit a parse tree produced by RustLikeParser#ExprFuncCall.
	VisitExprFuncCall(ctx *ExprFuncCallContext) interface{}

	// Visit a parse tree produced by RustLikeParser#ExprID.
	VisitExprID(ctx *ExprIDContext) interface{}

	// Visit a parse tree produced by RustLikeParser#funcCallList.
	VisitFuncCallList(ctx *FuncCallListContext) interface{}

	// Visit a parse tree produced by RustLikeParser#funcCallParam.
	VisitFuncCallParam(ctx *FuncCallParamContext) interface{}

	// Visit a parse tree produced by RustLikeParser#funcDeclaration.
	VisitFuncDeclaration(ctx *FuncDeclarationContext) interface{}

	// Visit a parse tree produced by RustLikeParser#funcSignature.
	VisitFuncSignature(ctx *FuncSignatureContext) interface{}

	// Visit a parse tree produced by RustLikeParser#funcDeclarationReturn.
	VisitFuncDeclarationReturn(ctx *FuncDeclarationReturnContext) interface{}

	// Visit a parse tree produced by RustLikeParser#funcParamsList.
	VisitFuncParamsList(ctx *FuncParamsListContext) interface{}

	// Visit a parse tree produced by RustLikeParser#funcParams.
	VisitFuncParams(ctx *FuncParamsContext) interface{}

	// Visit a parse tree produced by RustLikeParser#funcParam.
	VisitFuncParam(ctx *FuncParamContext) interface{}

	// Visit a parse tree produced by RustLikeParser#funcBlock.
	VisitFuncBlock(ctx *FuncBlockContext) interface{}

	// Visit a parse tree produced by RustLikeParser#rtype.
	VisitRtype(ctx *RtypeContext) interface{}

	// Visit a parse tree produced by RustLikeParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by RustLikeParser#StatBlock.
	VisitStatBlock(ctx *StatBlockContext) interface{}

	// Visit a parse tree produced by RustLikeParser#StatFuncReturn.
	VisitStatFuncReturn(ctx *StatFuncReturnContext) interface{}

	// Visit a parse tree produced by RustLikeParser#StatVarDeclare.
	VisitStatVarDeclare(ctx *StatVarDeclareContext) interface{}

	// Visit a parse tree produced by RustLikeParser#StatVarAssign.
	VisitStatVarAssign(ctx *StatVarAssignContext) interface{}

	// Visit a parse tree produced by RustLikeParser#StatExpr.
	VisitStatExpr(ctx *StatExprContext) interface{}

	// Visit a parse tree produced by RustLikeParser#StatIfElse.
	VisitStatIfElse(ctx *StatIfElseContext) interface{}

	// Visit a parse tree produced by RustLikeParser#StatWhile.
	VisitStatWhile(ctx *StatWhileContext) interface{}

	// Visit a parse tree produced by RustLikeParser#StatLoop.
	VisitStatLoop(ctx *StatLoopContext) interface{}

	// Visit a parse tree produced by RustLikeParser#StatEmpty.
	VisitStatEmpty(ctx *StatEmptyContext) interface{}

	// Visit a parse tree produced by RustLikeParser#ifBranch.
	VisitIfBranch(ctx *IfBranchContext) interface{}

	// Visit a parse tree produced by RustLikeParser#elifBranch.
	VisitElifBranch(ctx *ElifBranchContext) interface{}

	// Visit a parse tree produced by RustLikeParser#elseBranch.
	VisitElseBranch(ctx *ElseBranchContext) interface{}

	// Visit a parse tree produced by RustLikeParser#varType.
	VisitVarType(ctx *VarTypeContext) interface{}

	// Visit a parse tree produced by RustLikeParser#varInit.
	VisitVarInit(ctx *VarInitContext) interface{}
}
