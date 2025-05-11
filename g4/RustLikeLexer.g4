lexer grammar RustLikeLexer;

// comments
SL_COMMENT: '//' (~'\n')* '\n' -> skip;
ML_COMMENT: '/*' .*? '*/' -> skip;

// whitespace
WS: [ \t\r\n]+ -> skip;

// keyword
INT32: 'i32';
LET: 'let';
IF: 'if';
ELSE: 'else';
WHILE: 'while';
RETURN: 'return';
MUT: 'mut';
FN: 'fn';
LOOP: 'loop';
BREAK: 'break';
CONTINUE: 'continue';

// id
ID: (LETTER | '_') (LETTER | '_' | DIGIT)*;

// number
NUMBER: DIGITS;

// algorithm
PLUS: '+';
MINUS: '-';
MULT: '*';
DIV: '/';

// relation
EQ: '==';
NE: '!=';
LT: '<';
LE: '<=';
GT: '>';
GE: '>=';

// assignment
ASSIGN: '=';

// delimiter
SEMI: ';';
COLON: ':';
COMMA: ',';

// 
LPAREN: '(';
RPAREN: ')';
LBRAC: '[';
RBRAC: ']';
LBRACE: '{';
RBRACE: '}';

// some more specal
ARROW: '->';
DOT: '.';
DOT2: '..';

fragment DIGIT: [0-9];
fragment DIGITS: [0-9]+;
fragment LETTER: [a-zA-Z];
