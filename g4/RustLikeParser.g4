parser grammar RustLikeParser;
options {
	tokenVocab = RustLikeLexer;
}

prog: declaration;

declaration: funcDeclaration*;

expr:
	expr (MULT | DIV) expr						# ExprMulDiv
	| expr (PLUS | MINUS) expr					# ExprAddSub
	| expr (EQ | NE | LT | GT | LE | GE) expr	# ExprCmp
	| LPAREN expr RPAREN						# ExprParen
	| ID funcCallList							# ExprFuncCall
	| ID										# ExprID
	| NUMBER									# ExprNum;

funcCallList: LPAREN funcCallParam RPAREN;

funcCallParam: expr (COMMA expr)* |;

funcDeclaration:
	funcDeclarationHeader funcDeclarationReturn? funcBlock; // 函数声明

funcDeclarationHeader: FN ID funcParamsList;

funcDeclarationReturn: ARROW rtype;

funcParamsList: LPAREN funcParams RPAREN;

funcParams: funcParam (COMMA funcParam)* |; // fps可以为空

funcParam: MUT? ID COLON rtype;

funcBlock: LBRACE stat* RBRACE;

rtype: INT32;

block: LBRACE stat* RBRACE;

stat:
	block												# StatBlock
	| RETURN expr? SEMI									# StatFuncReturn
	| LET MUT? ID varType? varInit? SEMI				# StatVarDeclare
	| ID ASSIGN expr SEMI								# StatVarAssign
	| expr SEMI											# StatExpr
	| IF expr block (ELSE IF expr block)* (ELSE block)?	# StatIfElse
	| WHILE expr block									# StatWhile
	| LOOP block										# StatLoop
	| SEMI												# StatEmpty;

varType: COLON rtype;

varInit: ASSIGN expr;
