package cmd

import (
	"tj-compiler/g4/parser"

	"github.com/antlr4-go/antlr/v4"
)

func NewLexer(input antlr.CharStream) (*parser.RustLikeLexer, func() []string) {
	lexer := parser.NewRustLikeLexer(input)
	errorListener := &LexerErrorListener{}
	lexer.RemoveErrorListeners()          // 移除默认监听器
	lexer.AddErrorListener(errorListener) // 添加自定义监听器
	return lexer, errorListener.Errors
}

func NewParser(input antlr.TokenStream) (*parser.RustLikeParser, func() []string) {
	p := parser.NewRustLikeParser(input)
	errorListener := &ParserErrorListener{}
	p.RemoveErrorListeners()          // 非常重要：移除默认监听器
	p.AddErrorListener(errorListener) // 添加自定义监听器
	return p, errorListener.Errors
}
