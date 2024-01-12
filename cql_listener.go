package cql

import (
	"encoding/xml"
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
		ast:   newAST(),
		stack: newStack(),
	}
}

// func (l *Listener) VisitErrorNode(node antlr.ErrorNode) {}

func (l *Listener) EnterLibrary(ctx *parser.LibraryContext) {
	for _, v := range ctx.AllValuesetDefinition() {
		fmt.Println("valueset", v.GetText())
	}

	for _, v := range ctx.AllStatement() {
		if v.ContextDefinition() != nil {
			fmt.Println("context", v.ContextDefinition().GetText())
		}
	}

	l.stack.push(ctx)
}

func (l *Listener) ExitLibrary(ctx *parser.LibraryContext) {
	if len(l.errs) > 0 {
		fmt.Println("errors found")
		os.Exit(1)
	}

	l.stack.pop()

	// jsonstr, _ := l.ast.MarshalJSON()
	// fmt.Println(string(jsonstr))

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

func stripQuotes(s string) string {
	return s[1 : len(s)-1]
}

func line(ctx antlr.ParserRuleContext) {
	fmt.Printf("%d:%d %s\n", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), ctx.GetText())
}

// func (l *Listener) EnterLibraryDefinition(ctx *parser.LibraryDefinitionContext) {
// 	l.stack.push(ctx)
// }

// func (l *Listener) ExitLibraryDefinition(ctx *parser.LibraryDefinitionContext) {
// 	ident := LibraryDefinition{
// 		Identifier: ctx.QualifiedIdentifier().GetText(),
// 	}

// 	if ctx.VersionSpecifier() != nil {
// 		ident.Version = stripQuotes(ctx.VersionSpecifier().GetText())
// 	}

// 	l.stack.pop()
// }

// func (l *Listener) EnterUsingDefinition(ctx *parser.UsingDefinitionContext) {
// 	l.stack.push(ctx)
// }

// func (l *Listener) ExitUsingDefinition(ctx *parser.UsingDefinitionContext) {
// 	def := UsingDefinition{
// 		Name: ctx.ModelIdentifier().GetText(),
// 	}

// 	if ctx.VersionSpecifier() != nil {
// 		def.Version = stripQuotes(ctx.VersionSpecifier().GetText())
// 	}
// 	l.ast.Library.UsingDefinitions = append(l.ast.Library.UsingDefinitions, def)

// 	l.stack.pop()
// }

// func (l *Listener) EnterIncludeDefinition(ctx *parser.IncludeDefinitionContext) {
// 	l.stack.push(ctx)
// }

// func (l *Listener) ExitIncludeDefinition(ctx *parser.IncludeDefinitionContext) {
// 	def := IncludeDefinition{
// 		Name: ctx.QualifiedIdentifier().GetText(),
// 	}

// 	if ctx.VersionSpecifier() != nil {
// 		def.Version = stripQuotes(ctx.VersionSpecifier().GetText())
// 	}

// 	if ctx.LocalIdentifier() != nil {
// 		def.Alias = ctx.LocalIdentifier().GetText()
// 	}

// 	l.ast.Library.IncludeDefinitions = append(l.ast.Library.IncludeDefinitions, def)
// 	l.stack.pop()
// }

// func (l *Listener) EnterValuesetDefinition(ctx *parser.ValuesetDefinitionContext) {
// 	l.stack.push(ctx)
// }

// func (l *Listener) ExitValuesetDefinition(ctx *parser.ValuesetDefinitionContext) {
// 	set := ValuesetDefinition{
// 		Name:           stripQuotes(ctx.Identifier().GetText()),
// 		AccessModifier: string(PublicModifier),
// 	}

// 	if ctx.AccessModifier() != nil {
// 		set.AccessModifier = cases.Title(language.English).String(strings.ToLower(ctx.AccessModifier().GetText()))
// 	}

// 	if ctx.VersionSpecifier() != nil {
// 		set.Version = stripQuotes(ctx.VersionSpecifier().GetText())
// 	}

// 	if ctx.ValuesetId() != nil {
// 		ctx.ValuesetId().GetText()
// 		set.Id = stripQuotes(ctx.ValuesetId().GetText())
// 	}

// 	if ctx.Codesystems() != nil {
// 		for _, cs := range ctx.Codesystems().(*parser.CodesystemsContext).AllCodesystemIdentifier() {
// 			set.Codesystems = append(set.Codesystems, Codesystems{
// 				Identifier: cs.Identifier().GetText(),
// 			})

// 		}
// 	}
// 	l.ast.Library.ValuesetDefinitions = append(l.ast.Library.ValuesetDefinitions, set)
// 	l.stack.pop()
// }

// func (l *Listener) EnterParameterDefinition(ctx *parser.ParameterDefinitionContext) {
// 	l.stack.push(ctx)
// }

// func (l *Listener) ExitParameterDefinition(ctx *parser.ParameterDefinitionContext) {
// 	if ctx.TypeSpecifier() != nil {
// 		fmt.Println("type of param", ctx.TypeSpecifier().GetText())
// 	}

// 	param := Parameter{
// 		Name:           ctx.Identifier().GetText(),
// 		AccessModifier: string(PrivateModifier),
// 	}

// 	if ctx.AccessModifier() != nil {
// 		param.AccessModifier = cases.Title(language.English).String(strings.ToLower(ctx.AccessModifier().GetText()))
// 	}

// 	l.ast.Library.Parameters = append(l.ast.Library.Parameters, param)
// 	l.stack.pop()
// }

// func (l *Listener) EnterContextDefinition(ctx *parser.ContextDefinitionContext) {
// 	l.stack.push(ctx)
// }

// func (l *Listener) ExitContextDefinition(ctx *parser.ContextDefinitionContext) {
// 	l.ast.Library.Context = Context{
// 		Identifier: ctx.Identifier().GetText(),
// 	}

// 	l.stack.pop()
// }

// // func (l *Listener) EnterStatement(ctx *parser.StatementContext) {
// // 	l.stack.push(ctx)
// // }

// // func (l *Listener) ExitStatement(ctx *parser.StatementContext) {
// // 	fmt.Println("statement", ctx.GetText())
// // 	l.stack.pop()
// // }

// func (l *Listener) EnterCodeDefinition(ctx *parser.CodeDefinitionContext) {
// 	l.stack.push(ctx)
// }

// func (l *Listener) ExitCodeDefinition(ctx *parser.CodeDefinitionContext) {
// 	fmt.Println("code definition", ctx.GetText())
// 	l.stack.pop()
// }
