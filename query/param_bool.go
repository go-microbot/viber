package query

import "strconv"

// ParamBool represents bool query param.
type ParamBool struct {
	v *bool
}

// Set sets the query bool param value.
func (qpb *ParamBool) Set(v bool) {
	qpb.v = &v
}

// Value returns query bool param value as a string.
func (qpb ParamBool) Value() string {
	if qpb.v == nil {
		return ""
	}

	return strconv.FormatBool(*qpb.v)
}

// Lookup retrieves the value of the bool param.
// If the value is present the value (which may be 0) is returned and the boolean is true.
// Otherwise the returned value will be empty and the boolean will be false.
func (qpb ParamBool) Lookup() (string, bool) {
	if qpb.v == nil {
		return "", false
	}

	return qpb.Value(), true
}

// NewParamBool returns new ParamBool instance.
func NewParamBool(v bool) ParamBool {
	return ParamBool{
		v: &v,
	}
}
