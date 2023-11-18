package responses

import "toggler/models"

type AccountsResponse struct {
	Data []models.Account `json:"data"`
}
type AccountResponse struct {
	Data models.Account `json:"data"`
}
