package models

import "github.com/go-microbot/viber/models"

// AccountInfoResponse represents account info response model.
type AccountInfoResponse struct {
	MessageResponse
	models.AccountInfo
}
