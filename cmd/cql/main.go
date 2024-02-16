package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/hawyar/cql"
)

func showUsage() {
	usage := `Usage: cql [FLAGS] [COMMANDS]

COMMANDS:
  parse 		parse a CQL file
  repl 			start a CQL repl

FLAGS:
  -version	  	show version
  -help	 		show usage
  -input 		input file path
  -output 		output file path, defaults to stdout
  -format 		output format, either JSON or XML (default: json)

Example:
  $ cql -input example.cql -format xml parse
  $ cql repl`
	fmt.Println(usage)
}

func main() {
	version := flag.Bool("version", false, "Show version")
	help := flag.Bool("help", false, "Show usage information")

	input := flag.String("input", "", "File path to a CQL file")
	output := flag.String("output", "", "Output file path, defaults to stdout")
	format := flag.String("format", "json", "Output format, either JSON or XML")
	repl := flag.Bool("repl", false, "Interactive REPL")

	flag.Parse()

	if version != nil && *version {
		fmt.Println("cql version", "1.0.0")
		os.Exit(0)
	}

	if help != nil && *help {
		showUsage()
		os.Exit(0)
	}

	if repl != nil && *repl {
		cql.NewREPL(cql.DefaultOptions()).Start()
		return
	}

	opt := cql.DefaultOptions()

	if strings.ToLower(*format) == "xml" {
		opt.Format = cql.XMLFormat
	}

	if input == nil || *input == "" {
		showUsage()
		return
	}

	fstream, err := FileStream(*input)

	if err != nil {
		log.Fatal(err)
	}

	prog, err := cql.Parse(fstream, opt)

	if err != nil {
		log.Fatal(err)
	}

	json, err := prog.JSON()

	if err != nil {
		log.Fatal(err)
	}

	if opt.Output == os.Stdout {
		fmt.Println(string(json))
		return
	}

	outputFile, err := os.OpenFile(*output, os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer outputFile.Close()

	n, err := outputFile.Write(json)

	if err != nil || n != len(json) {
		log.Fatal(err)
	}
}

func FileStream(path string) (*antlr.FileStream, error) {
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

	fstream, err := antlr.NewFileStream(path)

	if err != nil {
		return nil, err
	}

	return fstream, nil
}
