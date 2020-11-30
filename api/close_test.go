package api

/*
import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

type close struct{}

func (h close) Test(ctx context.Context, t *testing.T) context.Context {
	prevToken := localAPI.token
	defer func() {
		localAPI.token = prevToken
	}()
	localAPI.token = "invalid"
	err := localAPI.Close(ctx)
	require.Error(t, err)

	return ctx
}
*/
