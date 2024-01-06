package cql

import (
	"encoding/json"
)

type Statement interface{}

type QualifiedIdentifier struct {
	Identifier string `json:"identifier"`
	Version    string `json:"version,omitempty"`
}

type Definition interface {
	Identifier() string
	Version() string
}

type LibraryDefinition struct {
	Idenitifier string   `json:"id"`
	Version     string   `json:"version,omitempty"`
	Path        []string `json:"path,omitempty"`
}

type UsingDefinition struct {
	Name    string `json:"name"`
	Version string `json:"version,omitempty"`
}

type IncludeDefinition struct {
	Name    string   `json:"name"`
	Version string   `json:"version,omitempty"`
	Alias   string   `json:"alias,omitempty"`
	Path    []string `json:"path,omitempty"`
}

type Library struct {
	Identifier         LibraryDefinition   `json:"idenitifier"`
	UsingDefinitions   []UsingDefinition   `json:"usings,omitempty"`
	IncludeDefinitions []IncludeDefinition `json:"includes,omitempty"`
	Definitions        []Definition        `json:"definitions,omitempty"`
	Statements         []Statement         `json:"statements,omitempty"`
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
