package models

// Member represents Viber chat member model.
type Member struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Role   string `json:"role"`
}
