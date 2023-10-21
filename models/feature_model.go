package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Feature struct {
	Id        primitive.ObjectID     `json:"id,omitempty"`
	AccountId string                 `json:"accountId" validate:"required"`
	Name      string                 `json:"name,omitempty" validate:"required"`
	Enabled   bool                   `json:"enabled,omitempty" validate:"required"`
	Flags     map[string]interface{} `json:"flags,omitempty" validate:"required"`
	CreatedAt string                 `json:"createdAt,omitempty"`
	CreatedBy string                 `json:"createdBy,omitempty" validate:"required"`
	UpdatedAt string                 `json:"updatedAt,omitempty"`
	UpdatedBy string                 `json:"updatedBy,omitempty"`
}
