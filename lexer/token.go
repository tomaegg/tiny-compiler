package lexer

import (
	"fmt"
	"sync"
)

type TokenDesc interface {
	Literal() string // token的字面量, 比如123, 456, <, >
	Type() TokenType
	SymbolicName() string
	String() string
}

type Token interface {
	TokenDesc
	Position() (line, col int)
}

type TokenType = int

const (
	// keyword
	TokenINT32 = iota
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

	// for count
	tokenCount
)

func TokenDescs() []TokenDesc {
	onceToken.Do(initTokens)
	return tokens
}

func NewToken(tokenType TokenType, literal string, line, col int) Token {
	return &token{
		TokenDesc: tokens[tokenType],
		line:      line, col: col,
	}
}

var (
	tokens             []TokenDesc
	tokenSymbolicNames []string
)

var onceToken sync.Once

func initTokens() {
	tokens = make([]TokenDesc, tokenCount)
	tokenSymbolicNames = make([]string, tokenCount)
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
}

type token struct {
	TokenDesc
	line, col int
}

var _ Token = (*token)(nil)

func (tk *token) Position() (int, int) {
	return tk.line, tk.col
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
	return fmt.Sprintf("token(%s):%s", tk.SymbolicName(), tk.Literal())
}

func newTokenDesc(tkType TokenType, text, symbol string) {
	tokens[tkType] = &tokenDesc{
		tkType: tkType,
		text:   text,
	}
	tokenSymbolicNames[tkType] = symbol
}
