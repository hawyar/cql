package main

import (
	"flag"
	"fmt"
	"os"
)

const usage = `Usage: cql [options] [command]

COMMANDS:
  repl 			Interactive REPL for CQL Library and FHIRPath expressions

OPTIONS:
  --version		-v   	Show version 
  --help		-h      Show usage information

EXAMPLES:
  $ cql --help
  $ cql --version`

func main() {
	// os.Args[1:]

	versionF := flag.Bool("version", false, "Show version")
	helpF := flag.Bool("help", false, "Show usage information")

	flag.Parse()

	if *versionF {
		fmt.Println("cql version 0.0.1")
		os.Exit(0)
	}

	if *helpF {
		fmt.Println(usage)
		os.Exit(0)
	}

	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(0)
	}

	// freader := cql.NewFileReader("./examples/basic.fp")

	// input, err := antlr.NewFileStream("./examples/basic.fp")

	// if err != nil {
	// 	panic(err)
	// }

	// lexer := fhirpath.NewfhirpathLexer(input)

	// stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// parser := fhirpath.NewfhirpathParser(stream)

	// astBuilder := cql.NewFHIRPathASTBuilder()
	// antlr.ParseTreeWalkerDefault.Walk(astBuilder, parser.Expression())

	// for _, token := range stream.GetAllTokens() {
	// 	println(token.GetText())
	// }
}
