package globaltags

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ito-company/jobsito-service/helper"
	"github.com/ito-company/jobsito-service/middleware"
)

type GlobalTagHandler interface {
	RegisterRoutes(router fiber.Router)
	FindAll(c *fiber.Ctx) error
}

type Handler struct {
	service GlobalTagsService
}

func NewHandler(service GlobalTagsService) GlobalTagHandler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	globalTagsGroup := router.Group("/global-tags").Use(middleware.JwtMiddleware())
	globalTagsGroup.Get("/", h.FindAll)
}

func (h *Handler) FindAll(c *fiber.Ctx) error {
	opts := helper.NewFindAllOptionsFromQuery(c)
	result, err := h.service.FindAll(opts)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
