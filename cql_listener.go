package cql

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/hawyar/cql/internal/parser"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type node struct {
	val  antlr.RuleContext
	next *node
}

type stack struct {
	curr  *node
	count int
}

type repl struct {
	scanner *bufio.Scanner
	buffer  *bytes.Buffer
}

type Listener struct {
	parser.BasecqlListener
	ast   *AST
	stack *stack
	errs  []antlr.ErrorNode
}

func newREPL() *repl {
	return &repl{
		scanner: bufio.NewScanner(os.Stdin),
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

func NewListener() *Listener {
	return &Listener{
		ast:   newAST(),
		stack: newStack(),
	}
}

func (l *Listener) VisitErrorNode(node antlr.ErrorNode) {
	l.errs = append(l.errs, node)
	fmt.Println("error node", node.GetText())
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
	l.ast.Library.Identifier.Identifier = ctx.QualifiedIdentifier().GetText()

	if ctx.VersionSpecifier() != nil {
		l.ast.Library.Identifier.Version = removeQuotes(ctx.VersionSpecifier().GetText())
	}
	l.stack.pop()
}

func removeQuotes(s string) string {
	return s[1 : len(s)-1]
}
func (l *Listener) EnterUsingDefinition(ctx *parser.UsingDefinitionContext) {
	l.stack.push(ctx)
}

func (l *Listener) ExitUsingDefinition(ctx *parser.UsingDefinitionContext) {
	def := UsingDefinition{
		Name: ctx.ModelIdentifier().GetText(),
	}

	if ctx.VersionSpecifier() != nil {
		def.Version = removeQuotes(ctx.VersionSpecifier().GetText())
	}
	l.ast.Library.UsingDefinitions = append(l.ast.Library.UsingDefinitions, def)

	l.stack.pop()
}

func (l *Listener) EnterIncludeDefinition(ctx *parser.IncludeDefinitionContext) {
	l.stack.push(ctx)
}

func (l *Listener) ExitIncludeDefinition(ctx *parser.IncludeDefinitionContext) {
	def := IncludeDefinition{
		Name: ctx.QualifiedIdentifier().GetText(),
	}

	if ctx.VersionSpecifier() != nil {
		def.Version = removeQuotes(ctx.VersionSpecifier().GetText())
	}

	if ctx.LocalIdentifier() != nil {
		def.Alias = ctx.LocalIdentifier().GetText()
	}

	l.ast.Library.IncludeDefinitions = append(l.ast.Library.IncludeDefinitions, def)
	l.stack.pop()
}

func (l *Listener) EnterValuesetDefinition(ctx *parser.ValuesetDefinitionContext) {
	l.stack.push(ctx)
}

func (l *Listener) ExitValuesetDefinition(ctx *parser.ValuesetDefinitionContext) {
	set := ValuesetDefinition{
		Name:           removeQuotes(ctx.Identifier().GetText()),
		AccessModifier: string(PublicModifier),
	}

	if ctx.AccessModifier() != nil {
		set.AccessModifier = cases.Title(language.English).String(strings.ToLower(ctx.AccessModifier().GetText()))
	}

	if ctx.VersionSpecifier() != nil {
		set.Version = removeQuotes(ctx.VersionSpecifier().GetText())
	}

	if ctx.ValuesetId() != nil {
		ctx.ValuesetId().GetText()
		set.Id = removeQuotes(ctx.ValuesetId().GetText())
	}

	if ctx.Codesystems() != nil {
		for _, cs := range ctx.Codesystems().(*parser.CodesystemsContext).AllCodesystemIdentifier() {
			set.Codesystems = append(set.Codesystems, Codesystems{
				Identifier: cs.Identifier().GetText(),
			})

		}
	}
	l.ast.Library.ValuesetDefinitions = append(l.ast.Library.ValuesetDefinitions, set)
	l.stack.pop()
}

func (l *Listener) EnterParameterDefinition(ctx *parser.ParameterDefinitionContext) {
	l.stack.push(ctx)
}

func (l *Listener) ExitParameterDefinition(ctx *parser.ParameterDefinitionContext) {
	if ctx.TypeSpecifier() != nil {
		fmt.Println("type of param", ctx.TypeSpecifier().GetText())
	}

	param := Parameter{
		Name:           ctx.Identifier().GetText(),
		AccessModifier: string(PrivateModifier),
	}

	if ctx.AccessModifier() != nil {
		param.AccessModifier = cases.Title(language.English).String(strings.ToLower(ctx.AccessModifier().GetText()))
	}

	l.ast.Library.Parameters = append(l.ast.Library.Parameters, param)
	l.stack.pop()
}

func (l *Listener) EnterContextDefinition(ctx *parser.ContextDefinitionContext) {
	l.stack.push(ctx)
}

func (l *Listener) ExitContextDefinition(ctx *parser.ContextDefinitionContext) {
	fmt.Println("context definition", ctx.Identifier().GetText())
	l.ast.Library.Context = Context{
		Identifier: ctx.Identifier().GetText(),
	}

	l.stack.pop()
}
