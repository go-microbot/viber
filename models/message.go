package models

// There are available message types.
const (
	MessageTypeText     MessageType = "text"
	MessageTypePicture  MessageType = "picture"
	MessageTypeVideo    MessageType = "video"
	MessageTypeFile     MessageType = "file"
	MessageTypeSticker  MessageType = "sticker"
	MessageTypeContact  MessageType = "contact"
	MessageTypeURL      MessageType = "url"
	MessageTypeLocation MessageType = "location"
)

// MessageType represents message type.
type MessageType string

// Message represents Viber message model.
type Message struct {
	// Message type.
	// "text", "picture", "video", "file", "sticker", "contact", "url" and "location".
	Type MessageType `json:"type"`
	// The message text.
	Text string `json:"text"`
	// URL of the message media - can be image, video, file and url.
	// Image/Video/File URLs will have a TTL of 1 hour.
	Media string `json:"media,omitempty"`
	// Location coordinates (lat & lon within valid ranges).
	Location *Location `json:"location,omitempty"`
	// Contact.
	Contact *Contact `json:"contact,omitempty"`
	// Tracking data sent with the last message to the user.
	TrackingData string `json:"tracking_data,omitempty"`
	// File name. Relevant for file type messages.
	FileName string `json:"file_name,omitempty"`
	// File size in bytes. Relevant for file type messages.
	FileSize int64 `json:"file_size,omitempty"`
	// Video length in seconds. Relevant for video type messages.
	Duration int64 `json:"duration,omitempty"`
	// Viber sticker id. Relevant for sticker type messages.
	StickerID int64 `json:"sticker_id,omitempty"`
}
