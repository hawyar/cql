package cql

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/hawyar/cql/internal/parser"
)

type Listener struct {
	parser.BasecqlListener
	ast   *AST
	stack *stack
	errs  []antlr.ErrorNode
}

func NewListener() *Listener {
	return &Listener{
		ast:   &AST{},
		stack: newStack(),
	}
}

func (l *Listener) EnterLibrary(ctx *parser.LibraryContext) {
	l.stack.push(ctx, ctx.GetChildCount())
}

func (l *Listener) ExitLibrary(ctx *parser.LibraryContext) {
	ast := l.ast

	if ctx.LibraryDefinition() != nil {
		def := LibraryDefinition{}

		if ctx.LibraryDefinition().QualifiedIdentifier().Identifier() != nil {
			// def.Identifier = ctx.LibraryDefinition().QualifiedIdentifier().GetText()

			// print the full
			fmt.Println("identifier:", ctx.LibraryDefinition().QualifiedIdentifier().GetText())

			// print the quoted identifier
			if ctx.LibraryDefinition().QualifiedIdentifier().Identifier().QUOTEDIDENTIFIER() != nil {
				for ctx.LibraryDefinition().QualifiedIdentifier().Identifier().QUOTEDIDENTIFIER().GetChildCount() > 0 {
					fmt.Println("quoted identifier:", ctx.LibraryDefinition().QualifiedIdentifier().Identifier().QUOTEDIDENTIFIER().GetText())
				}
			}

			// print the delimited identifier
			if ctx.LibraryDefinition().QualifiedIdentifier().Identifier().DELIMITEDIDENTIFIER() != nil {
				fmt.Println("delimited identifier:", ctx.LibraryDefinition().QualifiedIdentifier().Identifier().DELIMITEDIDENTIFIER().GetText())
			}

		}
		ast.Library.LibDef = def
	}

	if len(l.errs) > 0 {
		fmt.Println("errors:")
		for _, err := range l.errs {
			fmt.Println(err.GetText())
		}
		os.Exit(1)
	}

	l.stack.pop()
	jsonstr, _ := l.ast.MarshalJSON()
	fmt.Println(string(jsonstr))
}

func stripQuotes(s string) string {
	return s[1 : len(s)-1]
}

func line(ctx antlr.ParserRuleContext) {
	fmt.Printf("%d:%d %s\n", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), ctx.GetText())
}
