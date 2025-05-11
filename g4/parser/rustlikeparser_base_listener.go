// Code generated from RustLikeParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // RustLikeParser

import "github.com/antlr4-go/antlr/v4"

// BaseRustLikeParserListener is a complete listener for a parse tree produced by RustLikeParser.
type BaseRustLikeParserListener struct{}

var _ RustLikeParserListener = &BaseRustLikeParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRustLikeParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRustLikeParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRustLikeParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRustLikeParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseRustLikeParserListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseRustLikeParserListener) ExitProg(ctx *ProgContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseRustLikeParserListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseRustLikeParserListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseRustLikeParserListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseRustLikeParserListener) ExitExpr(ctx *ExprContext) {}

// EnterFunc_call is called when production func_call is entered.
func (s *BaseRustLikeParserListener) EnterFunc_call(ctx *Func_callContext) {}

// ExitFunc_call is called when production func_call is exited.
func (s *BaseRustLikeParserListener) ExitFunc_call(ctx *Func_callContext) {}

// EnterFunc_call_list is called when production func_call_list is entered.
func (s *BaseRustLikeParserListener) EnterFunc_call_list(ctx *Func_call_listContext) {}

// ExitFunc_call_list is called when production func_call_list is exited.
func (s *BaseRustLikeParserListener) ExitFunc_call_list(ctx *Func_call_listContext) {}

// EnterFunc_call_param is called when production func_call_param is entered.
func (s *BaseRustLikeParserListener) EnterFunc_call_param(ctx *Func_call_paramContext) {}

// ExitFunc_call_param is called when production func_call_param is exited.
func (s *BaseRustLikeParserListener) ExitFunc_call_param(ctx *Func_call_paramContext) {}

// EnterFunc_declaration is called when production func_declaration is entered.
func (s *BaseRustLikeParserListener) EnterFunc_declaration(ctx *Func_declarationContext) {}

// ExitFunc_declaration is called when production func_declaration is exited.
func (s *BaseRustLikeParserListener) ExitFunc_declaration(ctx *Func_declarationContext) {}

// EnterFunc_declaration_header is called when production func_declaration_header is entered.
func (s *BaseRustLikeParserListener) EnterFunc_declaration_header(ctx *Func_declaration_headerContext) {
}

// ExitFunc_declaration_header is called when production func_declaration_header is exited.
func (s *BaseRustLikeParserListener) ExitFunc_declaration_header(ctx *Func_declaration_headerContext) {
}

// EnterFunc_declaration_return is called when production func_declaration_return is entered.
func (s *BaseRustLikeParserListener) EnterFunc_declaration_return(ctx *Func_declaration_returnContext) {
}

// ExitFunc_declaration_return is called when production func_declaration_return is exited.
func (s *BaseRustLikeParserListener) ExitFunc_declaration_return(ctx *Func_declaration_returnContext) {
}

// EnterFps_list is called when production fps_list is entered.
func (s *BaseRustLikeParserListener) EnterFps_list(ctx *Fps_listContext) {}

// ExitFps_list is called when production fps_list is exited.
func (s *BaseRustLikeParserListener) ExitFps_list(ctx *Fps_listContext) {}

// EnterFps is called when production fps is entered.
func (s *BaseRustLikeParserListener) EnterFps(ctx *FpsContext) {}

// ExitFps is called when production fps is exited.
func (s *BaseRustLikeParserListener) ExitFps(ctx *FpsContext) {}

// EnterFp is called when production fp is entered.
func (s *BaseRustLikeParserListener) EnterFp(ctx *FpContext) {}

// ExitFp is called when production fp is exited.
func (s *BaseRustLikeParserListener) ExitFp(ctx *FpContext) {}

// EnterType is called when production type is entered.
func (s *BaseRustLikeParserListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseRustLikeParserListener) ExitType(ctx *TypeContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseRustLikeParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseRustLikeParserListener) ExitBlock(ctx *BlockContext) {}

// EnterStat is called when production stat is entered.
func (s *BaseRustLikeParserListener) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (s *BaseRustLikeParserListener) ExitStat(ctx *StatContext) {}

// EnterStat_return is called when production stat_return is entered.
func (s *BaseRustLikeParserListener) EnterStat_return(ctx *Stat_returnContext) {}

// ExitStat_return is called when production stat_return is exited.
func (s *BaseRustLikeParserListener) ExitStat_return(ctx *Stat_returnContext) {}

// EnterVar_declare is called when production var_declare is entered.
func (s *BaseRustLikeParserListener) EnterVar_declare(ctx *Var_declareContext) {}

// ExitVar_declare is called when production var_declare is exited.
func (s *BaseRustLikeParserListener) ExitVar_declare(ctx *Var_declareContext) {}

// EnterVar_type is called when production var_type is entered.
func (s *BaseRustLikeParserListener) EnterVar_type(ctx *Var_typeContext) {}

// ExitVar_type is called when production var_type is exited.
func (s *BaseRustLikeParserListener) ExitVar_type(ctx *Var_typeContext) {}

// EnterVar_init is called when production var_init is entered.
func (s *BaseRustLikeParserListener) EnterVar_init(ctx *Var_initContext) {}

// ExitVar_init is called when production var_init is exited.
func (s *BaseRustLikeParserListener) ExitVar_init(ctx *Var_initContext) {}

// EnterVar_assign is called when production var_assign is entered.
func (s *BaseRustLikeParserListener) EnterVar_assign(ctx *Var_assignContext) {}

// ExitVar_assign is called when production var_assign is exited.
func (s *BaseRustLikeParserListener) ExitVar_assign(ctx *Var_assignContext) {}

// EnterStat_if_else is called when production stat_if_else is entered.
func (s *BaseRustLikeParserListener) EnterStat_if_else(ctx *Stat_if_elseContext) {}

// ExitStat_if_else is called when production stat_if_else is exited.
func (s *BaseRustLikeParserListener) ExitStat_if_else(ctx *Stat_if_elseContext) {}

// EnterStat_while is called when production stat_while is entered.
func (s *BaseRustLikeParserListener) EnterStat_while(ctx *Stat_whileContext) {}

// ExitStat_while is called when production stat_while is exited.
func (s *BaseRustLikeParserListener) ExitStat_while(ctx *Stat_whileContext) {}

// EnterStat_loop is called when production stat_loop is entered.
func (s *BaseRustLikeParserListener) EnterStat_loop(ctx *Stat_loopContext) {}

// ExitStat_loop is called when production stat_loop is exited.
func (s *BaseRustLikeParserListener) ExitStat_loop(ctx *Stat_loopContext) {}
