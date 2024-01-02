package cql

import (
	"encoding/json"
)

type Statement interface{}

type Definition interface {
	Name() string
	Version() string
}

type LibraryDefinition struct {
	Name    string `json:"name"`
	Version string `json:"version,omitempty"`
}

type UsingDefinition struct {
	Name    string `json:"name"`
	Version string `json:"version,omitempty"`
}

type IncludeDefinition struct {
	Name    string `json:"name"`
	Version string `json:"version,omitempty"`
	Alias   string `json:"alias,omitempty"`
}

type Library struct {
	LibraryDefinition  `json:"definition"`
	UsingDefinitions   []UsingDefinition   `json:"using"`
	IncludeDefinitions []IncludeDefinition `json:"include"`
	Definitions        []Definition        `json:"definitions"`
	Statements         []Statement         `json:"statements"`
}

type AST struct {
	Library Library `json:"library"`
}

func NewAST() AST {
	return AST{
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
