package query

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewParamAny(t *testing.T) {
	p := NewParamAny("test")
	require.NotNil(t, p)
	require.NotNil(t, p.v)
	require.Equal(t, "test", *p.v)
}

func TestParamAny_Set(t *testing.T) {
	p := ParamAny{}
	p.Set(1234)
	require.NotNil(t, p.v)
	require.Equal(t, 1234, *p.v)
}

func TestParamAny_Value(t *testing.T) {
	p := ParamAny{}
	require.Equal(t, "", p.Value())
	p.Set([]string{"hello", "world"})
	require.Equal(t, "[hello world]", p.Value())
}

func TestParamAny_Lookup(t *testing.T) {
	p := ParamAny{}
	v, ok := p.Lookup()
	require.Equal(t, "", v)
	require.False(t, ok)
	p.Set("test")
	v, ok = p.Lookup()
	require.Equal(t, "test", v)
	require.True(t, ok)
}
