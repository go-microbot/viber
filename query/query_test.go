package query

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_AsMap(t *testing.T) {
	testStruct := struct {
		// int params.
		RequiredIntValue      ParamInt `query:"int_required"`
		RequiredIntEmptyValue ParamInt `query:"int_required_empty"`
		IntIgnoreValue        ParamInt `query:"int_ignore,omitempty"`
		IntValue              ParamInt
		// string slice params.
		RequiredStringSliceValue      ParamStringSlice `query:"string_slice_required"`
		RequiredStringSliceEmptyValue ParamStringSlice `query:"string_slice_required_empty"`
		StringSliceIgnoreValue        ParamStringSlice `query:"string_slice_ignore,omitempty"`
		StringSliceValue              ParamStringSlice
		// string params.
		RequiredStringValue      ParamString `query:"string_required"`
		RequiredStringEmptyValue ParamString `query:"string_required_empty"`
		StringIgnoreValue        ParamString `query:"string_ignore,omitempty"`
		StringValue              ParamString
		// any params.
		RequiredAnyValue      ParamAny `query:"any_required"`
		RequiredAnyEmptyValue ParamAny `query:"any_required_empty"`
		AnyIgnoreValue        ParamAny `query:"any_ignore,omitempty"`
		AnyValue              ParamAny
		// boolean params.
		RequiredBoolValue      ParamBool `query:"bool_required"`
		RequiredBoolEmptyValue ParamBool `query:"bool_required_empty"`
		BoolIgnoreValue        ParamBool `query:"bool_ignore,omitempty"`
		BoolValue              ParamBool
		// unparsed params.
		Hello string `query:"hello"`
		World int
	}{
		RequiredIntValue:         NewParamInt(100),
		IntValue:                 NewParamInt(200),
		RequiredStringSliceValue: NewParamStringSlice([]string{"1", "2", "3"}),
		StringSliceValue:         NewParamStringSlice([]string{"hello", "world"}),
		RequiredStringValue:      NewParamString("hello"),
		StringValue:              NewParamString("world"),
		RequiredAnyValue:         NewParamAny("hello"),
		AnyValue:                 NewParamAny(123),
		RequiredBoolValue:        NewParamBool(true),
		BoolValue:                NewParamBool(false),
	}

	expMap := map[string]string{
		"int_required":                "100",
		"IntValue":                    "200",
		"int_required_empty":          "",
		"string_slice_required":       "[\"1\",\"2\",\"3\"]",
		"StringSliceValue":            "[\"hello\",\"world\"]",
		"string_slice_required_empty": "[]",
		"string_required":             "hello",
		"StringValue":                 "world",
		"string_required_empty":       "",
		"any_required":                "hello",
		"AnyValue":                    "123",
		"any_required_empty":          "",
		"bool_required":               "true",
		"BoolValue":                   "false",
		"bool_required_empty":         "",
	}
	result := AsMap(testStruct)

	require.Equal(t, expMap, result)
}
