package form

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewPartFile(t *testing.T) {
	p := NewPartFile("hello")
	require.NotNil(t, p)
	require.NotNil(t, p.filePath)
	require.Equal(t, "hello", *p.filePath)
}

func TestPartFile_Marshal(t *testing.T) {
	t.Run("open file error", func(t *testing.T) {
		p := PartFile{}
		body := bytes.Buffer{}
		writer := multipart.NewWriter(&body)
		err := p.Marshal(writer, "test part", false)
		require.Error(t, err)
		_, ok := err.(*os.PathError)
		require.True(t, ok)
		require.NoError(t, writer.Close())
	})
	t.Run("with empty value and omitempty", func(t *testing.T) {
		p := PartFile{}
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
		fPath := path.Join(os.TempDir(), "test.txt")
		err := ioutil.WriteFile(fPath, []byte("test value"), 0667)
		require.NoError(t, err)
		defer func() {
			err := os.Remove(fPath)
			require.NoError(t, err)
		}()

		p := NewPartFile(fPath)
		body := bytes.Buffer{}
		writer := multipart.NewWriter(&body)
		err = p.Marshal(writer, "test part", false)
		require.NoError(t, err)
		require.NoError(t, writer.Close())
		require.NotNil(t, body)
		require.NotEmpty(t, body.Bytes())
		partStr := body.String()
		require.Contains(t, partStr, "test part")
		require.Contains(t, partStr, "test value")
	})
}
