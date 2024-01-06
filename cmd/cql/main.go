package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/hawyar/cql"
	"github.com/hawyar/cql/internal/parser"
)

const (
	usage = `Usage: cql [options] [command]

COMMANDS:
  repl 			Interactive REPL for CQL Library and FHIRPath expressions

OPTIONS:
  --version		-v   	Show version 
  --help		-h      Show usage information

EXAMPLES:
  $ cql repl # start the REPL
  $ cql file.cql # parse a CQL file`
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
		cql.NewREPL().Start()
	}

	source, err := ReadFile(os.Args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// errListener := cql.NewErrorListener()

	lexer := parser.NewcqlLexer(source)

	// lexer.RemoveErrorListeners()
	// lexer.AddErrorListener(errListener)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := parser.NewcqlParser(stream)

	// parser.RemoveErrorListeners()
	// parser.AddErrorListener(errListener)

	parser.BuildParseTrees = true

	listener := cql.NewListener()

	listener.Lexer = lexer
	listener.Parser = parser

	antlr.ParseTreeWalkerDefault.Walk(listener, parser.Library())
}

func ReadFile(path string) (*antlr.FileStream, error) {
	if path == "" {
		return nil, fmt.Errorf("empty file path")
	}

	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	file, err := antlr.NewFileStream(path)

	if err != nil {
		return nil, err
	}

	return file, nil
}
