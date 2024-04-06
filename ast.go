package cql

import (
	"encoding/json"
)

type AccessModifier string

const (
	PrivateAccessModifier AccessModifier = "Private"
	PublicAccessModifier  AccessModifier = "Public"
)

type AST struct {
	Library Library `json:"Library"`
}

type Definition interface{}

type Library struct {
	LibraryDefinition     LibraryDefinition
	ContextDefinition     ContextDefinition      `json:",omitempty"`
	IncludeDefinitions    []IncludeDefinition    `json:",omitempty"`
	UsingDefinitions      []UsingDefinition      `json:",omitempty"`
	ValuesetDefinitions   []ValuesetDefinition   `json:",omitempty"`
	CodesystemDefinitions []CodesystemDefinition `json:",omitempty"`
	CodeDefinitions       []CodeDefinition       `json:",omitempty"`
	ParamterDefinitions   []ParameterDefinition  `json:",omitempty"`
	ConceptDefinitions    []ConceptDefinition    `json:",omitempty"`
	// AllDefinition     []AllDefinition      `json:"AllDefinition"`
	// AllStatement      []Statement       `json:"AllStatement"`
}

// type Definition interface {
// 	QualifiedIdentifier() string
// 	Identifier() string
// 	Version() string
// }

type Statement interface{}

type Identifier interface {
	Identifier() string
}

type PlainIdentifier struct {
	Name string
}

func (s PlainIdentifier) Identifier() string { return s.Name }

type QoutedIdentifier struct {
	Name string
}

func (q QoutedIdentifier) Identifier() string { return q.Name }

type DelimtedIdentifier struct {
	Name string
}

func (d DelimtedIdentifier) Identifier() string { return d.Name }

type QualifiedIdentifier struct {
	Qualifiers []Identifier `json:",omitempty"`
	Identifier Identifier
	Resolved   string
}

type LibraryDefinition struct {
	QualifiedIdentifier QualifiedIdentifier
	Version             string `json:",omitempty"`
}

type IncludeDefinition struct {
	QualifiedIdentifier QualifiedIdentifier
	Version             string `json:",omitempty"`
	// back identifier
	Alias Identifier `json:",omitempty"`
}

type UsingDefinition struct {
	Identifier Identifier
	Version    string `json:",omitempty"`
}

type ValuesetDefinition struct {
	Identifier     Identifier
	ValuesetID     string         `json:",omitempty"`
	Version        string         `json:",omitempty"`
	AccessModifier AccessModifier `json:",omitempty"`
	Codesystems    Codesystems    `json:",omitempty"`
}

type CodesystemDefinition struct {
	AccessModifier AccessModifier `json:",omitempty"`
	Identifier     Identifier
	CodesystemID   string `json:",omitempty"`
	Version        string `json:",omitempty"`
}

type Codesystems struct {
	QualifiedIdentifier []QualifiedIdentifier
}

type ParameterDefinition struct {
	AccessModifier    AccessModifier `json:",omitempty"`
	Name              string
	TypeSpecifier     string     `json:",omitempty"`
	DefaultExpression Expression `json:",omitempty"`
}

type Expression interface{}

type ContextDefinition struct {
	QualifiedIdentifier
}

type DisplayClause string

type CodeDefinition struct {
	AccessModifier      AccessModifier `json:",omitempty"`
	Identifier          Identifier
	QualifiedIdentifier QualifiedIdentifier
	CodeID              string        `json:",omitempty"`
	DisplayClause       DisplayClause `json:",omitempty"`
}

type ConceptDefinition struct {
	AccessModifier      AccessModifier `json:",omitempty"`
	Identifier          string
	QualifiedIdentifier []QualifiedIdentifier
	DisplayClause       DisplayClause `json:",omitempty"`
}

type Function struct {
	AccessModifier AccessModifier `json:",omitempty"`
	Identifier     string
}

type TypeSpecifier interface {
	QualifiedIdentifier() QualifiedIdentifier
	TypeSpecifier() TypeSpecifier
	Identifier() string
	// TupleElementDefinition()
}

type NamedTypeSpecifier struct {
	QualifiedIdentifier QualifiedIdentifier
}

// type ListTypeSpecifier struct {
// 	TypeSpecifier []ITypeSpecifier
// }

// type IntervalTypeSpecifier struct {
// 	TypeSpecifier []ITypeSpecifier
// }

// type TupleTypeSpecifier struct{}
// type ChoiceTypeSpecifier struct{}

// // not part of the typeSpecifier interface
// type TupleElementDefinition struct {
// 	Identifier     string
// 	TypeSpecifier   ITypeSpecifier
// }

func (a *AST) JSON() ([]byte, error) {
	if a == nil {
		return nil, nil
	}

	jstr, err := json.MarshalIndent(a, "", "  ")

	if err != nil {
		return nil, err
	}

	return jstr, nil
}
