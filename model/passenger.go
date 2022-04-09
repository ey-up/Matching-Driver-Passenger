package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Passenger struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Latitude    float64            `bson:"Latitude"`
	Longtitude  float64            `bson:"Longtitude"`
	HasDriver   bool               `bson:"HasDriver"`
	IsActive    bool               `bson:"IsActive"`
	CreatedDate string             `bson:"CreatedDate"`
}
type FindDriverRequest struct {
	Latitude   float64 `json:"latitude" validate:"required,max=180,min=-180"`
	Longtitude float64 `json:"longtitude" validate:"required,max=90,min=-90"`
}
type MatchDriverWithPassengerRequest struct {
	Latitude   float64 `json:"latitude" validate:"required,max=180,min=-180"`
	Longtitude float64 `json:"longtitude" validate:"required,max=90,min=-90"`
}
type CreatePassengerRequest struct {
	Latitude   float64 `json:"latitude" validate:"required,max=180,min=-180"`
	Longtitude float64 `json:"longtitude" validate:"required,max=90,min=-90"`
}

type CreatePassengerResponse struct {
	Id          primitive.ObjectID `json:"id"`
	Latitude    float64            `json:"latitude"`
	Longtitude  float64            `json:"longtitude"`
	HasDriver   bool               `json:"HasDriver"`
	IsActive    bool               `json:"isActive"`
	CreatedDate string             `json:"createDate"`
}
