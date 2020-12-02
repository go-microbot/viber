package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/viber/api/models"
	"github.com/go-microbot/viber/models"
)

// SendTextMessage represents method to send text message.
func (api *ViberAPI) SendTextMessage(ctx context.Context, req apiModels.SendTextMessageRequest) (*apiModels.MessageResponse, error) {
	req.Type = models.MessageTypeText
	return api.sendMessage(ctx, req, len(req.BroadcastList) != 0)
}

// SendPictureMessage represents method to send picture message.
func (api *ViberAPI) SendPictureMessage(ctx context.Context, req apiModels.SendPictureMessageRequest) (*apiModels.MessageResponse, error) {
	req.Type = models.MessageTypePicture
	return api.sendMessage(ctx, req, len(req.BroadcastList) != 0)
}

// SendVideoMessage represents method to send video message.
func (api *ViberAPI) SendVideoMessage(ctx context.Context, req apiModels.SendVideoMessageRequest) (*apiModels.MessageResponse, error) {
	req.Type = models.MessageTypeVideo
	return api.sendMessage(ctx, req, len(req.BroadcastList) != 0)
}

// SendFileMessage represents method to send file message.
func (api *ViberAPI) SendFileMessage(ctx context.Context, req apiModels.SendFileMessageRequest) (*apiModels.MessageResponse, error) {
	req.Type = models.MessageTypeFile
	return api.sendMessage(ctx, req, len(req.BroadcastList) != 0)
}

// SendContactMessage represents method to send contact message.
func (api *ViberAPI) SendContactMessage(ctx context.Context, req apiModels.SendContactMessageRequest) (*apiModels.MessageResponse, error) {
	req.Type = models.MessageTypeContact
	return api.sendMessage(ctx, req, len(req.BroadcastList) != 0)
}

// SendLocationMessage represents method to send location message.
func (api *ViberAPI) SendLocationMessage(ctx context.Context, req apiModels.SendLocationMessageRequest) (*apiModels.MessageResponse, error) {
	req.Type = models.MessageTypeLocation
	return api.sendMessage(ctx, req, len(req.BroadcastList) != 0)
}

// SendURLMessage represents method to send URL message.
func (api *ViberAPI) SendURLMessage(ctx context.Context, req apiModels.SendURLMessageRequest) (*apiModels.MessageResponse, error) {
	req.Type = models.MessageTypeURL
	return api.sendMessage(ctx, req, len(req.BroadcastList) != 0)
}

// SendStickerMessage represents method to send sticker message.
func (api *ViberAPI) SendStickerMessage(ctx context.Context, req apiModels.SendStickerMessageRequest) (*apiModels.MessageResponse, error) {
	req.Type = models.MessageTypeSticker
	return api.sendMessage(ctx, req, len(req.BroadcastList) != 0)
}

// SendRichMediaMessage represents method to send rich media message.
func (api *ViberAPI) SendRichMediaMessage(ctx context.Context, req apiModels.SendRichMediaMessageRequest) (*apiModels.MessageResponse, error) {
	req.Type = models.MessageTypeRichMedia
	req.RichMedia.Type = string(models.MessageTypeRichMedia)
	return api.sendMessage(ctx, req, len(req.BroadcastList) != 0)
}

func (api *ViberAPI) sendMessage(ctx context.Context, req interface{}, asBroadcast bool) (*apiModels.MessageResponse, error) {
	apiMethod := "send_message"
	if asBroadcast {
		apiMethod = "broadcast_message"
	}

	resp, err := api.NewRequest(apiMethod).
		Body(NewJSONBody(req)).
		Method(http.MethodPost).
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
