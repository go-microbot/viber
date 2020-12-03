package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/viber/api/models"
	"github.com/go-microbot/viber/models"
	"github.com/stretchr/testify/require"
)

type getOnline struct{}

func (h getOnline) Test(ctx context.Context, t *testing.T) context.Context {
	memberID := ctx.Value(chatMemberIDCtxKey)
	require.NotNil(t, memberID)

	online, err := testAPI.GetOnline(ctx, apiModels.UserIDsRequest{
		IDs: []string{memberID.(string)},
	})
	require.NoError(t, err)
	require.NotNil(t, online)
	require.NotEmpty(t, online.Users)
	var found bool
	for i := range online.Users {
		if online.Users[i].ID == memberID.(string) {
			require.True(t,
				online.Users[i].OnlineStatus == models.UserOnlineStatusOnline ||
					online.Users[i].OnlineStatus == models.UserOnlineStatusOffline)
			found = true
			break
		}
	}
	require.True(t, found)

	return ctx
}
