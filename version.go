package cql

import "fmt"

type VersionOperator int

const (
	Equal VersionOperator = iota
	Greater
	Less
	LessOrEqual
	GreaterOrEqual
)

type Version struct {
	Operator VersionOperator
	Major    int
	Minor    int
	Patch    int
}

func (v VersionOperator) String() string {
	switch v {
	case Equal:
		return "="
	case Greater:
		return ">"
	case Less:
		return "<"
	case LessOrEqual:
		return "<="
	case GreaterOrEqual:
		return ">="
	default:
		return ""
	}
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}
