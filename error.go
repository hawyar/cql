package cql

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type CustomErr struct {
	Line, Column int
	Msg          string
}

func (c *CustomErr) Error() string {
	return fmt.Sprintf("%s %d:%d", c.Msg, c.Line, c.Column)
}

type ErrorListener struct {
	*antlr.DefaultErrorListener
	Errs []CustomErr
}

func (e *ErrorListener) SyntaxError(r antlr.Recognizer, d interface{}, line, column int, msg string, ex antlr.RecognitionException) {
	e.Errs = append(e.Errs, CustomErr{line, column, msg})
}

func (e *ErrorListener) ReportAmbiguity(r antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	line := r.GetTokenStream().LT(1).GetLine()
	fmt.Println("ReportAmbiguity: line: ", line)
}

func (e *ErrorListener) ReportAttemptingFullContext(r antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	fmt.Println("ReportAttemptingFullContext")
}

func (e *ErrorListener) ReportContextSensitivity(r antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
	fmt.Println("ReportContextSensitivity")

}
