package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/viber/api/models"
	"github.com/go-microbot/viber/models"
	"github.com/stretchr/testify/require"
)

type setWebhook struct{}

func (h setWebhook) Test(ctx context.Context, t *testing.T) context.Context {
	webhookURL := ctx.Value(webhookURLCtxKey)
	require.NotNil(t, webhookURL)

	resp, err := testAPI.SetWebhook(ctx, apiModels.SetWebhookRequest{
		URL: webhookURL.(string),
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, models.ResponseStatusCodeOK, resp.Status)
	require.NotEmpty(t, resp.EventTypes)

	return ctx
}
