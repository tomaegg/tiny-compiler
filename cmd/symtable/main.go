package main

import (
	"fmt"
	"os"
	"tj-compiler/g4/parser"
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

	lexer := parser.NewRustLikeLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := parser.NewRustLikeParser(tokens)
	tree := parser.Prog() // 假设起始规则是 `expr`

	log.Println("[Parsing Tree]")
	log.Println(tree.ToStringTree(nil, parser))

	checker := symtable.NewSemanticChecker(tree)
	symTable := checker.SymbolTable()
	dot := symTable.ToDotGraph()

	log.Println("[Dot Graph]")
	fmt.Println(string(dot))
}
