package repository

import (
	"DriverLocation/db"
	"DriverLocation/exception"
	"DriverLocation/model"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Insert(user model.User)
}

type userService struct {
	Collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userService{
		Collection: db.Collection("users"),
	}
}
func (u userService) Insert(user model.User) {
	ctx, cancel := db.NewMongoContext()
	defer cancel()

	result, err := u.Collection.InsertOne(ctx, bson.M{
		"name":      user.Name,
		"email":     user.Email,
		"createdAt": user.CreatedAt,
	})
	fmt.Println(result.InsertedID)
	exception.IsPanic(err)
}
