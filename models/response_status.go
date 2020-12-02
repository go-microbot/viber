package models

// There are available Viber response status codes.
const (
	ResponseStatusCodeOK                           ResponseStatusCode = iota
	ResponseStatusCodeInvalidURL                   ResponseStatusCode = iota
	ResponseStatusCodeInvalidStatusToken           ResponseStatusCode = iota
	ResponseStatusCodeBadData                      ResponseStatusCode = iota
	ResponseStatusCodeMissingData                  ResponseStatusCode = iota
	ResponseStatusCodeReceiverNotRegistered        ResponseStatusCode = iota
	ResponseStatusCodeReceiverNotSubscribed        ResponseStatusCode = iota
	ResponseStatusCodePublicAccountBlocked         ResponseStatusCode = iota
	ResponseStatusCodePublicAccountNotFound        ResponseStatusCode = iota
	ResponseStatusCodePublicAccountSuspended       ResponseStatusCode = iota
	ResponseStatusCodeWebhookNotSet                ResponseStatusCode = iota
	ResponseStatusCodeReceiverNoSuitableDevice     ResponseStatusCode = iota
	ResponseStatusCodeTooManyRequests              ResponseStatusCode = iota
	ResponseStatusCodeAPIVersionNotSupported       ResponseStatusCode = iota
	ResponseStatusCodeIncompatibleWithVersion      ResponseStatusCode = iota
	ResponseStatusCodePublicAccountNotAuthorized   ResponseStatusCode = iota
	ResponseStatusCodeInchatReplyMessageNotAllowed ResponseStatusCode = iota
	ResponseStatusCodePublicAccountIsNotInline     ResponseStatusCode = iota
	ResponseStatusCodeNoPublicChat                 ResponseStatusCode = iota
	ResponseStatusCodeCannotSendBroadcast          ResponseStatusCode = iota
	ResponseStatusCodeBroadcastNotAllowed          ResponseStatusCode = iota
)

// ResponseStatusCode represents response Viber status code.
type ResponseStatusCode int64
