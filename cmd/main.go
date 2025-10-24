package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/ito-company/jobsito-service/cmd/api"
	"github.com/ito-company/jobsito-service/config"
	"github.com/ito-company/jobsito-service/middleware"
	"github.com/ito-company/jobsito-service/src"
)

func main() {
	config.Load()

	app := fiber.New()
	app.Use(middleware.Logger())

	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("ALLOW_ORIGINS"),
		AllowMethods: "GET,POST,PATCH,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	c := src.SetupContainer()
	api.SetupApi(app, c)

	app.Listen("0.0.0.0:" + config.Port)
}
