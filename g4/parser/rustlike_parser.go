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
		"prog", "declaration", "expr", "func_call", "func_call_list", "func_call_param",
		"func_declaration", "func_declaration_header", "func_declaration_return",
		"fps_list", "fps", "fp", "type", "block", "stat", "stat_return", "var_declare",
		"var_type", "var_init", "var_assign", "stat_if_else", "stat_while",
		"stat_loop",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 39, 211, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 1, 0, 1, 0, 1, 1, 5, 1, 50, 8, 1, 10, 1, 12, 1,
		53, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 63, 8,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 74, 8, 2,
		10, 2, 12, 2, 77, 9, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 5, 5, 5, 89, 8, 5, 10, 5, 12, 5, 92, 9, 5, 1, 5, 3, 5, 95, 8,
		5, 1, 6, 1, 6, 3, 6, 99, 8, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 5, 10, 117, 8,
		10, 10, 10, 12, 10, 120, 9, 10, 1, 10, 3, 10, 123, 8, 10, 1, 11, 3, 11,
		126, 8, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 5,
		13, 136, 8, 13, 10, 13, 12, 13, 139, 9, 13, 1, 13, 1, 13, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 160, 8, 14, 1, 15, 1, 15, 3, 15,
		164, 8, 15, 1, 16, 1, 16, 3, 16, 168, 8, 16, 1, 16, 1, 16, 3, 16, 172,
		8, 16, 1, 16, 3, 16, 175, 8, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1,
		18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 5, 20, 195, 8, 20, 10, 20, 12, 20, 198, 9, 20, 1, 20, 1,
		20, 3, 20, 202, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22,
		1, 22, 0, 1, 4, 23, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
		28, 30, 32, 34, 36, 38, 40, 42, 44, 0, 3, 1, 0, 19, 20, 1, 0, 17, 18, 1,
		0, 21, 26, 215, 0, 46, 1, 0, 0, 0, 2, 51, 1, 0, 0, 0, 4, 62, 1, 0, 0, 0,
		6, 78, 1, 0, 0, 0, 8, 81, 1, 0, 0, 0, 10, 94, 1, 0, 0, 0, 12, 96, 1, 0,
		0, 0, 14, 102, 1, 0, 0, 0, 16, 106, 1, 0, 0, 0, 18, 109, 1, 0, 0, 0, 20,
		122, 1, 0, 0, 0, 22, 125, 1, 0, 0, 0, 24, 131, 1, 0, 0, 0, 26, 133, 1,
		0, 0, 0, 28, 159, 1, 0, 0, 0, 30, 161, 1, 0, 0, 0, 32, 165, 1, 0, 0, 0,
		34, 176, 1, 0, 0, 0, 36, 179, 1, 0, 0, 0, 38, 182, 1, 0, 0, 0, 40, 186,
		1, 0, 0, 0, 42, 203, 1, 0, 0, 0, 44, 207, 1, 0, 0, 0, 46, 47, 3, 2, 1,
		0, 47, 1, 1, 0, 0, 0, 48, 50, 3, 12, 6, 0, 49, 48, 1, 0, 0, 0, 50, 53,
		1, 0, 0, 0, 51, 49, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 3, 1, 0, 0, 0,
		53, 51, 1, 0, 0, 0, 54, 55, 6, 2, -1, 0, 55, 56, 5, 31, 0, 0, 56, 57, 3,
		4, 2, 0, 57, 58, 5, 32, 0, 0, 58, 63, 1, 0, 0, 0, 59, 63, 3, 6, 3, 0, 60,
		63, 5, 15, 0, 0, 61, 63, 5, 16, 0, 0, 62, 54, 1, 0, 0, 0, 62, 59, 1, 0,
		0, 0, 62, 60, 1, 0, 0, 0, 62, 61, 1, 0, 0, 0, 63, 75, 1, 0, 0, 0, 64, 65,
		10, 7, 0, 0, 65, 66, 7, 0, 0, 0, 66, 74, 3, 4, 2, 8, 67, 68, 10, 6, 0,
		0, 68, 69, 7, 1, 0, 0, 69, 74, 3, 4, 2, 7, 70, 71, 10, 5, 0, 0, 71, 72,
		7, 2, 0, 0, 72, 74, 3, 4, 2, 6, 73, 64, 1, 0, 0, 0, 73, 67, 1, 0, 0, 0,
		73, 70, 1, 0, 0, 0, 74, 77, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 75, 76, 1,
		0, 0, 0, 76, 5, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 78, 79, 5, 15, 0, 0, 79,
		80, 3, 8, 4, 0, 80, 7, 1, 0, 0, 0, 81, 82, 5, 31, 0, 0, 82, 83, 3, 10,
		5, 0, 83, 84, 5, 32, 0, 0, 84, 9, 1, 0, 0, 0, 85, 90, 3, 4, 2, 0, 86, 87,
		5, 30, 0, 0, 87, 89, 3, 4, 2, 0, 88, 86, 1, 0, 0, 0, 89, 92, 1, 0, 0, 0,
		90, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 95, 1, 0, 0, 0, 92, 90, 1,
		0, 0, 0, 93, 95, 1, 0, 0, 0, 94, 85, 1, 0, 0, 0, 94, 93, 1, 0, 0, 0, 95,
		11, 1, 0, 0, 0, 96, 98, 3, 14, 7, 0, 97, 99, 3, 16, 8, 0, 98, 97, 1, 0,
		0, 0, 98, 99, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 101, 3, 26, 13, 0,
		101, 13, 1, 0, 0, 0, 102, 103, 5, 11, 0, 0, 103, 104, 5, 15, 0, 0, 104,
		105, 3, 18, 9, 0, 105, 15, 1, 0, 0, 0, 106, 107, 5, 37, 0, 0, 107, 108,
		3, 24, 12, 0, 108, 17, 1, 0, 0, 0, 109, 110, 5, 31, 0, 0, 110, 111, 3,
		20, 10, 0, 111, 112, 5, 32, 0, 0, 112, 19, 1, 0, 0, 0, 113, 118, 3, 22,
		11, 0, 114, 115, 5, 30, 0, 0, 115, 117, 3, 22, 11, 0, 116, 114, 1, 0, 0,
		0, 117, 120, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119,
		123, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 121, 123, 1, 0, 0, 0, 122, 113,
		1, 0, 0, 0, 122, 121, 1, 0, 0, 0, 123, 21, 1, 0, 0, 0, 124, 126, 5, 10,
		0, 0, 125, 124, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0,
		127, 128, 5, 15, 0, 0, 128, 129, 5, 29, 0, 0, 129, 130, 3, 24, 12, 0, 130,
		23, 1, 0, 0, 0, 131, 132, 5, 4, 0, 0, 132, 25, 1, 0, 0, 0, 133, 137, 5,
		35, 0, 0, 134, 136, 3, 28, 14, 0, 135, 134, 1, 0, 0, 0, 136, 139, 1, 0,
		0, 0, 137, 135, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 140, 1, 0, 0, 0,
		139, 137, 1, 0, 0, 0, 140, 141, 5, 36, 0, 0, 141, 27, 1, 0, 0, 0, 142,
		160, 3, 26, 13, 0, 143, 144, 3, 30, 15, 0, 144, 145, 5, 28, 0, 0, 145,
		160, 1, 0, 0, 0, 146, 147, 3, 32, 16, 0, 147, 148, 5, 28, 0, 0, 148, 160,
		1, 0, 0, 0, 149, 150, 3, 38, 19, 0, 150, 151, 5, 28, 0, 0, 151, 160, 1,
		0, 0, 0, 152, 153, 3, 4, 2, 0, 153, 154, 5, 28, 0, 0, 154, 160, 1, 0, 0,
		0, 155, 160, 3, 40, 20, 0, 156, 160, 3, 42, 21, 0, 157, 160, 3, 44, 22,
		0, 158, 160, 5, 28, 0, 0, 159, 142, 1, 0, 0, 0, 159, 143, 1, 0, 0, 0, 159,
		146, 1, 0, 0, 0, 159, 149, 1, 0, 0, 0, 159, 152, 1, 0, 0, 0, 159, 155,
		1, 0, 0, 0, 159, 156, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 159, 158, 1, 0,
		0, 0, 160, 29, 1, 0, 0, 0, 161, 163, 5, 9, 0, 0, 162, 164, 3, 4, 2, 0,
		163, 162, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 31, 1, 0, 0, 0, 165, 167,
		5, 5, 0, 0, 166, 168, 5, 10, 0, 0, 167, 166, 1, 0, 0, 0, 167, 168, 1, 0,
		0, 0, 168, 169, 1, 0, 0, 0, 169, 171, 5, 15, 0, 0, 170, 172, 3, 34, 17,
		0, 171, 170, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 174, 1, 0, 0, 0, 173,
		175, 3, 36, 18, 0, 174, 173, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 33,
		1, 0, 0, 0, 176, 177, 5, 29, 0, 0, 177, 178, 3, 24, 12, 0, 178, 35, 1,
		0, 0, 0, 179, 180, 5, 27, 0, 0, 180, 181, 3, 4, 2, 0, 181, 37, 1, 0, 0,
		0, 182, 183, 5, 15, 0, 0, 183, 184, 5, 27, 0, 0, 184, 185, 3, 4, 2, 0,
		185, 39, 1, 0, 0, 0, 186, 187, 5, 6, 0, 0, 187, 188, 3, 4, 2, 0, 188, 196,
		3, 26, 13, 0, 189, 190, 5, 7, 0, 0, 190, 191, 5, 6, 0, 0, 191, 192, 3,
		4, 2, 0, 192, 193, 3, 26, 13, 0, 193, 195, 1, 0, 0, 0, 194, 189, 1, 0,
		0, 0, 195, 198, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0,
		197, 201, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0, 199, 200, 5, 7, 0, 0, 200,
		202, 3, 26, 13, 0, 201, 199, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 41,
		1, 0, 0, 0, 203, 204, 5, 8, 0, 0, 204, 205, 3, 4, 2, 0, 205, 206, 3, 26,
		13, 0, 206, 43, 1, 0, 0, 0, 207, 208, 5, 12, 0, 0, 208, 209, 3, 26, 13,
		0, 209, 45, 1, 0, 0, 0, 18, 51, 62, 73, 75, 90, 94, 98, 118, 122, 125,
		137, 159, 163, 167, 171, 174, 196, 201,
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
	RustLikeParserRULE_prog                    = 0
	RustLikeParserRULE_declaration             = 1
	RustLikeParserRULE_expr                    = 2
	RustLikeParserRULE_func_call               = 3
	RustLikeParserRULE_func_call_list          = 4
	RustLikeParserRULE_func_call_param         = 5
	RustLikeParserRULE_func_declaration        = 6
	RustLikeParserRULE_func_declaration_header = 7
	RustLikeParserRULE_func_declaration_return = 8
	RustLikeParserRULE_fps_list                = 9
	RustLikeParserRULE_fps                     = 10
	RustLikeParserRULE_fp                      = 11
	RustLikeParserRULE_type                    = 12
	RustLikeParserRULE_block                   = 13
	RustLikeParserRULE_stat                    = 14
	RustLikeParserRULE_stat_return             = 15
	RustLikeParserRULE_var_declare             = 16
	RustLikeParserRULE_var_type                = 17
	RustLikeParserRULE_var_init                = 18
	RustLikeParserRULE_var_assign              = 19
	RustLikeParserRULE_stat_if_else            = 20
	RustLikeParserRULE_stat_while              = 21
	RustLikeParserRULE_stat_loop               = 22
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

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitProg(s)
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
	AllFunc_declaration() []IFunc_declarationContext
	Func_declaration(i int) IFunc_declarationContext

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

func (s *DeclarationContext) AllFunc_declaration() []IFunc_declarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunc_declarationContext); ok {
			len++
		}
	}

	tst := make([]IFunc_declarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunc_declarationContext); ok {
			tst[i] = t.(IFunc_declarationContext)
			i++
		}
	}

	return tst
}

func (s *DeclarationContext) Func_declaration(i int) IFunc_declarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_declarationContext); ok {
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

	return t.(IFunc_declarationContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitDeclaration(s)
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
			p.Func_declaration()
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

	// Getter signatures
	LPAREN() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	RPAREN() antlr.TerminalNode
	Func_call() IFunc_callContext
	ID() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	MULT() antlr.TerminalNode
	DIV() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	MINUS() antlr.TerminalNode
	EQ() antlr.TerminalNode
	NE() antlr.TerminalNode
	LT() antlr.TerminalNode
	GT() antlr.TerminalNode
	LE() antlr.TerminalNode
	GE() antlr.TerminalNode

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

func (s *ExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(RustLikeParserLPAREN, 0)
}

func (s *ExprContext) AllExpr() []IExprContext {
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

func (s *ExprContext) Expr(i int) IExprContext {
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

func (s *ExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(RustLikeParserRPAREN, 0)
}

func (s *ExprContext) Func_call() IFunc_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_callContext)
}

func (s *ExprContext) ID() antlr.TerminalNode {
	return s.GetToken(RustLikeParserID, 0)
}

func (s *ExprContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(RustLikeParserNUMBER, 0)
}

func (s *ExprContext) MULT() antlr.TerminalNode {
	return s.GetToken(RustLikeParserMULT, 0)
}

func (s *ExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(RustLikeParserDIV, 0)
}

func (s *ExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(RustLikeParserPLUS, 0)
}

func (s *ExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(RustLikeParserMINUS, 0)
}

func (s *ExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(RustLikeParserEQ, 0)
}

func (s *ExprContext) NE() antlr.TerminalNode {
	return s.GetToken(RustLikeParserNE, 0)
}

func (s *ExprContext) LT() antlr.TerminalNode {
	return s.GetToken(RustLikeParserLT, 0)
}

func (s *ExprContext) GT() antlr.TerminalNode {
	return s.GetToken(RustLikeParserGT, 0)
}

func (s *ExprContext) LE() antlr.TerminalNode {
	return s.GetToken(RustLikeParserLE, 0)
}

func (s *ExprContext) GE() antlr.TerminalNode {
	return s.GetToken(RustLikeParserGE, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitExpr(s)
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
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(55)
			p.Match(RustLikeParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(56)
			p.expr(0)
		}
		{
			p.SetState(57)
			p.Match(RustLikeParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(59)
			p.Func_call()
		}

	case 3:
		{
			p.SetState(60)
			p.Match(RustLikeParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		{
			p.SetState(61)
			p.Match(RustLikeParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(75)
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
			p.SetState(73)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RustLikeParserRULE_expr)
				p.SetState(64)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(65)
					_la = p.GetTokenStream().LA(1)

					if !(_la == RustLikeParserMULT || _la == RustLikeParserDIV) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(66)
					p.expr(8)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RustLikeParserRULE_expr)
				p.SetState(67)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(68)
					_la = p.GetTokenStream().LA(1)

					if !(_la == RustLikeParserPLUS || _la == RustLikeParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(69)
					p.expr(7)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RustLikeParserRULE_expr)
				p.SetState(70)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(71)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&132120576) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(72)
					p.expr(6)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(77)
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

// IFunc_callContext is an interface to support dynamic dispatch.
type IFunc_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Func_call_list() IFunc_call_listContext

	// IsFunc_callContext differentiates from other interfaces.
	IsFunc_callContext()
}

type Func_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_callContext() *Func_callContext {
	var p = new(Func_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_func_call
	return p
}

func InitEmptyFunc_callContext(p *Func_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_func_call
}

func (*Func_callContext) IsFunc_callContext() {}

func NewFunc_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_callContext {
	var p = new(Func_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_func_call

	return p
}

func (s *Func_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_callContext) ID() antlr.TerminalNode {
	return s.GetToken(RustLikeParserID, 0)
}

func (s *Func_callContext) Func_call_list() IFunc_call_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_call_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_call_listContext)
}

func (s *Func_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterFunc_call(s)
	}
}

func (s *Func_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitFunc_call(s)
	}
}

func (p *RustLikeParser) Func_call() (localctx IFunc_callContext) {
	localctx = NewFunc_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RustLikeParserRULE_func_call)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(RustLikeParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(79)
		p.Func_call_list()
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

// IFunc_call_listContext is an interface to support dynamic dispatch.
type IFunc_call_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	Func_call_param() IFunc_call_paramContext
	RPAREN() antlr.TerminalNode

	// IsFunc_call_listContext differentiates from other interfaces.
	IsFunc_call_listContext()
}

type Func_call_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_call_listContext() *Func_call_listContext {
	var p = new(Func_call_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_func_call_list
	return p
}

func InitEmptyFunc_call_listContext(p *Func_call_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_func_call_list
}

func (*Func_call_listContext) IsFunc_call_listContext() {}

func NewFunc_call_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_call_listContext {
	var p = new(Func_call_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_func_call_list

	return p
}

func (s *Func_call_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_call_listContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(RustLikeParserLPAREN, 0)
}

func (s *Func_call_listContext) Func_call_param() IFunc_call_paramContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_call_paramContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_call_paramContext)
}

func (s *Func_call_listContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(RustLikeParserRPAREN, 0)
}

func (s *Func_call_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_call_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_call_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterFunc_call_list(s)
	}
}

func (s *Func_call_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitFunc_call_list(s)
	}
}

func (p *RustLikeParser) Func_call_list() (localctx IFunc_call_listContext) {
	localctx = NewFunc_call_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RustLikeParserRULE_func_call_list)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.Match(RustLikeParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)
		p.Func_call_param()
	}
	{
		p.SetState(83)
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

// IFunc_call_paramContext is an interface to support dynamic dispatch.
type IFunc_call_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFunc_call_paramContext differentiates from other interfaces.
	IsFunc_call_paramContext()
}

type Func_call_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_call_paramContext() *Func_call_paramContext {
	var p = new(Func_call_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_func_call_param
	return p
}

func InitEmptyFunc_call_paramContext(p *Func_call_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_func_call_param
}

func (*Func_call_paramContext) IsFunc_call_paramContext() {}

func NewFunc_call_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_call_paramContext {
	var p = new(Func_call_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_func_call_param

	return p
}

func (s *Func_call_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_call_paramContext) AllExpr() []IExprContext {
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

func (s *Func_call_paramContext) Expr(i int) IExprContext {
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

func (s *Func_call_paramContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(RustLikeParserCOMMA)
}

func (s *Func_call_paramContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(RustLikeParserCOMMA, i)
}

func (s *Func_call_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_call_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_call_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterFunc_call_param(s)
	}
}

func (s *Func_call_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitFunc_call_param(s)
	}
}

func (p *RustLikeParser) Func_call_param() (localctx IFunc_call_paramContext) {
	localctx = NewFunc_call_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RustLikeParserRULE_func_call_param)
	var _la int

	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RustLikeParserID, RustLikeParserNUMBER, RustLikeParserLPAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(85)
			p.expr(0)
		}
		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == RustLikeParserCOMMA {
			{
				p.SetState(86)
				p.Match(RustLikeParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(87)
				p.expr(0)
			}

			p.SetState(92)
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

// IFunc_declarationContext is an interface to support dynamic dispatch.
type IFunc_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Func_declaration_header() IFunc_declaration_headerContext
	Block() IBlockContext
	Func_declaration_return() IFunc_declaration_returnContext

	// IsFunc_declarationContext differentiates from other interfaces.
	IsFunc_declarationContext()
}

type Func_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_declarationContext() *Func_declarationContext {
	var p = new(Func_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_func_declaration
	return p
}

func InitEmptyFunc_declarationContext(p *Func_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_func_declaration
}

func (*Func_declarationContext) IsFunc_declarationContext() {}

func NewFunc_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_declarationContext {
	var p = new(Func_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_func_declaration

	return p
}

func (s *Func_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_declarationContext) Func_declaration_header() IFunc_declaration_headerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_declaration_headerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_declaration_headerContext)
}

func (s *Func_declarationContext) Block() IBlockContext {
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

func (s *Func_declarationContext) Func_declaration_return() IFunc_declaration_returnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_declaration_returnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_declaration_returnContext)
}

func (s *Func_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterFunc_declaration(s)
	}
}

func (s *Func_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitFunc_declaration(s)
	}
}

func (p *RustLikeParser) Func_declaration() (localctx IFunc_declarationContext) {
	localctx = NewFunc_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RustLikeParserRULE_func_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Func_declaration_header()
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == RustLikeParserARROW {
		{
			p.SetState(97)
			p.Func_declaration_return()
		}

	}
	{
		p.SetState(100)
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

// IFunc_declaration_headerContext is an interface to support dynamic dispatch.
type IFunc_declaration_headerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FN() antlr.TerminalNode
	ID() antlr.TerminalNode
	Fps_list() IFps_listContext

	// IsFunc_declaration_headerContext differentiates from other interfaces.
	IsFunc_declaration_headerContext()
}

type Func_declaration_headerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_declaration_headerContext() *Func_declaration_headerContext {
	var p = new(Func_declaration_headerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_func_declaration_header
	return p
}

func InitEmptyFunc_declaration_headerContext(p *Func_declaration_headerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_func_declaration_header
}

func (*Func_declaration_headerContext) IsFunc_declaration_headerContext() {}

func NewFunc_declaration_headerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_declaration_headerContext {
	var p = new(Func_declaration_headerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_func_declaration_header

	return p
}

func (s *Func_declaration_headerContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_declaration_headerContext) FN() antlr.TerminalNode {
	return s.GetToken(RustLikeParserFN, 0)
}

func (s *Func_declaration_headerContext) ID() antlr.TerminalNode {
	return s.GetToken(RustLikeParserID, 0)
}

func (s *Func_declaration_headerContext) Fps_list() IFps_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFps_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFps_listContext)
}

func (s *Func_declaration_headerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_declaration_headerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_declaration_headerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterFunc_declaration_header(s)
	}
}

func (s *Func_declaration_headerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitFunc_declaration_header(s)
	}
}

func (p *RustLikeParser) Func_declaration_header() (localctx IFunc_declaration_headerContext) {
	localctx = NewFunc_declaration_headerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, RustLikeParserRULE_func_declaration_header)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		p.Match(RustLikeParserFN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(103)
		p.Match(RustLikeParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(104)
		p.Fps_list()
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

// IFunc_declaration_returnContext is an interface to support dynamic dispatch.
type IFunc_declaration_returnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ARROW() antlr.TerminalNode
	Type_() ITypeContext

	// IsFunc_declaration_returnContext differentiates from other interfaces.
	IsFunc_declaration_returnContext()
}

type Func_declaration_returnContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_declaration_returnContext() *Func_declaration_returnContext {
	var p = new(Func_declaration_returnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_func_declaration_return
	return p
}

func InitEmptyFunc_declaration_returnContext(p *Func_declaration_returnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_func_declaration_return
}

func (*Func_declaration_returnContext) IsFunc_declaration_returnContext() {}

func NewFunc_declaration_returnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_declaration_returnContext {
	var p = new(Func_declaration_returnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_func_declaration_return

	return p
}

func (s *Func_declaration_returnContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_declaration_returnContext) ARROW() antlr.TerminalNode {
	return s.GetToken(RustLikeParserARROW, 0)
}

func (s *Func_declaration_returnContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *Func_declaration_returnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_declaration_returnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_declaration_returnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterFunc_declaration_return(s)
	}
}

func (s *Func_declaration_returnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitFunc_declaration_return(s)
	}
}

func (p *RustLikeParser) Func_declaration_return() (localctx IFunc_declaration_returnContext) {
	localctx = NewFunc_declaration_returnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RustLikeParserRULE_func_declaration_return)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.Match(RustLikeParserARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.Type_()
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

// IFps_listContext is an interface to support dynamic dispatch.
type IFps_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	Fps() IFpsContext
	RPAREN() antlr.TerminalNode

	// IsFps_listContext differentiates from other interfaces.
	IsFps_listContext()
}

type Fps_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFps_listContext() *Fps_listContext {
	var p = new(Fps_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_fps_list
	return p
}

func InitEmptyFps_listContext(p *Fps_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_fps_list
}

func (*Fps_listContext) IsFps_listContext() {}

func NewFps_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fps_listContext {
	var p = new(Fps_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_fps_list

	return p
}

func (s *Fps_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Fps_listContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(RustLikeParserLPAREN, 0)
}

func (s *Fps_listContext) Fps() IFpsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFpsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFpsContext)
}

func (s *Fps_listContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(RustLikeParserRPAREN, 0)
}

func (s *Fps_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fps_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fps_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterFps_list(s)
	}
}

func (s *Fps_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitFps_list(s)
	}
}

func (p *RustLikeParser) Fps_list() (localctx IFps_listContext) {
	localctx = NewFps_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, RustLikeParserRULE_fps_list)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Match(RustLikeParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(110)
		p.Fps()
	}
	{
		p.SetState(111)
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

// IFpsContext is an interface to support dynamic dispatch.
type IFpsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFp() []IFpContext
	Fp(i int) IFpContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFpsContext differentiates from other interfaces.
	IsFpsContext()
}

type FpsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFpsContext() *FpsContext {
	var p = new(FpsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_fps
	return p
}

func InitEmptyFpsContext(p *FpsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_fps
}

func (*FpsContext) IsFpsContext() {}

func NewFpsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FpsContext {
	var p = new(FpsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_fps

	return p
}

func (s *FpsContext) GetParser() antlr.Parser { return s.parser }

func (s *FpsContext) AllFp() []IFpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFpContext); ok {
			len++
		}
	}

	tst := make([]IFpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFpContext); ok {
			tst[i] = t.(IFpContext)
			i++
		}
	}

	return tst
}

func (s *FpsContext) Fp(i int) IFpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFpContext); ok {
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

	return t.(IFpContext)
}

func (s *FpsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(RustLikeParserCOMMA)
}

func (s *FpsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(RustLikeParserCOMMA, i)
}

func (s *FpsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FpsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FpsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterFps(s)
	}
}

func (s *FpsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitFps(s)
	}
}

func (p *RustLikeParser) Fps() (localctx IFpsContext) {
	localctx = NewFpsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, RustLikeParserRULE_fps)
	var _la int

	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RustLikeParserMUT, RustLikeParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(113)
			p.Fp()
		}
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == RustLikeParserCOMMA {
			{
				p.SetState(114)
				p.Match(RustLikeParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(115)
				p.Fp()
			}

			p.SetState(120)
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

// IFpContext is an interface to support dynamic dispatch.
type IFpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext
	MUT() antlr.TerminalNode

	// IsFpContext differentiates from other interfaces.
	IsFpContext()
}

type FpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFpContext() *FpContext {
	var p = new(FpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_fp
	return p
}

func InitEmptyFpContext(p *FpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_fp
}

func (*FpContext) IsFpContext() {}

func NewFpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FpContext {
	var p = new(FpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_fp

	return p
}

func (s *FpContext) GetParser() antlr.Parser { return s.parser }

func (s *FpContext) ID() antlr.TerminalNode {
	return s.GetToken(RustLikeParserID, 0)
}

func (s *FpContext) COLON() antlr.TerminalNode {
	return s.GetToken(RustLikeParserCOLON, 0)
}

func (s *FpContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FpContext) MUT() antlr.TerminalNode {
	return s.GetToken(RustLikeParserMUT, 0)
}

func (s *FpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterFp(s)
	}
}

func (s *FpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitFp(s)
	}
}

func (p *RustLikeParser) Fp() (localctx IFpContext) {
	localctx = NewFpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, RustLikeParserRULE_fp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == RustLikeParserMUT {
		{
			p.SetState(124)
			p.Match(RustLikeParserMUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(127)
		p.Match(RustLikeParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(128)
		p.Match(RustLikeParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
		p.Type_()
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

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT32() antlr.TerminalNode

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) INT32() antlr.TerminalNode {
	return s.GetToken(RustLikeParserINT32, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *RustLikeParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, RustLikeParserRULE_type)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Match(RustLikeParserINT32)
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

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *RustLikeParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, RustLikeParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.Match(RustLikeParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36775760736) != 0 {
		{
			p.SetState(134)
			p.Stat()
		}

		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(140)
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

	// Getter signatures
	Block() IBlockContext
	Stat_return() IStat_returnContext
	SEMI() antlr.TerminalNode
	Var_declare() IVar_declareContext
	Var_assign() IVar_assignContext
	Expr() IExprContext
	Stat_if_else() IStat_if_elseContext
	Stat_while() IStat_whileContext
	Stat_loop() IStat_loopContext

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

func (s *StatContext) Block() IBlockContext {
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

func (s *StatContext) Stat_return() IStat_returnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStat_returnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStat_returnContext)
}

func (s *StatContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RustLikeParserSEMI, 0)
}

func (s *StatContext) Var_declare() IVar_declareContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_declareContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_declareContext)
}

func (s *StatContext) Var_assign() IVar_assignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_assignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_assignContext)
}

func (s *StatContext) Expr() IExprContext {
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

func (s *StatContext) Stat_if_else() IStat_if_elseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStat_if_elseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStat_if_elseContext)
}

func (s *StatContext) Stat_while() IStat_whileContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStat_whileContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStat_whileContext)
}

func (s *StatContext) Stat_loop() IStat_loopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStat_loopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStat_loopContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterStat(s)
	}
}

func (s *StatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitStat(s)
	}
}

func (p *RustLikeParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, RustLikeParserRULE_stat)
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(142)
			p.Block()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(143)
			p.Stat_return()
		}
		{
			p.SetState(144)
			p.Match(RustLikeParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(146)
			p.Var_declare()
		}
		{
			p.SetState(147)
			p.Match(RustLikeParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(149)
			p.Var_assign()
		}
		{
			p.SetState(150)
			p.Match(RustLikeParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(152)
			p.expr(0)
		}
		{
			p.SetState(153)
			p.Match(RustLikeParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(155)
			p.Stat_if_else()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(156)
			p.Stat_while()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(157)
			p.Stat_loop()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(158)
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

// IStat_returnContext is an interface to support dynamic dispatch.
type IStat_returnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expr() IExprContext

	// IsStat_returnContext differentiates from other interfaces.
	IsStat_returnContext()
}

type Stat_returnContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStat_returnContext() *Stat_returnContext {
	var p = new(Stat_returnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_stat_return
	return p
}

func InitEmptyStat_returnContext(p *Stat_returnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_stat_return
}

func (*Stat_returnContext) IsStat_returnContext() {}

func NewStat_returnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Stat_returnContext {
	var p = new(Stat_returnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_stat_return

	return p
}

func (s *Stat_returnContext) GetParser() antlr.Parser { return s.parser }

func (s *Stat_returnContext) RETURN() antlr.TerminalNode {
	return s.GetToken(RustLikeParserRETURN, 0)
}

func (s *Stat_returnContext) Expr() IExprContext {
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

func (s *Stat_returnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stat_returnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Stat_returnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterStat_return(s)
	}
}

func (s *Stat_returnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitStat_return(s)
	}
}

func (p *RustLikeParser) Stat_return() (localctx IStat_returnContext) {
	localctx = NewStat_returnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, RustLikeParserRULE_stat_return)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.Match(RustLikeParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2147581952) != 0 {
		{
			p.SetState(162)
			p.expr(0)
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

// IVar_declareContext is an interface to support dynamic dispatch.
type IVar_declareContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LET() antlr.TerminalNode
	ID() antlr.TerminalNode
	MUT() antlr.TerminalNode
	Var_type() IVar_typeContext
	Var_init() IVar_initContext

	// IsVar_declareContext differentiates from other interfaces.
	IsVar_declareContext()
}

type Var_declareContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_declareContext() *Var_declareContext {
	var p = new(Var_declareContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_var_declare
	return p
}

func InitEmptyVar_declareContext(p *Var_declareContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_var_declare
}

func (*Var_declareContext) IsVar_declareContext() {}

func NewVar_declareContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_declareContext {
	var p = new(Var_declareContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_var_declare

	return p
}

func (s *Var_declareContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_declareContext) LET() antlr.TerminalNode {
	return s.GetToken(RustLikeParserLET, 0)
}

func (s *Var_declareContext) ID() antlr.TerminalNode {
	return s.GetToken(RustLikeParserID, 0)
}

func (s *Var_declareContext) MUT() antlr.TerminalNode {
	return s.GetToken(RustLikeParserMUT, 0)
}

func (s *Var_declareContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *Var_declareContext) Var_init() IVar_initContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_initContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_initContext)
}

func (s *Var_declareContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_declareContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_declareContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterVar_declare(s)
	}
}

func (s *Var_declareContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitVar_declare(s)
	}
}

func (p *RustLikeParser) Var_declare() (localctx IVar_declareContext) {
	localctx = NewVar_declareContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, RustLikeParserRULE_var_declare)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		p.Match(RustLikeParserLET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == RustLikeParserMUT {
		{
			p.SetState(166)
			p.Match(RustLikeParserMUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(169)
		p.Match(RustLikeParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == RustLikeParserCOLON {
		{
			p.SetState(170)
			p.Var_type()
		}

	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == RustLikeParserASSIGN {
		{
			p.SetState(173)
			p.Var_init()
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

// IVar_typeContext is an interface to support dynamic dispatch.
type IVar_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COLON() antlr.TerminalNode
	Type_() ITypeContext

	// IsVar_typeContext differentiates from other interfaces.
	IsVar_typeContext()
}

type Var_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_typeContext() *Var_typeContext {
	var p = new(Var_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_var_type
	return p
}

func InitEmptyVar_typeContext(p *Var_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_var_type
}

func (*Var_typeContext) IsVar_typeContext() {}

func NewVar_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_typeContext {
	var p = new(Var_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_var_type

	return p
}

func (s *Var_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_typeContext) COLON() antlr.TerminalNode {
	return s.GetToken(RustLikeParserCOLON, 0)
}

func (s *Var_typeContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *Var_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterVar_type(s)
	}
}

func (s *Var_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitVar_type(s)
	}
}

func (p *RustLikeParser) Var_type() (localctx IVar_typeContext) {
	localctx = NewVar_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, RustLikeParserRULE_var_type)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(176)
		p.Match(RustLikeParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(177)
		p.Type_()
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

// IVar_initContext is an interface to support dynamic dispatch.
type IVar_initContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ASSIGN() antlr.TerminalNode
	Expr() IExprContext

	// IsVar_initContext differentiates from other interfaces.
	IsVar_initContext()
}

type Var_initContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_initContext() *Var_initContext {
	var p = new(Var_initContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_var_init
	return p
}

func InitEmptyVar_initContext(p *Var_initContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_var_init
}

func (*Var_initContext) IsVar_initContext() {}

func NewVar_initContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_initContext {
	var p = new(Var_initContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_var_init

	return p
}

func (s *Var_initContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_initContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(RustLikeParserASSIGN, 0)
}

func (s *Var_initContext) Expr() IExprContext {
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

func (s *Var_initContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_initContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_initContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterVar_init(s)
	}
}

func (s *Var_initContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitVar_init(s)
	}
}

func (p *RustLikeParser) Var_init() (localctx IVar_initContext) {
	localctx = NewVar_initContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, RustLikeParserRULE_var_init)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.Match(RustLikeParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(180)
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

// IVar_assignContext is an interface to support dynamic dispatch.
type IVar_assignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expr() IExprContext

	// IsVar_assignContext differentiates from other interfaces.
	IsVar_assignContext()
}

type Var_assignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_assignContext() *Var_assignContext {
	var p = new(Var_assignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_var_assign
	return p
}

func InitEmptyVar_assignContext(p *Var_assignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_var_assign
}

func (*Var_assignContext) IsVar_assignContext() {}

func NewVar_assignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_assignContext {
	var p = new(Var_assignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_var_assign

	return p
}

func (s *Var_assignContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_assignContext) ID() antlr.TerminalNode {
	return s.GetToken(RustLikeParserID, 0)
}

func (s *Var_assignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(RustLikeParserASSIGN, 0)
}

func (s *Var_assignContext) Expr() IExprContext {
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

func (s *Var_assignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_assignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_assignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterVar_assign(s)
	}
}

func (s *Var_assignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitVar_assign(s)
	}
}

func (p *RustLikeParser) Var_assign() (localctx IVar_assignContext) {
	localctx = NewVar_assignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, RustLikeParserRULE_var_assign)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.Match(RustLikeParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(183)
		p.Match(RustLikeParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(184)
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

// IStat_if_elseContext is an interface to support dynamic dispatch.
type IStat_if_elseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIF() []antlr.TerminalNode
	IF(i int) antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	AllELSE() []antlr.TerminalNode
	ELSE(i int) antlr.TerminalNode

	// IsStat_if_elseContext differentiates from other interfaces.
	IsStat_if_elseContext()
}

type Stat_if_elseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStat_if_elseContext() *Stat_if_elseContext {
	var p = new(Stat_if_elseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_stat_if_else
	return p
}

func InitEmptyStat_if_elseContext(p *Stat_if_elseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_stat_if_else
}

func (*Stat_if_elseContext) IsStat_if_elseContext() {}

func NewStat_if_elseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Stat_if_elseContext {
	var p = new(Stat_if_elseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_stat_if_else

	return p
}

func (s *Stat_if_elseContext) GetParser() antlr.Parser { return s.parser }

func (s *Stat_if_elseContext) AllIF() []antlr.TerminalNode {
	return s.GetTokens(RustLikeParserIF)
}

func (s *Stat_if_elseContext) IF(i int) antlr.TerminalNode {
	return s.GetToken(RustLikeParserIF, i)
}

func (s *Stat_if_elseContext) AllExpr() []IExprContext {
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

func (s *Stat_if_elseContext) Expr(i int) IExprContext {
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

func (s *Stat_if_elseContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *Stat_if_elseContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
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

	return t.(IBlockContext)
}

func (s *Stat_if_elseContext) AllELSE() []antlr.TerminalNode {
	return s.GetTokens(RustLikeParserELSE)
}

func (s *Stat_if_elseContext) ELSE(i int) antlr.TerminalNode {
	return s.GetToken(RustLikeParserELSE, i)
}

func (s *Stat_if_elseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stat_if_elseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Stat_if_elseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterStat_if_else(s)
	}
}

func (s *Stat_if_elseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitStat_if_else(s)
	}
}

func (p *RustLikeParser) Stat_if_else() (localctx IStat_if_elseContext) {
	localctx = NewStat_if_elseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, RustLikeParserRULE_stat_if_else)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)
		p.Match(RustLikeParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(187)
		p.expr(0)
	}
	{
		p.SetState(188)
		p.Block()
	}
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(189)
				p.Match(RustLikeParserELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(190)
				p.Match(RustLikeParserIF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(191)
				p.expr(0)
			}
			{
				p.SetState(192)
				p.Block()
			}

		}
		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == RustLikeParserELSE {
		{
			p.SetState(199)
			p.Match(RustLikeParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(200)
			p.Block()
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

// IStat_whileContext is an interface to support dynamic dispatch.
type IStat_whileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHILE() antlr.TerminalNode
	Expr() IExprContext
	Block() IBlockContext

	// IsStat_whileContext differentiates from other interfaces.
	IsStat_whileContext()
}

type Stat_whileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStat_whileContext() *Stat_whileContext {
	var p = new(Stat_whileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_stat_while
	return p
}

func InitEmptyStat_whileContext(p *Stat_whileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_stat_while
}

func (*Stat_whileContext) IsStat_whileContext() {}

func NewStat_whileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Stat_whileContext {
	var p = new(Stat_whileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_stat_while

	return p
}

func (s *Stat_whileContext) GetParser() antlr.Parser { return s.parser }

func (s *Stat_whileContext) WHILE() antlr.TerminalNode {
	return s.GetToken(RustLikeParserWHILE, 0)
}

func (s *Stat_whileContext) Expr() IExprContext {
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

func (s *Stat_whileContext) Block() IBlockContext {
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

func (s *Stat_whileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stat_whileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Stat_whileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterStat_while(s)
	}
}

func (s *Stat_whileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitStat_while(s)
	}
}

func (p *RustLikeParser) Stat_while() (localctx IStat_whileContext) {
	localctx = NewStat_whileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, RustLikeParserRULE_stat_while)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(203)
		p.Match(RustLikeParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(204)
		p.expr(0)
	}
	{
		p.SetState(205)
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

// IStat_loopContext is an interface to support dynamic dispatch.
type IStat_loopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LOOP() antlr.TerminalNode
	Block() IBlockContext

	// IsStat_loopContext differentiates from other interfaces.
	IsStat_loopContext()
}

type Stat_loopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStat_loopContext() *Stat_loopContext {
	var p = new(Stat_loopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_stat_loop
	return p
}

func InitEmptyStat_loopContext(p *Stat_loopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_stat_loop
}

func (*Stat_loopContext) IsStat_loopContext() {}

func NewStat_loopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Stat_loopContext {
	var p = new(Stat_loopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_stat_loop

	return p
}

func (s *Stat_loopContext) GetParser() antlr.Parser { return s.parser }

func (s *Stat_loopContext) LOOP() antlr.TerminalNode {
	return s.GetToken(RustLikeParserLOOP, 0)
}

func (s *Stat_loopContext) Block() IBlockContext {
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

func (s *Stat_loopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stat_loopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Stat_loopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterStat_loop(s)
	}
}

func (s *Stat_loopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitStat_loop(s)
	}
}

func (p *RustLikeParser) Stat_loop() (localctx IStat_loopContext) {
	localctx = NewStat_loopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, RustLikeParserRULE_stat_loop)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(207)
		p.Match(RustLikeParserLOOP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(208)
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
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
