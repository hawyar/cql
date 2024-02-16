package cql

import (
	"fmt"
	"io"
	"os"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/hawyar/cql/internal/parser"
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

func Parse(chars antlr.CharStream, opt ParseOptions) (*Program, error) {
	if chars.Size() == 0 {
		return nil, fmt.Errorf("empty input")
	}

	lexer := parser.NewcqlLexer(chars)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := parser.NewcqlParser(stream)

	parser.BuildParseTrees = true

	entry := parser.GetTokenStream().LT(1)

	if entry.GetTokenType() != 1 {
		return nil, fmt.Errorf("invalid input")
	}

	listener := NewCQLListener(entry)

	antlr.ParseTreeWalkerDefault.Walk(listener, parser.Library())

	// TODO check for errors before return
	return listener.program, nil
}
