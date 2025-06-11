package cmd

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/sirupsen/logrus"
)

// 1. 自定义错误监听器
type LexerErrorListener struct {
	*antlr.DefaultErrorListener
	ErrMsg []string
}

func (l *LexerErrorListener) Errors() []string {
	return l.ErrMsg
}

func (l *LexerErrorListener) SyntaxError(recognizer antlr.Recognizer, _ any,
	line, column int, msg string, e antlr.RecognitionException,
) {
	errorMsg := fmt.Sprintf("LEXER ERROR line %d:%d - %s", line, column, msg)
	l.ErrMsg = append(l.ErrMsg, errorMsg)
	logrus.Error(errorMsg)
}

// 1. 自定义错误监听器
type ParserErrorListener struct {
	*antlr.DefaultErrorListener
	ErrMsg []string
}

func (l *ParserErrorListener) Errors() []string {
	return l.ErrMsg
}

func (l *ParserErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol any,
	line, column int, msg string, e antlr.RecognitionException,
) {
	// 获取违反规则的token（如果有）
	var tokenText string
	if token, ok := offendingSymbol.(antlr.Token); ok {
		tokenText = token.GetText()
	} else {
		tokenText = "<no token>"
	}

	errorMsg := fmt.Sprintf("PARSER ERROR at [%d:%d] token '%s' - %s",
		line, column, tokenText, msg)
	l.ErrMsg = append(l.ErrMsg, errorMsg)
	logrus.Error(errorMsg)
}
