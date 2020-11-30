package api

import (
	"net/http"
)

const baseURL = "https://chatapi.viber.com/pa/"

// ViberAPI represents Viber Bot API.
type ViberAPI struct {
	client *http.Client
	token  string
	url    string
}

// NewViberAPI returns new ViberAPI instance.
func NewViberAPI(token string) ViberAPI {
	return ViberAPI{
		token:  token,
		client: http.DefaultClient,
		url:    baseURL,
	}
}
