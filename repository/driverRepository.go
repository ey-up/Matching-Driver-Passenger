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

type DriverRepository interface {
	Create(driver model.Driver) primitive.ObjectID
	InsertFromCsv(drivers []interface{})
	FindAvailableDrivers(driver model.Driver) []model.Driver
	UpdateHasPassengerByDriverId(id primitive.ObjectID)
}

type driverRepository struct {
	Collection *mongo.Collection
}

func (d driverRepository) UpdateHasPassengerByDriverId(id primitive.ObjectID) {
	ctx, cancel := db.NewMongoContext()
	defer cancel()
	update := bson.D{
		{"$set", bson.D{{"HasPassenger", true}}},
	}
	result, err := d.Collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		update)
	exception.IsPanic(err)
	fmt.Println(result)
}

func (d driverRepository) FindAvailableDrivers(driver model.Driver) []model.Driver {
	ctx, cancel := db.NewMongoContext()
	defer cancel()
	filter := bson.D{
		{"HasPassenger", driver.HasPassenger},
		{"IsActive", driver.IsActive},
		{"Latitude", bson.M{"$lte": driver.Latitude + 0.02}},
		{"Latitude", bson.M{"$gt": driver.Latitude - 0.02}},
		{"Longtitude", bson.M{"$lte": driver.Longtitude + 0.02}},
		{"Longtitude", bson.M{"$gt": driver.Longtitude - 0.02}},
	}
	find, err := d.Collection.Find(ctx, filter)
	if err != nil {
		return nil
	}

	var results []bson.D
	find.All(ctx, &results)
	return model.ToModel(results)
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
