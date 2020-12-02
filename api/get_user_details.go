package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/viber/api/models"
)

// GetUserDetails represents method to fetch the details
// of a specific Viber user based on his unique user ID.
// This request can be sent twice during a 12 hours period for each user ID.
func (api *ViberAPI) GetUserDetails(ctx context.Context, req apiModels.UserIDRequest) (*apiModels.UserDetailsResponse, error) {
	resp, err := api.NewRequest("get_user_details").
		Method(http.MethodPost).
		Body(NewJSONBody(req)).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	var details apiModels.UserDetailsResponse
	if err := resp.Decode(&details); err != nil {
		return nil, err
	}

	return &details, err
}
