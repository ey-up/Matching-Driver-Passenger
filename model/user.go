package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id        primitive.ObjectID
	Name      string
	Email     string
	CreatedAt time.Time
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email" bson:"email"`
}

type CreateUserResponse struct {
	Id        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	Email     string             `json:"email" bson:"email"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}
