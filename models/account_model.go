package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	Id        primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string               `json:"name" validate:"required"`
	Users     []primitive.ObjectID `json:"users"`
	Features  []primitive.ObjectID `json:"features"`
	CreatedAt primitive.DateTime   `json:"createdAt,omitempty"`
	CreatedBy string               `json:"createdBy,omitempty"`
	UpdatedAt primitive.DateTime   `json:"updatedAt,omitempty"`
	UpdatedBy string               `json:"updatedBy,omitempty"`
}
