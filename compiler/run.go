package compiler

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"runtime"
	"sync"
	"tj-compiler/cmd"
	"tj-compiler/g4/parser"
	"tj-compiler/ir"
	"tj-compiler/symtable"
	"tj-compiler/target"
	"tj-compiler/utils"

	"github.com/antlr4-go/antlr/v4"
	log "github.com/sirupsen/logrus"
	"tinygo.org/x/go-llvm"
)

type CompileStage int

const (
	Lex CompileStage = iota
	Parse
	Semantic
	IRGen
	ASM
	Bin
	Exec
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
	releaseCallback                     []func()

	dotGraph bool
	checker  *symtable.SemanticChecker
	irGen    *ir.IRGenerator

	fin string

	fconfig *OutputConfig
}

type OutputConfig struct {
	FPath string
	Stage CompileStage
}

func (oc *OutputConfig) Path() string {
	stage := oc.Stage
	if stage > ASM && len(oc.FPath) == 0 {
		return "a.out"
	}
	return oc.FPath
}

func (oc *OutputConfig) Open() io.WriteCloser {
	stage := oc.Stage
	if stage <= ASM && len(oc.FPath) == 0 {
		return os.Stdout
	}

	if stage > ASM && len(oc.FPath) == 0 {
		oc.FPath = "a.out"
	}

	if err := utils.RemoveIfExists(oc.FPath); err != nil {
		log.Fatal(err)
	}

	// default out put to a.out
	perm := os.FileMode(0644)
	if stage == Exec {
		perm = 0755
	}
	flag := os.O_WRONLY | os.O_CREATE | os.O_TRUNC

	fout, err := os.OpenFile(oc.FPath, flag, perm)
	if err != nil {
		log.Fatal(err)
	}

	return fout
}

func NewUnitCompiler(c Config) *UnitCompiler {
	level, err := log.ParseLevel(c.LogLevel)
	if err != nil {
		log.Warn(fmt.Errorf("invalid log level %s, use default level InfoLevel", c.LogLevel))
	}
	log.SetLevel(level)

	stat, err := os.Stat(c.SrcPath)
	if err != nil {
		log.Fatal(err)
	}
	if stat.IsDir() {
		log.Fatalf("%s is a directory", c.SrcPath)
	}

	input, err := antlr.NewFileStream(c.SrcPath)
	if err != nil {
		log.Fatal(err)
	}

	ret := &UnitCompiler{
		stage:    c.Stage,
		fin:      c.SrcPath,
		dotGraph: c.Visualize,
		fconfig: &OutputConfig{
			FPath: c.OutPath,
			Stage: c.Stage,
		},
	}

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

func (uc *UnitCompiler) DotGraph() {
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

	wg.Wait()
	log.Info("dot graph generated")
}

func (uc *UnitCompiler) Semantic() {
	if errs := uc.checker.Check(); errs != 0 {
		if uc.dotGraph {
			uc.DotGraph()
		}
		log.Fatalf("total %d semantic error occurs", errs)
	} else {
		log.Info("semantic check passed")
	}

	if uc.dotGraph {
		uc.DotGraph()
	}
}

func (uc *UnitCompiler) IR() {
	irGen, llvmRelease := ir.NewIRGenerator(uc.fin, uc.checker.SymbolTable(), uc.parseTree)
	uc.irGen = irGen
	uc.releaseCallback = append(uc.releaseCallback, llvmRelease)
	result := irGen.IR()
	if uc.stage == IRGen {
		fout := uc.fconfig.Open()
		fout.Write(result)
		if err := fout.Close(); err != nil {
			log.Fatal(err)
		}
	}
	log.Info("ir generated done")
}

func (uc *UnitCompiler) Target() {
	bin := target.NewBinGenerator(uc.irGen.Module())
	codeGen := llvm.ObjectFile
	switch uc.stage {
	case Exec:
		tmpBinary, err := os.CreateTemp("/var/tmp/", "tiny-compiler-*")
		if err != nil {
			log.Fatal(err)
		}
		defer tmpBinary.Close()
		bin.Generate(tmpBinary, llvm.ObjectFile)
		uc.Exec(tmpBinary.Name())
		return

	case Bin:
		codeGen = llvm.ObjectFile
	case ASM:
		codeGen = llvm.AssemblyFile
	}

	fout := uc.fconfig.Open()
	defer fout.Close()
	bin.Generate(fout, codeGen)
	log.Info("target generated done")
}

func (uc *UnitCompiler) Exec(binaryIn string) {
	out := uc.fconfig.Path()
	cmd := exec.Command("clang", binaryIn, "-o", out)
	log.Debug(cmd.String())
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.WithError(err).Fatalf("link: %s", output)
	}
	log.Info("link stage done")
	log.Infof("target executable: %s", out)
}

func (uc *UnitCompiler) Compile() {
	if uc.stage <= Lex {
		uc.Lex()
	} else if uc.stage >= Parse {
		uc.Parse()
	}

	if uc.stage >= Semantic {
		uc.Semantic()
		runtime.GC()
	}

	if uc.stage >= IRGen {
		uc.IR()
		runtime.GC()
	}

	if uc.stage >= ASM {
		uc.Target()
	}

	uc.release()
}

func (uc *UnitCompiler) release() {
	for _, f := range uc.releaseCallback {
		f()
	}
}
