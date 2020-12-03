// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/go-microbot/viber/api/models"
	mock "github.com/stretchr/testify/mock"
)

// Bot is an autogenerated mock type for the Bot type
type Bot struct {
	mock.Mock
}

// GetAccountInfo provides a mock function with given fields: ctx
func (_m *Bot) GetAccountInfo(ctx context.Context) (*models.AccountInfoResponse, error) {
	ret := _m.Called(ctx)

	var r0 *models.AccountInfoResponse
	if rf, ok := ret.Get(0).(func(context.Context) *models.AccountInfoResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.AccountInfoResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOnline provides a mock function with given fields: ctx, req
func (_m *Bot) GetOnline(ctx context.Context, req models.UserIDsRequest) (*models.GetOnlineResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *models.GetOnlineResponse
	if rf, ok := ret.Get(0).(func(context.Context, models.UserIDsRequest) *models.GetOnlineResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.GetOnlineResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.UserIDsRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserDetails provides a mock function with given fields: ctx, req
func (_m *Bot) GetUserDetails(ctx context.Context, req models.UserIDRequest) (*models.UserDetailsResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *models.UserDetailsResponse
	if rf, ok := ret.Get(0).(func(context.Context, models.UserIDRequest) *models.UserDetailsResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.UserDetailsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.UserIDRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveWebhook provides a mock function with given fields: ctx
func (_m *Bot) RemoveWebhook(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendContactMessage provides a mock function with given fields: ctx, req
func (_m *Bot) SendContactMessage(ctx context.Context, req models.SendContactMessageRequest) (*models.MessageResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *models.MessageResponse
	if rf, ok := ret.Get(0).(func(context.Context, models.SendContactMessageRequest) *models.MessageResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.MessageResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.SendContactMessageRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendFileMessage provides a mock function with given fields: ctx, req
func (_m *Bot) SendFileMessage(ctx context.Context, req models.SendFileMessageRequest) (*models.MessageResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *models.MessageResponse
	if rf, ok := ret.Get(0).(func(context.Context, models.SendFileMessageRequest) *models.MessageResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.MessageResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.SendFileMessageRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendLocationMessage provides a mock function with given fields: ctx, req
func (_m *Bot) SendLocationMessage(ctx context.Context, req models.SendLocationMessageRequest) (*models.MessageResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *models.MessageResponse
	if rf, ok := ret.Get(0).(func(context.Context, models.SendLocationMessageRequest) *models.MessageResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.MessageResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.SendLocationMessageRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendPictureMessage provides a mock function with given fields: ctx, req
func (_m *Bot) SendPictureMessage(ctx context.Context, req models.SendPictureMessageRequest) (*models.MessageResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *models.MessageResponse
	if rf, ok := ret.Get(0).(func(context.Context, models.SendPictureMessageRequest) *models.MessageResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.MessageResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.SendPictureMessageRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendRichMediaMessage provides a mock function with given fields: ctx, req
func (_m *Bot) SendRichMediaMessage(ctx context.Context, req models.SendRichMediaMessageRequest) (*models.MessageResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *models.MessageResponse
	if rf, ok := ret.Get(0).(func(context.Context, models.SendRichMediaMessageRequest) *models.MessageResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.MessageResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.SendRichMediaMessageRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendStickerMessage provides a mock function with given fields: ctx, req
func (_m *Bot) SendStickerMessage(ctx context.Context, req models.SendStickerMessageRequest) (*models.MessageResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *models.MessageResponse
	if rf, ok := ret.Get(0).(func(context.Context, models.SendStickerMessageRequest) *models.MessageResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.MessageResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.SendStickerMessageRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendTextMessage provides a mock function with given fields: ctx, req
func (_m *Bot) SendTextMessage(ctx context.Context, req models.SendTextMessageRequest) (*models.MessageResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *models.MessageResponse
	if rf, ok := ret.Get(0).(func(context.Context, models.SendTextMessageRequest) *models.MessageResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.MessageResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.SendTextMessageRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendURLMessage provides a mock function with given fields: ctx, req
func (_m *Bot) SendURLMessage(ctx context.Context, req models.SendURLMessageRequest) (*models.MessageResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *models.MessageResponse
	if rf, ok := ret.Get(0).(func(context.Context, models.SendURLMessageRequest) *models.MessageResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.MessageResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.SendURLMessageRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendVideoMessage provides a mock function with given fields: ctx, req
func (_m *Bot) SendVideoMessage(ctx context.Context, req models.SendVideoMessageRequest) (*models.MessageResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *models.MessageResponse
	if rf, ok := ret.Get(0).(func(context.Context, models.SendVideoMessageRequest) *models.MessageResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.MessageResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.SendVideoMessageRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetWebhook provides a mock function with given fields: ctx, req
func (_m *Bot) SetWebhook(ctx context.Context, req models.SetWebhookRequest) (*models.SetWebhookResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *models.SetWebhookResponse
	if rf, ok := ret.Get(0).(func(context.Context, models.SetWebhookRequest) *models.SetWebhookResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.SetWebhookResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.SetWebhookRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
