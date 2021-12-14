package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Username string             `json:"username" binding:"required,min=6" bson:"userName"`
	Password string             `json:"password" binding:"required,min=6" bson:"password"`
}
