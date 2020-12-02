package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/viber/api/models"
)

// SetWebhook represents method to set your account’s webhook.
func (api *ViberAPI) SetWebhook(ctx context.Context, req apiModels.SetWebhookRequest) (*apiModels.SetWebhookResponse, error) {
	resp, err := api.NewRequest("set_webhook").
		Body(NewJSONBody(req)).
		Method(http.MethodPost).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	var whResp apiModels.SetWebhookResponse
	if err := resp.Decode(&whResp); err != nil {
		return nil, err
	}

	return &whResp, err
}
