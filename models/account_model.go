package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	Name      string             `json:"name,omitempty" validate:"required"`
	Users     []string           `json:"users"`
	Features  []string           `json:"features"`
	CreatedAt string             `json:"createdAt,omitempty"`
	CreatedBy string             `json:"createdBy,omitempty"`
	UpdatedAt string             `json:"updatedAt,omitempty"`
	UpdatedBy string             `json:"updatedBy,omitempty"`
}
