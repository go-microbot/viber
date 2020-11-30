package form

import (
	"errors"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/require"
)

type testPart struct {
	err error
}

func (p testPart) Marshal(w *multipart.Writer, partName string, omitempty bool) error {
	return p.err
}

func Test_Marshal(t *testing.T) {
	testCases := []struct {
		name string
		data interface{}
		exp  func(t *testing.T, res []byte, ct string, err error)
	}{
		{
			name: "nothing to marshal",
			data: struct{}{},
			exp: func(t *testing.T, res []byte, ct string, err error) {
				require.NoError(t, err)
			},
		},
		{
			name: "empty values",
			data: struct {
				Field1 int
				Field2 PartText `form:"-"`
			}{
				Field1: 11,
				Field2: NewPartText("test"),
			},
			exp: func(t *testing.T, res []byte, ct string, err error) {
				require.NoError(t, err)
			},
		},
		{
			name: "marshal error",
			data: struct {
				Field testPart
			}{
				Field: testPart{
					err: errors.New("error"),
				},
			},
			exp: func(t *testing.T, res []byte, ct string, err error) {
				require.EqualError(t, err, "could not marshal Field value: error")
			},
		},
		{
			name: "all ok. text-part values",
			data: struct {
				Field1 PartText `form:"-"`
				Field2 PartText `form:"test"`
				Field3 PartText `form:"test2,omitempty"`
				Field4 PartText
				Field5 Part `form:"test_test,omitempty"`
			}{
				Field1: NewPartText("test value 1"),
				Field2: NewPartText("test value 2"),
				Field4: NewPartText("test value 4"),
			},
			exp: func(t *testing.T, res []byte, ct string, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, res)
				require.NotEmpty(t, ct)
				resStr := string(res)
				require.Contains(t, resStr, "name=\"test\"")
				require.Contains(t, resStr, "test value 2")
				require.Contains(t, resStr, "name=\"Field4\"")
				require.Contains(t, resStr, "test value 4")
				require.NotContains(t, resStr, "name=\"test2\"")
				require.NotContains(t, resStr, "test value 1")
			},
		},
	}
	for i := range testCases {
		tc := &testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			var ct string
			res, err := Marshal(tc.data, &ct)
			tc.exp(t, res, ct, err)
		})
	}
}
