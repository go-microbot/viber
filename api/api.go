package api

import (
	"context"

	"github.com/go-microbot/viber/models"
)

// Bot represents general Bot API interface.
// https://developers.viber.com/docs/api/rest-bot-api/#get-started.
type Bot interface {
	// https://developers.viber.com/docs/api/rest-bot-api/#get-account-info.
	GetAccountInfo(ctx context.Context) (*models.AccountInfo, error)
	// https://core.telegram.org/bots/api#getupdates.
	/*GetUpdates(ctx context.Context, req apiModels.GetUpdatesRequest) ([]models.Update, error)
	GetPollUpdates(ctx context.Context, req apiModels.GetUpdatesRequest, client *http.Client) ([]models.Update, error)
	// https://core.telegram.org/bots/api#setwebhook.
	SetWebhook(ctx context.Context, req apiModels.SetWebhookRequest) error
	// https://core.telegram.org/bots/api#getwebhookinfo.
	GetWebhookInfo(ctx context.Context) (*models.WebhookInfo, error)
	// https://core.telegram.org/bots/api#deletewebhook.
	DeleteWebhook(ctx context.Context) error
	// https://core.telegram.org/bots/api#sendmessage.
	SendMessage(ctx context.Context, req apiModels.SendMessageRequest) (*models.Message, error)
	// https://core.telegram.org/bots/api#logout.
	Logout(ctx context.Context) error
	// https://core.telegram.org/bots/api#close.
	Close(ctx context.Context) error
	// https://core.telegram.org/bots/api#sendphoto.
	SendPhoto(ctx context.Context, req apiModels.SendPhotoRequest) (*models.Message, error)
	// https://core.telegram.org/bots/api#setchatpermissions
	SetChatPermissions(ctx context.Context, req apiModels.SetChatPermissionsRequest) error
	// https://core.telegram.org/bots/api#setchatphoto.
	SetChatPhoto(ctx context.Context, req apiModels.SetChatPhotoRequest) error
	// https://core.telegram.org/bots/api#setchattitle.
	SetChatTitle(ctx context.Context, req apiModels.SetChatTitleRequest) error
	// https://core.telegram.org/bots/api#getchat.
	GetChat(ctx context.Context, req apiModels.ChatID) (*models.Chat, error)
	// https://core.telegram.org/bots/api#leavechat.
	LeaveChat(ctx context.Context, req apiModels.ChatID) error
	// https://core.telegram.org/bots/api#forwardmessage.
	ForwardMessage(ctx context.Context, req apiModels.ForwardMessageRequest) (*models.Message, error)
	// https://core.telegram.org/bots/api#setmycommands.
	SetMyCommands(ctx context.Context, req apiModels.SetMyCommandsRequest) error
	// https://core.telegram.org/bots/api#getmycommands.
	GetMyCommands(ctx context.Context) ([]models.BotCommand, error)
	// https://core.telegram.org/bots/api#setchatdescription.
	SetChatDescription(ctx context.Context, req apiModels.SetChatDescriptionRequest) error
	// https://core.telegram.org/bots/api#deletechatphoto.
	DeleteChatPhoto(ctx context.Context, req apiModels.ChatID) error
	// https://core.telegram.org/bots/api#sendlocation.
	SendLocation(ctx context.Context, req apiModels.SendLocationRequest) (*models.Message, error)
	// https://core.telegram.org/bots/api#getfile.
	GetFile(ctx context.Context, req apiModels.FileID) (*models.File, error)
	// https://core.telegram.org/bots/api#pinchatmessage.
	PinChatMessage(ctx context.Context, req apiModels.PinChatMessageRequest) error
	// https://core.telegram.org/bots/api#unpinchatmessage.
	UnpinChatMessage(ctx context.Context, req apiModels.UnpinChatMessageRequest) error
	// https://core.telegram.org/bots/api#unpinallchatmessages.
	UnpinAllChatMessages(ctx context.Context, req apiModels.ChatID) error
	// https://core.telegram.org/bots/api#getchatmember.
	GetChatMember(ctx context.Context, req apiModels.GetChatMemberRequest) (*models.ChatMember, error)
	// https://core.telegram.org/bots/api#exportchatinvitelink.
	ExportChatInviteLink(ctx context.Context, req apiModels.ChatID) (string, error)
	// https://core.telegram.org/bots/api#setchatadministratorcustomtitle.
	SetChatAdministratorCustomTitle(ctx context.Context, req apiModels.SetChatAdminCustomTitleRequest) error
	// https://core.telegram.org/bots/api#getchatadministrators.
	GetChatAdministrators(ctx context.Context, req apiModels.ChatID) ([]models.ChatMember, error)
	// https://core.telegram.org/bots/api#getchatmemberscount.
	GetChatMembersCount(ctx context.Context, req apiModels.ChatID) (int32, error)
	// https://core.telegram.org/bots/api#sendaudio.
	SendAudio(ctx context.Context, req apiModels.SendAudioRequest) (*models.Message, error)
	// https://core.telegram.org/bots/api#senddocument.
	SendDocument(ctx context.Context, req apiModels.SendDocumentRequest) (*models.Message, error)
	// https://core.telegram.org/bots/api#sendvideo.
	SendVideo(ctx context.Context, req apiModels.SendVideoRequest) (*models.Message, error)
	// https://core.telegram.org/bots/api#sendanimation.
	SendAnimation(ctx context.Context, req apiModels.SendAnimationRequest) (*models.Message, error)
	// https://core.telegram.org/bots/api#sendvoice.
	SendVoice(ctx context.Context, req apiModels.SendVoiceRequest) (*models.Message, error)
	// https://core.telegram.org/bots/api#deletemessage.
	DeleteMessage(ctx context.Context, req apiModels.DeleteMessageRequest) error
	// https://core.telegram.org/bots/api#copymessage.
	CopyMessage(ctx context.Context, req apiModels.CopyMessageRequest) (*models.MessageID, error)
	// https://core.telegram.org/bots/api#sendchataction.
	SendChatAction(ctx context.Context, req apiModels.SendChatActionRequest) error
	// https://core.telegram.org/bots/api#getuserprofilephotos.
	GetUserProfilePhotos(ctx context.Context, req apiModels.GetUserProfilePhotosRequest) (*models.UserProfilePhotos, error)
	// https://core.telegram.org/bots/api#setchatstickerset.
	SetChatStickerSet(ctx context.Context, req apiModels.SetChatStickerSetRequest) error
	// https://core.telegram.org/bots/api#createnewstickerset.
	CreateNewStickerSet(ctx context.Context, req apiModels.CreateNewStickerSetRequest) error
	// https://core.telegram.org/bots/api#sendvideonote.
	SendVideoNote(ctx context.Context, req apiModels.SendVideoNoteRequest) (*models.Message, error)
	// https://core.telegram.org/bots/api#deletechatstickerset.
	DeleteChatStickerSet(ctx context.Context, req apiModels.ChatID) error
	// https://core.telegram.org/bots/api#sendmediagroup.
	SendMediaGroup(ctx context.Context, req apiModels.SendMediaGroupRequest) ([]models.Message, error)
	// https://core.telegram.org/bots/api#senddice.
	SendDice(ctx context.Context, req apiModels.SendDiceRequest) (*models.Message, error)
	// https://core.telegram.org/bots/api#sendvenue.
	SendVenue(ctx context.Context, req apiModels.SendVenueRequest) (*models.Message, error)
	// https://core.telegram.org/bots/api#sendcontact.
	SendContact(ctx context.Context, req apiModels.SendContactRequest) (*models.Message, error)
	// https://core.telegram.org/bots/api#sendpoll.
	SendPoll(ctx context.Context, req apiModels.SendPollRequest) (*models.Message, error)
	// https://core.telegram.org/bots/api#stoppoll.
	StopPoll(ctx context.Context, req apiModels.StopPollRequest) (*models.Poll, error)
	// https://core.telegram.org/bots/api#kickchatmember.
	KickChatMember(ctx context.Context, req apiModels.KickChatMemberRequest) error
	// https://core.telegram.org/bots/api#unbanchatmember.
	UnbanChatMember(ctx context.Context, req apiModels.UnbanChatMemberRequest) error
	// https://core.telegram.org/bots/api#restrictchatmember.
	RestrictChatMember(ctx context.Context, req apiModels.RestrictChatMemberRequest) error
	// https://core.telegram.org/bots/api#promotechatmember.
	PromoteChatMember(ctx context.Context, req apiModels.PromoteChatMemberRequest) error
	// https://core.telegram.org/bots/api#editmessagelivelocation.
	EditMessageLiveLocation(ctx context.Context, req apiModels.EditMessageLiveLocationRequest) error
	// https://core.telegram.org/bots/api#stopmessagelivelocation.
	StopMessageLiveLocation(ctx context.Context, req apiModels.StopMessageLiveLocationRequest) error
	// https://core.telegram.org/bots/api#getstickerset.
	GetStickerSet(ctx context.Context, req apiModels.GetStickerSetRequest) (*models.StickerSet, error)
	// https://core.telegram.org/bots/api#sendgame.
	SendGame(ctx context.Context, req apiModels.SendGameRequest) (*models.Message, error)
	// https://core.telegram.org/bots/api#setgamescore.
	SetGameScore(ctx context.Context, req apiModels.SetGameScoreRequest) (*models.Message, error)
	// https://core.telegram.org/bots/api#getgamehighscores.
	GetGameHighScores(ctx context.Context, req apiModels.GetGameHighScoresRequest) ([]models.GameHighScore, error)
	// https://core.telegram.org/bots/api#uploadstickerfile.
	UploadStickerFile(ctx context.Context, req apiModels.UploadStickerFileRequest) (*models.File, error)
	// https://core.telegram.org/bots/api#addstickertoset.
	AddStickerToSet(ctx context.Context, req apiModels.AddStickerToSetRequest) error
	// https://core.telegram.org/bots/api#sendsticker.
	SendSticker(ctx context.Context, req apiModels.SendStickerRequest) (*models.Message, error)
	// https://core.telegram.org/bots/api#setstickerpositioninset.
	SetStickerPositionInSet(ctx context.Context, req apiModels.SetStickerPositionInSetRequest) error
	// https://core.telegram.org/bots/api#setstickersetthumb.
	SetStickerSetThumb(ctx context.Context, req apiModels.SetStickerSetThumbRequest) error
	// https://core.telegram.org/bots/api#deletestickerfromset.
	DeleteStickerFromSet(ctx context.Context, req apiModels.DeleteStickerFromSetRequest) error
	// https://core.telegram.org/bots/api#editmessagetext.
	EditMessageText(ctx context.Context, req apiModels.EditMessageTextRequest) error
	// https://core.telegram.org/bots/api#editmessagecaption.
	EditMessageCaption(ctx context.Context, req apiModels.EditMessageCaptionRequest) error
	// https://core.telegram.org/bots/api#editmessagemedia.
	EditMessageMedia(ctx context.Context, req apiModels.EditMessageMediaRequest) error
	// https://core.telegram.org/bots/api#editmessagereplymarkup.
	EditMessageReplyMarkup(ctx context.Context, req apiModels.EditMessageReplyMarkupRequest) error
	// https://core.telegram.org/bots/api#answercallbackquery.
	AnswerCallbackQuery(ctx context.Context, req apiModels.AnswerCallbackQueryRequest) error
	// https://core.telegram.org/bots/api#answerinlinequery.
	AnswerInlineQuery(ctx context.Context, req apiModels.AnswerInlineQueryRequest) error
	// https://core.telegram.org/bots/api#sendinvoice.
	SendInvoice(ctx context.Context, req apiModels.SendInvoiceRequest) (*models.Message, error)
	// https://core.telegram.org/bots/api#setpassportdataerrors.
	SetPassportDataErrors(ctx context.Context, req apiModels.SetPassportDataErrorsRequest) error
	// https://core.telegram.org/bots/api#answershippingquery.
	AnswerShippingQuery(ctx context.Context, req apiModels.AnswerShippingQueryRequest) error
	// https://core.telegram.org/bots/api#answerprecheckoutquery.
	AnswerPreCheckoutQuery(ctx context.Context, req apiModels.AnswerPreCheckoutQueryRequest) error*/
}
