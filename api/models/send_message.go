package models

import "github.com/go-microbot/viber/models"

// GeneralMessageRequest represents general message request model.
type GeneralMessageRequest struct {
	// Unique Viber user id. Required, subscribed valid user ID.
	Receiver string `json:"receiver"`
	// Message type. required. Available message types: "text", "picture",
	// "video", "file", "location", "contact", "sticker", "carousel content" and "url".
	Type models.MessageType `json:"type"`
	// Sender.
	Sender MessageSender `json:"sender"`
	// Allow the account to track messages and user’s replies.
	// Sent tracking_data value will be passed back with user’s reply.
	// Optional. Max 4000 characters.
	TrackingData string `json:"tracking_data,omitempty"`
	// Minimal API version required by clients for this message (default 1).
	// Optional. Client version support the API version.
	// Certain features may not work as expected
	// if set to a number that’s below their requirements.
	MinAPIVersion int64 `json:"min_api_version,omitempty"`
}

// MessageSender represents message's sender model.
type MessageSender struct {
	// The sender’s name to display. Required. Max 28 characters.
	Name string `json:"name"`
	// The sender’s avatar URL. Optional.
	// Avatar size should be no more than 100 kb. Recommended 720x720.
	Avatar string `json:"avatar,omitempty"`
}

// SendTextMessageRequest represents model to send text message.
type SendTextMessageRequest struct {
	GeneralMessageRequest
	// The text of the message. Required. Max length 7,000 characters.
	Text string `json:"text"`
}

// MessageResponse represents default message response model.
type MessageResponse struct {
	Status        models.ResponseStatusCode `json:"status"`
	StatusMessage models.ResponseStatusName `json:"status_message"`
	MessageToken  int64                     `json:"message_token"`
	ChatHostname  string                    `json:"chat_hostname"`
}

// SendPictureMessageRequest represents model to send picture message.
type SendPictureMessageRequest struct {
	GeneralMessageRequest
	// Description of the photo. Can be an empty string if irrelevant.
	// Required. Max 120 characters.
	Text string `json:"text"`
	// URL of the image (JPEG, PNG, non-animated GIF).
	// Required. The URL must have a resource with a .jpeg, .png or .gif file extension
	// as the last path segment.
	// Example: http://www.example.com/path/image.jpeg.
	// Animated GIFs can be sent as URL messages or file messages.
	// Max image size: 1MB on iOS, 3MB on Android.
	Media string `json:"media"`
	// URL of a reduced size image (JPEG, PNG, GIF).
	// Optional. Recommended: 400x400. Max size: 100kb.
	Thumbnail string `json:"thumbnail,omitempty"`
}
