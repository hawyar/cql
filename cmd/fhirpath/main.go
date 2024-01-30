package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/hawyar/cql"
	"github.com/hawyar/cql/internal/fhirpath"
)

const (
	usage = `Usage: fhirpath [options] [command]

COMMANDS:
  repl 			Interactive REPL for CQL Library and FHIRPath expressions

OPTIONS:
  --version		-v   	Show version 
  --help		-h      Show usage information

EXAMPLES:
  $ fhirpath repl`
)

func main() {
	version := flag.Bool("version", false, "Show version")
	help := flag.Bool("help", false, "Show usage information")

	flag.Parse()

	if *version {
		fmt.Println("cql version", "0.0.1")
		os.Exit(0)
	}

	if *help {
		fmt.Println(usage)
		os.Exit(0)
	}

	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(0)
	}

	if os.Args[1] == "repl" {
		cql.NewFHIRPathRepl().Start()
	}

	// source, err := cql.ReadFile(os.Args[1])

	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	lexer := fhirpath.NewfhirpathLexer(antlr.NewInputStream(os.Args[1]))

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := fhirpath.NewfhirpathParser(tokenStream)

	parser.BuildParseTrees = true

	listener := cql.NewFHIRPathListener()

	antlr.ParseTreeWalkerDefault.Walk(listener, parser.Expression())
}
