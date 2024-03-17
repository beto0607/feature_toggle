package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Feature struct {
	Id        primitive.ObjectID     `json:"id,omitempty" bson:"_id"`
	AccountId string                 `json:"accountId" validate:"required" bson:"account_id"`
	Name      string                 `json:"name" validate:"required" bson:"name"`
	Enabled   bool                   `json:"enabled" validate:"required" bson:"enabled"`
	Flags     map[string]interface{} `json:"flags" validate:"required" bson:"flags"`
	CreatedAt string                 `json:"createdAt,omitempty" bson:"created_at"`
	UpdatedAt string                 `json:"updatedAt,omitempty" bson:"updated_at"`
}

type FeatureDto struct {
	Name    *string                 `bson:"name,omitempty" json:"name,omitempty"`
	Enabled *bool                   `bson:"enabled,omitempty" json:"enabled,omitempty"`
	Flags   *map[string]interface{} `bson:"flags,omitempty" json:"flags,omitempty"`
}
