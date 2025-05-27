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
		"prog", "declaration", "expr", "funcCallList", "funcCallParam", "funcDeclaration",
		"funcSignature", "funcDeclarationReturn", "funcParamsList", "funcParams",
		"funcParam", "funcBlock", "rtype", "block", "stat", "varType", "varInit",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 39, 195, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 1, 0, 1, 0, 1, 1, 5, 1, 38, 8, 1, 10, 1, 12, 1, 41, 9, 1,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 52, 8, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 63, 8, 2, 10,
		2, 12, 2, 66, 9, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 5, 4, 75,
		8, 4, 10, 4, 12, 4, 78, 9, 4, 1, 4, 3, 4, 81, 8, 4, 1, 5, 1, 5, 1, 5, 1,
		6, 1, 6, 1, 6, 1, 6, 3, 6, 90, 8, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 9, 5, 9, 102, 8, 9, 10, 9, 12, 9, 105, 9, 9, 1, 9,
		3, 9, 108, 8, 9, 1, 10, 3, 10, 111, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 11, 1, 11, 5, 11, 119, 8, 11, 10, 11, 12, 11, 122, 9, 11, 1, 11, 1,
		11, 1, 12, 1, 12, 1, 13, 1, 13, 5, 13, 130, 8, 13, 10, 13, 12, 13, 133,
		9, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 3, 14, 140, 8, 14, 1, 14, 1,
		14, 1, 14, 3, 14, 145, 8, 14, 1, 14, 1, 14, 3, 14, 149, 8, 14, 1, 14, 3,
		14, 152, 8, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 171,
		8, 14, 10, 14, 12, 14, 174, 9, 14, 1, 14, 1, 14, 3, 14, 178, 8, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 187, 8, 14, 1, 15,
		1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 0, 1, 4, 17, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 0, 3, 1, 0, 19, 20, 1, 0, 17,
		18, 1, 0, 21, 26, 206, 0, 34, 1, 0, 0, 0, 2, 39, 1, 0, 0, 0, 4, 51, 1,
		0, 0, 0, 6, 67, 1, 0, 0, 0, 8, 80, 1, 0, 0, 0, 10, 82, 1, 0, 0, 0, 12,
		85, 1, 0, 0, 0, 14, 91, 1, 0, 0, 0, 16, 94, 1, 0, 0, 0, 18, 107, 1, 0,
		0, 0, 20, 110, 1, 0, 0, 0, 22, 116, 1, 0, 0, 0, 24, 125, 1, 0, 0, 0, 26,
		127, 1, 0, 0, 0, 28, 186, 1, 0, 0, 0, 30, 188, 1, 0, 0, 0, 32, 191, 1,
		0, 0, 0, 34, 35, 3, 2, 1, 0, 35, 1, 1, 0, 0, 0, 36, 38, 3, 10, 5, 0, 37,
		36, 1, 0, 0, 0, 38, 41, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 39, 40, 1, 0, 0,
		0, 40, 3, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 42, 43, 6, 2, -1, 0, 43, 44,
		5, 31, 0, 0, 44, 45, 3, 4, 2, 0, 45, 46, 5, 32, 0, 0, 46, 52, 1, 0, 0,
		0, 47, 48, 5, 15, 0, 0, 48, 52, 3, 6, 3, 0, 49, 52, 5, 15, 0, 0, 50, 52,
		5, 16, 0, 0, 51, 42, 1, 0, 0, 0, 51, 47, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0,
		51, 50, 1, 0, 0, 0, 52, 64, 1, 0, 0, 0, 53, 54, 10, 7, 0, 0, 54, 55, 7,
		0, 0, 0, 55, 63, 3, 4, 2, 8, 56, 57, 10, 6, 0, 0, 57, 58, 7, 1, 0, 0, 58,
		63, 3, 4, 2, 7, 59, 60, 10, 5, 0, 0, 60, 61, 7, 2, 0, 0, 61, 63, 3, 4,
		2, 6, 62, 53, 1, 0, 0, 0, 62, 56, 1, 0, 0, 0, 62, 59, 1, 0, 0, 0, 63, 66,
		1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 5, 1, 0, 0, 0,
		66, 64, 1, 0, 0, 0, 67, 68, 5, 31, 0, 0, 68, 69, 3, 8, 4, 0, 69, 70, 5,
		32, 0, 0, 70, 7, 1, 0, 0, 0, 71, 76, 3, 4, 2, 0, 72, 73, 5, 30, 0, 0, 73,
		75, 3, 4, 2, 0, 74, 72, 1, 0, 0, 0, 75, 78, 1, 0, 0, 0, 76, 74, 1, 0, 0,
		0, 76, 77, 1, 0, 0, 0, 77, 81, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 79, 81,
		1, 0, 0, 0, 80, 71, 1, 0, 0, 0, 80, 79, 1, 0, 0, 0, 81, 9, 1, 0, 0, 0,
		82, 83, 3, 12, 6, 0, 83, 84, 3, 22, 11, 0, 84, 11, 1, 0, 0, 0, 85, 86,
		5, 11, 0, 0, 86, 87, 5, 15, 0, 0, 87, 89, 3, 16, 8, 0, 88, 90, 3, 14, 7,
		0, 89, 88, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 13, 1, 0, 0, 0, 91, 92,
		5, 37, 0, 0, 92, 93, 3, 24, 12, 0, 93, 15, 1, 0, 0, 0, 94, 95, 5, 31, 0,
		0, 95, 96, 3, 18, 9, 0, 96, 97, 5, 32, 0, 0, 97, 17, 1, 0, 0, 0, 98, 103,
		3, 20, 10, 0, 99, 100, 5, 30, 0, 0, 100, 102, 3, 20, 10, 0, 101, 99, 1,
		0, 0, 0, 102, 105, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0,
		0, 104, 108, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 106, 108, 1, 0, 0, 0, 107,
		98, 1, 0, 0, 0, 107, 106, 1, 0, 0, 0, 108, 19, 1, 0, 0, 0, 109, 111, 5,
		10, 0, 0, 110, 109, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 112, 1, 0, 0,
		0, 112, 113, 5, 15, 0, 0, 113, 114, 5, 29, 0, 0, 114, 115, 3, 24, 12, 0,
		115, 21, 1, 0, 0, 0, 116, 120, 5, 35, 0, 0, 117, 119, 3, 28, 14, 0, 118,
		117, 1, 0, 0, 0, 119, 122, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120, 121,
		1, 0, 0, 0, 121, 123, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 123, 124, 5, 36,
		0, 0, 124, 23, 1, 0, 0, 0, 125, 126, 5, 4, 0, 0, 126, 25, 1, 0, 0, 0, 127,
		131, 5, 35, 0, 0, 128, 130, 3, 28, 14, 0, 129, 128, 1, 0, 0, 0, 130, 133,
		1, 0, 0, 0, 131, 129, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 134, 1, 0,
		0, 0, 133, 131, 1, 0, 0, 0, 134, 135, 5, 36, 0, 0, 135, 27, 1, 0, 0, 0,
		136, 187, 3, 26, 13, 0, 137, 139, 5, 9, 0, 0, 138, 140, 3, 4, 2, 0, 139,
		138, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 187,
		5, 28, 0, 0, 142, 144, 5, 5, 0, 0, 143, 145, 5, 10, 0, 0, 144, 143, 1,
		0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 148, 5, 15, 0,
		0, 147, 149, 3, 30, 15, 0, 148, 147, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0,
		149, 151, 1, 0, 0, 0, 150, 152, 3, 32, 16, 0, 151, 150, 1, 0, 0, 0, 151,
		152, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 187, 5, 28, 0, 0, 154, 155,
		5, 15, 0, 0, 155, 156, 5, 27, 0, 0, 156, 157, 3, 4, 2, 0, 157, 158, 5,
		28, 0, 0, 158, 187, 1, 0, 0, 0, 159, 160, 3, 4, 2, 0, 160, 161, 5, 28,
		0, 0, 161, 187, 1, 0, 0, 0, 162, 163, 5, 6, 0, 0, 163, 164, 3, 4, 2, 0,
		164, 172, 3, 26, 13, 0, 165, 166, 5, 7, 0, 0, 166, 167, 5, 6, 0, 0, 167,
		168, 3, 4, 2, 0, 168, 169, 3, 26, 13, 0, 169, 171, 1, 0, 0, 0, 170, 165,
		1, 0, 0, 0, 171, 174, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 172, 173, 1, 0,
		0, 0, 173, 177, 1, 0, 0, 0, 174, 172, 1, 0, 0, 0, 175, 176, 5, 7, 0, 0,
		176, 178, 3, 26, 13, 0, 177, 175, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178,
		187, 1, 0, 0, 0, 179, 180, 5, 8, 0, 0, 180, 181, 3, 4, 2, 0, 181, 182,
		3, 26, 13, 0, 182, 187, 1, 0, 0, 0, 183, 184, 5, 12, 0, 0, 184, 187, 3,
		26, 13, 0, 185, 187, 5, 28, 0, 0, 186, 136, 1, 0, 0, 0, 186, 137, 1, 0,
		0, 0, 186, 142, 1, 0, 0, 0, 186, 154, 1, 0, 0, 0, 186, 159, 1, 0, 0, 0,
		186, 162, 1, 0, 0, 0, 186, 179, 1, 0, 0, 0, 186, 183, 1, 0, 0, 0, 186,
		185, 1, 0, 0, 0, 187, 29, 1, 0, 0, 0, 188, 189, 5, 29, 0, 0, 189, 190,
		3, 24, 12, 0, 190, 31, 1, 0, 0, 0, 191, 192, 5, 27, 0, 0, 192, 193, 3,
		4, 2, 0, 193, 33, 1, 0, 0, 0, 19, 39, 51, 62, 64, 76, 80, 89, 103, 107,
		110, 120, 131, 139, 144, 148, 151, 172, 177, 186,
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
	RustLikeParserRULE_funcCallList          = 3
	RustLikeParserRULE_funcCallParam         = 4
	RustLikeParserRULE_funcDeclaration       = 5
	RustLikeParserRULE_funcSignature         = 6
	RustLikeParserRULE_funcDeclarationReturn = 7
	RustLikeParserRULE_funcParamsList        = 8
	RustLikeParserRULE_funcParams            = 9
	RustLikeParserRULE_funcParam             = 10
	RustLikeParserRULE_funcBlock             = 11
	RustLikeParserRULE_rtype                 = 12
	RustLikeParserRULE_block                 = 13
	RustLikeParserRULE_stat                  = 14
	RustLikeParserRULE_varType               = 15
	RustLikeParserRULE_varInit               = 16
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
		p.SetState(34)
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
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == RustLikeParserFN {
		{
			p.SetState(36)
			p.FuncDeclaration()
		}

		p.SetState(41)
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

type ExprAddSubContext struct {
	ExprContext
	lhs IExprContext
	rhs IExprContext
}

func NewExprAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprAddSubContext {
	var p = new(ExprAddSubContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

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

func (s *ExprNumContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(RustLikeParserNUMBER, 0)
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
	rhs IExprContext
}

func NewExprMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprMulDivContext {
	var p = new(ExprMulDivContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

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
	rhs IExprContext
}

func NewExprCmpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprCmpContext {
	var p = new(ExprCmpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

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
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExprParenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(43)
			p.Match(RustLikeParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(44)
			p.expr(0)
		}
		{
			p.SetState(45)
			p.Match(RustLikeParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewExprFuncCallContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(47)
			p.Match(RustLikeParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(48)
			p.FuncCallList()
		}

	case 3:
		localctx = NewExprIDContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(49)
			p.Match(RustLikeParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewExprNumContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(50)
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
	p.SetState(64)
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
			p.SetState(62)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprMulDivContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ExprMulDivContext).lhs = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RustLikeParserRULE_expr)
				p.SetState(53)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(54)
					_la = p.GetTokenStream().LA(1)

					if !(_la == RustLikeParserMULT || _la == RustLikeParserDIV) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(55)

					var _x = p.expr(8)

					localctx.(*ExprMulDivContext).rhs = _x
				}

			case 2:
				localctx = NewExprAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ExprAddSubContext).lhs = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RustLikeParserRULE_expr)
				p.SetState(56)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(57)
					_la = p.GetTokenStream().LA(1)

					if !(_la == RustLikeParserPLUS || _la == RustLikeParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(58)

					var _x = p.expr(7)

					localctx.(*ExprAddSubContext).rhs = _x
				}

			case 3:
				localctx = NewExprCmpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ExprCmpContext).lhs = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RustLikeParserRULE_expr)
				p.SetState(59)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(60)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&132120576) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(61)

					var _x = p.expr(6)

					localctx.(*ExprCmpContext).rhs = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(66)
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
	p.EnterRule(localctx, 6, RustLikeParserRULE_funcCallList)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Match(RustLikeParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(68)
		p.FuncCallParam()
	}
	{
		p.SetState(69)
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
	p.EnterRule(localctx, 8, RustLikeParserRULE_funcCallParam)
	var _la int

	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RustLikeParserID, RustLikeParserNUMBER, RustLikeParserLPAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(71)
			p.expr(0)
		}
		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == RustLikeParserCOMMA {
			{
				p.SetState(72)
				p.Match(RustLikeParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(73)
				p.expr(0)
			}

			p.SetState(78)
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
	p.EnterRule(localctx, 10, RustLikeParserRULE_funcDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.FuncSignature()
	}
	{
		p.SetState(83)
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
	p.EnterRule(localctx, 12, RustLikeParserRULE_funcSignature)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.Match(RustLikeParserFN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(86)
		p.Match(RustLikeParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(87)
		p.FuncParamsList()
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == RustLikeParserARROW {
		{
			p.SetState(88)
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
	p.EnterRule(localctx, 14, RustLikeParserRULE_funcDeclarationReturn)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Match(RustLikeParserARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(92)
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
	p.EnterRule(localctx, 16, RustLikeParserRULE_funcParamsList)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.Match(RustLikeParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(95)
		p.FuncParams()
	}
	{
		p.SetState(96)
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
	p.EnterRule(localctx, 18, RustLikeParserRULE_funcParams)
	var _la int

	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RustLikeParserMUT, RustLikeParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(98)
			p.FuncParam()
		}
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == RustLikeParserCOMMA {
			{
				p.SetState(99)
				p.Match(RustLikeParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(100)
				p.FuncParam()
			}

			p.SetState(105)
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
	p.EnterRule(localctx, 20, RustLikeParserRULE_funcParam)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == RustLikeParserMUT {
		{
			p.SetState(109)
			p.Match(RustLikeParserMUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(112)
		p.Match(RustLikeParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)
		p.Match(RustLikeParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(114)
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
	p.EnterRule(localctx, 22, RustLikeParserRULE_funcBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.Match(RustLikeParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36775760736) != 0 {
		{
			p.SetState(117)
			p.Stat()
		}

		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(123)
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
	p.EnterRule(localctx, 24, RustLikeParserRULE_rtype)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
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
	p.EnterRule(localctx, 26, RustLikeParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Match(RustLikeParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36775760736) != 0 {
		{
			p.SetState(128)
			p.Stat()
		}

		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(134)
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

func (s *StatIfElseContext) AllIF() []antlr.TerminalNode {
	return s.GetTokens(RustLikeParserIF)
}

func (s *StatIfElseContext) IF(i int) antlr.TerminalNode {
	return s.GetToken(RustLikeParserIF, i)
}

func (s *StatIfElseContext) AllExpr() []IExprContext {
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

func (s *StatIfElseContext) Expr(i int) IExprContext {
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

func (s *StatIfElseContext) AllBlock() []IBlockContext {
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

func (s *StatIfElseContext) Block(i int) IBlockContext {
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

func (s *StatIfElseContext) AllELSE() []antlr.TerminalNode {
	return s.GetTokens(RustLikeParserELSE)
}

func (s *StatIfElseContext) ELSE(i int) antlr.TerminalNode {
	return s.GetToken(RustLikeParserELSE, i)
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
	p.EnterRule(localctx, 28, RustLikeParserRULE_stat)
	var _la int

	var _alt int

	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStatBlockContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(136)
			p.Block()
		}

	case 2:
		localctx = NewStatFuncReturnContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(137)
			p.Match(RustLikeParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2147581952) != 0 {
			{
				p.SetState(138)
				p.expr(0)
			}

		}
		{
			p.SetState(141)
			p.Match(RustLikeParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewStatVarDeclareContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(142)
			p.Match(RustLikeParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == RustLikeParserMUT {
			{
				p.SetState(143)
				p.Match(RustLikeParserMUT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(146)
			p.Match(RustLikeParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == RustLikeParserCOLON {
			{
				p.SetState(147)
				p.VarType()
			}

		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == RustLikeParserASSIGN {
			{
				p.SetState(150)
				p.VarInit()
			}

		}
		{
			p.SetState(153)
			p.Match(RustLikeParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewStatVarAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(154)
			p.Match(RustLikeParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(155)
			p.Match(RustLikeParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(156)
			p.expr(0)
		}
		{
			p.SetState(157)
			p.Match(RustLikeParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewStatExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(159)
			p.expr(0)
		}
		{
			p.SetState(160)
			p.Match(RustLikeParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewStatIfElseContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(162)
			p.Match(RustLikeParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(163)
			p.expr(0)
		}
		{
			p.SetState(164)
			p.Block()
		}
		p.SetState(172)
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
					p.SetState(165)
					p.Match(RustLikeParserELSE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(166)
					p.Match(RustLikeParserIF)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(167)
					p.expr(0)
				}
				{
					p.SetState(168)
					p.Block()
				}

			}
			p.SetState(174)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == RustLikeParserELSE {
			{
				p.SetState(175)
				p.Match(RustLikeParserELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(176)
				p.Block()
			}

		}

	case 7:
		localctx = NewStatWhileContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(179)
			p.Match(RustLikeParserWHILE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(180)
			p.expr(0)
		}
		{
			p.SetState(181)
			p.Block()
		}

	case 8:
		localctx = NewStatLoopContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(183)
			p.Match(RustLikeParserLOOP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(184)
			p.Block()
		}

	case 9:
		localctx = NewStatEmptyContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(185)
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
	p.EnterRule(localctx, 30, RustLikeParserRULE_varType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Match(RustLikeParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(189)
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
	p.EnterRule(localctx, 32, RustLikeParserRULE_varInit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(191)
		p.Match(RustLikeParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(192)
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
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
