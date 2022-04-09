package main

import (
	"DriverLocation/config"
	"DriverLocation/controller"
	"DriverLocation/db"
	_ "DriverLocation/docs"
	"DriverLocation/repository"
	"DriverLocation/service"
	"fmt"
	fiberSwagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
	"log"
)

// @title BiTaksi
// @version 2.0
// @termsOfService http://swagger.io/terms/
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	database := db.Connection()

	app := fiber.New(config.NewFiberConfig())
	app.Use(logger.New())
	app.Use(recover.New())
	app.Get("/swagger/*", fiberSwagger.HandlerDefault)

	driverRepository := repository.NewDriverRepository(database)
	driverService := service.NewDriverService(&driverRepository)
	driverController := controller.NewDriverController(&driverService)
	driverController.Route(app)

	passengerRepository := repository.NewPassengerRepository(database)
	passengerService := service.NewPassengerService(&passengerRepository)
	passengerController := controller.NewPassengerController(&passengerService)
	passengerController.Route(app)

	matcherRepository := repository.NewMatcherRepository(database)
	matcherService := service.NewMatcherService(&matcherRepository)
	matcherController := controller.NewMatcherController(&matcherService)
	matcherController.Route(app)
	//router.Route(app)

	fmt.Println(viper.GetString("appName") + " app started port: " + viper.GetString("port"))
	log.Fatal(app.Listen(":" + viper.GetString("port")))
}
