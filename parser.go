package cql

import (
	"github.com/antlr4-go/antlr/v4"
	parser "github.com/hawyar/cql/internal"
)

func Parse(chars antlr.CharStream) (*AST, []CustomErr) {
	if chars.Size() == 0 {
		return nil, []CustomErr{{0, 0, "empty input"}}
	}

	lexer := parser.NewcqlLexer(chars)

	// lexer.RemoveErrorListeners()
	// lexer.AddErrorListener()

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parseErr := &ErrorListener{}

	parser := parser.NewcqlParser(stream)

	parser.GetInterpreter().SetPredictionMode(antlr.PredictionModeLL)
	parser.BuildParseTrees = true

	parser.RemoveErrorListeners()
	parser.AddErrorListener(parseErr)

	// LibraryDefintion, Definition, or Statement followed by EOF otherwise panic
	entry := parser.GetTokenStream().LT(1)

	if entry.GetTokenType() != 1 {
		return nil, []CustomErr{{entry.GetLine(), entry.GetColumn(), "unexpected token: " + entry.GetText() + "expected LibraryDefinition, Definition, or Statement"}}
	}

	listener := NewCQLListener()

	antlr.ParseTreeWalkerDefault.Walk(listener, parser.Library())

	if len(parseErr.Errs) > 0 {
		return nil, parseErr.Errs
	}

	return listener.AST(), nil
}
