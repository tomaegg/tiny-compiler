package compiler

import (
	"fmt"
	"io"
	"os"
	"tj-compiler/cmd"
	"tj-compiler/g4/parser"
	"tj-compiler/ir"
	"tj-compiler/symtable"

	"github.com/antlr4-go/antlr/v4"
	log "github.com/sirupsen/logrus"
)

type CompileStage int

const (
	Lex CompileStage = iota
	Parse
	Symtable
	IRGen
	Bin
)

type Config struct {
	Stage    CompileStage
	SrcPath  string
	OutPath  string
	LogLevel string
}

type UnitCompiler struct {
	stage CompileStage

	parser                              *parser.RustLikeParser
	lexer                               *parser.RustLikeLexer
	lexerErrCallback, parserErrCallback func() []string

	checker *symtable.SemanticChecker

	irGen    *ir.IRGenerator
	irCancel func()

	fout io.WriteCloser
}

func NewUnitCompiler(c Config) *UnitCompiler {
	level, err := log.ParseLevel(c.LogLevel)
	if err != nil {
		log.Warn(fmt.Errorf("invalid log level %s, use default level InfoLevel", c.LogLevel))
	}
	log.SetLevel(level)

	input, err := antlr.NewFileStream(c.SrcPath)
	if err != nil {
		log.Fatal(err)
	}

	fout, err := os.Create(c.OutPath)
	if err != nil {
		log.Error(err)
		fout = os.Stdout
	}

	ret := &UnitCompiler{stage: c.Stage, fout: fout}

	switch {
	case c.Stage >= Lex:
		ret.lexer, ret.lexerErrCallback = cmd.NewLexer(input)
		fallthrough

	case c.Stage >= Parse:
		tokens := antlr.NewCommonTokenStream(ret.lexer, antlr.TokenDefaultChannel)
		ret.parser, ret.parserErrCallback = cmd.NewParser(tokens)
		fallthrough

	case c.Stage >= Symtable:
		ret.checker = symtable.NewSemanticChecker(ret.parser.Prog())
		fallthrough

	case c.Stage >= IRGen:
		ret.irGen, ret.irCancel = ir.NewIRGenerator(c.SrcPath, ret.checker.SymbolTable(), ret.parser.Prog())
		fallthrough

	default:
	}

	return ret
}

func (uc *UnitCompiler) Compile() {
	switch {
	case uc.stage <= Lex:
		for {
			token := uc.lexer.NextToken()
			tokenType := token.GetTokenType() // Token 类型（数值）
			if tokenType == antlr.TokenEOF {
				break
			}
			// 获取 Token 信息
			line := token.GetLine()     // 行号（从 1 开始）
			column := token.GetColumn() // 列号（从 0 开始）
			text := token.GetText()     // 文本内容
			tokenName := uc.lexer.SymbolicNames[tokenType]
			log.Debugf("token(%2d:%2d) type=%7s, text='%s'\n", line, column, tokenName, text)
		}
		if errs := uc.lexerErrCallback(); len(errs) != 0 {
			for _, err := range uc.lexerErrCallback() {
				log.Error(err)
			}
			log.Fatalf("total %d lexer error occurs", len(errs))
		} else {
			log.Info("lexer check passed")
		}

	case uc.stage <= Parse:
		antlr.ParseTreeWalkerDefault.Walk(&antlr.BaseParseTreeListener{}, uc.parser.Prog())
		if errs := uc.parserErrCallback(); len(errs) != 0 {
			for _, err := range uc.parserErrCallback() {
				log.Error(err)
			}
			log.Fatalf("total %d parser error occurs", len(errs))
		} else {
			log.Info("parser check passed")
		}

	case uc.stage <= Symtable:
		if errs := uc.checker.Check(); errs != 0 {
			log.Fatalf("total %d semantic error occurs", errs)
		} else {
			log.Info("semantic check passed")
		}

	case uc.stage <= IRGen:
		defer uc.irCancel()
		result := uc.irGen.IR()
		uc.fout.Write(result)
		if err := uc.fout.Close(); err != nil {
			log.Fatal(err)
		}

	}

}
