package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Driver struct {
	Id          primitive.ObjectID
	Latitude    float64
	Longtitude  float64
	HasCustomer bool
	IsActive    bool
	CreatedDate string
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
