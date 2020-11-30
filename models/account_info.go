package models

// AccountInfo represents account details info.
// https://developers.viber.com/docs/api/rest-bot-api/#response-1.
type AccountInfo struct {
	// Action result. 0 for success. In case of failure – appropriate failure status number.
	// See error codes (https://developers.viber.com/docs/api/rest-bot-api/#errorCodes)
	// table for additional information.
	Status ResponseStatusCode `json:"status"`
	// OK or failure reason. Success: ok.
	// Failure: invalidUrl, invalidAuthToken, badData, missingData and failure.
	// See error codes (https://developers.viber.com/docs/api/rest-bot-api/#errorCodes)
	// table for additional information.
	StatusMessage ResponseStatusName `json:"status_message"`
	// Unique numeric id of the account.
	ID string `json:"id"`
	// Account name. Max 75 characters.
	Name string `json:"name"`
	// Unique URI of the Account.
	URI string `json:"uri"`
	// Account icon URL. JPEG, 720x720, size no more than 512 kb.
	Icon string `json:"icon"`
	// Conversation background URL. JPEG, max 1920x1920, size no more than 512 kb.
	Background string `json:"background"`
	// Account category.
	Category string `json:"category"`
	// Account sub-category.
	Subcategory string `json:"subcategory"`
	// Account location (coordinates). Will be used for finding accounts near me
	// (lat & lon coordinates).
	Location Location `json:"location"`
	// Account country. 2 letters country code - ISO ALPHA-2 Code.
	Country string `json:"country"`
	// Account registered webhook (webhook URL).
	Webhook string `json:"webhook"`
	// Account registered events – as set by set_webhook request.
	EventTypes []EventType `json:"event_types"`
	// Number of subscribers.
	SubscribersCount int64 `json:"subscribers_count"`
	// Members of the bot’s public chat. id, name, avatar, role
	// for each Public Chat member (admin/participant). Deprecated.
	Members []Member `json:"members,omitempty"`
}
