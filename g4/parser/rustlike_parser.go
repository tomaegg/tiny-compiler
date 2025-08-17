// Code generated from RustLikeParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // RustLikeParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type RustLikeParser struct {
	*antlr.BaseParser
}

var RustLikeParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rustlikeparserParserInit() {
	staticData := &RustLikeParserParserStaticData
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
		"prog", "declaration", "expr", "exprNumber", "funcCallList", "funcCallParam",
		"funcDeclaration", "funcSignature", "funcDeclarationReturn", "funcParamsList",
		"funcParams", "funcParam", "funcBlock", "rtype", "arrayType", "arrayElems",
		"block", "stat", "ifBranch", "elifBranch", "elseBranch", "varType",
		"varInit",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 39, 246, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 1, 0, 1, 0, 1, 1, 5, 1, 50, 8, 1, 10, 1, 12, 1,
		53, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 73, 8, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 84, 8, 2, 10, 2, 12, 2,
		87, 9, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 5, 5, 98,
		8, 5, 10, 5, 12, 5, 101, 9, 5, 1, 5, 3, 5, 104, 8, 5, 1, 6, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 113, 8, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 5, 10, 125, 8, 10, 10, 10, 12, 10, 128,
		9, 10, 1, 10, 3, 10, 131, 8, 10, 1, 11, 3, 11, 134, 8, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 12, 1, 12, 5, 12, 142, 8, 12, 10, 12, 12, 12, 145, 9,
		12, 1, 12, 1, 12, 1, 13, 1, 13, 3, 13, 151, 8, 13, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 5, 15, 162, 8, 15, 10, 15, 12,
		15, 165, 9, 15, 1, 15, 3, 15, 168, 8, 15, 1, 16, 1, 16, 5, 16, 172, 8,
		16, 10, 16, 12, 16, 175, 9, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 3, 17,
		182, 8, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 191,
		8, 17, 1, 17, 1, 17, 3, 17, 195, 8, 17, 1, 17, 3, 17, 198, 8, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 5,
		17, 211, 8, 17, 10, 17, 12, 17, 214, 9, 17, 1, 17, 3, 17, 217, 8, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 226, 8, 17, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1,
		20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 0, 1, 4, 23, 0, 2,
		4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
		42, 44, 0, 3, 1, 0, 19, 20, 1, 0, 17, 18, 1, 0, 21, 26, 258, 0, 46, 1,
		0, 0, 0, 2, 51, 1, 0, 0, 0, 4, 72, 1, 0, 0, 0, 6, 88, 1, 0, 0, 0, 8, 90,
		1, 0, 0, 0, 10, 103, 1, 0, 0, 0, 12, 105, 1, 0, 0, 0, 14, 108, 1, 0, 0,
		0, 16, 114, 1, 0, 0, 0, 18, 117, 1, 0, 0, 0, 20, 130, 1, 0, 0, 0, 22, 133,
		1, 0, 0, 0, 24, 139, 1, 0, 0, 0, 26, 150, 1, 0, 0, 0, 28, 152, 1, 0, 0,
		0, 30, 167, 1, 0, 0, 0, 32, 169, 1, 0, 0, 0, 34, 225, 1, 0, 0, 0, 36, 227,
		1, 0, 0, 0, 38, 231, 1, 0, 0, 0, 40, 236, 1, 0, 0, 0, 42, 239, 1, 0, 0,
		0, 44, 242, 1, 0, 0, 0, 46, 47, 3, 2, 1, 0, 47, 1, 1, 0, 0, 0, 48, 50,
		3, 12, 6, 0, 49, 48, 1, 0, 0, 0, 50, 53, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0,
		51, 52, 1, 0, 0, 0, 52, 3, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 54, 55, 6, 2,
		-1, 0, 55, 56, 5, 33, 0, 0, 56, 57, 3, 30, 15, 0, 57, 58, 5, 34, 0, 0,
		58, 73, 1, 0, 0, 0, 59, 60, 5, 31, 0, 0, 60, 61, 3, 4, 2, 0, 61, 62, 5,
		32, 0, 0, 62, 73, 1, 0, 0, 0, 63, 64, 5, 15, 0, 0, 64, 65, 5, 33, 0, 0,
		65, 66, 3, 4, 2, 0, 66, 67, 5, 34, 0, 0, 67, 73, 1, 0, 0, 0, 68, 69, 5,
		15, 0, 0, 69, 73, 3, 8, 4, 0, 70, 73, 5, 15, 0, 0, 71, 73, 3, 6, 3, 0,
		72, 54, 1, 0, 0, 0, 72, 59, 1, 0, 0, 0, 72, 63, 1, 0, 0, 0, 72, 68, 1,
		0, 0, 0, 72, 70, 1, 0, 0, 0, 72, 71, 1, 0, 0, 0, 73, 85, 1, 0, 0, 0, 74,
		75, 10, 9, 0, 0, 75, 76, 7, 0, 0, 0, 76, 84, 3, 4, 2, 10, 77, 78, 10, 8,
		0, 0, 78, 79, 7, 1, 0, 0, 79, 84, 3, 4, 2, 9, 80, 81, 10, 7, 0, 0, 81,
		82, 7, 2, 0, 0, 82, 84, 3, 4, 2, 8, 83, 74, 1, 0, 0, 0, 83, 77, 1, 0, 0,
		0, 83, 80, 1, 0, 0, 0, 84, 87, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 85, 86,
		1, 0, 0, 0, 86, 5, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 88, 89, 5, 16, 0, 0,
		89, 7, 1, 0, 0, 0, 90, 91, 5, 31, 0, 0, 91, 92, 3, 10, 5, 0, 92, 93, 5,
		32, 0, 0, 93, 9, 1, 0, 0, 0, 94, 99, 3, 4, 2, 0, 95, 96, 5, 30, 0, 0, 96,
		98, 3, 4, 2, 0, 97, 95, 1, 0, 0, 0, 98, 101, 1, 0, 0, 0, 99, 97, 1, 0,
		0, 0, 99, 100, 1, 0, 0, 0, 100, 104, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 102,
		104, 1, 0, 0, 0, 103, 94, 1, 0, 0, 0, 103, 102, 1, 0, 0, 0, 104, 11, 1,
		0, 0, 0, 105, 106, 3, 14, 7, 0, 106, 107, 3, 24, 12, 0, 107, 13, 1, 0,
		0, 0, 108, 109, 5, 11, 0, 0, 109, 110, 5, 15, 0, 0, 110, 112, 3, 18, 9,
		0, 111, 113, 3, 16, 8, 0, 112, 111, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113,
		15, 1, 0, 0, 0, 114, 115, 5, 37, 0, 0, 115, 116, 3, 26, 13, 0, 116, 17,
		1, 0, 0, 0, 117, 118, 5, 31, 0, 0, 118, 119, 3, 20, 10, 0, 119, 120, 5,
		32, 0, 0, 120, 19, 1, 0, 0, 0, 121, 126, 3, 22, 11, 0, 122, 123, 5, 30,
		0, 0, 123, 125, 3, 22, 11, 0, 124, 122, 1, 0, 0, 0, 125, 128, 1, 0, 0,
		0, 126, 124, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 131, 1, 0, 0, 0, 128,
		126, 1, 0, 0, 0, 129, 131, 1, 0, 0, 0, 130, 121, 1, 0, 0, 0, 130, 129,
		1, 0, 0, 0, 131, 21, 1, 0, 0, 0, 132, 134, 5, 10, 0, 0, 133, 132, 1, 0,
		0, 0, 133, 134, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 136, 5, 15, 0, 0,
		136, 137, 5, 29, 0, 0, 137, 138, 3, 26, 13, 0, 138, 23, 1, 0, 0, 0, 139,
		143, 5, 35, 0, 0, 140, 142, 3, 34, 17, 0, 141, 140, 1, 0, 0, 0, 142, 145,
		1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 146, 1, 0,
		0, 0, 145, 143, 1, 0, 0, 0, 146, 147, 5, 36, 0, 0, 147, 25, 1, 0, 0, 0,
		148, 151, 5, 4, 0, 0, 149, 151, 3, 28, 14, 0, 150, 148, 1, 0, 0, 0, 150,
		149, 1, 0, 0, 0, 151, 27, 1, 0, 0, 0, 152, 153, 5, 33, 0, 0, 153, 154,
		3, 26, 13, 0, 154, 155, 5, 28, 0, 0, 155, 156, 3, 6, 3, 0, 156, 157, 5,
		34, 0, 0, 157, 29, 1, 0, 0, 0, 158, 163, 3, 4, 2, 0, 159, 160, 5, 30, 0,
		0, 160, 162, 3, 4, 2, 0, 161, 159, 1, 0, 0, 0, 162, 165, 1, 0, 0, 0, 163,
		161, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 168, 1, 0, 0, 0, 165, 163,
		1, 0, 0, 0, 166, 168, 1, 0, 0, 0, 167, 158, 1, 0, 0, 0, 167, 166, 1, 0,
		0, 0, 168, 31, 1, 0, 0, 0, 169, 173, 5, 35, 0, 0, 170, 172, 3, 34, 17,
		0, 171, 170, 1, 0, 0, 0, 172, 175, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 173,
		174, 1, 0, 0, 0, 174, 176, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 176, 177,
		5, 36, 0, 0, 177, 33, 1, 0, 0, 0, 178, 226, 3, 32, 16, 0, 179, 181, 5,
		9, 0, 0, 180, 182, 3, 4, 2, 0, 181, 180, 1, 0, 0, 0, 181, 182, 1, 0, 0,
		0, 182, 183, 1, 0, 0, 0, 183, 226, 5, 28, 0, 0, 184, 185, 5, 13, 0, 0,
		185, 226, 5, 28, 0, 0, 186, 187, 5, 14, 0, 0, 187, 226, 5, 28, 0, 0, 188,
		190, 5, 5, 0, 0, 189, 191, 5, 10, 0, 0, 190, 189, 1, 0, 0, 0, 190, 191,
		1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 194, 5, 15, 0, 0, 193, 195, 3, 42,
		21, 0, 194, 193, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 197, 1, 0, 0, 0,
		196, 198, 3, 44, 22, 0, 197, 196, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198,
		199, 1, 0, 0, 0, 199, 226, 5, 28, 0, 0, 200, 201, 5, 15, 0, 0, 201, 202,
		5, 27, 0, 0, 202, 203, 3, 4, 2, 0, 203, 204, 5, 28, 0, 0, 204, 226, 1,
		0, 0, 0, 205, 206, 3, 4, 2, 0, 206, 207, 5, 28, 0, 0, 207, 226, 1, 0, 0,
		0, 208, 212, 3, 36, 18, 0, 209, 211, 3, 38, 19, 0, 210, 209, 1, 0, 0, 0,
		211, 214, 1, 0, 0, 0, 212, 210, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213,
		216, 1, 0, 0, 0, 214, 212, 1, 0, 0, 0, 215, 217, 3, 40, 20, 0, 216, 215,
		1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 226, 1, 0, 0, 0, 218, 219, 5, 8,
		0, 0, 219, 220, 3, 4, 2, 0, 220, 221, 3, 32, 16, 0, 221, 226, 1, 0, 0,
		0, 222, 223, 5, 12, 0, 0, 223, 226, 3, 32, 16, 0, 224, 226, 5, 28, 0, 0,
		225, 178, 1, 0, 0, 0, 225, 179, 1, 0, 0, 0, 225, 184, 1, 0, 0, 0, 225,
		186, 1, 0, 0, 0, 225, 188, 1, 0, 0, 0, 225, 200, 1, 0, 0, 0, 225, 205,
		1, 0, 0, 0, 225, 208, 1, 0, 0, 0, 225, 218, 1, 0, 0, 0, 225, 222, 1, 0,
		0, 0, 225, 224, 1, 0, 0, 0, 226, 35, 1, 0, 0, 0, 227, 228, 5, 6, 0, 0,
		228, 229, 3, 4, 2, 0, 229, 230, 3, 32, 16, 0, 230, 37, 1, 0, 0, 0, 231,
		232, 5, 7, 0, 0, 232, 233, 5, 6, 0, 0, 233, 234, 3, 4, 2, 0, 234, 235,
		3, 32, 16, 0, 235, 39, 1, 0, 0, 0, 236, 237, 5, 7, 0, 0, 237, 238, 3, 32,
		16, 0, 238, 41, 1, 0, 0, 0, 239, 240, 5, 29, 0, 0, 240, 241, 3, 26, 13,
		0, 241, 43, 1, 0, 0, 0, 242, 243, 5, 27, 0, 0, 243, 244, 3, 4, 2, 0, 244,
		45, 1, 0, 0, 0, 22, 51, 72, 83, 85, 99, 103, 112, 126, 130, 133, 143, 150,
		163, 167, 173, 181, 190, 194, 197, 212, 216, 225,
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

// RustLikeParserInit initializes any static state used to implement RustLikeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRustLikeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RustLikeParserInit() {
	staticData := &RustLikeParserParserStaticData
	staticData.once.Do(rustlikeparserParserInit)
}

// NewRustLikeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRustLikeParser(input antlr.TokenStream) *RustLikeParser {
	RustLikeParserInit()
	this := new(RustLikeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &RustLikeParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "RustLikeParser.g4"

	return this
}

// RustLikeParser tokens.
const (
	RustLikeParserEOF        = antlr.TokenEOF
	RustLikeParserSL_COMMENT = 1
	RustLikeParserML_COMMENT = 2
	RustLikeParserWS         = 3
	RustLikeParserINT32      = 4
	RustLikeParserLET        = 5
	RustLikeParserIF         = 6
	RustLikeParserELSE       = 7
	RustLikeParserWHILE      = 8
	RustLikeParserRETURN     = 9
	RustLikeParserMUT        = 10
	RustLikeParserFN         = 11
	RustLikeParserLOOP       = 12
	RustLikeParserBREAK      = 13
	RustLikeParserCONTINUE   = 14
	RustLikeParserID         = 15
	RustLikeParserNUMBER     = 16
	RustLikeParserPLUS       = 17
	RustLikeParserMINUS      = 18
	RustLikeParserMULT       = 19
	RustLikeParserDIV        = 20
	RustLikeParserEQ         = 21
	RustLikeParserNE         = 22
	RustLikeParserLT         = 23
	RustLikeParserLE         = 24
	RustLikeParserGT         = 25
	RustLikeParserGE         = 26
	RustLikeParserASSIGN     = 27
	RustLikeParserSEMI       = 28
	RustLikeParserCOLON      = 29
	RustLikeParserCOMMA      = 30
	RustLikeParserLPAREN     = 31
	RustLikeParserRPAREN     = 32
	RustLikeParserLBRAC      = 33
	RustLikeParserRBRAC      = 34
	RustLikeParserLBRACE     = 35
	RustLikeParserRBRACE     = 36
	RustLikeParserARROW      = 37
	RustLikeParserDOT        = 38
	RustLikeParserDOT2       = 39
)

// RustLikeParser rules.
const (
	RustLikeParserRULE_prog                  = 0
	RustLikeParserRULE_declaration           = 1
	RustLikeParserRULE_expr                  = 2
	RustLikeParserRULE_exprNumber            = 3
	RustLikeParserRULE_funcCallList          = 4
	RustLikeParserRULE_funcCallParam         = 5
	RustLikeParserRULE_funcDeclaration       = 6
	RustLikeParserRULE_funcSignature         = 7
	RustLikeParserRULE_funcDeclarationReturn = 8
	RustLikeParserRULE_funcParamsList        = 9
	RustLikeParserRULE_funcParams            = 10
	RustLikeParserRULE_funcParam             = 11
	RustLikeParserRULE_funcBlock             = 12
	RustLikeParserRULE_rtype                 = 13
	RustLikeParserRULE_arrayType             = 14
	RustLikeParserRULE_arrayElems            = 15
	RustLikeParserRULE_block                 = 16
	RustLikeParserRULE_stat                  = 17
	RustLikeParserRULE_ifBranch              = 18
	RustLikeParserRULE_elifBranch            = 19
	RustLikeParserRULE_elseBranch            = 20
	RustLikeParserRULE_varType               = 21
	RustLikeParserRULE_varInit               = 22
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Declaration() IDeclarationContext

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) Declaration() IDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RustLikeParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RustLikeParserRULE_prog)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Declaration()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFuncDeclaration() []IFuncDeclarationContext
	FuncDeclaration(i int) IFuncDeclarationContext

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_declaration
	return p
}

func InitEmptyDeclarationContext(p *DeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_declaration
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) AllFuncDeclaration() []IFuncDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IFuncDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncDeclarationContext); ok {
			tst[i] = t.(IFuncDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *DeclarationContext) FuncDeclaration(i int) IFuncDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncDeclarationContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RustLikeParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RustLikeParserRULE_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == RustLikeParserFN {
		{
			p.SetState(48)
			p.FuncDeclaration()
		}

		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExprArrayAccessContext struct {
	ExprContext
}

func NewExprArrayAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprArrayAccessContext {
	var p = new(ExprArrayAccessContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprArrayAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprArrayAccessContext) ID() antlr.TerminalNode {
	return s.GetToken(RustLikeParserID, 0)
}

func (s *ExprArrayAccessContext) LBRAC() antlr.TerminalNode {
	return s.GetToken(RustLikeParserLBRAC, 0)
}

func (s *ExprArrayAccessContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprArrayAccessContext) RBRAC() antlr.TerminalNode {
	return s.GetToken(RustLikeParserRBRAC, 0)
}

func (s *ExprArrayAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitExprArrayAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprAddSubContext struct {
	ExprContext
	lhs IExprContext
	op  antlr.Token
	rhs IExprContext
}

func NewExprAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprAddSubContext {
	var p = new(ExprAddSubContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprAddSubContext) GetOp() antlr.Token { return s.op }

func (s *ExprAddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprAddSubContext) GetLhs() IExprContext { return s.lhs }

func (s *ExprAddSubContext) GetRhs() IExprContext { return s.rhs }

func (s *ExprAddSubContext) SetLhs(v IExprContext) { s.lhs = v }

func (s *ExprAddSubContext) SetRhs(v IExprContext) { s.rhs = v }

func (s *ExprAddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprAddSubContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprAddSubContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprAddSubContext) PLUS() antlr.TerminalNode {
	return s.GetToken(RustLikeParserPLUS, 0)
}

func (s *ExprAddSubContext) MINUS() antlr.TerminalNode {
	return s.GetToken(RustLikeParserMINUS, 0)
}

func (s *ExprAddSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitExprAddSub(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprParenContext struct {
	ExprContext
}

func NewExprParenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprParenContext {
	var p = new(ExprParenContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprParenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprParenContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(RustLikeParserLPAREN, 0)
}

func (s *ExprParenContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprParenContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(RustLikeParserRPAREN, 0)
}

func (s *ExprParenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitExprParen(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprNumContext struct {
	ExprContext
}

func NewExprNumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprNumContext {
	var p = new(ExprNumContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprNumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprNumContext) ExprNumber() IExprNumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprNumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprNumberContext)
}

func (s *ExprNumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitExprNum(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprMulDivContext struct {
	ExprContext
	lhs IExprContext
	op  antlr.Token
	rhs IExprContext
}

func NewExprMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprMulDivContext {
	var p = new(ExprMulDivContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprMulDivContext) GetOp() antlr.Token { return s.op }

func (s *ExprMulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprMulDivContext) GetLhs() IExprContext { return s.lhs }

func (s *ExprMulDivContext) GetRhs() IExprContext { return s.rhs }

func (s *ExprMulDivContext) SetLhs(v IExprContext) { s.lhs = v }

func (s *ExprMulDivContext) SetRhs(v IExprContext) { s.rhs = v }

func (s *ExprMulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprMulDivContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprMulDivContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprMulDivContext) MULT() antlr.TerminalNode {
	return s.GetToken(RustLikeParserMULT, 0)
}

func (s *ExprMulDivContext) DIV() antlr.TerminalNode {
	return s.GetToken(RustLikeParserDIV, 0)
}

func (s *ExprMulDivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitExprMulDiv(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprCmpContext struct {
	ExprContext
	lhs IExprContext
	op  antlr.Token
	rhs IExprContext
}

func NewExprCmpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprCmpContext {
	var p = new(ExprCmpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprCmpContext) GetOp() antlr.Token { return s.op }

func (s *ExprCmpContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprCmpContext) GetLhs() IExprContext { return s.lhs }

func (s *ExprCmpContext) GetRhs() IExprContext { return s.rhs }

func (s *ExprCmpContext) SetLhs(v IExprContext) { s.lhs = v }

func (s *ExprCmpContext) SetRhs(v IExprContext) { s.rhs = v }

func (s *ExprCmpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprCmpContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprCmpContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprCmpContext) EQ() antlr.TerminalNode {
	return s.GetToken(RustLikeParserEQ, 0)
}

func (s *ExprCmpContext) NE() antlr.TerminalNode {
	return s.GetToken(RustLikeParserNE, 0)
}

func (s *ExprCmpContext) LT() antlr.TerminalNode {
	return s.GetToken(RustLikeParserLT, 0)
}

func (s *ExprCmpContext) GT() antlr.TerminalNode {
	return s.GetToken(RustLikeParserGT, 0)
}

func (s *ExprCmpContext) LE() antlr.TerminalNode {
	return s.GetToken(RustLikeParserLE, 0)
}

func (s *ExprCmpContext) GE() antlr.TerminalNode {
	return s.GetToken(RustLikeParserGE, 0)
}

func (s *ExprCmpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitExprCmp(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprFuncCallContext struct {
	ExprContext
}

func NewExprFuncCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprFuncCallContext {
	var p = new(ExprFuncCallContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprFuncCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprFuncCallContext) ID() antlr.TerminalNode {
	return s.GetToken(RustLikeParserID, 0)
}

func (s *ExprFuncCallContext) FuncCallList() IFuncCallListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncCallListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncCallListContext)
}

func (s *ExprFuncCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitExprFuncCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprArrayContext struct {
	ExprContext
}

func NewExprArrayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprArrayContext {
	var p = new(ExprArrayContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprArrayContext) LBRAC() antlr.TerminalNode {
	return s.GetToken(RustLikeParserLBRAC, 0)
}

func (s *ExprArrayContext) ArrayElems() IArrayElemsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayElemsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayElemsContext)
}

func (s *ExprArrayContext) RBRAC() antlr.TerminalNode {
	return s.GetToken(RustLikeParserRBRAC, 0)
}

func (s *ExprArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitExprArray(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprIDContext struct {
	ExprContext
}

func NewExprIDContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprIDContext {
	var p = new(ExprIDContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprIDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprIDContext) ID() antlr.TerminalNode {
	return s.GetToken(RustLikeParserID, 0)
}

func (s *ExprIDContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitExprID(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RustLikeParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *RustLikeParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, RustLikeParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExprArrayContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(55)
			p.Match(RustLikeParserLBRAC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(56)
			p.ArrayElems()
		}
		{
			p.SetState(57)
			p.Match(RustLikeParserRBRAC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewExprParenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(59)
			p.Match(RustLikeParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(60)
			p.expr(0)
		}
		{
			p.SetState(61)
			p.Match(RustLikeParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewExprArrayAccessContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(63)
			p.Match(RustLikeParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(64)
			p.Match(RustLikeParserLBRAC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)
			p.expr(0)
		}
		{
			p.SetState(66)
			p.Match(RustLikeParserRBRAC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewExprFuncCallContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(68)
			p.Match(RustLikeParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(69)
			p.FuncCallList()
		}

	case 5:
		localctx = NewExprIDContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(70)
			p.Match(RustLikeParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewExprNumContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(71)
			p.ExprNumber()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(83)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprMulDivContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ExprMulDivContext).lhs = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RustLikeParserRULE_expr)
				p.SetState(74)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(75)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprMulDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == RustLikeParserMULT || _la == RustLikeParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprMulDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(76)

					var _x = p.expr(10)

					localctx.(*ExprMulDivContext).rhs = _x
				}

			case 2:
				localctx = NewExprAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ExprAddSubContext).lhs = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RustLikeParserRULE_expr)
				p.SetState(77)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(78)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprAddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == RustLikeParserPLUS || _la == RustLikeParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprAddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(79)

					var _x = p.expr(9)

					localctx.(*ExprAddSubContext).rhs = _x
				}

			case 3:
				localctx = NewExprCmpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ExprCmpContext).lhs = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RustLikeParserRULE_expr)
				p.SetState(80)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(81)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprCmpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&132120576) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprCmpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(82)

					var _x = p.expr(8)

					localctx.(*ExprCmpContext).rhs = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprNumberContext is an interface to support dynamic dispatch.
type IExprNumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode

	// IsExprNumberContext differentiates from other interfaces.
	IsExprNumberContext()
}

type ExprNumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprNumberContext() *ExprNumberContext {
	var p = new(ExprNumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_exprNumber
	return p
}

func InitEmptyExprNumberContext(p *ExprNumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_exprNumber
}

func (*ExprNumberContext) IsExprNumberContext() {}

func NewExprNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprNumberContext {
	var p = new(ExprNumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_exprNumber

	return p
}

func (s *ExprNumberContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprNumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(RustLikeParserNUMBER, 0)
}

func (s *ExprNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprNumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprNumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitExprNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RustLikeParser) ExprNumber() (localctx IExprNumberContext) {
	localctx = NewExprNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RustLikeParserRULE_exprNumber)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Match(RustLikeParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncCallListContext is an interface to support dynamic dispatch.
type IFuncCallListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	FuncCallParam() IFuncCallParamContext
	RPAREN() antlr.TerminalNode

	// IsFuncCallListContext differentiates from other interfaces.
	IsFuncCallListContext()
}

type FuncCallListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncCallListContext() *FuncCallListContext {
	var p = new(FuncCallListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcCallList
	return p
}

func InitEmptyFuncCallListContext(p *FuncCallListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcCallList
}

func (*FuncCallListContext) IsFuncCallListContext() {}

func NewFuncCallListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncCallListContext {
	var p = new(FuncCallListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_funcCallList

	return p
}

func (s *FuncCallListContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncCallListContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(RustLikeParserLPAREN, 0)
}

func (s *FuncCallListContext) FuncCallParam() IFuncCallParamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncCallParamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncCallParamContext)
}

func (s *FuncCallListContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(RustLikeParserRPAREN, 0)
}

func (s *FuncCallListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncCallListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitFuncCallList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RustLikeParser) FuncCallList() (localctx IFuncCallListContext) {
	localctx = NewFuncCallListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RustLikeParserRULE_funcCallList)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.Match(RustLikeParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(91)
		p.FuncCallParam()
	}
	{
		p.SetState(92)
		p.Match(RustLikeParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncCallParamContext is an interface to support dynamic dispatch.
type IFuncCallParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFuncCallParamContext differentiates from other interfaces.
	IsFuncCallParamContext()
}

type FuncCallParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncCallParamContext() *FuncCallParamContext {
	var p = new(FuncCallParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcCallParam
	return p
}

func InitEmptyFuncCallParamContext(p *FuncCallParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcCallParam
}

func (*FuncCallParamContext) IsFuncCallParamContext() {}

func NewFuncCallParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncCallParamContext {
	var p = new(FuncCallParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_funcCallParam

	return p
}

func (s *FuncCallParamContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncCallParamContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *FuncCallParamContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FuncCallParamContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(RustLikeParserCOMMA)
}

func (s *FuncCallParamContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(RustLikeParserCOMMA, i)
}

func (s *FuncCallParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncCallParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitFuncCallParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RustLikeParser) FuncCallParam() (localctx IFuncCallParamContext) {
	localctx = NewFuncCallParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RustLikeParserRULE_funcCallParam)
	var _la int

	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RustLikeParserID, RustLikeParserNUMBER, RustLikeParserLPAREN, RustLikeParserLBRAC:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(94)
			p.expr(0)
		}
		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == RustLikeParserCOMMA {
			{
				p.SetState(95)
				p.Match(RustLikeParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(96)
				p.expr(0)
			}

			p.SetState(101)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case RustLikeParserRPAREN:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncDeclarationContext is an interface to support dynamic dispatch.
type IFuncDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FuncSignature() IFuncSignatureContext
	FuncBlock() IFuncBlockContext

	// IsFuncDeclarationContext differentiates from other interfaces.
	IsFuncDeclarationContext()
}

type FuncDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDeclarationContext() *FuncDeclarationContext {
	var p = new(FuncDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcDeclaration
	return p
}

func InitEmptyFuncDeclarationContext(p *FuncDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcDeclaration
}

func (*FuncDeclarationContext) IsFuncDeclarationContext() {}

func NewFuncDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDeclarationContext {
	var p = new(FuncDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_funcDeclaration

	return p
}

func (s *FuncDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDeclarationContext) FuncSignature() IFuncSignatureContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncSignatureContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncSignatureContext)
}

func (s *FuncDeclarationContext) FuncBlock() IFuncBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncBlockContext)
}

func (s *FuncDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitFuncDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RustLikeParser) FuncDeclaration() (localctx IFuncDeclarationContext) {
	localctx = NewFuncDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RustLikeParserRULE_funcDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.FuncSignature()
	}
	{
		p.SetState(106)
		p.FuncBlock()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncSignatureContext is an interface to support dynamic dispatch.
type IFuncSignatureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FN() antlr.TerminalNode
	ID() antlr.TerminalNode
	FuncParamsList() IFuncParamsListContext
	FuncDeclarationReturn() IFuncDeclarationReturnContext

	// IsFuncSignatureContext differentiates from other interfaces.
	IsFuncSignatureContext()
}

type FuncSignatureContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncSignatureContext() *FuncSignatureContext {
	var p = new(FuncSignatureContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcSignature
	return p
}

func InitEmptyFuncSignatureContext(p *FuncSignatureContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcSignature
}

func (*FuncSignatureContext) IsFuncSignatureContext() {}

func NewFuncSignatureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncSignatureContext {
	var p = new(FuncSignatureContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_funcSignature

	return p
}

func (s *FuncSignatureContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncSignatureContext) FN() antlr.TerminalNode {
	return s.GetToken(RustLikeParserFN, 0)
}

func (s *FuncSignatureContext) ID() antlr.TerminalNode {
	return s.GetToken(RustLikeParserID, 0)
}

func (s *FuncSignatureContext) FuncParamsList() IFuncParamsListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncParamsListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncParamsListContext)
}

func (s *FuncSignatureContext) FuncDeclarationReturn() IFuncDeclarationReturnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDeclarationReturnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncDeclarationReturnContext)
}

func (s *FuncSignatureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncSignatureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncSignatureContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitFuncSignature(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RustLikeParser) FuncSignature() (localctx IFuncSignatureContext) {
	localctx = NewFuncSignatureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, RustLikeParserRULE_funcSignature)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(RustLikeParserFN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
		p.Match(RustLikeParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(110)
		p.FuncParamsList()
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == RustLikeParserARROW {
		{
			p.SetState(111)
			p.FuncDeclarationReturn()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncDeclarationReturnContext is an interface to support dynamic dispatch.
type IFuncDeclarationReturnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ARROW() antlr.TerminalNode
	Rtype() IRtypeContext

	// IsFuncDeclarationReturnContext differentiates from other interfaces.
	IsFuncDeclarationReturnContext()
}

type FuncDeclarationReturnContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDeclarationReturnContext() *FuncDeclarationReturnContext {
	var p = new(FuncDeclarationReturnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcDeclarationReturn
	return p
}

func InitEmptyFuncDeclarationReturnContext(p *FuncDeclarationReturnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcDeclarationReturn
}

func (*FuncDeclarationReturnContext) IsFuncDeclarationReturnContext() {}

func NewFuncDeclarationReturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDeclarationReturnContext {
	var p = new(FuncDeclarationReturnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_funcDeclarationReturn

	return p
}

func (s *FuncDeclarationReturnContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDeclarationReturnContext) ARROW() antlr.TerminalNode {
	return s.GetToken(RustLikeParserARROW, 0)
}

func (s *FuncDeclarationReturnContext) Rtype() IRtypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRtypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRtypeContext)
}

func (s *FuncDeclarationReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDeclarationReturnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDeclarationReturnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitFuncDeclarationReturn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RustLikeParser) FuncDeclarationReturn() (localctx IFuncDeclarationReturnContext) {
	localctx = NewFuncDeclarationReturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RustLikeParserRULE_funcDeclarationReturn)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(RustLikeParserARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(115)
		p.Rtype()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncParamsListContext is an interface to support dynamic dispatch.
type IFuncParamsListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	FuncParams() IFuncParamsContext
	RPAREN() antlr.TerminalNode

	// IsFuncParamsListContext differentiates from other interfaces.
	IsFuncParamsListContext()
}

type FuncParamsListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncParamsListContext() *FuncParamsListContext {
	var p = new(FuncParamsListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcParamsList
	return p
}

func InitEmptyFuncParamsListContext(p *FuncParamsListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcParamsList
}

func (*FuncParamsListContext) IsFuncParamsListContext() {}

func NewFuncParamsListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncParamsListContext {
	var p = new(FuncParamsListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_funcParamsList

	return p
}

func (s *FuncParamsListContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncParamsListContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(RustLikeParserLPAREN, 0)
}

func (s *FuncParamsListContext) FuncParams() IFuncParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncParamsContext)
}

func (s *FuncParamsListContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(RustLikeParserRPAREN, 0)
}

func (s *FuncParamsListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncParamsListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncParamsListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitFuncParamsList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RustLikeParser) FuncParamsList() (localctx IFuncParamsListContext) {
	localctx = NewFuncParamsListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, RustLikeParserRULE_funcParamsList)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(RustLikeParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(118)
		p.FuncParams()
	}
	{
		p.SetState(119)
		p.Match(RustLikeParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncParamsContext is an interface to support dynamic dispatch.
type IFuncParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFuncParam() []IFuncParamContext
	FuncParam(i int) IFuncParamContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFuncParamsContext differentiates from other interfaces.
	IsFuncParamsContext()
}

type FuncParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncParamsContext() *FuncParamsContext {
	var p = new(FuncParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcParams
	return p
}

func InitEmptyFuncParamsContext(p *FuncParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcParams
}

func (*FuncParamsContext) IsFuncParamsContext() {}

func NewFuncParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncParamsContext {
	var p = new(FuncParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_funcParams

	return p
}

func (s *FuncParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncParamsContext) AllFuncParam() []IFuncParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncParamContext); ok {
			len++
		}
	}

	tst := make([]IFuncParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncParamContext); ok {
			tst[i] = t.(IFuncParamContext)
			i++
		}
	}

	return tst
}

func (s *FuncParamsContext) FuncParam(i int) IFuncParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncParamContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncParamContext)
}

func (s *FuncParamsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(RustLikeParserCOMMA)
}

func (s *FuncParamsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(RustLikeParserCOMMA, i)
}

func (s *FuncParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitFuncParams(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RustLikeParser) FuncParams() (localctx IFuncParamsContext) {
	localctx = NewFuncParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, RustLikeParserRULE_funcParams)
	var _la int

	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RustLikeParserMUT, RustLikeParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.FuncParam()
		}
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == RustLikeParserCOMMA {
			{
				p.SetState(122)
				p.Match(RustLikeParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(123)
				p.FuncParam()
			}

			p.SetState(128)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case RustLikeParserRPAREN:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncParamContext is an interface to support dynamic dispatch.
type IFuncParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Rtype() IRtypeContext
	MUT() antlr.TerminalNode

	// IsFuncParamContext differentiates from other interfaces.
	IsFuncParamContext()
}

type FuncParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncParamContext() *FuncParamContext {
	var p = new(FuncParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcParam
	return p
}

func InitEmptyFuncParamContext(p *FuncParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcParam
}

func (*FuncParamContext) IsFuncParamContext() {}

func NewFuncParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncParamContext {
	var p = new(FuncParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_funcParam

	return p
}

func (s *FuncParamContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncParamContext) ID() antlr.TerminalNode {
	return s.GetToken(RustLikeParserID, 0)
}

func (s *FuncParamContext) COLON() antlr.TerminalNode {
	return s.GetToken(RustLikeParserCOLON, 0)
}

func (s *FuncParamContext) Rtype() IRtypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRtypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRtypeContext)
}

func (s *FuncParamContext) MUT() antlr.TerminalNode {
	return s.GetToken(RustLikeParserMUT, 0)
}

func (s *FuncParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitFuncParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RustLikeParser) FuncParam() (localctx IFuncParamContext) {
	localctx = NewFuncParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, RustLikeParserRULE_funcParam)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == RustLikeParserMUT {
		{
			p.SetState(132)
			p.Match(RustLikeParserMUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(135)
		p.Match(RustLikeParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(136)
		p.Match(RustLikeParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)
		p.Rtype()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncBlockContext is an interface to support dynamic dispatch.
type IFuncBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllStat() []IStatContext
	Stat(i int) IStatContext

	// IsFuncBlockContext differentiates from other interfaces.
	IsFuncBlockContext()
}

type FuncBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncBlockContext() *FuncBlockContext {
	var p = new(FuncBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcBlock
	return p
}

func InitEmptyFuncBlockContext(p *FuncBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcBlock
}

func (*FuncBlockContext) IsFuncBlockContext() {}

func NewFuncBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncBlockContext {
	var p = new(FuncBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_funcBlock

	return p
}

func (s *FuncBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncBlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(RustLikeParserLBRACE, 0)
}

func (s *FuncBlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(RustLikeParserRBRACE, 0)
}

func (s *FuncBlockContext) AllStat() []IStatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatContext); ok {
			len++
		}
	}

	tst := make([]IStatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatContext); ok {
			tst[i] = t.(IStatContext)
			i++
		}
	}

	return tst
}

func (s *FuncBlockContext) Stat(i int) IStatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *FuncBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitFuncBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RustLikeParser) FuncBlock() (localctx IFuncBlockContext) {
	localctx = NewFuncBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, RustLikeParserRULE_funcBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.Match(RustLikeParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&45365719904) != 0 {
		{
			p.SetState(140)
			p.Stat()
		}

		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(146)
		p.Match(RustLikeParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRtypeContext is an interface to support dynamic dispatch.
type IRtypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT32() antlr.TerminalNode
	ArrayType() IArrayTypeContext

	// IsRtypeContext differentiates from other interfaces.
	IsRtypeContext()
}

type RtypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRtypeContext() *RtypeContext {
	var p = new(RtypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_rtype
	return p
}

func InitEmptyRtypeContext(p *RtypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_rtype
}

func (*RtypeContext) IsRtypeContext() {}

func NewRtypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RtypeContext {
	var p = new(RtypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_rtype

	return p
}

func (s *RtypeContext) GetParser() antlr.Parser { return s.parser }

func (s *RtypeContext) INT32() antlr.TerminalNode {
	return s.GetToken(RustLikeParserINT32, 0)
}

func (s *RtypeContext) ArrayType() IArrayTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayTypeContext)
}

func (s *RtypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RtypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RtypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitRtype(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RustLikeParser) Rtype() (localctx IRtypeContext) {
	localctx = NewRtypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, RustLikeParserRULE_rtype)
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RustLikeParserINT32:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(148)
			p.Match(RustLikeParserINT32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RustLikeParserLBRAC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(149)
			p.ArrayType()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayTypeContext is an interface to support dynamic dispatch.
type IArrayTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRAC() antlr.TerminalNode
	Rtype() IRtypeContext
	SEMI() antlr.TerminalNode
	ExprNumber() IExprNumberContext
	RBRAC() antlr.TerminalNode

	// IsArrayTypeContext differentiates from other interfaces.
	IsArrayTypeContext()
}

type ArrayTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayTypeContext() *ArrayTypeContext {
	var p = new(ArrayTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_arrayType
	return p
}

func InitEmptyArrayTypeContext(p *ArrayTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_arrayType
}

func (*ArrayTypeContext) IsArrayTypeContext() {}

func NewArrayTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayTypeContext {
	var p = new(ArrayTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_arrayType

	return p
}

func (s *ArrayTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayTypeContext) LBRAC() antlr.TerminalNode {
	return s.GetToken(RustLikeParserLBRAC, 0)
}

func (s *ArrayTypeContext) Rtype() IRtypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRtypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRtypeContext)
}

func (s *ArrayTypeContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RustLikeParserSEMI, 0)
}

func (s *ArrayTypeContext) ExprNumber() IExprNumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprNumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprNumberContext)
}

func (s *ArrayTypeContext) RBRAC() antlr.TerminalNode {
	return s.GetToken(RustLikeParserRBRAC, 0)
}

func (s *ArrayTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitArrayType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RustLikeParser) ArrayType() (localctx IArrayTypeContext) {
	localctx = NewArrayTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, RustLikeParserRULE_arrayType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.Match(RustLikeParserLBRAC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Rtype()
	}
	{
		p.SetState(154)
		p.Match(RustLikeParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(155)
		p.ExprNumber()
	}
	{
		p.SetState(156)
		p.Match(RustLikeParserRBRAC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayElemsContext is an interface to support dynamic dispatch.
type IArrayElemsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArrayElemsContext differentiates from other interfaces.
	IsArrayElemsContext()
}

type ArrayElemsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayElemsContext() *ArrayElemsContext {
	var p = new(ArrayElemsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_arrayElems
	return p
}

func InitEmptyArrayElemsContext(p *ArrayElemsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_arrayElems
}

func (*ArrayElemsContext) IsArrayElemsContext() {}

func NewArrayElemsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayElemsContext {
	var p = new(ArrayElemsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_arrayElems

	return p
}

func (s *ArrayElemsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayElemsContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ArrayElemsContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArrayElemsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(RustLikeParserCOMMA)
}

func (s *ArrayElemsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(RustLikeParserCOMMA, i)
}

func (s *ArrayElemsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayElemsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayElemsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitArrayElems(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RustLikeParser) ArrayElems() (localctx IArrayElemsContext) {
	localctx = NewArrayElemsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, RustLikeParserRULE_arrayElems)
	var _la int

	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RustLikeParserID, RustLikeParserNUMBER, RustLikeParserLPAREN, RustLikeParserLBRAC:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(158)
			p.expr(0)
		}
		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == RustLikeParserCOMMA {
			{
				p.SetState(159)
				p.Match(RustLikeParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(160)
				p.expr(0)
			}

			p.SetState(165)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case RustLikeParserRBRAC:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllStat() []IStatContext
	Stat(i int) IStatContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(RustLikeParserLBRACE, 0)
}

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(RustLikeParserRBRACE, 0)
}

func (s *BlockContext) AllStat() []IStatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatContext); ok {
			len++
		}
	}

	tst := make([]IStatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatContext); ok {
			tst[i] = t.(IStatContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Stat(i int) IStatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RustLikeParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, RustLikeParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(RustLikeParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&45365719904) != 0 {
		{
			p.SetState(170)
			p.Stat()
		}

		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(176)
		p.Match(RustLikeParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_stat
	return p
}

func InitEmptyStatContext(p *StatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_stat
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) CopyAll(ctx *StatContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StatContinueContext struct {
	StatContext
}

func NewStatContinueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatContinueContext {
	var p = new(StatContinueContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *StatContinueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContinueContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(RustLikeParserCONTINUE, 0)
}

func (s *StatContinueContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RustLikeParserSEMI, 0)
}

func (s *StatContinueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitStatContinue(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatBlockContext struct {
	StatContext
}

func NewStatBlockContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatBlockContext {
	var p = new(StatBlockContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *StatBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatBlockContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitStatBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatBreakContext struct {
	StatContext
}

func NewStatBreakContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatBreakContext {
	var p = new(StatBreakContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *StatBreakContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatBreakContext) BREAK() antlr.TerminalNode {
	return s.GetToken(RustLikeParserBREAK, 0)
}

func (s *StatBreakContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RustLikeParserSEMI, 0)
}

func (s *StatBreakContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitStatBreak(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatVarAssignContext struct {
	StatContext
}

func NewStatVarAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatVarAssignContext {
	var p = new(StatVarAssignContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *StatVarAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatVarAssignContext) ID() antlr.TerminalNode {
	return s.GetToken(RustLikeParserID, 0)
}

func (s *StatVarAssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(RustLikeParserASSIGN, 0)
}

func (s *StatVarAssignContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StatVarAssignContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RustLikeParserSEMI, 0)
}

func (s *StatVarAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitStatVarAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatExprContext struct {
	StatContext
}

func NewStatExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatExprContext {
	var p = new(StatExprContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *StatExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StatExprContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RustLikeParserSEMI, 0)
}

func (s *StatExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitStatExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatFuncReturnContext struct {
	StatContext
}

func NewStatFuncReturnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatFuncReturnContext {
	var p = new(StatFuncReturnContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *StatFuncReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatFuncReturnContext) RETURN() antlr.TerminalNode {
	return s.GetToken(RustLikeParserRETURN, 0)
}

func (s *StatFuncReturnContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RustLikeParserSEMI, 0)
}

func (s *StatFuncReturnContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StatFuncReturnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitStatFuncReturn(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatVarDeclareContext struct {
	StatContext
}

func NewStatVarDeclareContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatVarDeclareContext {
	var p = new(StatVarDeclareContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *StatVarDeclareContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatVarDeclareContext) LET() antlr.TerminalNode {
	return s.GetToken(RustLikeParserLET, 0)
}

func (s *StatVarDeclareContext) ID() antlr.TerminalNode {
	return s.GetToken(RustLikeParserID, 0)
}

func (s *StatVarDeclareContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RustLikeParserSEMI, 0)
}

func (s *StatVarDeclareContext) MUT() antlr.TerminalNode {
	return s.GetToken(RustLikeParserMUT, 0)
}

func (s *StatVarDeclareContext) VarType() IVarTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarTypeContext)
}

func (s *StatVarDeclareContext) VarInit() IVarInitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarInitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarInitContext)
}

func (s *StatVarDeclareContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitStatVarDeclare(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatLoopContext struct {
	StatContext
}

func NewStatLoopContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatLoopContext {
	var p = new(StatLoopContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *StatLoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatLoopContext) LOOP() antlr.TerminalNode {
	return s.GetToken(RustLikeParserLOOP, 0)
}

func (s *StatLoopContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatLoopContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitStatLoop(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatIfElseContext struct {
	StatContext
}

func NewStatIfElseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatIfElseContext {
	var p = new(StatIfElseContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *StatIfElseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatIfElseContext) IfBranch() IIfBranchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfBranchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfBranchContext)
}

func (s *StatIfElseContext) AllElifBranch() []IElifBranchContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElifBranchContext); ok {
			len++
		}
	}

	tst := make([]IElifBranchContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElifBranchContext); ok {
			tst[i] = t.(IElifBranchContext)
			i++
		}
	}

	return tst
}

func (s *StatIfElseContext) ElifBranch(i int) IElifBranchContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElifBranchContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElifBranchContext)
}

func (s *StatIfElseContext) ElseBranch() IElseBranchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseBranchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseBranchContext)
}

func (s *StatIfElseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitStatIfElse(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatWhileContext struct {
	StatContext
}

func NewStatWhileContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatWhileContext {
	var p = new(StatWhileContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *StatWhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatWhileContext) WHILE() antlr.TerminalNode {
	return s.GetToken(RustLikeParserWHILE, 0)
}

func (s *StatWhileContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StatWhileContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatWhileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitStatWhile(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatEmptyContext struct {
	StatContext
}

func NewStatEmptyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatEmptyContext {
	var p = new(StatEmptyContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *StatEmptyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatEmptyContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RustLikeParserSEMI, 0)
}

func (s *StatEmptyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitStatEmpty(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RustLikeParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, RustLikeParserRULE_stat)
	var _la int

	var _alt int

	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStatBlockContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(178)
			p.Block()
		}

	case 2:
		localctx = NewStatFuncReturnContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(179)
			p.Match(RustLikeParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&10737516544) != 0 {
			{
				p.SetState(180)
				p.expr(0)
			}

		}
		{
			p.SetState(183)
			p.Match(RustLikeParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewStatBreakContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(184)
			p.Match(RustLikeParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(185)
			p.Match(RustLikeParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewStatContinueContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(186)
			p.Match(RustLikeParserCONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.Match(RustLikeParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewStatVarDeclareContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(188)
			p.Match(RustLikeParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == RustLikeParserMUT {
			{
				p.SetState(189)
				p.Match(RustLikeParserMUT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(192)
			p.Match(RustLikeParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == RustLikeParserCOLON {
			{
				p.SetState(193)
				p.VarType()
			}

		}
		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == RustLikeParserASSIGN {
			{
				p.SetState(196)
				p.VarInit()
			}

		}
		{
			p.SetState(199)
			p.Match(RustLikeParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewStatVarAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(200)
			p.Match(RustLikeParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(201)
			p.Match(RustLikeParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(202)
			p.expr(0)
		}
		{
			p.SetState(203)
			p.Match(RustLikeParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewStatExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(205)
			p.expr(0)
		}
		{
			p.SetState(206)
			p.Match(RustLikeParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewStatIfElseContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(208)
			p.IfBranch()
		}
		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(209)
					p.ElifBranch()
				}

			}
			p.SetState(214)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == RustLikeParserELSE {
			{
				p.SetState(215)
				p.ElseBranch()
			}

		}

	case 9:
		localctx = NewStatWhileContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(218)
			p.Match(RustLikeParserWHILE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(219)
			p.expr(0)
		}
		{
			p.SetState(220)
			p.Block()
		}

	case 10:
		localctx = NewStatLoopContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(222)
			p.Match(RustLikeParserLOOP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(223)
			p.Block()
		}

	case 11:
		localctx = NewStatEmptyContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(224)
			p.Match(RustLikeParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfBranchContext is an interface to support dynamic dispatch.
type IIfBranchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	Expr() IExprContext
	Block() IBlockContext

	// IsIfBranchContext differentiates from other interfaces.
	IsIfBranchContext()
}

type IfBranchContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfBranchContext() *IfBranchContext {
	var p = new(IfBranchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_ifBranch
	return p
}

func InitEmptyIfBranchContext(p *IfBranchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_ifBranch
}

func (*IfBranchContext) IsIfBranchContext() {}

func NewIfBranchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfBranchContext {
	var p = new(IfBranchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_ifBranch

	return p
}

func (s *IfBranchContext) GetParser() antlr.Parser { return s.parser }

func (s *IfBranchContext) IF() antlr.TerminalNode {
	return s.GetToken(RustLikeParserIF, 0)
}

func (s *IfBranchContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfBranchContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfBranchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfBranchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfBranchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitIfBranch(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RustLikeParser) IfBranch() (localctx IIfBranchContext) {
	localctx = NewIfBranchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, RustLikeParserRULE_ifBranch)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		p.Match(RustLikeParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(228)
		p.expr(0)
	}
	{
		p.SetState(229)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElifBranchContext is an interface to support dynamic dispatch.
type IElifBranchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ELSE() antlr.TerminalNode
	IF() antlr.TerminalNode
	Expr() IExprContext
	Block() IBlockContext

	// IsElifBranchContext differentiates from other interfaces.
	IsElifBranchContext()
}

type ElifBranchContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElifBranchContext() *ElifBranchContext {
	var p = new(ElifBranchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_elifBranch
	return p
}

func InitEmptyElifBranchContext(p *ElifBranchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_elifBranch
}

func (*ElifBranchContext) IsElifBranchContext() {}

func NewElifBranchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElifBranchContext {
	var p = new(ElifBranchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_elifBranch

	return p
}

func (s *ElifBranchContext) GetParser() antlr.Parser { return s.parser }

func (s *ElifBranchContext) ELSE() antlr.TerminalNode {
	return s.GetToken(RustLikeParserELSE, 0)
}

func (s *ElifBranchContext) IF() antlr.TerminalNode {
	return s.GetToken(RustLikeParserIF, 0)
}

func (s *ElifBranchContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ElifBranchContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElifBranchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElifBranchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElifBranchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitElifBranch(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RustLikeParser) ElifBranch() (localctx IElifBranchContext) {
	localctx = NewElifBranchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, RustLikeParserRULE_elifBranch)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(231)
		p.Match(RustLikeParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(232)
		p.Match(RustLikeParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(233)
		p.expr(0)
	}
	{
		p.SetState(234)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElseBranchContext is an interface to support dynamic dispatch.
type IElseBranchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ELSE() antlr.TerminalNode
	Block() IBlockContext

	// IsElseBranchContext differentiates from other interfaces.
	IsElseBranchContext()
}

type ElseBranchContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseBranchContext() *ElseBranchContext {
	var p = new(ElseBranchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_elseBranch
	return p
}

func InitEmptyElseBranchContext(p *ElseBranchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_elseBranch
}

func (*ElseBranchContext) IsElseBranchContext() {}

func NewElseBranchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseBranchContext {
	var p = new(ElseBranchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_elseBranch

	return p
}

func (s *ElseBranchContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseBranchContext) ELSE() antlr.TerminalNode {
	return s.GetToken(RustLikeParserELSE, 0)
}

func (s *ElseBranchContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElseBranchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseBranchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseBranchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitElseBranch(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RustLikeParser) ElseBranch() (localctx IElseBranchContext) {
	localctx = NewElseBranchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, RustLikeParserRULE_elseBranch)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		p.Match(RustLikeParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(237)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarTypeContext is an interface to support dynamic dispatch.
type IVarTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COLON() antlr.TerminalNode
	Rtype() IRtypeContext

	// IsVarTypeContext differentiates from other interfaces.
	IsVarTypeContext()
}

type VarTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarTypeContext() *VarTypeContext {
	var p = new(VarTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_varType
	return p
}

func InitEmptyVarTypeContext(p *VarTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_varType
}

func (*VarTypeContext) IsVarTypeContext() {}

func NewVarTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarTypeContext {
	var p = new(VarTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_varType

	return p
}

func (s *VarTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *VarTypeContext) COLON() antlr.TerminalNode {
	return s.GetToken(RustLikeParserCOLON, 0)
}

func (s *VarTypeContext) Rtype() IRtypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRtypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRtypeContext)
}

func (s *VarTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitVarType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RustLikeParser) VarType() (localctx IVarTypeContext) {
	localctx = NewVarTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, RustLikeParserRULE_varType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)
		p.Match(RustLikeParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(240)
		p.Rtype()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarInitContext is an interface to support dynamic dispatch.
type IVarInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ASSIGN() antlr.TerminalNode
	Expr() IExprContext

	// IsVarInitContext differentiates from other interfaces.
	IsVarInitContext()
}

type VarInitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarInitContext() *VarInitContext {
	var p = new(VarInitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_varInit
	return p
}

func InitEmptyVarInitContext(p *VarInitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_varInit
}

func (*VarInitContext) IsVarInitContext() {}

func NewVarInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarInitContext {
	var p = new(VarInitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_varInit

	return p
}

func (s *VarInitContext) GetParser() antlr.Parser { return s.parser }

func (s *VarInitContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(RustLikeParserASSIGN, 0)
}

func (s *VarInitContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VarInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarInitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RustLikeParserVisitor:
		return t.VisitVarInit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RustLikeParser) VarInit() (localctx IVarInitContext) {
	localctx = NewVarInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, RustLikeParserRULE_varInit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Match(RustLikeParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(243)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *RustLikeParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *RustLikeParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
