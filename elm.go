package cql

import (
	"encoding/json"
	"encoding/xml"
)

type OutputFormat int

const (
	JSONFormat OutputFormat = iota
	XMLFormat
)

type ELM struct {
	ast    *AST
	format OutputFormat
}

func NewELM(ast *AST) (*ELM, error) {
	return &ELM{
		ast:    ast,
		format: JSONFormat,
	}, nil
}

func (e *ELM) AST() *AST {
	return e.ast
}

func (e *ELM) MarshalJSON() ([]byte, error) {
	jsonMap := map[string]interface{}{"library": e.ast.Library}

	jstr, err := json.MarshalIndent(jsonMap, "", "  ")

	if err != nil {
		return nil, err
	}

	return jstr, nil
}

func (e *ELM) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	return encoder.EncodeElement(e.ast.Library, start)
}
