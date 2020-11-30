package query

import "strconv"

// ParamInt represents integer query param.
type ParamInt struct {
	v *int
}

// Set sets the query int param value.
func (qpi *ParamInt) Set(v int) {
	qpi.v = &v
}

// Value returns query int param value as a string.
func (qpi ParamInt) Value() string {
	if qpi.v == nil {
		return ""
	}

	return strconv.Itoa(*qpi.v)
}

// Lookup retrieves the value of the int param.
// If the value is present the value (which may be 0) is returned and the boolean is true.
// Otherwise the returned value will be empty and the boolean will be false.
func (qpi ParamInt) Lookup() (string, bool) {
	if qpi.v == nil {
		return "", false
	}

	return qpi.Value(), true
}

// NewParamInt returns new ParamInt instance.
func NewParamInt(v int) ParamInt {
	return ParamInt{
		v: &v,
	}
}
