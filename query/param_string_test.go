package query

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewParamString(t *testing.T) {
	p := NewParamString("test")
	require.NotNil(t, p)
	require.NotNil(t, p.v)
	require.Equal(t, "test", *p.v)
}

func TestParamString_Set(t *testing.T) {
	p := ParamString{}
	p.Set("test")
	require.NotNil(t, p.v)
	require.Equal(t, "test", *p.v)
}

func TestParamString_Value(t *testing.T) {
	p := ParamString{}
	require.Equal(t, "", p.Value())
	p.Set("test")
	require.Equal(t, "test", p.Value())
}

func TestParamString_Lookup(t *testing.T) {
	p := ParamString{}
	v, ok := p.Lookup()
	require.Equal(t, "", v)
	require.False(t, ok)
	p.Set("test")
	v, ok = p.Lookup()
	require.Equal(t, "test", v)
	require.True(t, ok)
}
