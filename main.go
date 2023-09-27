package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/hawyar/cql/parser"
)

type CqlListener struct {
	*parser.BasecqlListener
}

func (l *CqlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf("Enter: %s\n", ctx.GetText())
}

func (l *CqlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf("Exit: %s\n", ctx.GetText())
}

func main() {
	input, err := antlr.NewFileStream("./example/SimpleLib.cql")

	if err != nil {
		panic(err)
	}

	lexer := parser.NewcqlLexer(input)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewcqlParser(stream)

	listener := &CqlListener{}

	antlr.ParseTreeWalkerDefault.Walk(listener, p.Function())
}
