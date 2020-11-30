package models

// There are available event types.
const (
	EventTypeDelivered           EventType = "delivered"
	EventTypeSeen                EventType = "seen"
	EventTypeFailed              EventType = "failed"
	EventTypeConversationStarted EventType = "conversation_started"
	EventTypeSubscribed          EventType = "subscribed"
	EventTypeMessage             EventType = "message"
	EventTypeUnsubscribed        EventType = "unsubscribed"
)

// There are available conversation started event types.
const (
	ConversationStartedTypeOpen ConversationStartedType = "open"
)

// EventType represents event type.
type EventType string

// ConversationStartedType represents type of the conversation_started event.
type ConversationStartedType string
