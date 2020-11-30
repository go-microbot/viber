package query

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewParamBool(t *testing.T) {
	p := NewParamBool(true)
	require.NotNil(t, p)
	require.NotNil(t, p.v)
	require.Equal(t, true, *p.v)
}

func TestParamBool_Set(t *testing.T) {
	p := ParamBool{}
	p.Set(true)
	require.NotNil(t, p.v)
	require.Equal(t, true, *p.v)
}

func TestParamBool_Value(t *testing.T) {
	p := ParamBool{}
	require.Equal(t, "", p.Value())
	p.Set(false)
	require.Equal(t, "false", p.Value())
}

func TestParamBool_Lookup(t *testing.T) {
	p := ParamBool{}
	v, ok := p.Lookup()
	require.Equal(t, "", v)
	require.False(t, ok)
	p.Set(true)
	v, ok = p.Lookup()
	require.Equal(t, "true", v)
	require.True(t, ok)
}
