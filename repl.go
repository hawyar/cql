package cql

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/hawyar/cql/internal/parser"
)

const (
	CQLPrefix      = "cql> "
	FHIRPathPrefix = "fhirpath> "
	Exit           = ".exit"
)

type REPL struct {
	scanner *bufio.Scanner
	builder *Listener
	buff    *bytes.Buffer
}

func NewREPL() *REPL {
	return &REPL{
		scanner: bufio.NewScanner(os.Stdin),
		builder: NewListener(),
		buff:    bytes.NewBufferString(""),
	}
}

func (r *REPL) Start() {
	fmt.Println("Welcome to the CQL repl!\nType 'help' for help.")

	for {
		r.read()
		r.eval()
	}
}

func (r *REPL) read() {
	fmt.Print(CQLPrefix)

	if !r.scanner.Scan() {
		return
	}

	next := r.scanner.Text()

	if next == Exit {
		fmt.Println("Exiting...")
		os.Exit(0)
	}

	r.buff.WriteString(next)
}

func (r *REPL) eval() {
	input := r.buff.String()

	fmt.Println("current buffer:", input)

	lexer := parser.NewcqlLexer(antlr.NewInputStream(input))

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := parser.NewcqlParser(stream)

	antlr.ParseTreeWalkerDefault.Walk(r.builder, parser.Expression())
	// fmt.Println("last token:", stream.LT(-1).GetText())

	if stream.LT(1).GetTokenType() == antlr.TokenEOF {
		fmt.Println("complete statement")
		r.reset()
	}
	// otherwise, display another line of prompt and read more input
	// fmt.Println(CQLPrefix)
}

func (r *REPL) reset() {
	r.buff.Reset()
}
