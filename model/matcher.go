package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Matcher struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	PassengerId primitive.ObjectID `bson:"PassengerId"`
	DriverId    primitive.ObjectID `bson:"DriverId"`
	IsActive    bool               `bson:"IsActive"`
	CreatedDate string             `bson:"CreatedDate"`
}

type MatchRequest struct {
	PassengerId primitive.ObjectID `json:"passengerId"`
	DriverId    primitive.ObjectID `json:"driverId"`
}
