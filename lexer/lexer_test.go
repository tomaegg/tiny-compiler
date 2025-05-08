package lexer

import (
	"bytes"
	"fmt"
	"math/rand/v2"
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

	t.Run("RANDOM", TestNextTokenRand)
}

// TODO: 根据数组:	tokens []TokenDesc
// 生成随机token进行测试, 其中id/number/comments是需要自己生成的
// number的生成是显然的, id的生成已经给你写好了
// 只需要随机取tokens即可, 利用token.Literal()方法得到token
// 再次测试一下token的生成是否正确
// 注意写入[]byte时可以用fmt.Appendf方法
func TestNextTokenRand(t *testing.T) {
	const testCount = 1000 // 设置生成的随机token数量
	var b bytes.Buffer

	// 随机生成指定数量的token
	generatedTokens := make([]Token, 0, testCount)
	for i := 0; i < testCount; i++ {
		// 随机选择一个TokenDesc
		tokenIndex := rand.IntN(len(tokens))
		token := tokens[tokenIndex]

		if token == nil || token.Type() == TokenWS || token.Type() == TokenEOF || token.Type() == TokenUNKNONW {
			continue
		}

		// 根据token类型生成对应的字面量
		var generatedToken Token
		switch token.Type() {
		case TokenID:
			id := genTokenID()
			b.Write(id)
			generatedToken = NewTokenWithText(TokenID, string(id), 0, 0)
		case TokenNUMBER:
			number := fmt.Appendf(nil, "%d", rand.IntN(10000)) // 使用 fmt.Appendf 生成随机数字
			b.Write(number)
			generatedToken = NewTokenWithText(TokenNUMBER, string(number), 0, 0)
		case TokenSCOMMENT:
			comment := fmt.Appendf(nil, "// Random comment %d\n", rand.IntN(100))
			b.Write(comment)
			generatedToken = NewTokenWithText(TokenSCOMMENT, string(comment), 0, 0)
		case TokenMCOMMENT:
			comment := fmt.Appendf(nil, "/* Random multi-line comment %d */", rand.IntN(100))
			b.Write(comment)
			generatedToken = NewTokenWithText(TokenMCOMMENT, string(comment), 0, 0)
		default:
			b.WriteString(token.Literal())
			generatedToken = NewTokenWithText(token.Type(), token.Literal(), 0, 0)
		}

		// 添加生成的Token到列表
		generatedTokens = append(generatedTokens, generatedToken)

		b.WriteByte(' ')
	}

	// 创建Lexer并解析生成的随机token
	reader := bytes.NewReader(b.Bytes())
	lexer := NewRustLikeLexer(reader)

	for i, expectedToken := range generatedTokens {
		tk := lexer.NextToken()
		if tk.Type() == TokenWS {
			tk = lexer.NextToken() // 跳过空格
		}

		// 打印识别出的token信息
		t.Logf("Token %d: Recognized Token: %s", i, tk.String())

		// 验证生成的token是否正确
		if tk.Type() != expectedToken.Type() {
			t.Errorf("Token %d: expected type %s, got %s",
				i, tokenSymbolicNames[expectedToken.Type()], tokenSymbolicNames[tk.Type()])
		} else if tk.Literal() != expectedToken.Literal() && tk.Type() != TokenID && tk.Type() != TokenNUMBER {
			t.Errorf("Token %d: expected literal '%s', got '%s'",
				i, expectedToken.Literal(), tk.Literal())
		}
	}

	// Ensure the next token is EOF
	tk := lexer.NextToken()
	if tk.Type() == TokenWS {
		tk = lexer.NextToken()
	}
	if tk.Type() != TokenEOF {
		t.Errorf("expected EOF, got %d", tk.Type())
	}
}

func genTokenID() []byte {
	genDigit := func() byte {
		return byte(rand.IntN(10) + '0')
	}
	genLetter := func() byte {
		if rand.IntN(2) != 0 {
			return byte(rand.IntN(26) + 'a')
		}
		return byte(rand.IntN(26) + 'A')
	}
	genOthers := func() byte {
		switch rand.IntN(3) {
		case 2:
			return genLetter()
		case 1:
			return '_'
		default:
			return genDigit()
		}
	}
	genFirst := func() byte {
		switch rand.IntN(2) {
		case 0:
			return genLetter()
		default:
			return '_'
		}
	}

	// default gen lenth <= 8
	const maxIDLen = 8

	l := rand.IntN(maxIDLen) + 1
	var buffer bytes.Buffer
	buffer.WriteByte(genFirst())
	for range l - 1 {
		buffer.WriteByte(genOthers())
	}

	kwTable := GetKWTable()
	if kwTable.Exist(buffer.String()) != nil {
		return genTokenID()
	}

	return buffer.Bytes()
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
		`let i32 if else return mut fn loop for in break continue while identifier123 _anotherID`,
	)
	reader := bytes.NewReader(b)
	lexer := NewRustLikeLexer(reader)

	expectedTokens := []struct {
		tokenType TokenType
		literal   string
	}{
		{TokenLET, "let"},
		{TokenINT32, "i32"},
		{TokenIF, "if"},
		{TokenELSE, "else"},
		{TokenRETURN, "return"},
		{TokenMUT, "mut"},
		{TokenFN, "fn"},
		{TokenLOOP, "loop"},
		{TokenFOR, "for"},
		{TokenIN, "in"},
		{TokenBREAK, "break"},
		{TokenCONTINUE, "continue"},
		{TokenWhile, "while"},
		{TokenID, "identifier123"},
		{TokenID, "_anotherID"},
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
		{TokenNUMBER, "123"},
		{TokenNUMBER, "456"},
		{TokenNUMBER, "1"},
		{TokenNUMBER, "3"},
		{TokenNUMBER, "5"},
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
		{TokenLPAREN, "("},
		{TokenRPAREN, ")"},
		{TokenLBRACE, "{"},
		{TokenRBRACE, "}"},
		{TokenLBRAC, "["},
		{TokenRBRAC, "]"},
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
		{TokenPLUS, "+"},
		{TokenMINUS, "-"},
		{TokenMULTI, "*"},
		{TokenDIV, "/"},
		{TokenEQ, "=="},
		{TokenNE, "!="},
		{TokenLT, "<"},
		{TokenLE, "<="},
		{TokenGT, ">"},
		{TokenGE, ">="},
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
		{TokenLET, "let"},
		{TokenID, "x"},
		{TokenASSIGN, "="},
		{TokenNUMBER, "123"},
		{TokenSEMI, ";"},
		{TokenIF, "if"},
		{TokenID, "x"},
		{TokenGE, ">="},
		{TokenNUMBER, "100"},
		{TokenLBRACE, "{"},

		// Single-line comment
		{TokenSCOMMENT, "// Check if x is large"},

		{TokenRETURN, "return"},
		{TokenID, "x"},
		{TokenPLUS, "+"},
		{TokenNUMBER, "1"},
		{TokenSEMI, ";"},
		{TokenRBRACE, "}"},
		{TokenELSE, "else"},
		{TokenLBRACE, "{"},

		// Multi-line comment
		{TokenMCOMMENT, "/* Handle smaller values */"},

		{TokenLET, "let"},
		{TokenID, "y"},
		{TokenASSIGN, "="},
		{TokenID, "x"},
		{TokenMINUS, "-"},
		{TokenNUMBER, "1"},
		{TokenSEMI, ";"},
		{TokenID, "y"},
		{TokenRBRACE, "}"},

		// Special symbols and operators
		{TokenARROW, "->"},
		{Token2DOT, ".."},
		{TokenDOT, "."},
		{TokenEQ, "=="},
		{TokenNE, "!="},
		{TokenLE, "<="},
		{TokenGE, ">="},
		{TokenLPAREN, "("},
		{TokenRPAREN, ")"},
		{TokenLBRACE, "{"},
		{TokenRBRACE, "}"},
		{TokenLBRAC, "["},
		{TokenRBRAC, "]"},
		{TokenPLUS, "+"},
		{TokenMINUS, "-"},
		{TokenMULTI, "*"},
		{TokenDIV, "/"},
		{TokenCOMMA, ","},
		{TokenCOLON, ":"},
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
