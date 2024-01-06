package cql

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/hawyar/cql/internal/parser"
)

type node struct {
	val  antlr.RuleContext
	next *node
}

type stack struct {
	curr  *node
	count int
}

type Listener struct {
	parser.BasecqlListener
	ast    *AST
	stack  *stack
	errs   []antlr.ErrorNode
	Parser antlr.Parser
	Lexer  antlr.Lexer
}

func NewListener() *Listener {
	return &Listener{
		ast:   NewAST(),
		stack: newStack(),
	}
}

func newStack() *stack {
	return &stack{nil, 0}
}

func (s *stack) size() int {
	return s.count
}

func (s *stack) peek() antlr.RuleContext {
	if s.curr == nil {
		return nil
	}
	return s.curr.val
}

func (s *stack) push(val antlr.RuleContext) {
	if s.curr == nil {
		s.curr = &node{val, nil}
		s.count++
		return
	}

	s.curr = &node{val, s.curr}
	s.count++
}

func (s *stack) pop() antlr.RuleContext {
	if s.curr == nil {
		return nil
	}

	val := s.curr.val
	s.curr = s.curr.next
	s.count--
	return val
}

func (s *stack) isEmpty() bool {
	return s.curr == nil
}

func line(ctx antlr.ParserRuleContext) {
	fmt.Printf("%d:%d %s\n", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), ctx.GetText())
}

func (l *Listener) EnterLibrary(ctx *parser.LibraryContext) {
	l.stack.push(ctx)
}

func (l *Listener) ExitLibrary(ctx *parser.LibraryContext) {
	if len(l.errs) > 0 {
		fmt.Println("errors found")
		os.Exit(1)
	}

	l.stack.pop()

	jsonstr, _ := l.ast.MarshalJSON()

	fmt.Println(string(jsonstr))

	elm, _ := NewELM(l.ast)

	xmlF, err := os.Create("output.xml")

	if err != nil {
		fmt.Println(err)
	}

	encoder := xml.NewEncoder(xmlF)
	err = elm.MarshalXML(encoder, xml.StartElement{Name: xml.Name{Local: "library"}})

	if err != nil {
		fmt.Println(err)
	}

	jstr, err := elm.MarshalJSON()

	if err != nil {
		fmt.Println(err)
	}

	f, err := os.Create("output.json")

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	_, err = f.Write(jstr)

	if err != nil {
		fmt.Println(err)
	}
}

func (l *Listener) EnterLibraryDefinition(ctx *parser.LibraryDefinitionContext) {
	l.stack.push(ctx)
}

func (l *Listener) ExitLibraryDefinition(ctx *parser.LibraryDefinitionContext) {
	if l.stack.peek() == nil {
		return
	}

	switch l.stack.peek().(type) {
	case *parser.VersionSpecifierContext:
		version := l.stack.pop()
		l.ast.Library.Identifier.Version = strings.ReplaceAll(version.GetText(), "'", "")
		l.ExitLibraryDefinition(ctx)
	case *parser.IdentifierContext:
		ident := l.stack.pop()
		l.ast.Library.Identifier.Path = append(l.ast.Library.Identifier.Path, ident.GetText())
		l.ExitLibraryDefinition(ctx)

	case *parser.QualifiedIdentifierContext:
		ident := l.stack.pop()
		l.ast.Library.Identifier.Idenitifier = ident.GetText()
		l.ExitLibraryDefinition(ctx)
	default:
		l.stack.pop()
	}
}

func (l *Listener) EnterQualifiedIdentifier(ctx *parser.QualifiedIdentifierContext) {
	l.stack.push(ctx)
}

// func (l *Listener) ExitQualifiedIdentifier(ctx *parser.QualifiedIdentifierContext) {}

func (l *Listener) EnterVersionSpecifier(ctx *parser.VersionSpecifierContext) {
	l.stack.push(ctx)
}

func (l *Listener) EnterIdentifier(ctx *parser.IdentifierContext) {
	l.stack.push(ctx)
}

func (l *Listener) VisitErrorNode(node antlr.ErrorNode) {
	l.errs = append(l.errs, node)
	l.ExitLibrary(nil)
}

func (l *Listener) EnterUsingDefinition(ctx *parser.UsingDefinitionContext) {
	l.stack.push(ctx)
}

func (l *Listener) ExitUsingDefinition(ctx *parser.UsingDefinitionContext) {
	if l.stack.peek() == nil {
		return
	}

	usedef := UsingDefinition{}

	if _, ok := l.stack.peek().(*parser.VersionSpecifierContext); ok {
		version := l.stack.pop()
		usedef.Version = strings.ReplaceAll(version.GetText(), "'", "")
	}

	ident := l.stack.pop()
	usedef.Name = ident.GetText()
	l.ast.Library.UsingDefinitions = append(l.ast.Library.UsingDefinitions, usedef)
	l.stack.pop()

}

func (l *Listener) EnterIncludeDefinition(ctx *parser.IncludeDefinitionContext) {
	l.stack.push(ctx)
}

func (l *Listener) ExitIncludeDefinition(ctx *parser.IncludeDefinitionContext) {
	fmt.Println("exit include definition")

	if l.stack.peek() == nil {
		return
	}

	incdef := IncludeDefinition{}

	if _, ok := l.stack.peek().(*parser.IdentifierContext); ok {
		ident := l.stack.pop()
		incdef.Alias = ident.GetText()
	}

	if _, ok := l.stack.peek().(*parser.VersionSpecifierContext); ok {
		version := l.stack.pop()
		incdef.Version = strings.ReplaceAll(version.GetText(), "'", "")
	}

	if _, ok := l.stack.peek().(*parser.IdentifierContext); ok {
		for !l.stack.isEmpty() {
			if _, ok := l.stack.peek().(*parser.IdentifierContext); ok {
				ident := l.stack.pop()
				incdef.Path = append(incdef.Path, ident.GetText())
			} else {
				break
			}
		}
	}

	if _, ok := l.stack.peek().(*parser.QualifiedIdentifierContext); ok {
		ident := l.stack.pop()
		incdef.Name = ident.GetText()
	}

	l.ast.Library.IncludeDefinitions = append(l.ast.Library.IncludeDefinitions, incdef)
}

// skip *CodesystemDefinitionContext
