package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/viber/api/models"
)

// GetOnline represents method to fetch the online status of a given subscribed account members.
// The API supports up to 100 user ID per request and
// those users must be subscribed to the account.
func (api *ViberAPI) GetOnline(ctx context.Context, req apiModels.UserIDsRequest) (*apiModels.GetOnlineResponse, error) {
	resp, err := api.NewRequest("get_online").
		Method(http.MethodPost).
		Body(NewJSONBody(req)).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	var onlineResp apiModels.GetOnlineResponse
	if err := resp.Decode(&onlineResp); err != nil {
		return nil, err
	}

	return &onlineResp, err
}
