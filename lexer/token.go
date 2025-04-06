package lexer

import (
	"sync"
)

type Token interface {
	TokenDesc
	Position() (line, col int)
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

type token struct {
	TokenDesc
	line, col int
}

var _ Token = (*token)(nil)

func (tk *token) Position() (int, int) {
	return tk.line, tk.col
}
