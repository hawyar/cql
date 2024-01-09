package cql

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
	fhirpath "github.com/hawyar/cql/internal/fhirpath"
)

const (
	FHIRPathPrefix = "fhirpath> "
)

type FHIRPathRepl struct {
	scanner     *bufio.Scanner
	listener    *FHIRPathListener
	buff        *bytes.Buffer
	tokenStream *antlr.CommonTokenStream
}

func NewFHIRPathRepl() *FHIRPathRepl {
	return &FHIRPathRepl{
		scanner: bufio.NewScanner(os.Stdin),
		buff:    bytes.NewBufferString(""),
	}
}

func (r *FHIRPathRepl) Start() {
	fmt.Println("Welcome to the CQL repl!\nType 'help' for help.")

	for {
		r.read()
		r.eval()
	}
}

func (r *FHIRPathRepl) read() {
	fmt.Print(FHIRPathPrefix)

	if !r.scanner.Scan() {
		return
	}

	next := r.scanner.Text()

	if next == ExitCmd {
		fmt.Println("exiting...")
		os.Exit(0)
	}

	lexer := fhirpath.NewfhirpathLexer(antlr.NewInputStream(next))

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	r.tokenStream = tokenStream

	listener := NewFHIRPathListener()

	r.listener = listener

	r.buff.WriteString(next)
}

func (r *FHIRPathRepl) eval() {
	lexer := fhirpath.NewfhirpathLexer(antlr.NewInputStream(r.buff.String()))

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	r.tokenStream = tokenStream

	// add a EOF token to the end of the stream
	parser := fhirpath.NewfhirpathParser(tokenStream)

	// add a EOF token to the end of the stream

	// if ends with EOF, then we have a complete expression
	// if r.tokenStream.LT(1).GetTokenType() == antlr.TokenEOF {
	antlr.ParseTreeWalkerDefault.Walk(r.listener, parser.Expression())
	r.reset()
	// }
}

func (r *FHIRPathRepl) reset() {
	r.buff.Reset()
}
