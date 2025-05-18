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

	// EnterFunc_call_list is called when entering the func_call_list production.
	EnterFunc_call_list(c *Func_call_listContext)

	// EnterFunc_call_param is called when entering the func_call_param production.
	EnterFunc_call_param(c *Func_call_paramContext)

	// EnterFunc_declaration is called when entering the func_declaration production.
	EnterFunc_declaration(c *Func_declarationContext)

	// EnterFunc_declaration_header is called when entering the func_declaration_header production.
	EnterFunc_declaration_header(c *Func_declaration_headerContext)

	// EnterFunc_declaration_return is called when entering the func_declaration_return production.
	EnterFunc_declaration_return(c *Func_declaration_returnContext)

	// EnterFps_list is called when entering the fps_list production.
	EnterFps_list(c *Fps_listContext)

	// EnterFps is called when entering the fps production.
	EnterFps(c *FpsContext)

	// EnterFp is called when entering the fp production.
	EnterFp(c *FpContext)

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

	// EnterVar_type is called when entering the var_type production.
	EnterVar_type(c *Var_typeContext)

	// EnterVar_init is called when entering the var_init production.
	EnterVar_init(c *Var_initContext)

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

	// ExitFunc_call_list is called when exiting the func_call_list production.
	ExitFunc_call_list(c *Func_call_listContext)

	// ExitFunc_call_param is called when exiting the func_call_param production.
	ExitFunc_call_param(c *Func_call_paramContext)

	// ExitFunc_declaration is called when exiting the func_declaration production.
	ExitFunc_declaration(c *Func_declarationContext)

	// ExitFunc_declaration_header is called when exiting the func_declaration_header production.
	ExitFunc_declaration_header(c *Func_declaration_headerContext)

	// ExitFunc_declaration_return is called when exiting the func_declaration_return production.
	ExitFunc_declaration_return(c *Func_declaration_returnContext)

	// ExitFps_list is called when exiting the fps_list production.
	ExitFps_list(c *Fps_listContext)

	// ExitFps is called when exiting the fps production.
	ExitFps(c *FpsContext)

	// ExitFp is called when exiting the fp production.
	ExitFp(c *FpContext)

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

	// ExitVar_type is called when exiting the var_type production.
	ExitVar_type(c *Var_typeContext)

	// ExitVar_init is called when exiting the var_init production.
	ExitVar_init(c *Var_initContext)
}
