// Code generated from RustLikeParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // RustLikeParser

import "github.com/antlr4-go/antlr/v4"

// RustLikeParserListener is a complete listener for a parse tree produced by RustLikeParser.
type RustLikeParserListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterExprAddSub is called when entering the ExprAddSub production.
	EnterExprAddSub(c *ExprAddSubContext)

	// EnterExprParen is called when entering the ExprParen production.
	EnterExprParen(c *ExprParenContext)

	// EnterExprNum is called when entering the ExprNum production.
	EnterExprNum(c *ExprNumContext)

	// EnterExprMulDiv is called when entering the ExprMulDiv production.
	EnterExprMulDiv(c *ExprMulDivContext)

	// EnterExprCmp is called when entering the ExprCmp production.
	EnterExprCmp(c *ExprCmpContext)

	// EnterExprFuncCall is called when entering the ExprFuncCall production.
	EnterExprFuncCall(c *ExprFuncCallContext)

	// EnterExprID is called when entering the ExprID production.
	EnterExprID(c *ExprIDContext)

	// EnterFuncCallList is called when entering the funcCallList production.
	EnterFuncCallList(c *FuncCallListContext)

	// EnterFuncCallParam is called when entering the funcCallParam production.
	EnterFuncCallParam(c *FuncCallParamContext)

	// EnterFuncDeclaration is called when entering the funcDeclaration production.
	EnterFuncDeclaration(c *FuncDeclarationContext)

	// EnterFuncDeclarationHeader is called when entering the funcDeclarationHeader production.
	EnterFuncDeclarationHeader(c *FuncDeclarationHeaderContext)

	// EnterFuncDeclarationReturn is called when entering the funcDeclarationReturn production.
	EnterFuncDeclarationReturn(c *FuncDeclarationReturnContext)

	// EnterFuncParamsList is called when entering the funcParamsList production.
	EnterFuncParamsList(c *FuncParamsListContext)

	// EnterFuncParams is called when entering the funcParams production.
	EnterFuncParams(c *FuncParamsContext)

	// EnterFuncParam is called when entering the funcParam production.
	EnterFuncParam(c *FuncParamContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStatBlock is called when entering the StatBlock production.
	EnterStatBlock(c *StatBlockContext)

	// EnterStatFuncReturn is called when entering the StatFuncReturn production.
	EnterStatFuncReturn(c *StatFuncReturnContext)

	// EnterStatVarDeclare is called when entering the StatVarDeclare production.
	EnterStatVarDeclare(c *StatVarDeclareContext)

	// EnterStatVarAssign is called when entering the StatVarAssign production.
	EnterStatVarAssign(c *StatVarAssignContext)

	// EnterStatExpr is called when entering the StatExpr production.
	EnterStatExpr(c *StatExprContext)

	// EnterStatIfElse is called when entering the StatIfElse production.
	EnterStatIfElse(c *StatIfElseContext)

	// EnterStatWhile is called when entering the StatWhile production.
	EnterStatWhile(c *StatWhileContext)

	// EnterStatLoop is called when entering the StatLoop production.
	EnterStatLoop(c *StatLoopContext)

	// EnterStatEmpty is called when entering the StatEmpty production.
	EnterStatEmpty(c *StatEmptyContext)

	// EnterVarType is called when entering the varType production.
	EnterVarType(c *VarTypeContext)

	// EnterVarInit is called when entering the varInit production.
	EnterVarInit(c *VarInitContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitExprAddSub is called when exiting the ExprAddSub production.
	ExitExprAddSub(c *ExprAddSubContext)

	// ExitExprParen is called when exiting the ExprParen production.
	ExitExprParen(c *ExprParenContext)

	// ExitExprNum is called when exiting the ExprNum production.
	ExitExprNum(c *ExprNumContext)

	// ExitExprMulDiv is called when exiting the ExprMulDiv production.
	ExitExprMulDiv(c *ExprMulDivContext)

	// ExitExprCmp is called when exiting the ExprCmp production.
	ExitExprCmp(c *ExprCmpContext)

	// ExitExprFuncCall is called when exiting the ExprFuncCall production.
	ExitExprFuncCall(c *ExprFuncCallContext)

	// ExitExprID is called when exiting the ExprID production.
	ExitExprID(c *ExprIDContext)

	// ExitFuncCallList is called when exiting the funcCallList production.
	ExitFuncCallList(c *FuncCallListContext)

	// ExitFuncCallParam is called when exiting the funcCallParam production.
	ExitFuncCallParam(c *FuncCallParamContext)

	// ExitFuncDeclaration is called when exiting the funcDeclaration production.
	ExitFuncDeclaration(c *FuncDeclarationContext)

	// ExitFuncDeclarationHeader is called when exiting the funcDeclarationHeader production.
	ExitFuncDeclarationHeader(c *FuncDeclarationHeaderContext)

	// ExitFuncDeclarationReturn is called when exiting the funcDeclarationReturn production.
	ExitFuncDeclarationReturn(c *FuncDeclarationReturnContext)

	// ExitFuncParamsList is called when exiting the funcParamsList production.
	ExitFuncParamsList(c *FuncParamsListContext)

	// ExitFuncParams is called when exiting the funcParams production.
	ExitFuncParams(c *FuncParamsContext)

	// ExitFuncParam is called when exiting the funcParam production.
	ExitFuncParam(c *FuncParamContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStatBlock is called when exiting the StatBlock production.
	ExitStatBlock(c *StatBlockContext)

	// ExitStatFuncReturn is called when exiting the StatFuncReturn production.
	ExitStatFuncReturn(c *StatFuncReturnContext)

	// ExitStatVarDeclare is called when exiting the StatVarDeclare production.
	ExitStatVarDeclare(c *StatVarDeclareContext)

	// ExitStatVarAssign is called when exiting the StatVarAssign production.
	ExitStatVarAssign(c *StatVarAssignContext)

	// ExitStatExpr is called when exiting the StatExpr production.
	ExitStatExpr(c *StatExprContext)

	// ExitStatIfElse is called when exiting the StatIfElse production.
	ExitStatIfElse(c *StatIfElseContext)

	// ExitStatWhile is called when exiting the StatWhile production.
	ExitStatWhile(c *StatWhileContext)

	// ExitStatLoop is called when exiting the StatLoop production.
	ExitStatLoop(c *StatLoopContext)

	// ExitStatEmpty is called when exiting the StatEmpty production.
	ExitStatEmpty(c *StatEmptyContext)

	// ExitVarType is called when exiting the varType production.
	ExitVarType(c *VarTypeContext)

	// ExitVarInit is called when exiting the varInit production.
	ExitVarInit(c *VarInitContext)
}
