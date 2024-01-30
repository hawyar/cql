package cql

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type Range struct {
	// Line  int `json:"line"`
	Start int `json:"start"`
	End   int `json:"end"`
}

func NewRange(ctx antlr.ParserRuleContext) Range {
	return Range{
		// Line:  ctx.GetStart().GetLine(),
		Start: ctx.GetStart().GetStart(),
		End:   ctx.GetStop().GetStop(),
	}
}

func (l *Range) String() string {
	return fmt.Sprintf("%d:%d", l.Start, l.End)
}
