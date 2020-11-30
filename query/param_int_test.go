package query

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewParamInt(t *testing.T) {
	p := NewParamInt(123)
	require.NotNil(t, p)
	require.NotNil(t, p.v)
	require.Equal(t, 123, *p.v)
}

func TestParamInt_Set(t *testing.T) {
	p := ParamInt{}
	p.Set(123)
	require.NotNil(t, p.v)
	require.Equal(t, 123, *p.v)
}

func TestParamInt_Value(t *testing.T) {
	p := ParamInt{}
	require.Equal(t, "", p.Value())
	p.Set(100)
	require.Equal(t, "100", p.Value())
}

func TestParamInt_Lookup(t *testing.T) {
	p := ParamInt{}
	v, ok := p.Lookup()
	require.Equal(t, "", v)
	require.False(t, ok)
	p.Set(123)
	v, ok = p.Lookup()
	require.Equal(t, "123", v)
	require.True(t, ok)
}
