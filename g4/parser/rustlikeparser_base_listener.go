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

// EnterExprAddSub is called when production ExprAddSub is entered.
func (s *BaseRustLikeParserListener) EnterExprAddSub(ctx *ExprAddSubContext) {}

// ExitExprAddSub is called when production ExprAddSub is exited.
func (s *BaseRustLikeParserListener) ExitExprAddSub(ctx *ExprAddSubContext) {}

// EnterExprParen is called when production ExprParen is entered.
func (s *BaseRustLikeParserListener) EnterExprParen(ctx *ExprParenContext) {}

// ExitExprParen is called when production ExprParen is exited.
func (s *BaseRustLikeParserListener) ExitExprParen(ctx *ExprParenContext) {}

// EnterExprNum is called when production ExprNum is entered.
func (s *BaseRustLikeParserListener) EnterExprNum(ctx *ExprNumContext) {}

// ExitExprNum is called when production ExprNum is exited.
func (s *BaseRustLikeParserListener) ExitExprNum(ctx *ExprNumContext) {}

// EnterExprMulDiv is called when production ExprMulDiv is entered.
func (s *BaseRustLikeParserListener) EnterExprMulDiv(ctx *ExprMulDivContext) {}

// ExitExprMulDiv is called when production ExprMulDiv is exited.
func (s *BaseRustLikeParserListener) ExitExprMulDiv(ctx *ExprMulDivContext) {}

// EnterExprCmp is called when production ExprCmp is entered.
func (s *BaseRustLikeParserListener) EnterExprCmp(ctx *ExprCmpContext) {}

// ExitExprCmp is called when production ExprCmp is exited.
func (s *BaseRustLikeParserListener) ExitExprCmp(ctx *ExprCmpContext) {}

// EnterExprFuncCall is called when production ExprFuncCall is entered.
func (s *BaseRustLikeParserListener) EnterExprFuncCall(ctx *ExprFuncCallContext) {}

// ExitExprFuncCall is called when production ExprFuncCall is exited.
func (s *BaseRustLikeParserListener) ExitExprFuncCall(ctx *ExprFuncCallContext) {}

// EnterExprID is called when production ExprID is entered.
func (s *BaseRustLikeParserListener) EnterExprID(ctx *ExprIDContext) {}

// ExitExprID is called when production ExprID is exited.
func (s *BaseRustLikeParserListener) ExitExprID(ctx *ExprIDContext) {}

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

// EnterStatBlock is called when production StatBlock is entered.
func (s *BaseRustLikeParserListener) EnterStatBlock(ctx *StatBlockContext) {}

// ExitStatBlock is called when production StatBlock is exited.
func (s *BaseRustLikeParserListener) ExitStatBlock(ctx *StatBlockContext) {}

// EnterStatFuncReturn is called when production StatFuncReturn is entered.
func (s *BaseRustLikeParserListener) EnterStatFuncReturn(ctx *StatFuncReturnContext) {}

// ExitStatFuncReturn is called when production StatFuncReturn is exited.
func (s *BaseRustLikeParserListener) ExitStatFuncReturn(ctx *StatFuncReturnContext) {}

// EnterStatVarDeclare is called when production StatVarDeclare is entered.
func (s *BaseRustLikeParserListener) EnterStatVarDeclare(ctx *StatVarDeclareContext) {}

// ExitStatVarDeclare is called when production StatVarDeclare is exited.
func (s *BaseRustLikeParserListener) ExitStatVarDeclare(ctx *StatVarDeclareContext) {}

// EnterStatVarAssign is called when production StatVarAssign is entered.
func (s *BaseRustLikeParserListener) EnterStatVarAssign(ctx *StatVarAssignContext) {}

// ExitStatVarAssign is called when production StatVarAssign is exited.
func (s *BaseRustLikeParserListener) ExitStatVarAssign(ctx *StatVarAssignContext) {}

// EnterStatExpr is called when production StatExpr is entered.
func (s *BaseRustLikeParserListener) EnterStatExpr(ctx *StatExprContext) {}

// ExitStatExpr is called when production StatExpr is exited.
func (s *BaseRustLikeParserListener) ExitStatExpr(ctx *StatExprContext) {}

// EnterStatIfElse is called when production StatIfElse is entered.
func (s *BaseRustLikeParserListener) EnterStatIfElse(ctx *StatIfElseContext) {}

// ExitStatIfElse is called when production StatIfElse is exited.
func (s *BaseRustLikeParserListener) ExitStatIfElse(ctx *StatIfElseContext) {}

// EnterStatWhile is called when production StatWhile is entered.
func (s *BaseRustLikeParserListener) EnterStatWhile(ctx *StatWhileContext) {}

// ExitStatWhile is called when production StatWhile is exited.
func (s *BaseRustLikeParserListener) ExitStatWhile(ctx *StatWhileContext) {}

// EnterStatLoop is called when production StatLoop is entered.
func (s *BaseRustLikeParserListener) EnterStatLoop(ctx *StatLoopContext) {}

// ExitStatLoop is called when production StatLoop is exited.
func (s *BaseRustLikeParserListener) ExitStatLoop(ctx *StatLoopContext) {}

// EnterStatEmpty is called when production StatEmpty is entered.
func (s *BaseRustLikeParserListener) EnterStatEmpty(ctx *StatEmptyContext) {}

// ExitStatEmpty is called when production StatEmpty is exited.
func (s *BaseRustLikeParserListener) ExitStatEmpty(ctx *StatEmptyContext) {}

// EnterVar_type is called when production var_type is entered.
func (s *BaseRustLikeParserListener) EnterVar_type(ctx *Var_typeContext) {}

// ExitVar_type is called when production var_type is exited.
func (s *BaseRustLikeParserListener) ExitVar_type(ctx *Var_typeContext) {}

// EnterVar_init is called when production var_init is entered.
func (s *BaseRustLikeParserListener) EnterVar_init(ctx *Var_initContext) {}

// ExitVar_init is called when production var_init is exited.
func (s *BaseRustLikeParserListener) ExitVar_init(ctx *Var_initContext) {}
