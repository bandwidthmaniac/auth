package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username" binding:"required,min=4" bson:"username"`
	Password string             `json:"password" binding:"required,min=6" bson:"password,omitempty"`
}
