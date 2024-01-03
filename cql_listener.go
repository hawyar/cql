package cql

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/hawyar/cql/internal/parser"
)

type errorListener struct {
	*antlr.DefaultErrorListener
}

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
	ast   *AST
	stack *stack
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

func (s *stack) clear() {
	s.curr = nil
	s.count = 0
}

func (s *stack) print() {
	curr := s.curr
	for curr != nil {
		println(curr.val.GetText())
		curr = curr.next
	}
}

func line(ctx antlr.ParserRuleContext) {
	fmt.Printf("%d:%d %s\n", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), ctx.GetText())
}

func NewErrorListener() *errorListener {
	return &errorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
	}
}

func (l *Listener) EnterLibrary(ctx *parser.LibraryContext) {
	l.stack.push(ctx)
}

func (l *Listener) ExitLibrary(ctx *parser.LibraryContext) {
	l.stack.pop()

	// if !l.stack.isEmpty() {
	// 	// TODO: fix this
	// 	for {
	// 		fmt.Println(l.stack.pop())
	// 	}
	// }

	jsonstr, err := l.ast.MarshalJSON()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonstr))
}

func (l *Listener) EnterLibraryDefinition(ctx *parser.LibraryDefinitionContext) {
	l.stack.push(ctx)
}

func (l *Listener) EnterQualifiedIdentifier(ctx *parser.QualifiedIdentifierContext) {
	fmt.Println("qualified identifier", ctx.GetText())
	l.stack.push(ctx)
}

func (l *Listener) EnterIdentifier(ctx *parser.IdentifierContext) {
	fmt.Println("identifier", ctx.GetText())
	l.stack.push(ctx)
}

// func (l *Listener) ExitLibraryDefinition(ctx *parser.LibraryDefinitionContext) {
// 	if l.stack.size() == 1 {
// 		return
// 	}

// 	current := l.stack.peek()

// 	if current == nil {
// 		return
// 	}

// 	definition := LibraryDefinition{}

// 	switch current.(type) {
// 	case *parser.VersionSpecifierContext:
// 		version := l.stack.pop()
// 		definition.Version = strings.ReplaceAll(version.GetText(), "'", "")
// 		ident := l.stack.pop()
// 		definition.Name = ident.GetText()
// 	case *parser.QualifiedIdentifierContext:
// 		ident := l.stack.pop()
// 		definition.Name = ident.GetText()
// 	default:
// 		fmt.Println("something is wrong, check the inserts in the previous branch")
// 	}
// 	l.ast.Library.LibraryDefinition = definition
// 	l.stack.pop()
// }

// func (l *Listener) EnterVersionSpecifier(ctx *parser.VersionSpecifierContext) {
// 	l.stack.push(ctx)
// }

// func (l *Listener) EnterQualifiedIdentifier(ctx *parser.QualifiedIdentifierContext) {
// 	l.stack.push(ctx)
// }

// func (l *Listener) EnterUsingDefinition(ctx *parser.UsingDefinitionContext) {
// 	l.stack.push(ctx)
// }

// func (l *Listener) ExitUsingDefinition(ctx *parser.UsingDefinitionContext) {
// 	if l.stack.size() == 1 {
// 		return
// 	}

// 	current := l.stack.peek()

// 	if current == nil {
// 		return
// 	}

// 	definition := UsingDefinition{}

// 	switch current.(type) {
// 	case *parser.VersionSpecifierContext:
// 		version := l.stack.pop()
// 		definition.Version = strings.ReplaceAll(version.GetText(), "'", "")
// 		ident := l.stack.pop()
// 		definition.Name = ident.GetText()
// 	case *parser.IdentifierContext:
// 		ident := l.stack.pop()
// 		definition.Name = ident.GetText()
// 	default:
// 		fmt.Println("something is wrong, check the inserts in the previous branch")
// 	}
// 	l.ast.Library.UsingDefinitions = append(l.ast.Library.UsingDefinitions, definition)
// 	l.stack.pop()
// }

// func (l *Listener) EnterIdentifier(ctx *parser.IdentifierContext) {
// 	l.stack.push(ctx)
// }

// func (l *Listener) EnterIncludeDefinition(ctx *parser.IncludeDefinitionContext) {
// 	l.stack.push(ctx)
// }

// func (l *Listener) ExitIncludeDefinition(ctx *parser.IncludeDefinitionContext) {
// 	if l.stack.size() == 1 {
// 		return
// 	}

// 	current := l.stack.peek()

// 	if current == nil {
// 		return
// 	}

// 	definition := IncludeDefinition{}

// 	switch current.(type) {
// 	case *parser.QualifiedIdentifierContext:
// 		ident := l.stack.pop()
// 		definition.Name = ident.GetText()
// 	case *parser.VersionSpecifierContext:
// 		version := l.stack.pop()
// 		definition.Version = strings.ReplaceAll(version.GetText(), "'", "")
// 		ident := l.stack.pop()
// 		definition.Name = ident.GetText()
// 	case *parser.IdentifierContext:
// 		alias := l.stack.pop()
// 		definition.Alias = alias.GetText()

// 		current = l.stack.peek()
// 		if current == nil {
// 			return
// 		}

// 		if _, ok := current.(*parser.VersionSpecifierContext); ok {
// 			version := l.stack.pop()
// 			definition.Version = strings.ReplaceAll(version.GetText(), "'", "")
// 		}

// 		ident := l.stack.pop()
// 		definition.Name = ident.GetText()
// 	default:
// 		fmt.Println("something is wrong, check the inserts in the previous branch")
// 	}
// 	l.ast.Library.IncludeDefinitions = append(l.ast.Library.IncludeDefinitions, definition)
// 	l.stack.pop()
// }

// func (l *Listener) EnterAccessModifier(ctx *parser.AccessModifierContext) {
// 	l.stack.push(ctx)
// }

// func (l *Listener) EnterCodesystemDefinition(ctx *parser.CodesystemDefinitionContext) {
// 	l.stack.push(ctx)
// }

// func (l *Listener) EnterValuesetDefinition(ctx *parser.ValuesetDefinitionContext) {
// 	l.stack.push(ctx)
// }

// func (l *Listener) EnterCodesystems(ctx *parser.CodesystemsContext) {
// 	l.stack.push(ctx)
// }
