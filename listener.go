package cql

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/hawyar/cql/internal/parser"
)

type node struct {
	value antlr.RuleContext
	next  *node
	depth int
}

type stack struct {
	curr *node
	cap  int
}

func (s *stack) size() int {
	return s.cap
}

func (s *stack) peek() antlr.RuleContext {
	if s.curr == nil {
		return nil
	}
	return s.curr.value
}

func (s *stack) push(n node) {
	if s.curr == nil {
		s.curr = &node{n.value, nil, n.depth}
		s.cap++
		return
	}

	s.curr = &node{n.value, s.curr, n.depth}
	s.cap++
}

func (s *stack) pop() antlr.RuleContext {
	if s.curr == nil {
		return nil
	}

	val := s.curr.value
	s.curr = s.curr.next
	s.cap--
	return val
}

func (s *stack) isEmpty() bool {
	return s.curr == nil
}

type Listener interface {
	*antlr.ParseTreeListener
	Entry() antlr.Token
	AST() *AST
	Program() *Program
}
type CQLListener struct {
	*parser.BasecqlListener
	ast     *AST
	program *Program
	stack   *stack
	entry   antlr.Token
}

func NewCQLListener(entry antlr.Token) *CQLListener {
	return &CQLListener{
		entry:   entry,
		ast:     &AST{},
		stack:   &stack{},
		program: &Program{},
	}
}

func (c *CQLListener) AST() *AST {
	return c.ast
}

func (c *CQLListener) EnterLibrary(ctx *parser.LibraryContext) {
	c.stack.push(node{ctx, nil, 0})
}

func (c *CQLListener) ExitLibrary(ctx *parser.LibraryContext) {
	if ctx.LibraryDefinition() != nil {
		fmt.Println("- LibraryDefinition: ", ctx.LibraryDefinition().GetText(), "start:", ctx.LibraryDefinition().GetStart().GetStart(), "end:", ctx.LibraryDefinition().GetStop().GetStop())

		if ctx.LibraryDefinition().QualifiedIdentifier() == nil {
			return
		}

		if !ctx.LibraryDefinition().QualifiedIdentifier().IsEmpty() {
			// libdef := LibraryDefinition{}
			// c.ast.Library.LibDef.Identifier
			c.ast.Library.Loc = LocFromContext(ctx.LibraryDefinition().QualifiedIdentifier())
		}

		c.stack.pop()

		if !c.stack.isEmpty() {
			panic("yo check yo stack")
		}
	}
}

func removeQuotes(s string) string {
	if len(s) > 0 {
		if s[0] == '"' && s[len(s)-1] == '"' {
			return s[1 : len(s)-1]
		}
	}
	return s
}

type Loc struct {
	Start loc
	End   loc
}
type loc struct {
	Line   int `json:"line"`
	Column int `json:"column"`
	Offset int `json:"offset"`
}

func LocFromContext(ctx antlr.ParserRuleContext) Loc {
	start := ctx.GetStart()
	stop := ctx.GetStop()
	return Loc{
		Start: loc{
			Line:   start.GetLine(),
			Column: start.GetColumn(),
			Offset: start.GetStart(),
		},
		End: loc{
			Line:   stop.GetLine(),
			Column: stop.GetColumn(),
			Offset: stop.GetStop(),
		},
	}
}
