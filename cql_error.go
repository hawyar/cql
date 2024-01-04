package cql

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type ErrorListener struct {
	*antlr.DefaultErrorListener
}

func NewErrorListener() *antlr.DefaultErrorListener {
	return antlr.NewDefaultErrorListener()
}

func NewConsoleErrorListener() *antlr.ConsoleErrorListener {
	return antlr.NewConsoleErrorListener()
}

func (e *ErrorListener) SyntaxError(r antlr.Recognizer, d interface{}, line, column int, msg string, ex antlr.RecognitionException) {
	fmt.Println("Syntax error:", msg)
}

func (e *ErrorListener) ReportAmbiguity(r antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	fmt.Println("ReportAmbiguity")
}

func (e *ErrorListener) ReportAttemptingFullContext(r antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	fmt.Println("ReportAttemptingFullContext")
}

func (e *ErrorListener) ReportContextSensitivity(r antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
	fmt.Println("ReportContextSensitivity")
}
