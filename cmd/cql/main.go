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

func main() {
	version := flag.Bool("version", false, "Show version")
	help := flag.Bool("help", false, "Show usage information")

	input := flag.String("input", "", "File path to a CQL file ")
	format := flag.String("format", "json", "Output format, either JSON or XML")
	debug := flag.Bool("debug", false, "Print AST to stdout")
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
		cql.NewREPL(cql.DefaultConfig()).Start()
		return
	}

	config := cql.DefaultConfig()

	if strings.ToLower(*format) == "xml" {
		config.Format = cql.XMLFormat
	}

	if debug != nil && *debug {
		config.Debug = true
	}

	if input == nil || *input == "" {
		showUsage()
		os.Exit(1)
	}

	fstream, err := ReadFile(*input)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("input:", fstream.InputStream.GetText(0, fstream.Size()))

	// _, err = cql.Parse(fstream, config)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// os.Exit(0)
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

	fstream, err := antlr.NewFileStream(path)

	if err != nil {
		return nil, err
	}

	return fstream, nil
}

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
