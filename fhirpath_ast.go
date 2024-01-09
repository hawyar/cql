package cql

type FHIRPathAST struct {
	Expression string `json:"expression"`
}

func newFHIRPathAST() *FHIRPathAST {
	return &FHIRPathAST{}
}
