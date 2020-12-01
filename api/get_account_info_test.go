package api

import (
	"context"
	"testing"

	"github.com/go-microbot/viber/models"
	"github.com/stretchr/testify/require"
)

type getAccountInfo struct{}

func (h getAccountInfo) Test(ctx context.Context, t *testing.T) context.Context {
	info, err := testAPI.GetAccountInfo(ctx)
	require.NoError(t, err)
	require.NotNil(t, info)
	require.Equal(t, models.ResponseStatusCodeOK, info.Status)
	require.Equal(t, models.ResponseStatusNameOK, info.StatusMessage)
	require.NotEmpty(t, info.Name)
	require.NotEmpty(t, info.Members)

	ctx = context.WithValue(ctx, TestDataKey(botNameCtxKey), info.Name)
	return context.WithValue(ctx, TestDataKey(chatMemberIDCtxKey), info.Members[0].ID)
}
