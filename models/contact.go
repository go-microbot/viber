package models

// Contact represents contact model.
type Contact struct {
	// Contact’s username.
	Name string `json:"name"`
	// Contact’s phone number.
	PhoneNumber string `json:"phone_number"`
	// Avatar URL.
	Avatar string `json:"avatar"`
}
