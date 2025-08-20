parser grammar RustLikeParser;
options {
	tokenVocab = RustLikeLexer;
}

prog: declaration;

declaration: funcDeclaration*;

expr:
	lhs = expr op = (MULT | DIV) rhs = expr						# ExprMulDiv
	| lhs = expr op = (PLUS | MINUS) rhs = expr					# ExprAddSub
	| lhs = expr op = (EQ | NE | LT | GT | LE | GE) rhs = expr	# ExprCmp
	| LBRAC arrayElems RBRAC									# ExprArray
	| LPAREN expr RPAREN										# ExprParen
	| lhs = expr LBRAC rhs = expr RBRAC					# ExprArrayAccess
	| ID funcCallList											# ExprFuncCall
	| ID														# ExprID
	| exprNumber													# ExprNum;

exprNumber: NUMBER;

funcCallList: LPAREN funcCallParam RPAREN;

funcCallParam: expr (COMMA expr)* |;

funcDeclaration: funcSignature funcBlock; // 函数声明

funcSignature: FN ID funcParamsList funcDeclarationReturn?;

funcDeclarationReturn: ARROW rtype;

funcParamsList: LPAREN funcParams RPAREN;

funcParams: funcParam (COMMA funcParam)* |; // fps可以为空

funcParam: MUT? ID COLON rtype;

funcBlock: LBRACE stat* RBRACE;

rtype: INT32 | arrayType;

arrayType: LBRAC rtype SEMI exprNumber RBRAC;

arrayElems: expr (COMMA expr)* |; // can be empty

block: LBRACE stat* RBRACE;

stat:
	block									# StatBlock
	| RETURN expr? SEMI						# StatFuncReturn
	| BREAK SEMI							# StatBreak
	| CONTINUE SEMI							# StatContinue
	| LET MUT? ID varType? varInit? SEMI	# StatVarDeclare
	| ID ASSIGN expr SEMI					# StatVarAssign
	| expr SEMI								# StatExpr
	| ifBranch elifBranch* elseBranch?		# StatIfElse
	| WHILE expr block						# StatWhile
	| LOOP block							# StatLoop
	| SEMI									# StatEmpty;

ifBranch: IF expr block;

elifBranch: ELSE IF expr block;

elseBranch: ELSE block;

varType: COLON rtype;

varInit: ASSIGN expr;
