package cql

import (
	"github.com/hawyar/cql/internal/fhirpath"
)

type FHIRPathListener struct {
	*fhirpath.BasefhirpathListener
	stack *stack
	ast   *FHIRPathAST
}

func NewFHIRPathListener() *FHIRPathListener {
	return &FHIRPathListener{
		stack: newStack(),
		ast:   newFHIRPathAST(),
	}
}
