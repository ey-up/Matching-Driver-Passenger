package controller

import (
	"DriverLocation/exception"
	"DriverLocation/model"
	"DriverLocation/service"
	"github.com/gofiber/fiber/v2"
)

type DriverController struct {
	DriverService service.DriverService
}

func NewDriverController(driverService *service.DriverService) DriverController {
	return DriverController{DriverService: *driverService}
}

func (driverController DriverController) Route(app *fiber.App) {
	app.Post("/api/driver", driverController.Create)
}

func (driverController *DriverController) Create(ctx *fiber.Ctx) error {
	var request model.CreateDriverRequest
	err := ctx.BodyParser(&request)
	exception.IsPanic(err)
	response := driverController.DriverService.Create(request)
	return ctx.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
