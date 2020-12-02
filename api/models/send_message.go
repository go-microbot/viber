package models

import "github.com/go-microbot/viber/models"

// There are available button background image types.
const (
	BgButtonMediaTypePicture BgButtonMediaType = "picture"
	BgButtonMediaTypeGif     BgButtonMediaType = "gif"
)

// There are available button background scale types.
const (
	BgButtonMediaScaleTypeCrop BgButtonMediaScaleType = "crop"
	BgButtonMediaScaleTypeFill BgButtonMediaScaleType = "fill"
	BgButtonMediaScaleTypeFit  BgButtonMediaScaleType = "fit"
)

// There are available button action types.
const (
	ButtonActionTypeReply          ButtonActionType = "reply"
	ButtonActionTypeOpenURL        ButtonActionType = "open-url"
	ButtonActionTypeLocationPicker ButtonActionType = "location-picker"
	ButtonActionTypeSharePhone     ButtonActionType = "share-phone"
	ButtonActionTypeNone           ButtonActionType = "none"
)

// There are available button vertical align text types.
const (
	ButtonTextVAlignTop    ButtonTextVAlign = "top"
	ButtonTextVAlignMiddle ButtonTextVAlign = "middle"
	ButtonTextVAlignBottom ButtonTextVAlign = "bottom"
)

// There are available button horizontal align text types.
const (
	ButtonTextHAlignLeft   ButtonTextHAlign = "left"
	ButtonTextHAlignCenter ButtonTextHAlign = "center"
	ButtonTextHAlignRight  ButtonTextHAlign = "right"
)

// There are available button text sizes.
const (
	ButtonTextSizeSmall   ButtonTextSize = "small"
	ButtonTextSizeRegular ButtonTextSize = "regular"
	ButtonTextSizeLarge   ButtonTextSize = "large"
)

// There are available button open URL types.
const (
	ButtonOpenURLTypeInternal ButtonOpenURLType = "internal"
	ButtonOpenURLTypeExternal ButtonOpenURLType = "external"
)

// There are available button open URL media types.
const (
	ButtonOpenURLMediaTypeNotMedia ButtonOpenURLMediaType = "not-media"
	ButtonOpenURLMediaTypeVideo    ButtonOpenURLMediaType = "video"
	ButtonOpenURLMediaTypeGif      ButtonOpenURLMediaType = "gif"
	ButtonOpenURLMediaTypePicture  ButtonOpenURLMediaType = "picture"
)

// There are availbale internal browser action buttons.
const (
	InternalBrowserActionButtonForward        InternalBrowserActionButton = "forward"
	InternalBrowserActionButtonSend           InternalBrowserActionButton = "send"
	InternalBrowserActionButtonOpenExternally InternalBrowserActionButton = "open-externally"
	InternalBrowserActionButtonSendToBot      InternalBrowserActionButton = "send-to-bot"
	InternalBrowserActionButtonNone           InternalBrowserActionButton = "none"
)

// There are available internal browser title types.
const (
	InternalBrowserTitleTypeDomain  InternalBrowserTitleType = "domain"
	InternalBrowserTitleTypeDefault InternalBrowserTitleType = "default"
)

// There are available internal browser mods.
const (
	InternalBrowserModeFullscreen          InternalBrowserMode = "fullscreen"
	InternalBrowserModeFullscreenPortrait  InternalBrowserMode = "fullscreen-portrait"
	InternalBrowserModeFullscreenLandscape InternalBrowserMode = "fullscreen-landscape"
	InternalBrowserModePartialSize         InternalBrowserMode = "partial-size"
)

// There are available internal browser footer types.
const (
	InternalBrowserFooterTypeDefault InternalBrowserFooterType = "default"
	InternalBrowserFooterTypeHidden  InternalBrowserFooterType = "hidden"
)

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

// SendVideoMessageRequest represents model to send video message.
type SendVideoMessageRequest struct {
	GeneralMessageRequest
	// URL of the video (MP4, H264).
	// Required. Max size 26 MB. Only MP4 and H264 are supported.
	// The URL must have a resource with a .mp4 file extension as the last path segment.
	// Example: http://www.example.com/path/video.mp4.
	Media string `json:"media"`
	// Size of the video in bytes. Required.
	Size int64 `json:"size"`
	// Video duration in seconds; will be displayed to the receiver.
	// Optional. Max 180 seconds.
	Duration int64 `json:"duration,omitempty"`
	// URL of a reduced size image (JPEG).
	// Optional. Max size 100 kb.
	// Recommended: 400x400. Only JPEG format is supported.
	Thumbnail string `json:"thumbnail,omitempty"`
}

// SendFileMessageRequest represents model to send file message.
type SendFileMessageRequest struct {
	GeneralMessageRequest
	// URL of the file. Required. Max size 50 MB.
	// See forbidden file formats
	// (https://developers.viber.com/docs/api/rest-bot-api/#forbiddenFileFormats)
	// for unsupported file types.
	Media string `json:"media"`
	// Size of the file in bytes. Required.
	Size int64 `json:"size"`
	// Name of the file. Required. File name should include extension.
	// Max 256 characters (including file extension).
	// Sending a file without extension or with the wrong extension
	// might cause the client to be unable to open the file.
	FileName string `json:"file_name"`
}

// SendContactMessageRequest represents model to send contact message.
type SendContactMessageRequest struct {
	GeneralMessageRequest
	// Contact to send.
	Contact MessageContact `json:"contact"`
}

// MessageContact represents message contact request model.
type MessageContact struct {
	// Name of the contact. Required. Max 28 characters.
	Name string `json:"name"`
	// Phone number of the contact. Required. Max 18 characters.
	PhoneNumber string `json:"phone_number"`
}

// SendLocationMessageRequest represents model to send location message.
type SendLocationMessageRequest struct {
	GeneralMessageRequest
	// Location coordinates. Required.
	// Latitude (±90°) & Longitude (±180°) within valid ranges.
	Location models.Location `json:"location"`
}

// SendURLMessageRequest represents model to send URL message.
type SendURLMessageRequest struct {
	GeneralMessageRequest
	// URL. Required. Max 2,000 characters.
	Media string `json:"media"`
}

// SendStickerMessageRequest represents model to send sticker message.
type SendStickerMessageRequest struct {
	GeneralMessageRequest
	// Unique Viber sticker ID. For examples visit the
	// sticker IDs (https://developers.viber.com/docs/tools/sticker-ids/) page.
	StickerID int64 `json:"sticker_id"`
}

// SendRichMediaMessageRequest represents model to send rich media message.
type SendRichMediaMessageRequest struct {
	GeneralMessageRequest
	// Backward compatibility text, limited to 7,000 characters.
	AltText string `json:"alt_text,omitempty"`
	// Rich Media request.
	RichMedia MessageRichMedia `json:"rich_media"`
}

// MessageRichMedia represents message rich media request.
type MessageRichMedia struct {
	// Type.
	Type string `json:"Type"`
	// Number of columns per carousel content block. Default 6 columns. 1-6.
	ButtonsGroupColumns int64 `json:"ButtonsGroupColumns,omitempty"`
	// Number of rows per carousel content block. Default 7 rows. 1-7.
	ButtonsGroupRows int64 `json:"ButtonsGroupRows,omitempty"`
	// Background color.
	BgColor string `json:"BgColor,omitempty"`
	// Array of buttons. Max of 6 * ButtonsGroupColumns * ButtonsGroupRows.
	Buttons []MessageButton `json:"Buttons,omitempty"`
}

// MessageButton represents message button request.
type MessageButton struct {
	// Button column span. Default ButtonsGroupColumns. 1..ButtonsGroupColumns.
	Columns int64 `json:"Columns,omitempty"`
	// Button row span. Default ButtonsGroupRows. 1..ButtonsGroupRows.
	Rows int64 `json:"Rows,omitempty"`
	// Optional. Background color of button. Valid color HEX value.
	// Default Viber button color.
	BgColor string `json:"BgColor,omitempty"`
	// Optional. Determine whether the user action is presented in the conversation.
	// True/False, default is False.
	Silent bool `json:"Silent,omitempty"`
	// Optional. Type of the background media.
	BgMediaType BgButtonMediaType `json:"BgMediaType,omitempty"`
	// Optional. URL for background media content (picture or gif).
	// Will be placed with aspect to fill logic.
	BgMedia string `json:"BgMedia,omitempty"`
	// Optional (api level 6).
	// Options for scaling the bounds of the background to the bounds of this view:
	// crop - contents scaled to fill with fixed aspect. Some portion of content may be clipped.
	// fill - contents scaled to fill without saving fixed aspect.
	// fit - at least one axis (X or Y) will fit exactly, aspect is saved.
	BgMediaScaleType BgButtonMediaScaleType `json:"BgMediaScaleType,omitempty"`
	// Optional (api level 6).
	// Options for scaling the bounds of an image to the bounds of this view:
	// crop - contents scaled to fill with fixed aspect. Some portion of content may be clipped.
	// fill - contents scaled to fill without saving fixed aspect.
	// fit - at least one axis (X or Y) will fit exactly, aspect is saved.
	ImageScaleType BgButtonMediaScaleType `json:"ImageScaleType,omitempty"`
	// Optional. When true - animated background media (gif) will loop continuously.
	// When false - animated background media will play once and stop. Default is true.
	BgLoop bool `json:"BgLoop,omitempty"`
	// Optional. Text to be displayed on the button.
	// Can contain some HTML tags - see keyboard design
	// (https://developers.viber.com/docs/tools/keyboards/#keyboardDesign) for more details.
	// Free text. Valid and allowed HTML tags Max 250 characters.
	// If the text is too long to display on the button it will be cropped and ended with "…".
	Text string `json:"Text,omitempty"`
	// Optional. Type of action pressing the button will perform.
	// reply - will send a reply to the PA.
	// open-url - will open the specified URL and send the URL as reply to the PA.
	// See reply logic (https://developers.viber.com/docs/tools/keyboards/#replyLogic)
	// for more details. Note: location-picker and share-phone are not supported on desktop,
	// and require adding any text in the ActionBody parameter. Default is reply.
	ActionType ButtonActionType `json:"ActionType,omitempty"`
	// Required. Text for reply and none.
	// ActionType or URL for open-url.
	// See reply logic (https://developers.viber.com/docs/tools/keyboards/#replyLogic)
	// for more details. Possible values: for ActionType reply - text,
	// for ActionType open-url - Valid URL.
	ActionBody string `json:"ActionBody"`
	// Optional. URL of image to place on top of background (if any).
	// Can be a partially transparent image that will allow showing some of the background.
	// Will be placed with aspect to fill logic. Valid URL.
	// JPEG and PNG files are supported. Max size: 500 kb.
	Image string `json:"Image,omitempty"`
	// Optional. Text size out of 3 available options. Default is regular.
	TextSize ButtonTextSize `json:"TextSize,omitempty"`
	// Optional. Vertical alignment of the text. Default is middle.
	TextVAlign ButtonTextVAlign `json:"TextVAlign,omitempty"`
	// Optional. Horizontal align of the text. Default is center.
	TextHAlign ButtonTextHAlign `json:"TextHAlign,omitempty"`
	// Optional (api level 4). Custom paddings for the text in points.
	// The value is an array of Integers [top, left, bottom, right].
	// Default is [12,12,12,12].
	TextPaddings []int64 `json:"TextPaddings,omitempty"`
	// Optional. Text opacity. 0-100. Default is 100.
	TextOpacity int64 `json:"TextOpacity,omitempty"`
	// Optional. Determine the open-url action result, in app or external browser.
	// Default is internal.
	OpenURLType ButtonOpenURLType `json:"OpenURLType,omitempty"`
	// Optional. Determine the url media type.
	// not-media - force browser usage.
	// video - will be opened via media player.
	// gif - client will play the gif in full screen mode.
	// picture - client will open the picture in full screen mode.
	// Default is not-media.
	OpenURLMediaType ButtonOpenURLMediaType `json:"OpenURLMediaType,omitempty"`
	// Optional. Background gradient to use under text, works only when TextVAlign
	// is equal to top or bottom. Hex value (6 characters).
	TextBgGradientColor string `json:"TextBgGradientColor,omitempty"`
	// Optional. (api level 6) If true the size of text will decreased to fit (minimum size is 12).
	// Default is false.
	TextShouldFit bool `json:"TextShouldFit,omitempty"`
	// Optional (api level 3).
	// JSON Object, which includes internal browser configuration
	// for open-url action with internal type.
	InternalBrowser *InternalBrowserConfig `json:"InternalBrowser,omitempty"`
	// Optional (api level 6).
	// JSON Object, which includes map configuration for open-map action with internal type.
	Map *MapConfig `json:"Map,omitempty"`
	// Optional (api level 6). JSON Object.
	// Draw frame above the background on the button,
	// the size will be equal the size of the button.
	Frame *FrameConfig `json:"Frame,omitempty"`
	// Optional (api level 6). JSON Object.
	// Specifies media player options. Will be ignored if OpenURLMediaType is not video or audio.
	MediaPlayer *MediaPlayerConfig `json:"MediaPlayer,omitempty"`
}

// BgButtonMediaType represents type of the background media button.
type BgButtonMediaType string

// BgButtonMediaScaleType represents type of the background media button scale type.
type BgButtonMediaScaleType string

// ButtonActionType represents button action type.
type ButtonActionType string

// ButtonTextVAlign represents button text vertical align.
type ButtonTextVAlign string

// ButtonTextHAlign represents button text horizontal align.
type ButtonTextHAlign string

// ButtonTextSize represents button text size.
type ButtonTextSize string

// ButtonOpenURLType represents button open URL type.
type ButtonOpenURLType string

// ButtonOpenURLMediaType represents button open URL media type.
type ButtonOpenURLMediaType string

// InternalBrowserConfig represents internal browser configuration.
type InternalBrowserConfig struct {
	// Optional (api level 3). Action button in internal’s browser navigation bar.
	// forward - will open the forward via Viber screen and share current URL or predefined URL.
	// send - sends the currently opened URL as an URL message, or predefined URL
	// if property ActionPredefinedURL is not empty.
	// open-externally - opens external browser with the current URL.
	// send-to-bot - (api level 6) sends reply data in msgInfo to bot in order to receive message.
	// none - will not display any button.
	// Default is forward.
	ActionButton InternalBrowserActionButton `json:"ActionButton,omitempty"`
	// Optional (api level 3).
	// If ActionButton is send or forward then the value from this property will be used
	// to be sent as message, otherwise ignored.
	ActionPredefinedURL string `json:"ActionPredefinedURL,omitempty"`
	// Optional (api level 3).
	// Type of title for internal browser if has no CustomTitle field.
	// default means the content in the page’s <OG:title> element or in <title> tag.
	// domain means the top level domain.
	// Default is default.
	TitleType InternalBrowserTitleType `json:"TitleType,omitempty"`
	// Optional (api level 3). Custom text for internal’s browser title,
	// TitleType will be ignored in case this key is presented. String up to 15 characters.
	CustomTitle string `json:"CustomTitle,omitempty"`
	// Optional (api level 3).
	// Indicates that browser should be opened in a full screen or in partial size
	// (50% of screen height). Full screen mode can be with orientation lock
	// (both orientations supported, only landscape or only portrait).
	// Default is fullscreen.
	Mode InternalBrowserMode `json:"Mode,omitempty"`
	// Optional (api level 3).
	// Should the browser’s footer will be displayed (default) or not (hidden).
	// Default is default.
	FooterType InternalBrowserFooterType `json:"FooterType,omitempty"`
	// Optional (api level 6).
	// Custom reply data for send-to-bot action that will be resent in msgInfo.
	ActionReplyData string `json:"ActionReplyData,omitempty"`
}

// InternalBrowserActionButton represents internal browser action button.
type InternalBrowserActionButton string

// InternalBrowserTitleType represents internal browser title type.
type InternalBrowserTitleType string

// InternalBrowserMode represents internal browser mode.
type InternalBrowserMode string

// InternalBrowserFooterType represents internal browser footer type.
type InternalBrowserFooterType string

// MapConfig represents map configuration.
type MapConfig struct {
	// Optional (api level 6). Location latitude (format: "12.12345").
	Latitude float64 `json:"Latitude,omitempty"`
	// Optional (api level 6). Location longitude (format: "3.12345").
	Longitude float64 `json:"Longitude,omitempty"`
}

// FrameConfig represents frame configuration.
type FrameConfig struct {
	// Optional (api level 6). Width of border. 0..10. Default is 1.
	BorderWidth int64 `json:"BorderWidth,omitempty"`
	// Optional (api level 6). Color of border. Hex color #XXXXXX. Default is #000000.
	BorderColor string `json:"BorderColor,omitempty"`
	// Optional (api level 6). The border will be drawn with rounded corners. 0..10. Default is 0.
	CornerRadius int64 `json:"CornerRadius,omitempty"`
}

// MediaPlayerConfig represents media player configuration.
type MediaPlayerConfig struct {
	// Optional (api level 6). Media player’s title (first line).
	Title string `json:"Title,omitempty"`
	// Optional (api level 6). Media player’s subtitle (second line).
	Subtitle string `json:"Subtitle,omitempty"`
	// Optional (api level 6). The URL for player’s thumbnail (background).
	ThumbnailURL string `json:"ThumbnailURL,omitempty"`
	// Optional (api level 6). Whether the media player should be looped forever or not.
	// Default is false.
	Loop bool `json:"Loop,omitempty"`
}
