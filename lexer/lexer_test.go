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

	t.Run("BOUND", testNextTokenBound)

	t.Run("DIVIDE", testDivide)

	t.Run("SPECIAL", testSpecial)

	t.Run("CALC", testCalc)

	t.Run("COMPREHENSIVE", TestComprehensiveTokenParsing)

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
			t.Errorf("expected token type %s with literal '%s', got type %s with literal '%s'",
				tokenSymbolicNames[expected.tokenType], expected.literal, tokenSymbolicNames[tk.Type()], tk.Literal())
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
			t.Errorf("expected token type %s with literal '%s', got type %s with literal '%s'",
				tokenSymbolicNames[expected.tokenType], expected.literal, tokenSymbolicNames[tk.Type()], tk.Literal())
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
			t.Errorf("expected token type %s with literal '%s', got type %s with literal '%s'",
				tokenSymbolicNames[expected.tokenType], expected.literal, tokenSymbolicNames[tk.Type()], tk.Literal())
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

func testNextTokenBound(t *testing.T) {
	b := []byte(
		`( )   { }    [ ]`,
	)
	reader := bytes.NewReader(b)
	lexer := NewRustLikeLexer(reader)

	expectedTokens := []struct {
		tokenType TokenType
		literal   string
	}{
		{TokenLPAREN, "("}, {TokenRPAREN, ")"}, {TokenLBRACE, "{"}, {TokenRBRACE, "}"},
		{TokenLBRAC, "["}, {TokenRBRAC, "]"},
	}

	for _, expected := range expectedTokens {
		tk := lexer.NextToken()
		if tk.Type() == TokenWS {
			tk = lexer.NextToken()
		}
		if tk.Type() != expected.tokenType {
			t.Errorf("expected token type %s with literal '%s', got type %s with literal '%s'",
				tokenSymbolicNames[expected.tokenType], expected.literal, tokenSymbolicNames[tk.Type()], tk.Literal())
		} else {
			t.Logf("token: %s", tk)
		}
	}

	tk := lexer.NextToken()
	if tk.Type() == TokenWS {
		tk = lexer.NextToken()
	}
	if tk.Type() != TokenEOF {
		t.Errorf("expected EOF, got %d", tk.Type())
	}
}

func testDivide(t *testing.T) {
	b := []byte(
		`; : ,`,
	)
	reader := bytes.NewReader(b)
	lexer := NewRustLikeLexer(reader)

	expectedTokens := []struct {
		tokenType TokenType
		literal   string
	}{
		{TokenSEMI, ";"}, {TokenCOLON, ":"}, {TokenCOMMA, ","},
	}

	for _, expected := range expectedTokens {
		tk := lexer.NextToken()
		if tk.Type() == TokenWS {
			tk = lexer.NextToken() // Skip whitespace tokens
		}
		if tk.Type() != expected.tokenType {
			t.Errorf("expected token type %s with literal '%s', got type %s with literal '%s'",
				tokenSymbolicNames[expected.tokenType], expected.literal, tokenSymbolicNames[tk.Type()], tk.Literal())
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

func testSpecial(t *testing.T) {
	b := []byte(
		`-> .. .`,
	)
	reader := bytes.NewReader(b)
	lexer := NewRustLikeLexer(reader)

	expectedTokens := []struct {
		tokenType TokenType
		literal   string
	}{
		{TokenARROW, "->"}, {Token2DOT, ".."}, {TokenDOT, "."},
	}

	for _, expected := range expectedTokens {
		tk := lexer.NextToken()
		if tk.Type() == TokenWS {
			tk = lexer.NextToken() // Skip whitespace tokens
		}
		if tk.Type() != expected.tokenType {
			t.Errorf("expected token type %s with literal '%s', got type %s with literal '%s'",
				tokenSymbolicNames[expected.tokenType], expected.literal, tokenSymbolicNames[tk.Type()], tk.Literal())
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

func testCalc(t *testing.T) {
	b := []byte(
		`+ - * / == != < <= > >=`,
	)
	reader := bytes.NewReader(b)
	lexer := NewRustLikeLexer(reader)

	expectedTokens := []struct {
		tokenType TokenType
		literal   string
	}{
		{TokenPLUS, "+"}, {TokenMINUS, "-"}, {TokenMULTI, "*"}, {TokenDIV, "/"},
		{TokenEQ, "=="}, {TokenNE, "!="}, {TokenLT, "<"}, {TokenLE, "<="},
		{TokenGT, ">"}, {TokenGE, ">="},
	}

	for _, expected := range expectedTokens {
		tk := lexer.NextToken()
		if tk.Type() == TokenWS {
			tk = lexer.NextToken() // Skip whitespace tokens
		}
		if tk.Type() != expected.tokenType {
			t.Errorf("expected token type %s with literal '%s', got type %s with literal '%s'",
				tokenSymbolicNames[expected.tokenType], expected.literal, tokenSymbolicNames[tk.Type()], tk.Literal())
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
func TestComprehensiveTokenParsing(t *testing.T) {
	b := []byte(
		`let x = 123; if x >= 100 {
			// Check if x is large
			return x + 1;
		} else {
			/* Handle smaller values */
			let y = x - 1;
			y
		}
		-> .. . == != <= >= ( ) { } [ ] + - * / , :`,
	)
	reader := bytes.NewReader(b)
	lexer := NewRustLikeLexer(reader)

	expectedTokens := []struct {
		tokenType TokenType
		literal   string
	}{
		// Keywords and identifiers
		{TokenLET, "let"}, {TokenID, "x"}, {TokenASSIGN, "="}, {TokenNUMBER, "123"}, {TokenSEMI, ";"},
		{TokenIF, "if"}, {TokenID, "x"}, {TokenGE, ">="}, {TokenNUMBER, "100"}, {TokenLBRACE, "{"},

		// Single-line comment
		{TokenSCOMMENT, "// Check if x is large"},

		{TokenRETURN, "return"}, {TokenID, "x"}, {TokenPLUS, "+"}, {TokenNUMBER, "1"}, {TokenSEMI, ";"},
		{TokenRBRACE, "}"}, {TokenELSE, "else"}, {TokenLBRACE, "{"},

		// Multi-line comment
		{TokenMCOMMENT, "/* Handle smaller values */"},

		{TokenLET, "let"}, {TokenID, "y"}, {TokenASSIGN, "="}, {TokenID, "x"}, {TokenMINUS, "-"},
		{TokenNUMBER, "1"}, {TokenSEMI, ";"}, {TokenID, "y"}, {TokenRBRACE, "}"},

		// Special symbols and operators
		{TokenARROW, "->"}, {Token2DOT, ".."}, {TokenDOT, "."}, {TokenEQ, "=="}, {TokenNE, "!="},
		{TokenLE, "<="}, {TokenGE, ">="}, {TokenLPAREN, "("}, {TokenRPAREN, ")"}, {TokenLBRACE, "{"},
		{TokenRBRACE, "}"}, {TokenLBRAC, "["}, {TokenRBRAC, "]"}, {TokenPLUS, "+"}, {TokenMINUS, "-"},
		{TokenMULTI, "*"}, {TokenDIV, "/"}, {TokenCOMMA, ","}, {TokenCOLON, ":"},
	}

	for _, expected := range expectedTokens {
		tk := lexer.NextToken()
		if tk.Type() == TokenWS {
			tk = lexer.NextToken() // Skip whitespace tokens
		}
		if tk.Type() != expected.tokenType {
			t.Errorf("expected token type %s with literal '%s', got type %s with literal '%s'",
				tokenSymbolicNames[expected.tokenType], expected.literal, tokenSymbolicNames[tk.Type()], tk.Literal())
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
