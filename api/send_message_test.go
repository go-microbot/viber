package api

import (
	"context"
	"math/rand"
	"strings"
	"testing"

	apiModels "github.com/go-microbot/viber/api/models"
	"github.com/go-microbot/viber/models"
	"github.com/stretchr/testify/require"
)

type sendTextMessage struct{}

func (h sendTextMessage) Test(ctx context.Context, t *testing.T) context.Context {
	memberID := ctx.Value(chatMemberIDCtxKey)
	require.NotNil(t, memberID)

	resp, err := testAPI.SendTextMessage(ctx, apiModels.SendTextMessageRequest{
		GeneralMessageRequest: apiModels.GeneralMessageRequest{
			Receiver: memberID.(string),
			Sender: apiModels.MessageSender{
				Name: "Bot",
			},
		},
		Text: "This is a test message",
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, models.ResponseStatusCodeOK, resp.Status)

	return ctx
}

type sendPictureMessage struct{}

func (h sendPictureMessage) Test(ctx context.Context, t *testing.T) context.Context {
	memberID := ctx.Value(chatMemberIDCtxKey)
	require.NotNil(t, memberID)
	pictureURL := ctx.Value(pictureToSendCtxKey)
	require.NotNil(t, pictureURL)

	resp, err := testAPI.SendPictureMessage(ctx, apiModels.SendPictureMessageRequest{
		GeneralMessageRequest: apiModels.GeneralMessageRequest{
			Receiver: memberID.(string),
			Sender: apiModels.MessageSender{
				Name: "Bot",
			},
		},
		Text:  "This is a test picture",
		Media: pictureURL.(string),
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, models.ResponseStatusCodeOK, resp.Status)

	return ctx
}

type sendVideoMessage struct{}

func (h sendVideoMessage) Test(ctx context.Context, t *testing.T) context.Context {
	memberID := ctx.Value(chatMemberIDCtxKey)
	require.NotNil(t, memberID)
	videoURL := ctx.Value(videoToSendCtxKey)
	require.NotNil(t, videoURL)
	videoSize := ctx.Value(videoToSendSizeCtxKey)
	require.NotNil(t, videoSize)

	resp, err := testAPI.SendVideoMessage(ctx, apiModels.SendVideoMessageRequest{
		GeneralMessageRequest: apiModels.GeneralMessageRequest{
			Receiver: memberID.(string),
			Sender: apiModels.MessageSender{
				Name: "Bot",
			},
		},
		Media: videoURL.(string),
		Size:  videoSize.(int64),
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, models.ResponseStatusCodeOK, resp.Status)

	return ctx
}

type sendFileMessage struct{}

func (h sendFileMessage) Test(ctx context.Context, t *testing.T) context.Context {
	memberID := ctx.Value(chatMemberIDCtxKey)
	require.NotNil(t, memberID)
	fileURL := ctx.Value(fileToSendCtxKey)
	require.NotNil(t, fileURL)
	fileSize := ctx.Value(fileToSendSizeCtxKey)
	require.NotNil(t, fileSize)

	index := strings.LastIndex(fileURL.(string), "/")
	fileName := string([]rune(fileURL.(string))[index+1:])
	resp, err := testAPI.SendFileMessage(ctx, apiModels.SendFileMessageRequest{
		GeneralMessageRequest: apiModels.GeneralMessageRequest{
			Receiver: memberID.(string),
			Sender: apiModels.MessageSender{
				Name: "Bot",
			},
		},
		Media:    fileURL.(string),
		Size:     fileSize.(int64),
		FileName: fileName,
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, models.ResponseStatusCodeOK, resp.Status)

	return ctx
}

type sendContactMessage struct{}

func (h sendContactMessage) Test(ctx context.Context, t *testing.T) context.Context {
	memberID := ctx.Value(chatMemberIDCtxKey)
	require.NotNil(t, memberID)

	resp, err := testAPI.SendContactMessage(ctx, apiModels.SendContactMessageRequest{
		GeneralMessageRequest: apiModels.GeneralMessageRequest{
			Receiver: memberID.(string),
			Sender: apiModels.MessageSender{
				Name: "Bot",
			},
		},
		Contact: apiModels.MessageContact{
			Name:        "Test contact",
			PhoneNumber: "+375250000000",
		},
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, models.ResponseStatusCodeOK, resp.Status)

	return ctx
}

type sendLocationMessage struct{}

func (h sendLocationMessage) Test(ctx context.Context, t *testing.T) context.Context {
	memberID := ctx.Value(chatMemberIDCtxKey)
	require.NotNil(t, memberID)

	resp, err := testAPI.SendLocationMessage(ctx, apiModels.SendLocationMessageRequest{
		GeneralMessageRequest: apiModels.GeneralMessageRequest{
			Receiver: memberID.(string),
			Sender: apiModels.MessageSender{
				Name: "Bot",
			},
		},
		Location: models.Location{
			Longitude: rand.Float64() * 80,
			Latitude:  rand.Float64() * 80,
		},
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, models.ResponseStatusCodeOK, resp.Status)

	return ctx
}

type sendURLMessage struct{}

func (h sendURLMessage) Test(ctx context.Context, t *testing.T) context.Context {
	memberID := ctx.Value(chatMemberIDCtxKey)
	require.NotNil(t, memberID)

	resp, err := testAPI.SendURLMessage(ctx, apiModels.SendURLMessageRequest{
		GeneralMessageRequest: apiModels.GeneralMessageRequest{
			Receiver: memberID.(string),
			Sender: apiModels.MessageSender{
				Name: "Bot",
			},
		},
		Media: "https://play.golang.org/",
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, models.ResponseStatusCodeOK, resp.Status)

	return ctx
}

type sendStickerMessage struct{}

func (h sendStickerMessage) Test(ctx context.Context, t *testing.T) context.Context {
	memberID := ctx.Value(chatMemberIDCtxKey)
	require.NotNil(t, memberID)
	stickerID := ctx.Value(stickerIDCtxKey)
	require.NotNil(t, stickerID)

	resp, err := testAPI.SendStickerMessage(ctx, apiModels.SendStickerMessageRequest{
		GeneralMessageRequest: apiModels.GeneralMessageRequest{
			Receiver: memberID.(string),
			Sender: apiModels.MessageSender{
				Name: "Bot",
			},
		},
		StickerID: stickerID.(int64),
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, models.ResponseStatusCodeOK, resp.Status)

	return ctx
}

type sendRichMediaMessage struct{}

func (h sendRichMediaMessage) Test(ctx context.Context, t *testing.T) context.Context {
	memberID := ctx.Value(chatMemberIDCtxKey)
	require.NotNil(t, memberID)
	pictureURL := ctx.Value(pictureToSendCtxKey)
	require.NotNil(t, pictureURL)

	resp, err := testAPI.SendRichMediaMessage(ctx, apiModels.SendRichMediaMessageRequest{
		GeneralMessageRequest: apiModels.GeneralMessageRequest{
			Receiver: memberID.(string),
			Sender: apiModels.MessageSender{
				Name: "Bot",
			},
			MinAPIVersion: 6,
		},
		RichMedia: apiModels.MessageRichMedia{
			ButtonsGroupColumns: 6,
			ButtonsGroupRows:    7,
			BgColor:             "#FFFFFF",
			Buttons: []apiModels.MessageButton{
				{
					Columns:    6,
					Rows:       3,
					ActionType: apiModels.ButtonActionTypeOpenURL,
					ActionBody: "https://www.google.com",
					Image:      pictureURL.(string),
				},
				{
					Columns:    6,
					Rows:       2,
					Text:       "<font color=#323232><b>Headphones with Microphone, On-ear Wired earphones</b></font><font color=#777777><br>Sound Intone </font><font color=#6fc133>$17.99</font>",
					ActionType: apiModels.ButtonActionTypeOpenURL,
					ActionBody: "https://www.google.com",
					TextSize:   apiModels.ButtonTextSizeRegular,
					TextVAlign: apiModels.ButtonTextVAlignMiddle,
					TextHAlign: apiModels.ButtonTextHAlignLeft,
				},
				{
					Columns:    6,
					Rows:       1,
					Text:       "<font color=#ffffff>Buy</font>",
					ActionType: apiModels.ButtonActionTypeReply,
					ActionBody: "https://www.google.com",
					TextSize:   apiModels.ButtonTextSizeLarge,
					TextVAlign: apiModels.ButtonTextVAlignMiddle,
					TextHAlign: apiModels.ButtonTextHAlignCenter,
					Image:      pictureURL.(string),
				},
				{
					Columns:    6,
					Rows:       1,
					Text:       "<font color=#8367db>MORE DETAILS</font>",
					ActionType: apiModels.ButtonActionTypeReply,
					ActionBody: "https://www.google.com",
					TextSize:   apiModels.ButtonTextSizeSmall,
					TextVAlign: apiModels.ButtonTextVAlignMiddle,
					TextHAlign: apiModels.ButtonTextHAlignCenter,
				},
				{
					Columns:    6,
					Rows:       3,
					Text:       "<font color=#8367db>MORE DETAILS</font>",
					ActionType: apiModels.ButtonActionTypeOpenURL,
					ActionBody: "https://www.google.com",
					Image:      pictureURL.(string),
				},
				{
					Columns:    6,
					Rows:       2,
					Text:       "<font color=#323232><b>Hanes Men's Humor Graphic T-Shirt</b></font><font color=#777777><br>Hanes</font><font color=#6fc133>$10.99</font>",
					ActionType: apiModels.ButtonActionTypeOpenURL,
					ActionBody: "https://www.google.com",
					TextSize:   apiModels.ButtonTextSizeRegular,
					TextVAlign: apiModels.ButtonTextVAlignMiddle,
					TextHAlign: apiModels.ButtonTextHAlignLeft,
				},
				{
					Columns:    6,
					Rows:       1,
					Text:       "<font color=#ffffff>Buy</font>",
					ActionType: apiModels.ButtonActionTypeReply,
					ActionBody: "https://www.google.com",
					TextSize:   apiModels.ButtonTextSizeLarge,
					TextVAlign: apiModels.ButtonTextVAlignMiddle,
					TextHAlign: apiModels.ButtonTextHAlignCenter,
					Image:      pictureURL.(string),
				},
				{
					Columns:    6,
					Rows:       1,
					Text:       "<font color=#8367db>MORE DETAILS</font>",
					ActionType: apiModels.ButtonActionTypeReply,
					ActionBody: "https://www.google.com",
					TextSize:   apiModels.ButtonTextSizeSmall,
					TextVAlign: apiModels.ButtonTextVAlignMiddle,
					TextHAlign: apiModels.ButtonTextHAlignCenter,
				},
			},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, models.ResponseStatusCodeOK, resp.Status)

	return ctx
}

type sendMessageWithKeyboard struct{}

func (h sendMessageWithKeyboard) Test(ctx context.Context, t *testing.T) context.Context {
	memberID := ctx.Value(chatMemberIDCtxKey)
	require.NotNil(t, memberID)

	resp, err := testAPI.SendTextMessage(ctx, apiModels.SendTextMessageRequest{
		GeneralMessageRequest: apiModels.GeneralMessageRequest{
			Receiver: memberID.(string),
			Sender: apiModels.MessageSender{
				Name: "Bot",
			},
			MinAPIVersion: 7,
			Keyboard: &apiModels.MessageKeyboard{
				DefaultHeight: true,
				Buttons: []apiModels.MessageButton{
					{
						ActionType: apiModels.ButtonActionTypeReply,
						ActionBody: "reply to me",
						Text:       "Key text",
						TextSize:   apiModels.ButtonTextSizeRegular,
					},
				},
			},
		},
		Text: "Test message with keyboard",
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, models.ResponseStatusCodeOK, resp.Status)

	return ctx
}

type sendBroadcastMessage struct{}

func (h sendBroadcastMessage) Test(ctx context.Context, t *testing.T) context.Context {
	memberID := ctx.Value(chatMemberIDCtxKey)
	require.NotNil(t, memberID)

	resp, err := testAPI.SendTextMessage(ctx, apiModels.SendTextMessageRequest{
		GeneralMessageRequest: apiModels.GeneralMessageRequest{
			Sender: apiModels.MessageSender{
				Name: "Bot",
			},
			Receiver:      memberID.(string),
			BroadcastList: []string{memberID.(string), "invalid"},
		},
		Text: "Test broadcast message",
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, models.ResponseStatusCodeOK, resp.Status)
	require.NotEmpty(t, resp.FailedList)
	var found bool
	for i := range resp.FailedList {
		if resp.FailedList[i].Receiver == "invalid" {
			require.Equal(t, resp.FailedList[i].Status, models.ResponseStatusCodeReceiverNotRegistered)
			found = true
			break
		}
	}
	require.True(t, found)

	return ctx
}
