package models

// CallbackEvent represents general Viber Callback event.
type CallbackEvent struct {
	// Callback type - which event triggered the callback.
	Event EventType `json:"event"`
	// Time of the event that triggered the callback (Epoch time).
	Timestamp int64 `json:"timestamp"`
	// Unique ID of the message.
	MessageToken int64 `json:"message_token"`
	// Viber user info.
	User *User `json:"user,omitempty"`
	// Viber unique ID of the user.
	UserID string `json:"user_id,omitempty"`
	// The specific type of conversation_started event.
	// `open`. Additional types may be added in the future.
	Type ConversationStartedType `json:"type,omitempty"`
	// Any additional parameters added to the deep link used to access the conversation
	// passed as a string. See deep link (https://developers.viber.com/docs/tools/deep-links/)
	// section for additional information.
	Context string `json:"context,omitempty"`
	// Indicated whether a user is already subscribed.
	// True if subscribed and False otherwise.
	Subscribed bool `json:"subscribed,omitempty"`
	// A string describing the failure.
	Desc string `json:"desc,omitempty"`
	// Sender.
	Sender *User `json:"sender,omitempty"`
	// Message.
	Message *Message `json:"message,omitempty"`
}
