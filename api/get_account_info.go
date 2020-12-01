package api

import (
	"context"

	"github.com/go-microbot/viber/models"
)

// GetAccountInfo represents method to fetch the account’s details as registered in Viber.
func (api *ViberAPI) GetAccountInfo(ctx context.Context) (*models.AccountInfo, error) {
	resp, err := api.NewRequest("get_account_info").Do(ctx)
	if err != nil {
		return nil, err
	}

	var info models.AccountInfo
	if err := resp.Decode(&info); err != nil {
		return nil, err
	}

	return &info, err
}
