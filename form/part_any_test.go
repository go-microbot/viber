package form

import (
	"bytes"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewPartAny(t *testing.T) {
	p := NewPartAny("hello")
	require.NotNil(t, p)
	require.NotNil(t, p.v)
	require.Equal(t, "hello", p.v)
}

func TestPartAny_Marshal(t *testing.T) {
	t.Run("with empty value", func(t *testing.T) {
		p := PartAny{}
		body := bytes.Buffer{}
		writer := multipart.NewWriter(&body)
		err := p.Marshal(writer, "test part", false)
		require.NoError(t, err)
		require.NoError(t, writer.Close())
		require.NotNil(t, body)
		require.NotEmpty(t, body.Bytes())
		partStr := body.String()
		require.Contains(t, partStr, "test part")
	})
	t.Run("with empty value and omitempty", func(t *testing.T) {
		p := PartAny{}
		body := bytes.Buffer{}
		writer := multipart.NewWriter(&body)
		err := p.Marshal(writer, "test part", true)
		require.NoError(t, err)
		require.NoError(t, writer.Close())
		require.NotNil(t, body)
		require.NotEmpty(t, body.Bytes())
		partStr := body.String()
		require.NotContains(t, partStr, "test part")
	})
	t.Run("with non empty value", func(t *testing.T) {
		p := NewPartAny("test value")
		body := bytes.Buffer{}
		writer := multipart.NewWriter(&body)
		err := p.Marshal(writer, "test part", false)
		require.NoError(t, err)
		require.NoError(t, writer.Close())
		require.NotNil(t, body)
		require.NotEmpty(t, body.Bytes())
		partStr := body.String()
		require.Contains(t, partStr, "test part")
		require.Contains(t, partStr, "test value")
	})
}
