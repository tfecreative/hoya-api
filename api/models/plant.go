package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Plant struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name"`
}
