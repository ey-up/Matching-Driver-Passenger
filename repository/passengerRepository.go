package repository

import (
	"DriverLocation/db"
	"DriverLocation/exception"
	"DriverLocation/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PassengerRepository interface {
	Create(passenger model.Passenger) primitive.ObjectID
}

type passengerRepository struct {
	Collection *mongo.Collection
}

func (p passengerRepository) Create(passenger model.Passenger) primitive.ObjectID {
	ctx, cancel := db.NewMongoContext()
	defer cancel()
	result, err := p.Collection.InsertOne(ctx, passenger)
	exception.IsPanic(err)
	return result.InsertedID.(primitive.ObjectID)
}

func NewPassengerRepository(database *mongo.Database) PassengerRepository {
	return &passengerRepository{Collection: database.Collection("passenger")}
}
