package repository

import (
	"DriverLocation/db"
	"DriverLocation/exception"
	"DriverLocation/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Insert(user model.User) primitive.ObjectID
}

type userRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{
		Collection: db.Collection("users"),
	}
}
func (u userRepository) Insert(user model.User) primitive.ObjectID {
	ctx, cancel := db.NewMongoContext()
	defer cancel()

	result, err := u.Collection.InsertOne(ctx, bson.M{
		"name":      user.Name,
		"email":     user.Email,
		"createdAt": user.CreatedAt,
	})
	exception.IsPanic(err)
	return result.InsertedID.(primitive.ObjectID)
}
