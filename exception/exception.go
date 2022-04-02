package exception

import (
	"DriverLocation/model"
	"github.com/go-playground/validator/v10"
)

func DriverCreateRequestException(err error) (model.WebResponse, error) {
	validationErrors := err.(validator.ValidationErrors)
	if validationErrors != nil {
		return model.WebResponse{
			Code:   400,
			Status: "Bad request",
			Data:   err.Error(),
		}, err
	}
	return model.WebResponse{}, nil
}
