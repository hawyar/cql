package cql

import "encoding/json"

type AccessModifier string

const (
	PrivateAccessModifier AccessModifier = "Private"
	PublicAccessModifier  AccessModifier = "Public"
)

type AST struct {
	Library Library `json:"Library"`
}

// type Definition interface {
// 	QualifiedIdentifier() string
// 	Identifier() string
// 	Version() string
// }

type Statement interface{}

type Library struct {
	LibraryDefinition     LibraryDefinition     `json:"LibraryDefinition"`
	ContextDefinition     ContextDefinition     `json:",omitempty"`
	IncludeDefinitions    []IncludeDefinition   `json:",omitempty"`
	UsingDefinitions      []UsingDefinition     `json:",omitempty"`
	ValuesetDefinitions   []ValuesetDefinition  `json:",omitempty"`
	CodesystemDefinitions []Codesystems         `json:",omitempty"`
	CodeDefinitions       []CodeDefinition      `json:",omitempty"`
	ParamterDefinitions   []ParameterDefinition `json:",omitempty"`
	ConceptDefinitions    []ConceptDefinition   `json:",omitempty"`
	// AllDefinition     []AllDefinition      `json:"AllDefinition"`
	// AllStatement      []Statement       `json:"AllStatement"`
	// Loc Loc
}

type QualifiedIdentifier struct {
	Qualifiers []string `json:",omitempty"`
	Identifier string
}

type LibraryDefinition struct {
	QualifiedIdentifier QualifiedIdentifier
	Version             string `json:",omitempty"`
}

type IncludeDefinition struct {
	Name    string
	Version string `json:",omitempty"`
	Alias   string `json:",omitempty"`
}

type UsingDefinition struct {
	Identifier string
	Version    string `json:",omitempty"`
}

type ValuesetDefinition struct {
	Name           string
	ID             string
	Version        string        `json:",omitempty"`
	AccessModifier string        `json:",omitempty"`
	Codesystems    []Codesystems `json:",omitempty"`
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
type ContextDefinition string
type DisplayClause string

type CodeDefinition struct {
	AccessModifier      AccessModifier `json:",omitempty"`
	Identifier          string
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
