package lexer

import (
	"fmt"
	"strconv"
)

type TokenDesc interface {
	Literal() string // token的字面量, 比如123, 456, <, >
	Type() TokenType
	SymbolicName() string
	String() string
}

type TokenType = int

const (
	TokenEOF = iota
	TokenUNKNONW
	// keyword
	TokenINT32
	TokenLET
	TokenIF
	TokenELSE
	TokenWhile
	TokenRETURN
	TokenMUT
	TokenFN
	TokenFOR
	TokenIN
	TokenLOOP
	TokenBREAK
	TokenCONTINUE

	// 标识符, 与关键字不相同
	TokenID

	// 数值
	TokenNUMBER

	// 四则运算
	TokenPLUS
	TokenMINUS
	TokenMULTI
	TokenDIV

	// 关系运算
	TokenEQ
	TokenNE
	TokenLT
	TokenLE
	TokenGT
	TokenGE

	// 赋值
	TokenASSIGN

	// 分隔符
	TokenSEMI
	TokenCOLON
	TokenCOMMA

	// 分界符
	TokenLPAREN // (
	TokenRPAREN // )
	TokenLBRAC  // [
	TokenRBRAC  // ]
	TokenLBRACE // {
	TokenRBRACE // }

	// for count
	tokenCount
)

func TokenDescs() []TokenDesc {
	onceToken.Do(initTokens)
	return tokens
}

func initTokens() {
	tokens = make([]TokenDesc, tokenCount)
	tokenSymbolicNames = make([]string, tokenCount)
	newTokenDesc(TokenEOF, "", "EOF")
	newTokenDesc(TokenUNKNONW, "", "UNKNOWN")
	// 关键字
	newTokenDesc(TokenINT32, "i32", "INT32")
	newTokenDesc(TokenLET, "let", "LET")
	newTokenDesc(TokenIF, "if", "IF")
	newTokenDesc(TokenELSE, "else", "ELSE")
	newTokenDesc(TokenRETURN, "return", "RETURN")
	newTokenDesc(TokenWhile, "while", "WHILE")
	newTokenDesc(TokenMUT, "mut", "MUT")
	newTokenDesc(TokenFN, "fn", "FN")
	newTokenDesc(TokenFOR, "for", "FOR")
	newTokenDesc(TokenIN, "in", "INT")
	newTokenDesc(TokenLOOP, "loop", "LOOP")
	newTokenDesc(TokenBREAK, "break", "BREAK")
	newTokenDesc(TokenCONTINUE, "continue", "CONTINUE")
	// 标识符（无固定字符串）
	newTokenDesc(TokenID, "", "ID") // 或占位符如 "ID"
	// number
	newTokenDesc(TokenNUMBER, "", "NUMBER") // 或占位符如 "ID"
	// 四则
	newTokenDesc(TokenPLUS, "+", "PLUS")
	newTokenDesc(TokenMINUS, "-", "MINUS")
	newTokenDesc(TokenMULTI, "*", "MULT")
	newTokenDesc(TokenDIV, "/", "DIV")
	// 关系
	newTokenDesc(TokenEQ, "==", "EQ") // 等于
	newTokenDesc(TokenNE, "!=", "NE") // 不等于
	newTokenDesc(TokenLT, "<", "LT")  // 小于
	newTokenDesc(TokenLE, "<=", "LE") // 小于等于
	newTokenDesc(TokenGT, ">", "GT")  // 大于
	newTokenDesc(TokenGE, ">=", "GE") // 大于等于
	// 赋值
	newTokenDesc(TokenASSIGN, "=", "ASSIGN") // 赋值符号
	// 分隔符
	newTokenDesc(TokenSEMI, ";", "SEMICOLON")
	newTokenDesc(TokenCOLON, ":", "COLON")
	newTokenDesc(TokenCOMMA, ",", "COMMA")
	// 分界符（补充符号定义）
	newTokenDesc(TokenLPAREN, "(", "LPAREN")
	newTokenDesc(TokenRPAREN, ")", "RPAREN")
	newTokenDesc(TokenLBRAC, "[", "LBRAC")
	newTokenDesc(TokenRBRAC, "]", "RBRAC")
	newTokenDesc(TokenLBRACE, "{", "LBRACE")
	newTokenDesc(TokenRBRACE, "}", "RBRACE")
}

type tokenDesc struct {
	tkType TokenType
	text   string // token字面量
}

var _ TokenDesc = (*tokenDesc)(nil)

func (tk *tokenDesc) Literal() string {
	return tk.text
}

func (tk *tokenDesc) Type() TokenType {
	return tk.tkType
}

func (tk *tokenDesc) SymbolicName() string {
	return tokenSymbolicNames[tk.tkType]
}

func (tk *tokenDesc) String() string {
	return fmt.Sprintf("token(%s): %s",
		tk.SymbolicName(),
		strconv.Quote(tk.Literal()),
	)
}

func newTokenDesc(tkType TokenType, text, symbol string) {
	tokens[tkType] = &tokenDesc{
		tkType: tkType,
		text:   text,
	}
	tokenSymbolicNames[tkType] = symbol
}
