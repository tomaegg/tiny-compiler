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
	"ir":    compiler.IRGen,
	"lex":   compiler.Lex,
	"parse": compiler.Parse,
	"bin":   compiler.Bin,
}
var (
	validLoglevels = []string{log.InfoLevel.String(), log.DebugLevel.String(), log.FatalLevel.String(), log.ErrorLevel.String()}
	validStage     []string
)

func init() {
	for s := range stageMap {
		validStage = append(validStage, s)
	}
}

func ParseFlag() compiler.Config {

	// 定义可选参数
	out := flag.String("o", "", "output file (default STDOUT)")
	stage := flag.String("stage", "bin", fmt.Sprintf("stage: %s", strings.Join(validStage, ", ")))
	loglevel := flag.String("loglevel", "info", fmt.Sprintf("loglevel: %s", strings.Join(validLoglevels, ", ")))

	// 解析参数
	flag.Parse()

	// 获取位置参数
	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s src.rs [-o out] [--stage=ir] [--loglevel=debug]\n", os.Args[0])
		flag.Usage()
		os.Exit(1)
	}
	src := args[0]

	// 检查 stage 合法性
	if !slices.Contains(validStage, *stage) {
		fmt.Printf("Invalid stage: %s\n", *stage)
		os.Exit(1)
	}

	// 检查 loglevel 合法性
	if !slices.Contains(validLoglevels, *loglevel) {
		fmt.Printf("Invalid loglevel: %s\n", *loglevel)
		os.Exit(1)
	}

	level, _ := log.ParseLevel(*loglevel)
	log.SetLevel(level)

	// 打印结果
	log.Debugf("src: %s\n", src)
	log.Debugf("out: %s\n", *out)
	log.Debugf("stage: %s\n", *stage)
	log.Debugf("loglevel: %s\n", *loglevel)

	return compiler.Config{
		Stage:    stageMap[*stage],
		SrcPath:  src,
		OutPath:  *out,
		LogLevel: *loglevel,
	}
}

func main() {
	ParseFlag()
}
