package models

// UserIDRequest represents user ID request model.
type UserIDRequest struct {
	// Unique Viber user id. Required. Subscribed valid user ID.
	ID string `json:"id"`
}
