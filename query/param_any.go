package query

import (
	"fmt"
)

// ParamAny represents any type query param.
type ParamAny struct {
	v *interface{}
}

// Set sets the query any type param value.
func (qpa *ParamAny) Set(v interface{}) {
	qpa.v = &v
}

// Value returns query any type param value as a string.
func (qpa ParamAny) Value() string {
	if qpa.v == nil {
		return ""
	}

	return fmt.Sprintf("%v", *qpa.v)
}

// Lookup retrieves the value of the any type param.
// If the value is present the value (which may be nil) is returned and the boolean is true.
// Otherwise the returned value will be empty and the boolean will be false.
func (qpa ParamAny) Lookup() (string, bool) {
	if qpa.v == nil {
		return "", false
	}

	return qpa.Value(), true
}

// NewParamAny returns new ParamAny instance.
func NewParamAny(v interface{}) ParamAny {
	return ParamAny{
		v: &v,
	}
}
