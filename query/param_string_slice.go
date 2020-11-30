package query

import (
	"fmt"
	"strings"
)

// ParamStringSlice represents string slice query param.
type ParamStringSlice struct {
	v []string
}

// Set sets the query string slice param value.
func (qpss *ParamStringSlice) Set(v []string) {
	qpss.v = v
}

// Value returns query string slice param value as a string.
func (qpss ParamStringSlice) Value() string {
	values := make([]string, len(qpss.v))
	for i := range qpss.v {
		values[i] = fmt.Sprintf("%q", qpss.v[i])
	}

	return fmt.Sprintf("[%s]", strings.Join(values, ","))
}

// Lookup retrieves the value of the string slice param.
// If the value is present the value (which may be empty) is returned and the boolean is true.
// Otherwise the returned value will be empty and the boolean will be false.
func (qpss ParamStringSlice) Lookup() (string, bool) {
	ok := len(qpss.v) != 0
	return qpss.Value(), ok
}

// NewParamStringSlice returns new ParamStringSlice instance.
func NewParamStringSlice(v []string) ParamStringSlice {
	return ParamStringSlice{
		v: v,
	}
}
