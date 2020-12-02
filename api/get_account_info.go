package api

import (
	"context"

	apiModels "github.com/go-microbot/viber/api/models"
)

// GetAccountInfo represents method to fetch the account’s details as registered in Viber.
func (api *ViberAPI) GetAccountInfo(ctx context.Context) (*apiModels.AccountInfoResponse, error) {
	resp, err := api.NewRequest("get_account_info").Do(ctx)
	if err != nil {
		return nil, err
	}

	var info apiModels.AccountInfoResponse
	if err := resp.Decode(&info); err != nil {
		return nil, err
	}

	return &info, err
}
