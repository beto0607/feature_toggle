package responses

import (
	"toggler/models"
)

type FeaturesResponse struct {
	Data []models.Feature `json:"data"`
}

type FeatureResponse struct {
	Data models.Feature `json:"data"`
}
