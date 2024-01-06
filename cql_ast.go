package cql

import (
	"encoding/json"
)

type Statement interface{}

type Definition interface {
	Identifier() string
	Version() string
}

type LibraryDefinition struct {
	Identifier string `json:"id"`
	Version    string `json:"version,omitempty"`
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

type Library struct {
	Identifier          LibraryDefinition    `json:"idenitifier"`
	UsingDefinitions    []UsingDefinition    `json:"usings,omitempty"`
	IncludeDefinitions  []IncludeDefinition  `json:"includes,omitempty"`
	ValuesetDefinitions []ValuesetDefinition `json:"valueSets,omitempty"`
	Parameters          []Parameter          `json:"parameters,omitempty"`
	Context             Context              `json:"contexts,omitempty"`
	Definitions         []Definition         `json:"definitions,omitempty"`
	Statements          []Statement          `json:"statements,omitempty"`
}

type AST struct {
	Library Library `json:"library"`
}

func NewAST() *AST {
	return &AST{
		Library: Library{
			UsingDefinitions: []UsingDefinition{},
		},
	}
}

func (a *AST) MarshalJSON() ([]byte, error) {
	jstr, err := json.MarshalIndent(a.Library, "", "  ")

	if err != nil {
		return nil, err
	}

	return jstr, nil
}
