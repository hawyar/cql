package cql

import (
	fhirpath "github.com/hawyar/cql/internal/fhirpath"
)

type PathListener struct {
	*fhirpath.BasefhirpathListener
}

type Expression interface{}

type PathAST struct {
	Expression Expression
}

func NewPathAST() *PathAST {
	return &PathAST{}
}
