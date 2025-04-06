package lexer

import (
	"bytes"
	"testing"
)

func TestNextToken(t *testing.T) {
	t.Run("WS", testNextTokenWS)

	t.Run("COMMENT", testNextTokenCOMMENT)
}

func testNextTokenWS(t *testing.T) {
	b := []byte(
		`     `,
	)
	reader := bytes.NewReader(b)
	lexer := NewRustLikeLexer(reader)
	tk := lexer.NextToken()
	t.Log(tk)
	tk = lexer.NextToken()
	if tk.Type() != TokenEOF {
		t.Errorf("should get token EOF, get %s", tk)
	}
}

func testNextTokenCOMMENT(t *testing.T) {
	b := []byte(
		`     // 你知道我要说什么
// 我也知道你要说什么`,
	)
	reader := bytes.NewReader(b)
	lexer := NewRustLikeLexer(reader)
	tk := lexer.NextToken()
	t.Log(tk)
	tk = lexer.NextToken()
	t.Log(tk)
	tk = lexer.NextToken()
	t.Log(tk)
}
