package cql

import (
	"fmt"
	"io"
	"os"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/hawyar/cql/internal/parser"
)

type OutputFormat int

const (
	JSONFormat OutputFormat = iota
	XMLFormat
)

type ParseOptions struct {
	Format OutputFormat
	Output io.Writer
}

func DefaultOptions() ParseOptions {
	return ParseOptions{
		Format: JSONFormat,
		Output: os.Stdout,
	}
}

func Parse(chars antlr.CharStream, opt ParseOptions) (*AST, error) {
	if chars.Size() == 0 {
		return nil, fmt.Errorf("empty input")
	}

	lexer := parser.NewcqlLexer(chars)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := parser.NewcqlParser(stream)

	parser.BuildParseTrees = true

	// LibraryDefintion, Definition, or Statement followed by EOF otherwise panic
	entry := parser.GetTokenStream().LT(1)

	if entry.GetTokenType() != 1 {
		return nil, fmt.Errorf("invalid entry statement; expected LibraryDefinition, Definition, or Statement")
	}

	listener := NewCQLListener()

	antlr.ParseTreeWalkerDefault.Walk(listener, parser.Library())

	// TODO: check for errors before return

	return listener.AST(), nil
}
