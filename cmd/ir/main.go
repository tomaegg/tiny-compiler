package main

import (
	"fmt"
	"os"
	"tj-compiler/cmd"
	"tj-compiler/ir"
	"tj-compiler/symtable"

	"github.com/antlr4-go/antlr/v4"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		DisableTimestamp: true, // 禁用时间戳
	})
	log.SetLevel(log.DebugLevel)

	if len(os.Args) < 2 {
		log.Printf("usage:\n   %v <source_file>\n", os.Args[0])
		os.Exit(-1)
	}

	filename := os.Args[1]
	input, err := antlr.NewFileStream(filename)
	if err != nil {
		log.Fatal(err)
	}

	lexer, lexerErrMsg := cmd.NewLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser, parserErrMsg := cmd.NewParser(tokens)
	tree := parser.Prog()
	if len(lexerErrMsg()) != 0 || len(parserErrMsg()) != 0 {
		os.Exit(-1)
	}
	log.Info("basic check passed")

	checker := symtable.NewSemanticChecker(tree)
	if total := checker.TotalErrors(); total != 0 {
		log.Fatalf("total %d errors occurs, semantic check done", total)
		return
	}
	log.Info("semantic check passed")

	symTable := checker.SymbolTable()

	irGenerator, cancel := ir.NewIRGenerator(filename, symTable)
	defer cancel()

	s := irGenerator.IR()
	fmt.Printf("%s", s)
}
