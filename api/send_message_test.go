package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/viber/api/models"
	"github.com/go-microbot/viber/models"
	"github.com/stretchr/testify/require"
)

type sendTextMessage struct{}

func (h sendTextMessage) Test(ctx context.Context, t *testing.T) context.Context {
	memberID := ctx.Value(chatMemberIDCtxKey)
	require.NotNil(t, memberID)

	resp, err := testAPI.SendTextMessage(ctx, apiModels.SendTextMessageRequest{
		GeneralMessageRequest: apiModels.GeneralMessageRequest{
			Receiver: memberID.(string),
			Sender: apiModels.MessageSender{
				Name: "Bot",
			},
		},
		Text: "This is a test message",
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, models.ResponseStatusCodeOK, resp.Status)

	return ctx
}

type sendPictureMessage struct{}

func (h sendPictureMessage) Test(ctx context.Context, t *testing.T) context.Context {
	memberID := ctx.Value(chatMemberIDCtxKey)
	require.NotNil(t, memberID)
	pictureURL := ctx.Value(pictureToSendCtxKey)
	require.NotNil(t, pictureURL)

	resp, err := testAPI.SendPictureMessage(ctx, apiModels.SendPictureMessageRequest{
		GeneralMessageRequest: apiModels.GeneralMessageRequest{
			Receiver: memberID.(string),
			Sender: apiModels.MessageSender{
				Name: "Bot",
			},
		},
		Text:  "This is a test picture",
		Media: pictureURL.(string),
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, models.ResponseStatusCodeOK, resp.Status)

	return ctx
}
