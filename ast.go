package cql

import (
	"encoding/json"
)

type AccessModifier string

const (
	PrivateModifier AccessModifier = "Private"
	PublicModifier  AccessModifier = "Public"
)

type Program struct {
	Library Library `json:"Library"`
	File    string  `json:"File"`
	Loc     Loc     `json:"Range"`
}

type Statement interface {
	IsContext() bool
}

type Library struct {
	LibDef LibraryDefinition `json:"LibraryDefinition"`
	Loc    Loc               `json:"Loc"`
}

type AST struct {
	Library Library `json:"library"`
	Loc     Loc     `json:"range"`
}

type Idenitifier interface {
	Name() string
	Loc() Loc
}

type BaseIdentifier struct {
	name string `json:"name"`
	loc  Loc    `json:"range"`
}

func (b *BaseIdentifier) Name() string {
	return b.name
}

func (b *BaseIdentifier) Loc() Loc {
	return b.loc
}

type QuotedIdentifier struct {
	name string `json:"name"`
	loc  Loc    `json:"range"`
}

func (q *QuotedIdentifier) Name() string {
	return q.name
}

func (q *QuotedIdentifier) Loc() Loc {
	return q.loc
}

type LibraryDefinition struct {
	Identifier Idenitifier `json:"identifier"`
	Version    string      `json:"version,omitempty"`
}

type UsingDefinition struct {
	Name    string `json:"name"`
	Version string `json:"version,omitempty"`
}

type ValuesetDefinition struct {
	Name           string        `json:"name"`
	Id             string        `json:"id,omitempty"`
	Version        string        `json:"version,omitempty"`
	AccessModifier string        `json:"accessModifier,omitempty"`
	Codesystems    []Codesystems `json:"codesystems,omitempty"`
}

type Context struct {
	Identifier string `json:"identifier"`
}

type Parameter struct {
	AccessModifier string `json:"accessModifier,omitempty"`
	Name           string `json:"name"`
}

type IncludeDefinition struct {
	Name    string `json:"name"`
	Version string `json:"version,omitempty"`
	Alias   string `json:"alias,omitempty"`
}

type Codesystems struct {
	Identifier string `json:"identifier"`
}

// type Library struct {
// 	Identifier          LibraryDefinition    `json:"idenitifier"`
// 	UsingDefinitions    []UsingDefinition    `json:"usings,omitempty"`
// 	IncludeDefinitions  []IncludeDefinition  `json:"includes,omitempty"`
// 	ValuesetDefinitions []ValuesetDefinition `json:"valueSets,omitempty"`
// 	Parameters          []Parameter          `json:"parameters,omitempty"`
// 	Context             Context              `json:"contexts,omitempty"`
// 	Statements          []Statement          `json:"statements,omitempty"`
// 	Range               Range                `json:"range"`
// }

func (a *Program) JSON() ([]byte, error) {
	jstr, err := json.MarshalIndent(a, "", "  ")

	if err != nil {
		return nil, err
	}

	return jstr, nil
}
