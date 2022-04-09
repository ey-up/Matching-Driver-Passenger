package exception

import (
	"DriverLocation/model"
)

func DriverCreateRequestException(err error) (model.WebResponse, error) {
	if err != nil {
		return model.WebResponse{
			Code:   400,
			Status: "Bad request",
			Data:   err.Error(),
		}, err
	}
	return model.WebResponse{}, nil
}
func PassengerCreateRequestException(err error) (model.WebResponse, error) {
	if err != nil {
		return model.WebResponse{
			Code:   400,
			Status: "Bad request",
			Data:   err.Error(),
		}, err
	}
	return model.WebResponse{}, nil
}
