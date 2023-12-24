package main

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/hawyar/cql/cql"
)

type CQLListener struct {
	*parser.BasecqlListener
}

type Library struct {
	Name       string
	Version    string
	Statements []Statement
}

type Statement interface{}

type AST struct {
	Library *Library
}

func NewAST() *AST {
	return &AST{}
}

func (l *CQLListener) EnterLibrary(ctx *parser.LibraryContext) {
	fmt.Printf("Enter Library: %s\n", ctx.GetText())
}

func (l *CQLListener) EnterLibraryDefinition(ctx *parser.LibraryDefinitionContext) {
	fmt.Printf("Enter LibraryDefinition: %s\n", ctx.GetText())
}

func (l *CQLListener) EnterLibraryIdentifier(ctx *parser.LibraryIdentifierContext) {
	fmt.Printf("Enter LibraryIdentifier: %s\n", ctx.GetText())
}

func (l *CQLListener) EnterIncludeDefinition(ctx *parser.IncludeDefinitionContext) {
	fmt.Printf("Enter IncludeDefinition: %s\n", ctx.GetText())
}

func (l *CQLListener) EnterUsingDefinition(ctx *parser.UsingDefinitionContext) {
	fmt.Printf("Enter UsingDefinition: %s\n", ctx.GetText())
}

func (l *CQLListener) EnterVersionSpecifier(ctx *parser.VersionSpecifierContext) {
	fmt.Printf("Enter VersionSpecifier: %s\n", ctx.GetText())
}

func (l *CQLListener) EnterLocalIdentifier(ctx *parser.LocalIdentifierContext) {
	fmt.Printf("Enter LocalIdentifier: %s\n", ctx.GetText())
}

func (l *CQLListener) EnterFunctionDefinition(ctx *parser.FunctionDefinitionContext) {
	fmt.Printf("Enter FunctionDefinition: %s\n", ctx.GetText())
}

func (l *CQLListener) EnterExpressionDefinition(ctx *parser.ExpressionDefinitionContext) {
	fmt.Printf("Enter ExpressionDefinition: %s\n", ctx.GetText())
}

func (l *CQLListener) EnterContextDefinition(ctx *parser.ContextDefinitionContext) {
	line := ctx.GetStart().GetLine()
	fmt.Printf("Enter ContextDefinition: %s\n", ctx.GetText())
	fmt.Println("Line: ", line)
}

func (l *CQLListener) ExitContextDefinition(ctx *parser.ContextDefinitionContext) {
	fmt.Printf("Exit ContextDefinition: %s\n", ctx.GetText())
}

func main() {
	input, err := antlr.NewFileStream("./example/simple.cql")

	if err != nil {
		panic(err)
	}

	lexer := parser.NewcqlLexer(input)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewcqlParser(stream)

	listener := &CQLListener{}

	antlr.ParseTreeWalkerDefault.Walk(listener, p.Library())

	// tokens := stream.GetAllTokens()

	// for _, token := range tokens {
	// 	fmt.Println(token)
	// }
}
