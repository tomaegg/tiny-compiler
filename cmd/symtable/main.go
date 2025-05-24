package main

import (
	"fmt"
	"log"
	"os"
	"tj-compiler/g4/parser"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage:\n   %v <source_file>\n", os.Args[0])
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

	fmt.Println("Parsing Tree:")
	fmt.Println(tree.ToStringTree(nil, parser))
}
