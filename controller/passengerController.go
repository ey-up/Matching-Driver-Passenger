package controller

import (
	"DriverLocation/exception"
	"DriverLocation/model"
	"DriverLocation/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type PassengerController struct {
	PassengerService service.PassengerService
}

func NewPassengerController(passengerService *service.PassengerService) PassengerController {
	return PassengerController{PassengerService: *passengerService}
}

func (passengerController PassengerController) Route(app *fiber.App) {
	passenger := app.Group("/api/passenger")
	passenger.Post("/", passengerController.Create)
}

// Create ShowAccount godoc
// @Summary Create a passenger
// @Tags Passenger
// @Description Create a new passenger
// @Accept  json
// @Param        passenger  body      model.CreatePassengerRequest  true  "Add Passenger"
// @Produce  json
// @Success 200 {object} model.WebResponse
// @Failure 400 {object} model.WebResponse
// @Router /api/passenger [post]
func (passengerController PassengerController) Create(ctx *fiber.Ctx) error {
	var request model.CreatePassengerRequest
	err := ctx.BodyParser(&request)
	exception.IsPanic(err)
	validate := validator.New()
	err = validate.Struct(request)
	requestControl, err := exception.PassengerCreateRequestException(err)
	if err != nil {
		return ctx.JSON(requestControl)
	}
	response := passengerController.PassengerService.Create(request)
	return ctx.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
