package api

import (
	"context"

	apiModels "github.com/go-microbot/viber/api/models"
)

// RemoveWebhook represents method to remove existing webhook.
func (api *ViberAPI) RemoveWebhook(ctx context.Context) error {
	_, err := api.NewRequest("set_webhook").
		Body(NewJSONBody(apiModels.SetWebhookRequest{})).
		Do(ctx)

	return err
}
