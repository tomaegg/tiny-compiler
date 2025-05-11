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

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterFunc_call is called when entering the func_call production.
	EnterFunc_call(c *Func_callContext)

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

	// EnterStat is called when entering the stat production.
	EnterStat(c *StatContext)

	// EnterStat_return is called when entering the stat_return production.
	EnterStat_return(c *Stat_returnContext)

	// EnterVar_declare is called when entering the var_declare production.
	EnterVar_declare(c *Var_declareContext)

	// EnterVar_type is called when entering the var_type production.
	EnterVar_type(c *Var_typeContext)

	// EnterVar_init is called when entering the var_init production.
	EnterVar_init(c *Var_initContext)

	// EnterVar_assign is called when entering the var_assign production.
	EnterVar_assign(c *Var_assignContext)

	// EnterStat_if_else is called when entering the stat_if_else production.
	EnterStat_if_else(c *Stat_if_elseContext)

	// EnterStat_while is called when entering the stat_while production.
	EnterStat_while(c *Stat_whileContext)

	// EnterStat_loop is called when entering the stat_loop production.
	EnterStat_loop(c *Stat_loopContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitFunc_call is called when exiting the func_call production.
	ExitFunc_call(c *Func_callContext)

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

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitStat_return is called when exiting the stat_return production.
	ExitStat_return(c *Stat_returnContext)

	// ExitVar_declare is called when exiting the var_declare production.
	ExitVar_declare(c *Var_declareContext)

	// ExitVar_type is called when exiting the var_type production.
	ExitVar_type(c *Var_typeContext)

	// ExitVar_init is called when exiting the var_init production.
	ExitVar_init(c *Var_initContext)

	// ExitVar_assign is called when exiting the var_assign production.
	ExitVar_assign(c *Var_assignContext)

	// ExitStat_if_else is called when exiting the stat_if_else production.
	ExitStat_if_else(c *Stat_if_elseContext)

	// ExitStat_while is called when exiting the stat_while production.
	ExitStat_while(c *Stat_whileContext)

	// ExitStat_loop is called when exiting the stat_loop production.
	ExitStat_loop(c *Stat_loopContext)
}
