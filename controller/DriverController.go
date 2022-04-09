package controller

import (
	"DriverLocation/exception"
	"DriverLocation/model"
	"DriverLocation/service"
	"github.com/go-playground/validator/v10"
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
	app.Put("/api/driver", driverController.InsertFromCsv)
}

// Create ShowAccount godoc
// @Summary Create a driver
// @Description Create a new driver
// @Tags Driver
// @Accept  json
// @Param        driver  body      model.CreateDriverRequest  true  "Add Driver"
// @Produce  json
// @Success 200 {object} model.WebResponse
// @Failure 400 {object} model.WebResponse
// @Router /api/driver [post]
func (driverController *DriverController) Create(ctx *fiber.Ctx) error {
	var request model.CreateDriverRequest
	err := ctx.BodyParser(&request)
	exception.IsPanic(err)
	validate := validator.New()
	err = validate.Struct(request)
	requestControl, err := exception.DriverCreateRequestException(err)
	if err != nil {
		return ctx.JSON(requestControl)
	}
	response := driverController.DriverService.Create(request)
	return ctx.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

// InsertFromCsv  ShowAccount godoc
// @Summary Insert from csv
// @Description Insert from csv
// @Tags Driver
// @Accept  json
// @Produce  json
// @Success 200 {object} model.WebResponse
// @Failure 400 {object} model.WebResponse
// @Router /api/driver [put]
func (driverController *DriverController) InsertFromCsv(ctx *fiber.Ctx) error {
	driverController.DriverService.InsertFromCsv()
	return ctx.Status(200).JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Inserted to db from csv file",
	})
}
