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

// There are available Viber response status names.
const (
	ResponseStatusNameOK                           ResponseStatusName = "ok"
	ResponseStatusNameInvalidURL                   ResponseStatusName = "invalidUrl"
	ResponseStatusNameInvalidStatusToken           ResponseStatusName = "invalidAuthToken"
	ResponseStatusNameBadData                      ResponseStatusName = "badData"
	ResponseStatusNameMissingData                  ResponseStatusName = "missingData"
	ResponseStatusNameReceiverNotRegistered        ResponseStatusName = "receiverNotRegistered"
	ResponseStatusNameReceiverNotSubscribed        ResponseStatusName = "receiverNotSubscribed"
	ResponseStatusNamePublicAccountBlocked         ResponseStatusName = "publicAccountBlocked"
	ResponseStatusNamePublicAccountNotFound        ResponseStatusName = "publicAccountNotFound"
	ResponseStatusNamePublicAccountSuspended       ResponseStatusName = "publicAccountSuspended"
	ResponseStatusNameWebhookNotSet                ResponseStatusName = "webhookNotSet"
	ResponseStatusNameReceiverNoSuitableDevice     ResponseStatusName = "receiverNoSuitableDevice"
	ResponseStatusNameTooManyRequests              ResponseStatusName = "tooManyRequests"
	ResponseStatusNameAPIVersionNotSupported       ResponseStatusName = "apiVersionNotSupported"
	ResponseStatusNameIncompatibleWithVersion      ResponseStatusName = "incompatibleWithVersion"
	ResponseStatusNamePublicAccountNotAuthorized   ResponseStatusName = "publicAccountNotAuthorized"
	ResponseStatusNameInchatReplyMessageNotAllowed ResponseStatusName = "inchatReplyMessageNotAllowed"
	ResponseStatusNamePublicAccountIsNotInline     ResponseStatusName = "publicAccountIsNotInline"
	ResponseStatusNameNoPublicChat                 ResponseStatusName = "noPublicChat"
	ResponseStatusNameCannotSendBroadcast          ResponseStatusName = "cannotSendBroadcast"
	ResponseStatusNameBroadcastNotAllowed          ResponseStatusName = "broadcastNotAllowed"
)

// ResponseStatusCode represents response Viber status code.
type ResponseStatusCode int64

// ResponseStatusName represents response Viber status name.
type ResponseStatusName string
