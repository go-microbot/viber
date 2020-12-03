package models

import "github.com/go-microbot/viber/models"

// UserDetailsResponse represents user details response model.
type UserDetailsResponse struct {
	MessageResponse
	User DetailsResponse `json:"user"`
}

// DetailsResponse represents details response model.
type DetailsResponse struct {
	models.User
	// The operating system type and version of the user’s primary device.
	PrimaryDeviceOS string `json:"primary_device_os,omitempty"`
	// The Viber version installed on the user’s primary device.
	ViberVersion string `json:"viber_version,omitempty"`
	// Mobile country code.
	MCC int64 `json:"mcc,omitempty"`
	// Mobile network code.
	MNC int64 `json:"mnc,omitempty"`
	// The user’s device type.
	DeviceType string `json:"device_type,omitempty"`
}
