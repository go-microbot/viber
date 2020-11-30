package form

import (
	"bytes"
	"errors"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/require"
)

type invalidJSON struct {
	err error
}

func (r invalidJSON) MarshalJSON() ([]byte, error) {
	return nil, r.err
}

func Test_NewPartJSON(t *testing.T) {
	p := NewPartJSON("hello")
	require.NotNil(t, p)
	require.NotNil(t, p.v)
	require.Equal(t, "hello", p.v)
}

func TestPartJSON_Marshal(t *testing.T) {
	t.Run("with empty value", func(t *testing.T) {
		p := PartJSON{}
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
		p := PartJSON{}
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
	t.Run("parse JSON error", func(t *testing.T) {
		p := NewPartJSON(invalidJSON{
			err: errors.New("decode error"),
		})
		body := bytes.Buffer{}
		writer := multipart.NewWriter(&body)
		err := p.Marshal(writer, "test part", true)
		require.Error(t, err)
		require.Contains(t, err.Error(), "decode error")
		require.NoError(t, writer.Close())
	})
	t.Run("with non empty value", func(t *testing.T) {
		p := NewPartJSON(struct {
			Data string `json:"test_data"`
		}{
			Data: "test",
		})
		body := bytes.Buffer{}
		writer := multipart.NewWriter(&body)
		err := p.Marshal(writer, "test part", false)
		require.NoError(t, err)
		require.NoError(t, writer.Close())
		require.NotNil(t, body)
		require.NotEmpty(t, body.Bytes())
		partStr := body.String()
		require.Contains(t, partStr, "test part")
		require.Contains(t, partStr, `{"test_data":"test"}`)
	})
}
