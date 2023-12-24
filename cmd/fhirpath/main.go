package main

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/hawyar/cql/fhirpath"
	"github.com/hawyar/cql/repl"
)

type FHIRPathListener struct {
	*fhirpath.BasefhirpathListener
}

func main() {
	input, err := antlr.NewFileStream("./examples/basic_fhirpath.fp")

	if err != nil {
		panic(err)
	}

	lexer := fhirpath.NewfhirpathLexer(input)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := fhirpath.NewfhirpathParser(stream)

	listener := &FHIRPathListener{}

	antlr.ParseTreeWalkerDefault.Walk(listener, p.Expression())

	// tokens := stream.GetAllTokens()

	// for _, token := range tokens {
	// 	fmt.Println(token)
	// }

	repl := repl.NewREPL()
	repl.Start()
}
