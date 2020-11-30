package query

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewParamStringSlice(t *testing.T) {
	p := NewParamStringSlice([]string{"value1", "value2"})
	require.NotNil(t, p)
	require.NotEmpty(t, p.v)
	require.Equal(t, []string{"value1", "value2"}, p.v)
}

func TestParamStringSlice_Set(t *testing.T) {
	p := ParamStringSlice{}
	p.Set([]string{"value1", "value2"})
	require.NotEmpty(t, p.v)
	require.Equal(t, []string{"value1", "value2"}, p.v)
}

func TestParamStringSlice_Value(t *testing.T) {
	p := ParamStringSlice{}
	require.Equal(t, "[]", p.Value())
	p.Set([]string{"value1", "value2"})
	require.Equal(t, "[\"value1\",\"value2\"]", p.Value())
}

func TestParamStringSlice_Lookup(t *testing.T) {
	p := ParamStringSlice{}
	v, ok := p.Lookup()
	require.Equal(t, "[]", v)
	require.False(t, ok)
	p.Set([]string{"value1", "value2"})
	v, ok = p.Lookup()
	require.Equal(t, "[\"value1\",\"value2\"]", v)
	require.True(t, ok)
}
