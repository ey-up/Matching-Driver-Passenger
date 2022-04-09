package repository

import (
	"DriverLocation/db"
	"DriverLocation/exception"
	"DriverLocation/model"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PassengerRepository interface {
	Create(passenger model.Passenger) primitive.ObjectID
	UpdateHasDriverByPassengerId(id primitive.ObjectID)
}

type passengerRepository struct {
	Collection *mongo.Collection
}

func (p passengerRepository) UpdateHasDriverByPassengerId(id primitive.ObjectID) {
	ctx, cancel := db.NewMongoContext()
	defer cancel()
	update := bson.D{
		{"$set", bson.D{{"HasDriver", true}}},
	}
	result, err := p.Collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		update)
	exception.IsPanic(err)
	fmt.Println(result)
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
