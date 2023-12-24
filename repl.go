package cql

import (
	"bufio"
	"fmt"
	"os"
)

const (
	CQLPrefix      = "cql> "
	FHIRPathPrefix = "fhirpath> "
	Exit           = "exit"
)

type REPL struct {
	Scanner *bufio.Scanner
}

func NewREPL() *REPL {
	return &REPL{
		Scanner: bufio.NewScanner(os.Stdin),
	}
}

func (r *REPL) Start() {
	fmt.Println("Welcome to the CQL REPL!")
	fmt.Println("Type 'exit' to exit.")
	r.loop()
}

func (r *REPL) loop() {
	for {
		r.read()
	}
}

func (r *REPL) read() {
	fmt.Print(CQLPrefix)

	if !r.Scanner.Scan() {
		return
	}

	next := r.Scanner.Text()

	if next == Exit {
		fmt.Println("Exit command received, exiting...")
		os.Exit(0)
	}

	r.eval(next)
}

func (r *REPL) eval(input string) {
	fmt.Println("Evaluated:", input)
}
