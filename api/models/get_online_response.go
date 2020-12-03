package models

import "github.com/go-microbot/viber/models"

// GetOnlineResponse represents `get_online` response model.
type GetOnlineResponse struct {
	MessageResponse
	Users []GetOnlineUser `json:"users"`
}

// GetOnlineUser represents user online status.
type GetOnlineUser struct {
	// Unique Viber user id.
	ID string `json:"id"`
	// Online status code.
	// 0 for online,
	// 1 for offline,
	// 2 for undisclosed - user set Viber to hide status,
	// 3 for try later - internal error,
	// 4 for unavailable - not a Viber user/unsubscribed/unregistered.
	OnlineStatus models.UserOnlineStatus `json:"online_status"`
	// Online status message.
	OnlineStatusMessage string `json:"online_status_message"`
}
