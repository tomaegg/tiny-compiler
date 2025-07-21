package main

import (
	"fmt"
	"os"
	"tj-compiler/cmd"
	"tj-compiler/symtable"

	"github.com/antlr4-go/antlr/v4"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		DisableTimestamp: true, // 禁用时间戳
	})

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

	log.Println("[Parsing Tree]")
	log.Println(tree.ToStringTree(nil, parser))

	checker := symtable.NewSemanticChecker(tree)
	if total := checker.Check(); total != 0 {
		log.Errorf("total %d errors occurs, semantic check done", total)
	}

	symTable := checker.SymbolTable()
	dot := symTable.ToDotGraph()

	log.Println("[Dot Graph]")
	fmt.Println(string(dot))
}
