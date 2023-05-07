package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/hawyar/cql/parser"
)

func main() {
	// f, err := os.Open("./example/ChlamydiaScreening_CQM.cql")

	// input, err := antlr.NewFileStream("./example/ChlamydiaScreening_CQM.cql")
	input, err := antlr.NewFileStream("./example/SimpleLib.cql")

	if err != nil {
		panic(err)
	}

	lexer := parser.NewcqlLexer(input)

	// Read all tokens
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("type: %s, text: %s\n", lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}