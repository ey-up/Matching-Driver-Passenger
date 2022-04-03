package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Driver struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Latitude    float64            `bson:"Latitude"`
	Longtitude  float64            `bson:"Longtitude"`
	HasCustomer bool               `bson:"HasCustomer"`
	IsActive    bool               `bson:"IsActive"`
	CreatedDate string             `bson:"CreatedDate"`
}

type CreateDriverRequest struct {
	Latitude   float64 `json:"latitude" validate:"required,max=180,min=-180"`
	Longtitude float64 `json:"longtitude" validate:"required,max=90,min=-90"`
}

type CreateDriverResponse struct {
	Id          primitive.ObjectID `json:"id"`
	Latitude    float64            `json:"latitude"`
	Longtitude  float64            `json:"longtitude"`
	HasCustomer bool               `json:"hasCustomer"`
	IsActive    bool               `json:"isActive"`
	CreatedDate string             `json:"createDate"`
}
