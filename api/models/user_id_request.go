package models

// UserIDRequest represents user ID request model.
type UserIDRequest struct {
	// Unique Viber user id. Required. Subscribed valid user ID.
	ID string `json:"id"`
}

// UserIDsRequest represents user IDs request model.
type UserIDsRequest struct {
	// Unique Viber user IDs. Required. 100 IDs per request.
	IDs []string `json:"ids"`
}
