package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/hawyar/cql"
	"github.com/hawyar/cql/internal/parser"
)

type Config struct {
	Input  string
	Output io.Writer
	Format cql.OutputFormat
}

const (
	usage = `Usage: cql [options] [command]

COMMANDS:
  repl 			interactive repl for CQL Library and FHIRPath expressions

OPTIONS:
  -version	  	show version
  -help	 		show usage
  -format 		output format either JSON or XML case insensitive (default: json)
  -input 		input file path, (e.g. example.cql)
  -output 		output file path, (e.g. output.json)

EXAMPLES:
  $ cql repl # start the REPL
  $ cql -input example.cql # parse a CQL file
  $ cql -input example.cql -format xml # set output to xml`
)

func main() {
	version := flag.Bool("version", false, "Show version")
	help := flag.Bool("help", false, "Show usage information")
	input := flag.String("input", "", "File path to a CQL file ")
	format := flag.String("format", "json", "Output format, either JSON or XML")
	// output := flag.String("output", "", "The output file path")

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

	if *input == "" {
		fmt.Println("invalid input file: ", *input)
		os.Exit(1)
	}

	source, err := ReadFile(*input)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	config := &Config{
		Input:  *input,
		Output: os.Stdout,
	}

	if *format == "xml" || *format == "XML" {
		config.Format = cql.XMLFormat
	}

	lexer := parser.NewcqlLexer(source)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := parser.NewcqlParser(stream)

	// parser.RemoveErrorListeners()
	// parser.AddErrorListener(errListener)

	parser.BuildParseTrees = true

	listener := cql.NewListener()

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
