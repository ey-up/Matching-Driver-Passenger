package repository

import (
	"DriverLocation/db"
	"DriverLocation/exception"
	"DriverLocation/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DriverRepository interface {
	Create(user model.Driver) primitive.ObjectID
}

type driverRepository struct {
	Collection *mongo.Collection
}

func (d driverRepository) Create(user model.Driver) primitive.ObjectID {
	ctx, cancel := db.NewMongoContext()
	defer cancel()
	result, err := d.Collection.InsertOne(ctx, bson.M{
		"longtitude":  user.Longtitude,
		"latitude":    user.Latitude,
		"hasCustomer": user.HasCustomer,
		"createDate":  user.CreatedDate,
	})
	exception.IsPanic(err)
	return result.InsertedID.(primitive.ObjectID)
}

func NewDriverRepository(database *mongo.Database) DriverRepository {
	return &driverRepository{Collection: database.Collection("asd")}
}
