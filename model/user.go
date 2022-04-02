package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID
	Name      string
	Email     string
	CreatedAt string
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email" bson:"email"`
}

type CreateUserResponse struct {
	Id        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	Email     string             `json:"email" bson:"email"`
	CreatedAt string             `json:"created_at"`
}
