parser grammar RustLikeParser;
options {
	tokenVocab = RustLikeLexer;
}

prog: declaration;

declaration: func_declaration*;

expr:
	expr (MULT | DIV) expr						# ExprMulDiv
	| expr (PLUS | MINUS) expr					# ExprAddSub
	| expr (EQ | NE | LT | GT | LE | GE) expr	# ExprCmp
	| LPAREN expr RPAREN						# ExprParen
	| ID func_call_list							# ExprFuncCall
	| ID										# ExprID
	| NUMBER									# ExprNum;

func_call_list: LPAREN func_call_param RPAREN;

func_call_param: expr (COMMA expr)* |;

func_declaration:
	func_declaration_header func_declaration_return? block; // 函数声明

func_declaration_header: FN ID fps_list;

func_declaration_return: ARROW type;

fps_list: LPAREN fps RPAREN;

fps: fp (COMMA fp)* |; // fps可以为空

fp: MUT? ID COLON type;

type: INT32;

block: LBRACE stat* RBRACE;

stat:
	block												# StatBlock
	| RETURN expr? SEMI									# StatFuncReturn
	| LET MUT? ID var_type? var_init? SEMI				# StatVarDeclare
	| ID ASSIGN expr SEMI								# StatVarAssign
	| expr SEMI											# StatExpr
	| IF expr block (ELSE IF expr block)* (ELSE block)?	# StatIfElse
	| WHILE expr block									# StatWhile
	| LOOP block										# StatLoop
	| SEMI												# StatEmpty;

var_type: COLON type;

var_init: ASSIGN expr;