package main

import (
	"DriverLocation/config"
	"DriverLocation/controller"
	"DriverLocation/db"
	"DriverLocation/repository"
	"DriverLocation/service"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	database := db.Connection()

	app := fiber.New(config.NewFiberConfig())
	app.Use(logger.New())
	app.Use(recover.New())

	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserService(&userRepository)
	userController := controller.NewUserController(&userService)
	userController.Route(app)

	driverRepository := repository.NewDriverRepository(database)
	driverService := service.NewDriverService(&driverRepository)
	driverController := controller.NewDriverController(&driverService)
	driverController.Route(app)

	//router.Route(app)

	fmt.Println(viper.GetString("appName") + " app started port: " + viper.GetString("port"))
	log.Fatal(app.Listen(":" + viper.GetString("port")))
}
