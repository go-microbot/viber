package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/viber/api/models"
	"github.com/go-microbot/viber/models"
	"github.com/stretchr/testify/require"
)

type getUserDetails struct{}

func (h getUserDetails) Test(ctx context.Context, t *testing.T) context.Context {
	memberID := ctx.Value(chatMemberIDCtxKey)
	require.NotNil(t, memberID)

	details, err := testAPI.GetUserDetails(ctx, apiModels.UserIDRequest{
		ID: memberID.(string),
	})
	require.NoError(t, err)
	require.NotNil(t, details)
	if details.Status != models.ResponseStatusCodeOK {
		require.Contains(t, details.StatusMessage, "maximum get user info requests exceeded")
		return ctx
	}
	require.NotEmpty(t, details.User.Name)
	require.Equal(t, memberID, details.User.ID)

	return ctx
}
