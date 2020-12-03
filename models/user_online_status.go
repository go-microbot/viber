package models

// There are available user online statuses.
const (
	UserOnlineStatusOnline        UserOnlineStatus = 0
	UserOnlineStatusOffline       UserOnlineStatus = 1
	UserOnlineStatusUndisclosed   UserOnlineStatus = 2
	UserOnlineStatusInternalError UserOnlineStatus = 3
	UserOnlineStatusUnavailable   UserOnlineStatus = 4
)

// UserOnlineStatus represents user online status.
type UserOnlineStatus int64
