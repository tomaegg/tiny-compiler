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

// EnterFuncCallList is called when production funcCallList is entered.
func (s *BaseRustLikeParserListener) EnterFuncCallList(ctx *FuncCallListContext) {}

// ExitFuncCallList is called when production funcCallList is exited.
func (s *BaseRustLikeParserListener) ExitFuncCallList(ctx *FuncCallListContext) {}

// EnterFuncCallParam is called when production funcCallParam is entered.
func (s *BaseRustLikeParserListener) EnterFuncCallParam(ctx *FuncCallParamContext) {}

// ExitFuncCallParam is called when production funcCallParam is exited.
func (s *BaseRustLikeParserListener) ExitFuncCallParam(ctx *FuncCallParamContext) {}

// EnterFuncDeclaration is called when production funcDeclaration is entered.
func (s *BaseRustLikeParserListener) EnterFuncDeclaration(ctx *FuncDeclarationContext) {}

// ExitFuncDeclaration is called when production funcDeclaration is exited.
func (s *BaseRustLikeParserListener) ExitFuncDeclaration(ctx *FuncDeclarationContext) {}

// EnterFuncDeclarationHeader is called when production funcDeclarationHeader is entered.
func (s *BaseRustLikeParserListener) EnterFuncDeclarationHeader(ctx *FuncDeclarationHeaderContext) {}

// ExitFuncDeclarationHeader is called when production funcDeclarationHeader is exited.
func (s *BaseRustLikeParserListener) ExitFuncDeclarationHeader(ctx *FuncDeclarationHeaderContext) {}

// EnterFuncDeclarationReturn is called when production funcDeclarationReturn is entered.
func (s *BaseRustLikeParserListener) EnterFuncDeclarationReturn(ctx *FuncDeclarationReturnContext) {}

// ExitFuncDeclarationReturn is called when production funcDeclarationReturn is exited.
func (s *BaseRustLikeParserListener) ExitFuncDeclarationReturn(ctx *FuncDeclarationReturnContext) {}

// EnterFuncParamsList is called when production funcParamsList is entered.
func (s *BaseRustLikeParserListener) EnterFuncParamsList(ctx *FuncParamsListContext) {}

// ExitFuncParamsList is called when production funcParamsList is exited.
func (s *BaseRustLikeParserListener) ExitFuncParamsList(ctx *FuncParamsListContext) {}

// EnterFuncParams is called when production funcParams is entered.
func (s *BaseRustLikeParserListener) EnterFuncParams(ctx *FuncParamsContext) {}

// ExitFuncParams is called when production funcParams is exited.
func (s *BaseRustLikeParserListener) ExitFuncParams(ctx *FuncParamsContext) {}

// EnterFuncParam is called when production funcParam is entered.
func (s *BaseRustLikeParserListener) EnterFuncParam(ctx *FuncParamContext) {}

// ExitFuncParam is called when production funcParam is exited.
func (s *BaseRustLikeParserListener) ExitFuncParam(ctx *FuncParamContext) {}

// EnterRtype is called when production rtype is entered.
func (s *BaseRustLikeParserListener) EnterRtype(ctx *RtypeContext) {}

// ExitRtype is called when production rtype is exited.
func (s *BaseRustLikeParserListener) ExitRtype(ctx *RtypeContext) {}

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

// EnterVarType is called when production varType is entered.
func (s *BaseRustLikeParserListener) EnterVarType(ctx *VarTypeContext) {}

// ExitVarType is called when production varType is exited.
func (s *BaseRustLikeParserListener) ExitVarType(ctx *VarTypeContext) {}

// EnterVarInit is called when production varInit is entered.
func (s *BaseRustLikeParserListener) EnterVarInit(ctx *VarInitContext) {}

// ExitVarInit is called when production varInit is exited.
func (s *BaseRustLikeParserListener) ExitVarInit(ctx *VarInitContext) {}
