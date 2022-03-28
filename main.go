package main

import (
	"DriverLocation/config"
	"DriverLocation/db"
	"DriverLocation/exception"
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
	exception.IsPanic(database)

	app := fiber.New(config.NewFiberConfig())
	app.Use(logger.New())
	app.Use(recover.New())

	fmt.Println("App started port: " + viper.GetString("port"))
	log.Fatal(app.Listen(":" + viper.GetString("port")))
}
