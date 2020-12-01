package api

import (
	"context"

	apiModels "github.com/go-microbot/viber/api/models"
	"github.com/go-microbot/viber/models"
)

// Bot represents general Bot API interface.
// https://developers.viber.com/docs/api/rest-bot-api/#get-started.
type Bot interface {
	// https://developers.viber.com/docs/api/rest-bot-api/#get-account-info.
	GetAccountInfo(ctx context.Context) (*models.AccountInfo, error)
	// https://developers.viber.com/docs/api/rest-bot-api/#setting-a-webhook.
	SetWebhook(ctx context.Context, req apiModels.SetWebhookRequest) (*apiModels.SetWebhookResponse, error)
	// https://developers.viber.com/docs/api/rest-bot-api/#removing-your-webhook.
	RemoveWebhook(ctx context.Context) error
}
