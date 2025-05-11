// Code generated from RustLikeLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type RustLikeLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var RustLikeLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rustlikelexerLexerInit() {
	staticData := &RustLikeLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "", "", "'i32'", "'let'", "'if'", "'else'", "'while'", "'return'",
		"'mut'", "'fn'", "'loop'", "'break'", "'continue'", "", "", "'+'", "'-'",
		"'*'", "'/'", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'='", "';'",
		"':'", "','", "'('", "')'", "'['", "']'", "'{'", "'}'", "'->'", "'.'",
		"'..'",
	}
	staticData.SymbolicNames = []string{
		"", "SL_COMMENT", "ML_COMMENT", "WS", "INT32", "LET", "IF", "ELSE",
		"WHILE", "RETURN", "MUT", "FN", "LOOP", "BREAK", "CONTINUE", "ID", "NUMBER",
		"PLUS", "MINUS", "MULT", "DIV", "EQ", "NE", "LT", "LE", "GT", "GE",
		"ASSIGN", "SEMI", "COLON", "COMMA", "LPAREN", "RPAREN", "LBRAC", "RBRAC",
		"LBRACE", "RBRACE", "ARROW", "DOT", "DOT2",
	}
	staticData.RuleNames = []string{
		"SL_COMMENT", "ML_COMMENT", "WS", "INT32", "LET", "IF", "ELSE", "WHILE",
		"RETURN", "MUT", "FN", "LOOP", "BREAK", "CONTINUE", "ID", "NUMBER",
		"PLUS", "MINUS", "MULT", "DIV", "EQ", "NE", "LT", "LE", "GT", "GE",
		"ASSIGN", "SEMI", "COLON", "COMMA", "LPAREN", "RPAREN", "LBRAC", "RBRAC",
		"LBRACE", "RBRACE", "ARROW", "DOT", "DOT2", "DIGIT", "DIGITS", "LETTER",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 39, 250, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 90, 8, 0, 10, 0, 12, 0, 93, 9, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 103, 8, 1, 10, 1, 12,
		1, 106, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 4, 2, 114, 8, 2, 11,
		2, 12, 2, 115, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 3, 14, 178, 8, 14, 1, 14, 1, 14,
		1, 14, 5, 14, 183, 8, 14, 10, 14, 12, 14, 186, 9, 14, 1, 15, 1, 15, 1,
		16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20,
		1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1,
		25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29,
		1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1,
		35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 39,
		1, 39, 1, 40, 4, 40, 245, 8, 40, 11, 40, 12, 40, 246, 1, 41, 1, 41, 1,
		104, 0, 42, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37,
		19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55,
		28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73,
		37, 75, 38, 77, 39, 79, 0, 81, 0, 83, 0, 1, 0, 4, 1, 0, 10, 10, 3, 0, 9,
		10, 13, 13, 32, 32, 1, 0, 48, 57, 2, 0, 65, 90, 97, 122, 254, 0, 1, 1,
		0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1,
		0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17,
		1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0,
		25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0,
		0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0,
		0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0,
		0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1,
		0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63,
		1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0,
		71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0,
		1, 85, 1, 0, 0, 0, 3, 98, 1, 0, 0, 0, 5, 113, 1, 0, 0, 0, 7, 119, 1, 0,
		0, 0, 9, 123, 1, 0, 0, 0, 11, 127, 1, 0, 0, 0, 13, 130, 1, 0, 0, 0, 15,
		135, 1, 0, 0, 0, 17, 141, 1, 0, 0, 0, 19, 148, 1, 0, 0, 0, 21, 152, 1,
		0, 0, 0, 23, 155, 1, 0, 0, 0, 25, 160, 1, 0, 0, 0, 27, 166, 1, 0, 0, 0,
		29, 177, 1, 0, 0, 0, 31, 187, 1, 0, 0, 0, 33, 189, 1, 0, 0, 0, 35, 191,
		1, 0, 0, 0, 37, 193, 1, 0, 0, 0, 39, 195, 1, 0, 0, 0, 41, 197, 1, 0, 0,
		0, 43, 200, 1, 0, 0, 0, 45, 203, 1, 0, 0, 0, 47, 205, 1, 0, 0, 0, 49, 208,
		1, 0, 0, 0, 51, 210, 1, 0, 0, 0, 53, 213, 1, 0, 0, 0, 55, 215, 1, 0, 0,
		0, 57, 217, 1, 0, 0, 0, 59, 219, 1, 0, 0, 0, 61, 221, 1, 0, 0, 0, 63, 223,
		1, 0, 0, 0, 65, 225, 1, 0, 0, 0, 67, 227, 1, 0, 0, 0, 69, 229, 1, 0, 0,
		0, 71, 231, 1, 0, 0, 0, 73, 233, 1, 0, 0, 0, 75, 236, 1, 0, 0, 0, 77, 238,
		1, 0, 0, 0, 79, 241, 1, 0, 0, 0, 81, 244, 1, 0, 0, 0, 83, 248, 1, 0, 0,
		0, 85, 86, 5, 47, 0, 0, 86, 87, 5, 47, 0, 0, 87, 91, 1, 0, 0, 0, 88, 90,
		8, 0, 0, 0, 89, 88, 1, 0, 0, 0, 90, 93, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0,
		91, 92, 1, 0, 0, 0, 92, 94, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 94, 95, 5,
		10, 0, 0, 95, 96, 1, 0, 0, 0, 96, 97, 6, 0, 0, 0, 97, 2, 1, 0, 0, 0, 98,
		99, 5, 47, 0, 0, 99, 100, 5, 42, 0, 0, 100, 104, 1, 0, 0, 0, 101, 103,
		9, 0, 0, 0, 102, 101, 1, 0, 0, 0, 103, 106, 1, 0, 0, 0, 104, 105, 1, 0,
		0, 0, 104, 102, 1, 0, 0, 0, 105, 107, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0,
		107, 108, 5, 42, 0, 0, 108, 109, 5, 47, 0, 0, 109, 110, 1, 0, 0, 0, 110,
		111, 6, 1, 0, 0, 111, 4, 1, 0, 0, 0, 112, 114, 7, 1, 0, 0, 113, 112, 1,
		0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0,
		0, 116, 117, 1, 0, 0, 0, 117, 118, 6, 2, 0, 0, 118, 6, 1, 0, 0, 0, 119,
		120, 5, 105, 0, 0, 120, 121, 5, 51, 0, 0, 121, 122, 5, 50, 0, 0, 122, 8,
		1, 0, 0, 0, 123, 124, 5, 108, 0, 0, 124, 125, 5, 101, 0, 0, 125, 126, 5,
		116, 0, 0, 126, 10, 1, 0, 0, 0, 127, 128, 5, 105, 0, 0, 128, 129, 5, 102,
		0, 0, 129, 12, 1, 0, 0, 0, 130, 131, 5, 101, 0, 0, 131, 132, 5, 108, 0,
		0, 132, 133, 5, 115, 0, 0, 133, 134, 5, 101, 0, 0, 134, 14, 1, 0, 0, 0,
		135, 136, 5, 119, 0, 0, 136, 137, 5, 104, 0, 0, 137, 138, 5, 105, 0, 0,
		138, 139, 5, 108, 0, 0, 139, 140, 5, 101, 0, 0, 140, 16, 1, 0, 0, 0, 141,
		142, 5, 114, 0, 0, 142, 143, 5, 101, 0, 0, 143, 144, 5, 116, 0, 0, 144,
		145, 5, 117, 0, 0, 145, 146, 5, 114, 0, 0, 146, 147, 5, 110, 0, 0, 147,
		18, 1, 0, 0, 0, 148, 149, 5, 109, 0, 0, 149, 150, 5, 117, 0, 0, 150, 151,
		5, 116, 0, 0, 151, 20, 1, 0, 0, 0, 152, 153, 5, 102, 0, 0, 153, 154, 5,
		110, 0, 0, 154, 22, 1, 0, 0, 0, 155, 156, 5, 108, 0, 0, 156, 157, 5, 111,
		0, 0, 157, 158, 5, 111, 0, 0, 158, 159, 5, 112, 0, 0, 159, 24, 1, 0, 0,
		0, 160, 161, 5, 98, 0, 0, 161, 162, 5, 114, 0, 0, 162, 163, 5, 101, 0,
		0, 163, 164, 5, 97, 0, 0, 164, 165, 5, 107, 0, 0, 165, 26, 1, 0, 0, 0,
		166, 167, 5, 99, 0, 0, 167, 168, 5, 111, 0, 0, 168, 169, 5, 110, 0, 0,
		169, 170, 5, 116, 0, 0, 170, 171, 5, 105, 0, 0, 171, 172, 5, 110, 0, 0,
		172, 173, 5, 117, 0, 0, 173, 174, 5, 101, 0, 0, 174, 28, 1, 0, 0, 0, 175,
		178, 3, 83, 41, 0, 176, 178, 5, 95, 0, 0, 177, 175, 1, 0, 0, 0, 177, 176,
		1, 0, 0, 0, 178, 184, 1, 0, 0, 0, 179, 183, 3, 83, 41, 0, 180, 183, 5,
		95, 0, 0, 181, 183, 3, 79, 39, 0, 182, 179, 1, 0, 0, 0, 182, 180, 1, 0,
		0, 0, 182, 181, 1, 0, 0, 0, 183, 186, 1, 0, 0, 0, 184, 182, 1, 0, 0, 0,
		184, 185, 1, 0, 0, 0, 185, 30, 1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 187, 188,
		3, 81, 40, 0, 188, 32, 1, 0, 0, 0, 189, 190, 5, 43, 0, 0, 190, 34, 1, 0,
		0, 0, 191, 192, 5, 45, 0, 0, 192, 36, 1, 0, 0, 0, 193, 194, 5, 42, 0, 0,
		194, 38, 1, 0, 0, 0, 195, 196, 5, 47, 0, 0, 196, 40, 1, 0, 0, 0, 197, 198,
		5, 61, 0, 0, 198, 199, 5, 61, 0, 0, 199, 42, 1, 0, 0, 0, 200, 201, 5, 33,
		0, 0, 201, 202, 5, 61, 0, 0, 202, 44, 1, 0, 0, 0, 203, 204, 5, 60, 0, 0,
		204, 46, 1, 0, 0, 0, 205, 206, 5, 60, 0, 0, 206, 207, 5, 61, 0, 0, 207,
		48, 1, 0, 0, 0, 208, 209, 5, 62, 0, 0, 209, 50, 1, 0, 0, 0, 210, 211, 5,
		62, 0, 0, 211, 212, 5, 61, 0, 0, 212, 52, 1, 0, 0, 0, 213, 214, 5, 61,
		0, 0, 214, 54, 1, 0, 0, 0, 215, 216, 5, 59, 0, 0, 216, 56, 1, 0, 0, 0,
		217, 218, 5, 58, 0, 0, 218, 58, 1, 0, 0, 0, 219, 220, 5, 44, 0, 0, 220,
		60, 1, 0, 0, 0, 221, 222, 5, 40, 0, 0, 222, 62, 1, 0, 0, 0, 223, 224, 5,
		41, 0, 0, 224, 64, 1, 0, 0, 0, 225, 226, 5, 91, 0, 0, 226, 66, 1, 0, 0,
		0, 227, 228, 5, 93, 0, 0, 228, 68, 1, 0, 0, 0, 229, 230, 5, 123, 0, 0,
		230, 70, 1, 0, 0, 0, 231, 232, 5, 125, 0, 0, 232, 72, 1, 0, 0, 0, 233,
		234, 5, 45, 0, 0, 234, 235, 5, 62, 0, 0, 235, 74, 1, 0, 0, 0, 236, 237,
		5, 46, 0, 0, 237, 76, 1, 0, 0, 0, 238, 239, 5, 46, 0, 0, 239, 240, 5, 46,
		0, 0, 240, 78, 1, 0, 0, 0, 241, 242, 7, 2, 0, 0, 242, 80, 1, 0, 0, 0, 243,
		245, 7, 2, 0, 0, 244, 243, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 244,
		1, 0, 0, 0, 246, 247, 1, 0, 0, 0, 247, 82, 1, 0, 0, 0, 248, 249, 7, 3,
		0, 0, 249, 84, 1, 0, 0, 0, 8, 0, 91, 104, 115, 177, 182, 184, 246, 1, 6,
		0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// RustLikeLexerInit initializes any static state used to implement RustLikeLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewRustLikeLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func RustLikeLexerInit() {
	staticData := &RustLikeLexerLexerStaticData
	staticData.once.Do(rustlikelexerLexerInit)
}

// NewRustLikeLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewRustLikeLexer(input antlr.CharStream) *RustLikeLexer {
	RustLikeLexerInit()
	l := new(RustLikeLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &RustLikeLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "RustLikeLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// RustLikeLexer tokens.
const (
	RustLikeLexerSL_COMMENT = 1
	RustLikeLexerML_COMMENT = 2
	RustLikeLexerWS         = 3
	RustLikeLexerINT32      = 4
	RustLikeLexerLET        = 5
	RustLikeLexerIF         = 6
	RustLikeLexerELSE       = 7
	RustLikeLexerWHILE      = 8
	RustLikeLexerRETURN     = 9
	RustLikeLexerMUT        = 10
	RustLikeLexerFN         = 11
	RustLikeLexerLOOP       = 12
	RustLikeLexerBREAK      = 13
	RustLikeLexerCONTINUE   = 14
	RustLikeLexerID         = 15
	RustLikeLexerNUMBER     = 16
	RustLikeLexerPLUS       = 17
	RustLikeLexerMINUS      = 18
	RustLikeLexerMULT       = 19
	RustLikeLexerDIV        = 20
	RustLikeLexerEQ         = 21
	RustLikeLexerNE         = 22
	RustLikeLexerLT         = 23
	RustLikeLexerLE         = 24
	RustLikeLexerGT         = 25
	RustLikeLexerGE         = 26
	RustLikeLexerASSIGN     = 27
	RustLikeLexerSEMI       = 28
	RustLikeLexerCOLON      = 29
	RustLikeLexerCOMMA      = 30
	RustLikeLexerLPAREN     = 31
	RustLikeLexerRPAREN     = 32
	RustLikeLexerLBRAC      = 33
	RustLikeLexerRBRAC      = 34
	RustLikeLexerLBRACE     = 35
	RustLikeLexerRBRACE     = 36
	RustLikeLexerARROW      = 37
	RustLikeLexerDOT        = 38
	RustLikeLexerDOT2       = 39
)
