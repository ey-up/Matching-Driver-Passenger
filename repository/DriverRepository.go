package repository

import (
	"DriverLocation/db"
	"DriverLocation/exception"
	"DriverLocation/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DriverRepository interface {
	Create(user model.Driver) primitive.ObjectID
	InsertFromCsv(drivers []interface{})
}

type driverRepository struct {
	Collection *mongo.Collection
}

func (d driverRepository) InsertFromCsv(drivers []interface{}) {
	ctx, cancel := db.NewMongoContext()
	defer cancel()
	_, err := d.Collection.InsertMany(ctx, drivers)
	exception.IsPanic(err)
}

func (d driverRepository) Create(driver model.Driver) primitive.ObjectID {
	ctx, cancel := db.NewMongoContext()
	defer cancel()
	result, err := d.Collection.InsertOne(ctx, driver)
	exception.IsPanic(err)
	return result.InsertedID.(primitive.ObjectID)
}

func NewDriverRepository(database *mongo.Database) DriverRepository {
	return &driverRepository{Collection: database.Collection("driver")}
}
