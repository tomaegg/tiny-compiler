package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
	"tj-compiler/compiler"

	log "github.com/sirupsen/logrus"
)

var stageMap = map[string]compiler.CompileStage{
	"ir":       compiler.IRGen,
	"lex":      compiler.Lex,
	"parse":    compiler.Parse,
	"bin":      compiler.Bin,
	"asm":      compiler.ASM,
	"semantic": compiler.Semantic,
}

var (
	validLoglevels = []string{
		log.InfoLevel.String(),
		log.DebugLevel.String(),
		log.FatalLevel.String(),
		log.ErrorLevel.String(),
		log.TraceLevel.String(),
	}
	validStage []string
)

func init() {
	for s := range stageMap {
		validStage = append(validStage, s)
	}
}

func ParseFlag() compiler.Config {
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	// 定义可选参数
	out := fs.String("o", "", "output file (default STDOUT)")
	stage := fs.String("stage", "bin", fmt.Sprintf("stage: %s", strings.Join(validStage, ", ")))
	loglevel := fs.String("loglevel", "info", fmt.Sprintf("loglevel: %s", strings.Join(validLoglevels, ", ")))
	visualize := fs.Bool("visualize", false, "enable visualization mode")

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <src.rs> [-o out] [-stage=ir] [-loglevel=debug]\n", os.Args[0])
		fs.Usage()
		os.Exit(1)
	}
	src := os.Args[1]

	// 解析参数
	err := fs.Parse(os.Args[2:])
	if err != nil {
		log.Error(err)
	}

	// 检查 stage 合法性
	if !slices.Contains(validStage, *stage) {
		fmt.Printf("Invalid stage: %s\n", *stage)
		fs.Usage()
		os.Exit(1)
	}

	// 检查 loglevel 合法性
	if !slices.Contains(validLoglevels, *loglevel) {
		fmt.Printf("Invalid loglevel: %s\n", *loglevel)
		fs.Usage()
		os.Exit(1)
	}

	level, err := log.ParseLevel(*loglevel)
	if err != nil {
		log.Panicf("should not get invalid log level: %s", *loglevel)
	}
	log.SetLevel(level)
	log.SetFormatter(&log.TextFormatter{DisableTimestamp: true})

	// 打印结果
	log.Debugf("src: %s\n", src)
	log.Debugf("out: %s\n", *out)
	log.Debugf("stage: %s\n", *stage)
	log.Debugf("loglevel: %s\n", *loglevel)

	return compiler.Config{
		Stage:     stageMap[*stage],
		SrcPath:   src,
		OutPath:   *out,
		LogLevel:  *loglevel,
		Visualize: *visualize,
	}
}

func main() {
	conf := ParseFlag()
	cl := compiler.NewUnitCompiler(conf)
	cl.Compile()
}
