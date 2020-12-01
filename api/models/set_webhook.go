package models

import "github.com/go-microbot/viber/models"

// SetWebhookRequest represents `set_webhook` request body.
type SetWebhookRequest struct {
	// Required. Account webhook URL to receive callbacks & messages from users.
	// Webhook URL must use SSL.
	// Note: Viber doesn’t support self signed certificates.
	URL string `json:"url"`
	// Optional. Indicates the types of Viber events that the account owner
	// would like to be notified about.
	// Don’t include this parameter in your request to get all events.
	// Possible values: "delivered", "seen", "failed", "subscribed",
	// "unsubscribed" and "conversation_started".
	EventTypes []models.EventType `json:"event_types,omitempty"`
	// Optional. Indicates whether or not the bot should receive the user name.
	// Default false. Possible values: true, false.
	SendName bool `json:"send_name,omitempty"`
	// Optional. Indicates whether or not the bot should receive the user photo.
	// Default false. Possible values: true, false.
	SendPhoto bool `json:"send_photo,omitempty"`
}

// SetWebhookResponse represents `set_webhook` response body.
type SetWebhookResponse struct {
	// Action result. 0 for success.
	// In case of failure – appropriate failure status number.
	// See error codes (https://developers.viber.com/docs/api/rest-bot-api/#errorCodes)
	// table for additional information.
	Status models.ResponseStatusCode `json:"status"`
	// OK or failure reason. Success: ok.
	// Failure: "invalidUrl", "invalidAuthToken", "badData", "missingData" and "failure".
	// See error codes (https://developers.viber.com/docs/api/rest-bot-api/#errorCodes)
	// table for additional information.
	StatusMessage models.ResponseStatusName `json:"status_message"`
	// List of event types you will receive a callback for.
	// Should return the same values sent in the request.
	// "delivered", "seen", "failed", "subscribed", "unsubscribed" and "conversation_started".
	EventTypes []models.EventType `json:"event_types,omitempty"`
}
