package repository

import (
	"DriverLocation/db"
	"DriverLocation/exception"
	"DriverLocation/model"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MatcherRepository interface {
	Match(matcher model.Matcher) primitive.ObjectID
}

type matcherRepository struct {
	Collection *mongo.Collection
}

func (m matcherRepository) Match(matcher model.Matcher) primitive.ObjectID {
	ctx, cancel := db.NewMongoContext()
	defer cancel()
	result, err := m.Collection.InsertOne(ctx, matcher)
	exception.IsPanic(err)
	fmt.Println(result)
	return result.InsertedID.(primitive.ObjectID)
}

func NewMatcherRepository(database *mongo.Database) MatcherRepository {
	return &matcherRepository{Collection: database.Collection("matcher")}
}
