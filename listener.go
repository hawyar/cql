package cql

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/hawyar/cql/internal/parser"
)

type CQLListener struct {
	*parser.BasecqlListener
	ast *AST
}

type Loc struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

func NewCQLListener() *CQLListener {
	return &CQLListener{
		ast: &AST{},
	}
}

func (c *CQLListener) AST() *AST {
	return c.ast
}

func (c *CQLListener) EnterLibrary(ctx *parser.LibraryContext) {
	c.ast.Library.Loc = LocFromContext(ctx)
}

func (c *CQLListener) ExitLibrary(ctx *parser.LibraryContext) {
	if c.ast == nil {
		return
	}

	astj, _ := c.ast.JSON()
	fmt.Println(string(astj))
}

func (c *CQLListener) EnterLibraryDefinition(ctx *parser.LibraryDefinitionContext) {
	if ctx.IsEmpty() || ctx == nil {
		return
	}

	def := LibraryDefinition{}

	if ctx.QualifiedIdentifier() != nil {
		def.QualifiedIdentifier = QualifiedIdentifier{
			Identifier: strings.Trim(ctx.QualifiedIdentifier().GetText(), "\""),
		}
	}

	if ctx.VersionSpecifier() != nil {
		def.Version = strings.Trim(ctx.VersionSpecifier().GetText(), "'")
	}

	c.ast.Library.LibraryDefinition = def
}

func (c *CQLListener) EnterContextDefinition(ctx *parser.ContextDefinitionContext) {
	if ctx.Identifier() != nil {
		c.ast.Library.ContextDefinition = Context(strings.Trim(ctx.Identifier().GetText(), "\""))
	}
}

func (c *CQLListener) EnterIncludeDefinition(ctx *parser.IncludeDefinitionContext) {
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

func (c *CQLListener) EnterUsingDefinition(ctx *parser.UsingDefinitionContext) {
	def := UsingDefinition{}

	if ctx.ModelIdentifier() != nil {
		def.Identifier = strings.Trim(ctx.ModelIdentifier().GetText(), "\"")
	}

	if ctx.VersionSpecifier() != nil {
		def.Version = strings.Trim(ctx.VersionSpecifier().GetText(), "'")
	}

	c.ast.Library.UsingDefinitions = append(c.ast.Library.UsingDefinitions, def)
}

func (c *CQLListener) EnterValuesetDefinition(ctx *parser.ValuesetDefinitionContext) {
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
		def.Id = strings.Trim(ctx.ValuesetId().GetText(), "\"")
	}

	if ctx.Codesystems() != nil {
		codesystem := Codesystems{
			QualifiedIdentifier: make([]QualifiedIdentifier, 0),
		}

		for _, qid := range ctx.Codesystems().AllCodesystemIdentifier() {
			qi := QualifiedIdentifier{
				Identifier: strings.Trim(qid.GetText(), "\""),
			}
			codesystem.QualifiedIdentifier = append(codesystem.QualifiedIdentifier, qi)
		}
		def.Codesystems = append(def.Codesystems, codesystem)
	}

	c.ast.Library.ValuesetDefinitions = append(c.ast.Library.ValuesetDefinitions, def)
}

func (c *CQLListener) EnterParameterDefinition(ctx *parser.ParameterDefinitionContext) {
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

func (c *CQLListener) EnterCodeDefinition(ctx *parser.CodeDefinitionContext) {
	def := CodeDefinition{}

	if ctx.AccessModifier() != nil {
		def.AccessModifier = AccessModifier(strings.Trim(ctx.AccessModifier().GetText(), "\""))
	}

	if ctx.Identifier() != nil {
		def.Identifier = strings.Trim(ctx.Identifier().GetText(), "\"")
	}

	if ctx.CodeId() != nil {
		def.CodeId = strings.Trim(ctx.CodeId().GetText(), "'")
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

func (c *CQLListener) EnterConceptDefinition(ctx *parser.ConceptDefinitionContext) {
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

func (c *CQLListener) EnterCastExpression(ctx *parser.CastExpressionContext) {
	fmt.Println("cast expression: ", ctx.GetText())
}

func LocFromContext(ctx antlr.ParserRuleContext) Loc {
	return Loc{
		Start: ctx.GetStart().GetStart(),
		End:   ctx.GetStop().GetStop() + 1,
	}
}

func LocFromToken(token antlr.Token) Loc {
	return Loc{
		Start: token.GetStart(),
		End:   token.GetStop() + 1,
	}
}
