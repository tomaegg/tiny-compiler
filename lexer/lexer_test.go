package lexer

import (
	"bytes"
	"testing"
)

func TestNextToken(t *testing.T) {
	t.Run("WS", testNextTokenWS)

	t.Run("COMMENT", testNextTokenCOMMENT)

	t.Run("KEYWORD", testNextTokenKeywordorID)

	t.Run("NUMBER", testNextTokenNumber)

	t.Run("ASSIGNorEQ", testNextTokenAssignOrEQ)

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

func testNextTokenKeywordorID(t *testing.T) {
	b := []byte(
		`let i32 if else return mut fn loop for in break continue identifier123 _anotherID`,
	)
	reader := bytes.NewReader(b)
	lexer := NewRustLikeLexer(reader)

	expectedTokens := []struct {
		tokenType TokenType
		literal   string
	}{
		{TokenLET, "let"}, {TokenINT32, "i32"}, {TokenIF, "if"}, {TokenELSE, "else"},
		{TokenRETURN, "return"}, {TokenMUT, "mut"}, {TokenFN, "fn"}, {TokenLOOP, "loop"},
		{TokenFOR, "for"}, {TokenIN, "in"}, {TokenBREAK, "break"}, {TokenCONTINUE, "continue"},
		{TokenID, "identifier123"}, {TokenID, "_anotherID"},
	}

	for _, expected := range expectedTokens {
		tk := lexer.NextToken()
		if tk.Type() == TokenWS {
			tk = lexer.NextToken() // 跳过空格
		}
		if tk.Type() != expected.tokenType {
			t.Errorf("expected token type %d with literal '%s', got type %d with literal '%s'",
				expected.tokenType, expected.literal, tk.Type(), tk.Literal())
		} else {
			t.Logf("token: %s", tk)
		}
	}

	// Ensure the next token is EOF
	tk := lexer.NextToken()
	if tk.Type() == TokenWS {
		tk = lexer.NextToken() // Skip whitespace tokens
	}
	if tk.Type() != TokenEOF {
		t.Errorf("expected EOF, got %d", tk.Type())
	}
}

func testNextTokenNumber(t *testing.T) {
	b := []byte(
		`123 456 1 3 5 123456789`,
	)
	reader := bytes.NewReader(b)
	lexer := NewRustLikeLexer(reader)

	expectedTokens := []struct {
		tokenType TokenType
		literal   string
	}{
		{TokenNUMBER, "123"}, {TokenNUMBER, "456"}, {TokenNUMBER, "1"}, {TokenNUMBER, "3"}, {TokenNUMBER, "5"},
		{TokenNUMBER, "123456789"},
	}

	for _, expected := range expectedTokens {
		tk := lexer.NextToken()
		if tk.Type() == TokenWS {
			tk = lexer.NextToken() // 跳过空格
		}
		if tk.Type() != expected.tokenType {
			t.Errorf("expected token type %d with literal '%s', got type %d with literal '%s'",
				expected.tokenType, expected.literal, tk.Type(), tk.Literal())
		} else {
			t.Logf("token: %s", tk)
		}
	}

	tk := lexer.NextToken()
	if tk.Type() == TokenWS {
		tk = lexer.NextToken() // Skip whitespace tokens
	}
	if tk.Type() != TokenEOF {
		t.Errorf("expected EOF, got %d", tk.Type())
	}
}

func testNextTokenAssignOrEQ(t *testing.T) {
	b := []byte(
		`= ==`,
	)
	reader := bytes.NewReader(b)
	lexer := NewRustLikeLexer(reader)

	expectedTokens := []struct {
		tokenType TokenType
		literal   string
	}{
		{TokenASSIGN, "="}, {TokenEQ, "=="},
	}

	for _, expected := range expectedTokens {
		tk := lexer.NextToken()
		if tk.Type() == TokenWS {
			tk = lexer.NextToken() // 跳过空格
		}
		if tk.Type() != expected.tokenType {
			t.Errorf("expected token type %d with literal '%s', got type %d with literal '%s'",
				expected.tokenType, expected.literal, tk.Type(), tk.Literal())
		} else {
			t.Logf("token: %s", tk)
		}
	}

	tk := lexer.NextToken()
	if tk.Type() == TokenWS {
		tk = lexer.NextToken() // Skip whitespace tokens
	}
	if tk.Type() != TokenEOF {
		t.Errorf("expected EOF, got %d", tk.Type())
	}
}
