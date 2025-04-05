package lexer

import "sync"

type Token interface {
	Literal() string // token的字面量, 比如123, 456, <, >
	Type() TokenType
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

var constTokens []Token

var onceToken sync.Once

func ConstTokens() []Token {
	onceToken.Do(initTokens)
	return constTokens
}

func initTokens() {
	constTokens = make([]Token, tokenCount)
	// 关键字
	constTokens[TokenINT32] = newToken(TokenINT32, "i32")
	constTokens[TokenLET] = newToken(TokenLET, "let")
	constTokens[TokenIF] = newToken(TokenIF, "if")
	constTokens[TokenELSE] = newToken(TokenELSE, "else")
	constTokens[TokenRETURN] = newToken(TokenRETURN, "return")
	constTokens[TokenMUT] = newToken(TokenMUT, "mut")
	constTokens[TokenFN] = newToken(TokenFN, "fn")
	constTokens[TokenFOR] = newToken(TokenFOR, "for")
	constTokens[TokenIN] = newToken(TokenIN, "in")
	constTokens[TokenLOOP] = newToken(TokenLOOP, "loop")
	constTokens[TokenBREAK] = newToken(TokenBREAK, "break")
	constTokens[TokenCONTINUE] = newToken(TokenCONTINUE, "continue")
	// 标识符（无固定字符串）
	constTokens[TokenID] = newToken(TokenID, "") // 或占位符如 "ID"
}

type token struct {
	tkType TokenType
	text   string // token字面量
}

var _ Token = (*token)(nil)

func (tk *token) Literal() string {
	return tk.text
}

func (tk *token) Type() TokenType {
	return tk.tkType
}

func NewToken(tkType TokenType, text string) Token {
	return newToken(tkType, text)
}

func newToken(tkType TokenType, text string) *token {
	return &token{
		tkType: tkType,
		text:   text,
	}
}
