parser grammar RustLikeParserRules;
options {
	tokenVocab = RustLikeLexerRules;
}

prog: declaration;

declaration: func_declaration;

expr:
	expr (MULT | DIV) expr // 乘除
	| expr (PLUS | MINUS) expr // 加减
	| expr (EQ | NE | LT | GT | LE | GE) expr // 比较
	| LPAREN expr RPAREN // 加括号
	| func_call // 函数调用
	| ID
	| NUMBER;

func_call: ID func_call_list;

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
	block
	| stat_return SEMI
	| var_declare SEMI
	| var_assign SEMI
	| stat_return SEMI
	| expr SEMI
	| stat_if_else
	| stat_while
	| stat_loop
	| SEMI;

stat_return: RETURN expr?;

var_declare:
	LET MUT? ID var_type? var_init?; // 变量声明

var_type: COLON type;

var_init: ASSIGN expr;

var_assign: ID ASSIGN expr; // 变量赋值

stat_if_else: IF expr block (ELSE IF expr block)* (ELSE block)?;

stat_while: WHILE expr block;

stat_loop: LOOP block;
