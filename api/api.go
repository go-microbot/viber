package api

import (
	"context"

	apiModels "github.com/go-microbot/viber/api/models"
)

// Bot represents general Bot API interface.
// https://developers.viber.com/docs/api/rest-bot-api/#get-started.
type Bot interface {
	// https://developers.viber.com/docs/api/rest-bot-api/#get-account-info.
	GetAccountInfo(ctx context.Context) (*apiModels.AccountInfoResponse, error)
	// https://developers.viber.com/docs/api/rest-bot-api/#setting-a-webhook.
	SetWebhook(ctx context.Context, req apiModels.SetWebhookRequest) (*apiModels.SetWebhookResponse, error)
	// https://developers.viber.com/docs/api/rest-bot-api/#removing-your-webhook.
	RemoveWebhook(ctx context.Context) error
	// https://developers.viber.com/docs/api/rest-bot-api/#text-message.
	SendTextMessage(ctx context.Context, req apiModels.SendTextMessageRequest) (*apiModels.MessageResponse, error)
	// https://developers.viber.com/docs/api/rest-bot-api/#picture-message.
	SendPictureMessage(ctx context.Context, req apiModels.SendPictureMessageRequest) (*apiModels.MessageResponse, error)
	// https://developers.viber.com/docs/api/rest-bot-api/#video-message.
	SendVideoMessage(ctx context.Context, req apiModels.SendVideoMessageRequest) (*apiModels.MessageResponse, error)
	// https://developers.viber.com/docs/api/rest-bot-api/#file-message.
	SendFileMessage(ctx context.Context, req apiModels.SendFileMessageRequest) (*apiModels.MessageResponse, error)
	// https://developers.viber.com/docs/api/rest-bot-api/#contact-message.
	SendContactMessage(ctx context.Context, req apiModels.SendContactMessageRequest) (*apiModels.MessageResponse, error)
	// https://developers.viber.com/docs/api/rest-bot-api/#location-message.
	SendLocationMessage(ctx context.Context, req apiModels.SendLocationMessageRequest) (*apiModels.MessageResponse, error)
	// https://developers.viber.com/docs/api/rest-bot-api/#url-message.
	SendURLMessage(ctx context.Context, req apiModels.SendURLMessageRequest) (*apiModels.MessageResponse, error)
	// https://developers.viber.com/docs/api/rest-bot-api/#sticker-message.
	SendStickerMessage(ctx context.Context, req apiModels.SendStickerMessageRequest) (*apiModels.MessageResponse, error)
	// https://developers.viber.com/docs/api/rest-bot-api/#rich-media-message--carousel-content-message.
	SendRichMediaMessage(ctx context.Context, req apiModels.SendRichMediaMessageRequest) (*apiModels.MessageResponse, error)
	// https://developers.viber.com/docs/api/rest-bot-api/#get-user-details.
	GetUserDetails(ctx context.Context, req apiModels.UserIDRequest) (*apiModels.UserDetailsResponse, error)
	// https://developers.viber.com/docs/api/rest-bot-api/#get-online.
	GetOnline(ctx context.Context, req apiModels.UserIDsRequest) (*apiModels.GetOnlineResponse, error)
}
