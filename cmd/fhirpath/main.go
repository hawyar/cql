package main

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	fhirpath "github.com/hawyar/cql/fhirpath/parser"
)

type FHIRPathListener struct {
	*fhirpath.BasefhirpathListener
}

func (f *FHIRPathListener) Report(ctx antlr.ParserRuleContext, msg string) {
	fmt.Printf("Line %d:%d %s\n", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), msg)
}

func main() {
	input, err := antlr.NewFileStream("./examples/basic.fp")

	if err != nil {
		panic(err)
	}

	lexer := fhirpath.NewfhirpathLexer(input)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := fhirpath.NewfhirpathParser(stream)

	listener := &FHIRPathListener{}

	antlr.ParseTreeWalkerDefault.Walk(listener, p.Expression())

	for _, token := range stream.GetAllTokens() {
		fmt.Println(token.GetText())
	}
	// cql.NewREPL().Start()
}
