package compiler

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"sync"
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
	Semantic
	IRGen
	Bin
)

type Config struct {
	Visualize bool
	Stage     CompileStage
	SrcPath   string
	OutPath   string
	LogLevel  string
}

type UnitCompiler struct {
	stage CompileStage

	parser    *parser.RustLikeParser
	lexer     *parser.RustLikeLexer
	tokens    *antlr.CommonTokenStream
	parseTree antlr.ParseTree

	lexerErrCallback, parserErrCallback func() []string

	checker  *symtable.SemanticChecker
	dotGraph bool

	fout io.WriteCloser
	fin  string
}

func NewUnitCompiler(c Config) *UnitCompiler {
	level, err := log.ParseLevel(c.LogLevel)
	if err != nil {
		log.Warn(fmt.Errorf("invalid log level %s, use default level InfoLevel", c.LogLevel))
	}
	log.SetLevel(level)

	stat, err := os.Stat(c.SrcPath)
	if stat.IsDir() {
		log.Fatalf("%s is a directory", c.SrcPath)
	}

	if err != nil {
		log.Fatal(err)
	}

	input, err := antlr.NewFileStream(c.SrcPath)
	if err != nil {
		log.Fatal(err)
	}

	var fout io.WriteCloser
	if len(c.OutPath) == 0 {
		fout = os.Stdout
	} else {
		fout, err = os.Create(c.OutPath)
		if err != nil {
			log.Fatal(err)
		}
	}

	ret := &UnitCompiler{stage: c.Stage, fout: fout, fin: c.SrcPath, dotGraph: c.Visualize}

	if c.Stage >= Lex {
		ret.lexer, ret.lexerErrCallback = cmd.NewLexer(input)
		ret.tokens = antlr.NewCommonTokenStream(ret.lexer, antlr.TokenDefaultChannel)
	}

	if c.Stage >= Parse {
		ret.parser, ret.parserErrCallback = cmd.NewParser(ret.tokens)
		ret.parseTree = ret.parser.Prog()
	}

	if c.Stage >= Semantic {
		ret.checker = symtable.NewSemanticChecker(ret.parseTree)
	}

	return ret
}

func (uc *UnitCompiler) Lex() {
	// uc.tokens.Fill()
	for _, token := range uc.tokens.GetAllTokens() {
		tokenType := token.GetTokenType() // Token 类型（数值）
		if tokenType == antlr.TokenEOF {
			break
		}
		// 获取 Token 信息
		line := token.GetLine()     // 行号（从 1 开始）
		column := token.GetColumn() // 列号（从 0 开始）
		text := token.GetText()     // 文本内容
		tokenName := uc.lexer.SymbolicNames[tokenType]
		log.Infof("token(%2d:%2d) type=%7s, text='%s'\n", line, column, tokenName, text)
	}
	if errs := uc.lexerErrCallback(); len(errs) != 0 {
		for _, err := range uc.lexerErrCallback() {
			log.Error(err)
		}
		log.Fatalf("total %d lexer error occurs", len(errs))
	} else {
		log.Info("lexer check passed")
	}
}

func (uc *UnitCompiler) Parse() {
	antlr.ParseTreeWalkerDefault.Walk(&antlr.BaseParseTreeListener{}, uc.parseTree)
	if errs := uc.parserErrCallback(); len(errs) != 0 {
		for _, err := range uc.parserErrCallback() {
			log.Error(err)
		}
		log.Fatalf("total %d parser error occurs", len(errs))
	} else {
		log.Info("parser check passed")
	}
}

func (uc *UnitCompiler) Semantic() {
	if errs := uc.checker.Check(); errs != 0 {
		log.Fatalf("total %d semantic error occurs", errs)
	} else {
		log.Info("semantic check passed")
	}
	if !uc.dotGraph {
		return
	}

	// 创建临时 dot 文件
	tmpFile, err := os.CreateTemp("", "graphviz-*.dot")
	if err != nil {
		log.Fatalf("failed to create temp file: %v", err)
	}
	tmpFile.Write(uc.checker.SymbolTable().ToDotGraph())
	defer os.Remove(tmpFile.Name())
	// dot -Tsvg symtable.gv -o symtable.svg
	// dot -Tpng symtable.gv -o symtable.png
	// dot -Tpdf symtable.gv -o symtable.pdf
	outputName := fmt.Sprintf("%s.gv", path.Base(uc.fin))
	parent := os.Getenv("DOT_DIR")
	cmds := []*exec.Cmd{
		exec.Command("dot", "-Tpng", tmpFile.Name(), "-o", path.Join(parent, outputName+".png")),
		exec.Command("dot", "-Tsvg", tmpFile.Name(), "-o", path.Join(parent, outputName+".svg")),
		exec.Command("dot", "-Tpdf", tmpFile.Name(), "-o", path.Join(parent, outputName+".pdf")),
	}
	var wg sync.WaitGroup
	wg.Add(len(cmds))
	for _, cmd := range cmds {
		go func() {
			log.Debug(cmd.String())
			if err := cmd.Run(); err != nil {
				log.Error(err)
			}
			wg.Done()
		}()
	}

	log.Info("dot graph generated")
	wg.Wait()
}

func (uc *UnitCompiler) IR() {
	irGen, irCancel := ir.NewIRGenerator(uc.fin, uc.checker.SymbolTable(), uc.parseTree)
	defer irCancel()
	result := irGen.IR()
	uc.fout.Write(result)
	if err := uc.fout.Close(); err != nil {
		log.Fatal(err)
	}
	log.Info("ir generated done")
}

func (uc *UnitCompiler) Compile() {
	if uc.stage <= Lex {
		uc.Lex()
	} else if uc.stage >= Parse {
		uc.Parse()
	}

	if uc.stage >= Semantic {
		uc.Semantic()
	}

	if uc.stage >= IRGen {
		uc.IR()
	}

}
