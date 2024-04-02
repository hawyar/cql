package cql

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/hawyar/cql/internal"
)

const (
	CQLPrefix = "cql> "
	ExitCmd   = ".exit"
)

type Repl struct {
	scanner  *bufio.Scanner
	buff     *bytes.Buffer
	listener CQLListener
}

func NewREPL() *Repl {
	return &Repl{
		scanner:  bufio.NewScanner(os.Stdin),
		listener: *NewCQLListener(),
		buff:     new(bytes.Buffer),
	}
}

func (r *Repl) Start() {
	fmt.Println("Welcome to the CQL repl!\nType 'help' for help.")

	for {
		r.read()
		r.eval()
	}
}

func (r *Repl) read() {
	fmt.Print(CQLPrefix)

	if !r.scanner.Scan() {
		return
	}

	next := r.scanner.Text()

	if next == ExitCmd {
		fmt.Println("Exiting...")
		os.Exit(0)
	}

	r.buff.WriteString(next)
}

func (r *Repl) eval() {
	input := r.buff.String()

	fmt.Println("current buffer:", input)

	lexer := parser.NewcqlLexer(antlr.NewInputStream(input))

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := parser.NewcqlParser(stream)

	antlr.ParseTreeWalkerDefault.Walk(r.listener, parser.Library())

	if stream.LT(1).GetTokenType() == antlr.TokenEOF {
		fmt.Println("complete statement")
		r.reset()
	}
}

func (r *Repl) reset() {
	r.buff.Reset()
}
