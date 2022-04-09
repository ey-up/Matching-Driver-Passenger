package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Driver struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	Latitude     float64            `bson:"Latitude"`
	Longtitude   float64            `bson:"Longtitude"`
	HasPassenger bool               `bson:"HasPassenger"`
	IsActive     bool               `bson:"IsActive"`
	CreatedDate  string             `bson:"CreatedDate"`
}

type CreateDriverRequest struct {
	Latitude   float64 `json:"latitude" validate:"required,max=180,min=-180"`
	Longtitude float64 `json:"longtitude" validate:"required,max=90,min=-90"`
}

type TheNearestDriverResponse struct {
	Id           primitive.ObjectID `json:"id"`
	Latitude     float64            `json:"latitude"`
	Longtitude   float64            `json:"longtitude"`
	HasPassenger bool               `json:"hasPassenger"`
	IsActive     bool               `json:"isActive"`
	CreatedDate  string             `json:"createDate"`
	Distance     float64            `json:"distance"`
}

type CreateDriverResponse struct {
	Id           primitive.ObjectID `json:"id"`
	Latitude     float64            `json:"latitude"`
	Longtitude   float64            `json:"longtitude"`
	HasPassenger bool               `json:"hasPassenger"`
	IsActive     bool               `json:"isActive"`
	CreatedDate  string             `json:"createDate"`
}

func ToModel(results []bson.D) (drivers []Driver) {
	for _, result := range results {
		drivers = append(drivers, Driver{
			Id:           result.Map()["_id"].(primitive.ObjectID),
			Latitude:     result.Map()["Latitude"].(float64),
			Longtitude:   result.Map()["Longtitude"].(float64),
			HasPassenger: result.Map()["HasPassenger"].(bool),
			IsActive:     result.Map()["IsActive"].(bool),
			CreatedDate:  result.Map()["CreatedDate"].(string),
		})
	}
	return drivers
}
