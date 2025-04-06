package lexer

import (
	"fmt"
	"strconv"
	"sync"
)

type Token interface {
	TokenDesc
	Text() string
	Position() (line, col int)
}

func NewTokenWithText(tokenType TokenType, text string, line, col int) Token {
	ret := &token{
		TokenDesc: tokens[tokenType],
		text:      text,
		line:      line, col: col,
	}
	return ret
}

func NewToken(tokenType TokenType, line, col int) Token {
	ret := &token{
		TokenDesc: tokens[tokenType],
		line:      line, col: col,
	}
	ret.text = ret.Literal()
	return ret
}

var (
	tokens             []TokenDesc
	tokenSymbolicNames []string
)

var onceToken sync.Once

type token struct {
	TokenDesc
	text      string
	line, col int
}

var _ Token = (*token)(nil)

func (tk *token) Position() (int, int) {
	return tk.line, tk.col
}

func (tk *token) Text() string {
	return tk.text
}

func (tk *token) String() string {
	return fmt.Sprintf("%d:%d: token(%s): %s",
		tk.line, tk.col,
		tk.SymbolicName(),
		strconv.Quote(tk.Text()),
	)
}
