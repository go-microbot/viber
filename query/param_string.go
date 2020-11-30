package query

// ParamString represents string query param.
type ParamString struct {
	v *string
}

// Set sets the query string param value.
func (qps *ParamString) Set(v string) {
	qps.v = &v
}

// Value returns query string param value as a string.
func (qps ParamString) Value() string {
	if qps.v == nil {
		return ""
	}

	return *qps.v
}

// Lookup retrieves the value of the string param.
// If the value is present the value (which may be 0) is returned and the boolean is true.
// Otherwise the returned value will be empty and the boolean will be false.
func (qps ParamString) Lookup() (string, bool) {
	if qps.v == nil {
		return "", false
	}

	return qps.Value(), true
}

// NewParamString returns new ParamString instance.
func NewParamString(v string) ParamString {
	return ParamString{
		v: &v,
	}
}
