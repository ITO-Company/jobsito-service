package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ito-company/jobsito-service/src"
)

func SetupApi(app *fiber.App, c *src.Container) {
	v1 := app.Group("/api/v1")

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Aloha")
	})

	handlers := []func(fiber.Router){
		c.JobSeekerHandler.RegisterRoutes,
		c.CompanyHandler.RegisterRoutes,
		c.GlobalTagHandler.RegisterRoutes,
		c.JobPostingHandler.RegisterRoutes,
		c.ApplicationHandler.RegisterRoutes,
		c.IntershipHandler.RegisterRoutes,
	}

	for _, register := range handlers {
		register(v1)
	}
}
