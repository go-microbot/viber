package api

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestErr_Unwrap(t *testing.T) {
	err := newErr(ErrDecodeBody, errors.New("some error"))
	require.Error(t, err)
	require.Equal(t, ErrDecodeBody, errors.Unwrap(err))
}

func TestErr_Error(t *testing.T) {
	err := newErr(ErrDecodeBody, errors.New("some error"))
	require.Error(t, err)
	require.Equal(t, fmt.Sprintf("%v: %v", ErrDecodeBody, errors.New("some error")), err.Error())
}
