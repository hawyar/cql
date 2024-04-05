package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/hawyar/cql"
)

func main() {
	version := flag.Bool("version", false, "Show version")
	help := flag.Bool("help", false, "Show usage information")
	expr := flag.String("e", "", "CQL expression to evaluate")
	fileinput := flag.String("f", "", "CQL file to parse")

	flag.Parse()

	if version != nil && *version {
		fmt.Println("cql version", "1.0.0")
		os.Exit(0)
	}

	if help != nil && *help {
		usage()
		os.Exit(0)
	}

	if expr != nil && *expr != "" {
		in := antlr.NewInputStream(*expr)

		if in.Size() == 0 {
			log.Fatal("empty input")
		}

		ast, perr := cql.Parse(in)

		if perr != nil {
			log.Fatal(perr)
		}

		_, err := ast.JSON()

		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println(astj)
		return
	}

	if fileinput != nil && *fileinput != "" {
		fstream, err := FileStream(*fileinput)

		if err != nil {
			log.Fatal(err)
		}

		ast, parseErr := cql.Parse(fstream)

		if len(parseErr) > 0 {
			for _, e := range parseErr {
				fmt.Println("Error:", e.Error())
			}
		}

		astj, err := ast.JSON()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(astj))
		return
	}

	usage()
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

func usage() {
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
