package lexer

import (
	"bytes"
	"io"
	"log"
	"math"
	"strings"
	"unicode"
)

type Lexer interface {
	Reset(offset int)
	Advance()
	NextToken() Token
}

type RustLikeLexer struct {
	reader    *bytes.Reader // reader
	peek_     byte          // 当前的字符
	lastPeek  byte          // 上一个字符
	line, col int           // 当前的行与列

	longestValidPrefixType   TokenType // token type
	longestValidPrefixOffset int       // 上一次匹配的位置的下一个字符的位置
	lastMatchOffset          int
}

var _ Lexer = (*RustLikeLexer)(nil)

const (
	invalidChar = math.MaxUint8
	peekEOF     = 0
)

func NewRustLikeLexer(reader *bytes.Reader) *RustLikeLexer {
	ret := &RustLikeLexer{
		reader:   reader,
		peek_:    invalidChar,
		lastPeek: invalidChar,
		line:     0,
		col:      -1,
	}
	ret.Advance()
	return ret
}

func (rl *RustLikeLexer) Reset(offset int) {
	_, err := rl.reader.Seek(int64(offset), io.SeekStart)
	if err != nil {
		panic(err)
	}
}

func (rl *RustLikeLexer) Advance() {
	var err error
	rl.lastPeek = rl.peek_
	rl.peek_, err = rl.reader.ReadByte()
	if err == io.EOF {
		rl.peek_ = peekEOF
		// msg := fmt.Sprintf("unexpected EOF at %d:%d", rl.line, rl.col)
		// panic(msg)
		return
	}

	// NOTE:这里不需要考虑\r的情况,因为在预处理环节,就可以把\r\n替换为\n
	if rl.lastPeek == '\n' {
		// 如果上一个是换行, 那么当前的就是另起一行
		rl.line++
		rl.col = 0
	} else {
		rl.col++
	}
}

func (rl *RustLikeLexer) NextToken() Token {
	if rl.reader.Len() == 0 {
		// 已经读取结束, 那么返回EOF
		return NewToken(TokenEOF, rl.line, rl.col)
	}

	// 开启状态机的转移
	var ret Token
	r := rl.peek()
	switch {
	// 处理空白
	case unicode.IsSpace(r):
		ret = rl.tokenWS()
		// 处理注释
	case r == '/':
		ret = rl.tokenCOMMENT()
	}

	log.Printf("current peek: %c(%d)", rl.peek(), rl.peek())

	return ret
}

func (rl *RustLikeLexer) offset() int {
	return int(rl.reader.Size()) - rl.reader.Len()
}

func (rl *RustLikeLexer) peek() rune {
	return rune(rl.peek_)
}

func (rl *RustLikeLexer) tokenWS() Token {
	var text strings.Builder
	sline, scol := rl.line, rl.col
	for unicode.IsSpace(rl.peek()) {
		text.WriteByte(byte(rl.peek()))
		rl.Advance()
	}
	return NewTokenWithText(TokenWS, text.String(), sline, scol)
}

func (rl *RustLikeLexer) tokenCOMMENT() Token {
	// 参照comment处理的DFA
	const (
		q0 = iota
		q1
		q2
		q3
		q4    // 多行注释的接受态
		q5    // 单行注释的接受态
		qDead // unknown token
	)

	sline, scol := rl.line, rl.col
	// 把当前的/消耗掉
	var text strings.Builder
	text.WriteByte(byte(rl.peek()))
	rl.Advance()
	state := q0

	for {
		r := rl.peek()

		if r != peekEOF {
			switch state {
			case q4, q5, qDead:
			// pass此时属于终态，不写入
			default:
				text.WriteByte(byte(r))
			}
		}

		switch state {
		case q0:
			switch r {
			case '/':
				state = q1
				rl.Advance()
			case '*':
				state = q2
				rl.Advance()
			default: // unknown token
				state = qDead
				rl.Advance()
			}

		case q1:
			switch r {
			case '\n':
				state = q5
				rl.Advance()
			case peekEOF:
				state = qDead
			default:
				// 保持当前
				rl.Advance()
			}

		case q2:
			switch r {
			case '*':
				state = q3
				rl.Advance()
			case peekEOF:
				state = qDead
			default:
				rl.Advance()
			}

		case q3:
			if r == '/' {
				state = q4
				rl.Advance()
			} else {
				state = qDead
				rl.Advance()
			}

		case q4:
			// 此时已经到达多行注释的接受态
			return NewTokenWithText(TokenMCOMMENT, text.String(), sline, scol)
		case q5:
			// 此时已经到达单行注释的接受态
			return NewTokenWithText(TokenSCOMMENT, text.String(), sline, scol)

		case qDead:
			return NewTokenWithText(TokenUNKNONW, text.String(), sline, scol)

		}
	}
}
