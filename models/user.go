package models

// User represents Viber user model.
type User struct {
	// Unique Viber user ID.
	ID string `json:"id"`
	// User’s Viber name.
	Name string `json:"name"`
	// URL of user’s avatar.
	Avatar string `json:"avatar"`
	// User’s 2 letter country code. ISO ALPHA-2 Code.
	Country string `json:"country"`
	// User’s phone language. Will be returned according to the device language. ISO 639-1.
	Language string `json:"language"`
	// The maximal Viber version that is supported by all of the user’s devices.
	APIVersion int64 `json:"api_version"`
}
