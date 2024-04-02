package cql

import (
	"fmt"
	"log"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/hawyar/cql/internal"
)

type CQLListener struct {
	*parser.BasecqlListener
	ast *AST
	// antlr already generates a parse tree, but we use this stack
	// to make sure we enter and exit some terminal nodes correctly.
	stack stack
}

type Loc struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

func NewCQLListener() *CQLListener {
	return &CQLListener{
		ast:   &AST{},
		stack: stack{},
	}
}

func (c *CQLListener) AST() *AST {
	return c.ast
}

func (c *CQLListener) EnterLibrary(ctx *parser.LibraryContext) {
	// c.ast.Library.Loc = locFromContext(ctx)
	c.stack.push(ctx)
}

func (c *CQLListener) ExitLibrary(ctx *parser.LibraryContext) {
	c.stack.pop()

	if !c.stack.isEmpty() {
		log.Fatal("stack is not empty, make sure terminal nodes are visited correctly")
	}
}

func (c *CQLListener) EnterLibraryDefinition(ctx *parser.LibraryDefinitionContext) {
	c.stack.push(ctx)

	if ctx.IsEmpty() || ctx == nil {
		return
	}

	def := LibraryDefinition{}

	// no qualifier with single indentifier
	if ctx.QualifiedIdentifier().AllQualifier() != nil && len(ctx.QualifiedIdentifier().AllQualifier()) == 0 {
		def.QualifiedIdentifier = QualifiedIdentifier{
			Identifier: strings.Trim(ctx.QualifiedIdentifier().GetText(), "\""),
		}
	}

	// single identifier with possibly more than one qualifier
	if ctx.QualifiedIdentifier().AllQualifier() != nil {
		def.QualifiedIdentifier = QualifiedIdentifier{
			Identifier: ctx.QualifiedIdentifier().Identifier().GetText(),
		}

		qualifiers := make([]string, 0)

		for _, q := range ctx.QualifiedIdentifier().AllQualifier() {
			if q.Identifier().IDENTIFIER() != nil {
				qualifiers = append(qualifiers, q.Identifier().IDENTIFIER().GetText())
			}

			if q.Identifier().QUOTEDIDENTIFIER() != nil {
				qualifiers = append(qualifiers, q.Identifier().QUOTEDIDENTIFIER().GetText())
			}

			if q.Identifier().DELIMITEDIDENTIFIER() != nil {
				qualifiers = append(qualifiers, q.Identifier().DELIMITEDIDENTIFIER().GetText())
			}
		}

		if len(qualifiers) != 0 {
			def.QualifiedIdentifier.Qualifiers = qualifiers
		}
	}

	if ctx.VersionSpecifier() != nil {
		def.Version = strings.Trim(ctx.VersionSpecifier().GetText(), "'")
	}

	c.ast.Library.LibraryDefinition = def
}

func (c *CQLListener) ExitLibraryDefinition(ctx *parser.LibraryDefinitionContext) {
	c.stack.pop()
}

func (c *CQLListener) EnterContextDefinition(ctx *parser.ContextDefinitionContext) {
	c.stack.push(ctx)

	if ctx.Identifier() != nil {
		c.ast.Library.ContextDefinition = ContextDefinition((strings.Trim(ctx.Identifier().GetText(), "\"")))
	}
}

func (c *CQLListener) ExitContextDefinition(ctx *parser.ContextDefinitionContext) {
	// if no context is specified in the library, and the model has not declared a default context
	// then set default context to Unfiltered.
	if c.ast.Library.ContextDefinition == "" {
		c.ast.Library.ContextDefinition = "Unfiltered"
	}
	c.stack.pop()
}

func (c *CQLListener) EnterIncludeDefinition(ctx *parser.IncludeDefinitionContext) {
	c.stack.push(ctx)

	def := IncludeDefinition{}

	if ctx.QualifiedIdentifier() != nil {
		def.Name = strings.Trim(ctx.QualifiedIdentifier().GetText(), "\"")
	}

	if ctx.VersionSpecifier() != nil {
		def.Version = strings.Trim(ctx.VersionSpecifier().GetText(), "'")
	}

	if ctx.LocalIdentifier() != nil {
		def.Alias = strings.Trim(ctx.LocalIdentifier().GetText(), "\"")
	}

	c.ast.Library.IncludeDefinitions = append(c.ast.Library.IncludeDefinitions, def)
}

func (c *CQLListener) ExitIncludeDefinition(ctx *parser.IncludeDefinitionContext) {
	c.stack.pop()
}

func (c *CQLListener) EnterUsingDefinition(ctx *parser.UsingDefinitionContext) {
	c.stack.push(ctx)

	def := UsingDefinition{}

	if ctx.ModelIdentifier() != nil {
		def.Identifier = strings.Trim(ctx.ModelIdentifier().GetText(), "\"")
	}

	if ctx.VersionSpecifier() != nil {
		def.Version = strings.Trim(ctx.VersionSpecifier().GetText(), "'")
	}

	c.ast.Library.UsingDefinitions = append(c.ast.Library.UsingDefinitions, def)
}

func (c *CQLListener) ExitUsingDefinition(ctx *parser.UsingDefinitionContext) {
	c.stack.pop()
}

func (c *CQLListener) EnterValuesetDefinition(ctx *parser.ValuesetDefinitionContext) {
	c.stack.push(ctx)

	def := ValuesetDefinition{}

	if ctx.Identifier() != nil {
		def.Name = strings.Trim(ctx.Identifier().GetText(), "\"")
	}

	if ctx.VersionSpecifier() != nil {
		def.Version = strings.Trim(ctx.VersionSpecifier().GetText(), "\"")
	}

	if ctx.AccessModifier() != nil {
		def.AccessModifier = strings.Trim(ctx.AccessModifier().GetText(), "\"")
	}

	if ctx.ValuesetId() != nil {
		def.ID = strings.Trim(ctx.ValuesetId().GetText(), "\"")
	}

	if ctx.Codesystems() != nil {
		codesystem := Codesystems{
			QualifiedIdentifier: make([]QualifiedIdentifier, 0),
		}

		for _, id := range ctx.Codesystems().AllCodesystemIdentifier() {
			qi := QualifiedIdentifier{
				Identifier: strings.Trim(id.GetText(), "\""),
			}
			codesystem.QualifiedIdentifier = append(codesystem.QualifiedIdentifier, qi)
		}
		def.Codesystems = append(def.Codesystems, codesystem)
	}

	c.ast.Library.ValuesetDefinitions = append(c.ast.Library.ValuesetDefinitions, def)
}

func (c *CQLListener) ExitValuesetDefinition(ctx *parser.ValuesetDefinitionContext) {
	c.stack.pop()
}

// codesystems may appear without a valueset definition
func (c *CQLListener) EnterCodesystems(ctx *parser.CodesystemsContext) {
	c.stack.push(ctx)

	codesystem := Codesystems{}

	for _, id := range ctx.AllCodesystemIdentifier() {
		qi := QualifiedIdentifier{
			Identifier: strings.Trim(id.GetText(), "\""),
		}
		codesystem.QualifiedIdentifier = append(codesystem.QualifiedIdentifier, qi)
	}

	c.ast.Library.CodesystemDefinitions = append(c.ast.Library.CodesystemDefinitions, codesystem)
}

func (c *CQLListener) ExitCodesystems(ctx *parser.CodesystemsContext) {
	c.stack.pop()
}

func (c *CQLListener) EnterParameterDefinition(ctx *parser.ParameterDefinitionContext) {
	c.stack.push(ctx)

	def := ParameterDefinition{}

	if ctx.AccessModifier() != nil {
		def.AccessModifier = AccessModifier(strings.Trim(ctx.AccessModifier().GetText(), "\""))
	}

	if ctx.Identifier() != nil {
		def.Name = strings.Trim(ctx.Identifier().GetText(), "\"")
	}

	if ctx.TypeSpecifier() != nil {
		// fmt.Println("type specifier: ", ctx.TypeSpecifier().GetText())
		def.TypeSpecifier = strings.Trim(ctx.TypeSpecifier().GetText(), "\"")
	}

	// TODO: handle expression
	if ctx.Expression() != nil {
		// fmt.Println("expression: ", ctx.Expression().GetText())
	}

	c.ast.Library.ParamterDefinitions = append(c.ast.Library.ParamterDefinitions, def)
}

func (c *CQLListener) ExitParameterDefinition(ctx *parser.ParameterDefinitionContext) {
	c.stack.pop()
}

func (c *CQLListener) EnterCodeDefinition(ctx *parser.CodeDefinitionContext) {
	c.stack.push(ctx)

	def := CodeDefinition{}

	if ctx.AccessModifier() != nil {
		def.AccessModifier = AccessModifier(strings.Trim(ctx.AccessModifier().GetText(), "\""))
	}

	if ctx.Identifier() != nil {
		def.Identifier = strings.Trim(ctx.Identifier().GetText(), "\"")
	}

	if ctx.CodeId() != nil {
		def.CodeID = strings.Trim(ctx.CodeId().GetText(), "'")
	}

	if ctx.CodesystemIdentifier() != nil {
		def.QualifiedIdentifier = QualifiedIdentifier{
			Identifier: strings.Trim(ctx.CodesystemIdentifier().GetText(), "\""),
		}
	}

	if ctx.DisplayClause() != nil {
		if strings.HasPrefix(ctx.DisplayClause().GetText(), "display") {
			def.DisplayClause = DisplayClause(strings.Trim(strings.Trim(ctx.DisplayClause().GetText(), "display"), "'"))
		}
	}

	c.ast.Library.CodeDefinitions = append(c.ast.Library.CodeDefinitions, def)
}

func (c *CQLListener) ExitCodeDefinition(ctx *parser.CodeDefinitionContext) {
	c.stack.pop()
}

func (c *CQLListener) EnterConceptDefinition(ctx *parser.ConceptDefinitionContext) {
	c.stack.push(ctx)

	concept := ConceptDefinition{}

	if ctx.AccessModifier() != nil {
		concept.AccessModifier = AccessModifier(strings.Trim(ctx.AccessModifier().GetText(), "\""))
	}

	if ctx.Identifier() != nil {
		concept.Identifier = strings.Trim(ctx.Identifier().GetText(), "\"")
	}

	if ctx.DisplayClause() != nil {
		if strings.HasPrefix(ctx.DisplayClause().GetText(), "display") {
			concept.DisplayClause = DisplayClause(strings.Trim(strings.Trim(ctx.DisplayClause().GetText(), "display"), "'"))
		}
	}

	if ctx.AllCodeIdentifier() != nil {
		for _, ident := range ctx.AllCodeIdentifier() {
			qi := QualifiedIdentifier{
				Identifier: strings.Trim(ident.GetText(), "\""),
			}

			concept.QualifiedIdentifier = append(concept.QualifiedIdentifier, qi)
		}
	}

	c.ast.Library.ConceptDefinitions = append(c.ast.Library.ConceptDefinitions, concept)
}

func (c *CQLListener) ExitConceptDefinition(ctx *parser.ConceptDefinitionContext) {
	c.stack.pop()
}

func (c *CQLListener) EnterExpressionDefinition(ctx *parser.ExpressionDefinitionContext) {
	// c.stack.push(ctx)
	fmt.Println("expression definition: ", ctx.GetText())
}

type node struct {
	ctx antlr.ParserRuleContext
}

type stack struct {
	items []node
	size  int
}

func (s *stack) push(ctx antlr.ParserRuleContext) {
	s.items = append(s.items, node{ctx: ctx})
	s.size++
}

func (s *stack) pop() node {
	if s.size == 0 {
		return node{}
	}

	n := s.items[s.size-1]
	s.items = s.items[:s.size-1]
	s.size--

	return n
}

func (s *stack) peek() node {
	if s.size == 0 {
		return node{}
	}

	return s.items[s.size-1]
}

func (s *stack) isEmpty() bool {
	return s.size == 0
}

func (s *stack) clear() {
	s.items = []node{}
	s.size = 0
}
