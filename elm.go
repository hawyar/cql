package cql

type ELM struct {
	json string
	ast  *AST // the underlying AST
}

func NewELM(ast *AST) (*ELM, error) {
	jsontrs, err := ast.MarshalJSON()

	if err != nil {
		return nil, err
	}

	return &ELM{
		json: string(jsontrs),
		ast:  ast,
	}, nil
}

func (e *ELM) JSON() string {
	return e.json
}

func (e *ELM) AST() *AST {
	return e.ast
}
