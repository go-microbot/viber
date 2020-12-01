package api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

type removeWebhook struct{}

func (h removeWebhook) Test(ctx context.Context, t *testing.T) context.Context {
	err := testAPI.RemoveWebhook(ctx)
	require.NoError(t, err)

	return ctx
}
