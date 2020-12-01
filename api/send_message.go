package api

import (
	"context"

	apiModels "github.com/go-microbot/viber/api/models"
	"github.com/go-microbot/viber/models"
)

// SendTextMessage represents method to send text message.
func (api *ViberAPI) SendTextMessage(ctx context.Context, req apiModels.SendTextMessageRequest) (*apiModels.MessageResponse, error) {
	req.Type = models.MessageTypeText
	return api.sendMessage(ctx, req)
}

// SendPictureMessage represents method to send picture message.
func (api *ViberAPI) SendPictureMessage(ctx context.Context, req apiModels.SendPictureMessageRequest) (*apiModels.MessageResponse, error) {
	req.Type = models.MessageTypePicture
	return api.sendMessage(ctx, req)
}

func (api *ViberAPI) sendMessage(ctx context.Context, req interface{}) (*apiModels.MessageResponse, error) {
	resp, err := api.NewRequest("send_message").
		Body(NewJSONBody(req)).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	var msgResp apiModels.MessageResponse
	if err := resp.Decode(&msgResp); err != nil {
		return nil, err
	}

	return &msgResp, err
}
