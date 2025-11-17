package savedjob

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ito-company/jobsito-service/helper"
	"github.com/ito-company/jobsito-service/middleware"
	"github.com/ito-company/jobsito-service/src/enum"
)

type SavedJobHandler interface {
	RegisterRoutes(router fiber.Router)
	Create(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	FindAllByJobSeeker(c *fiber.Ctx) error
}

type Handler struct {
	service SavedJobService
}

func NewHandler(service SavedJobService) SavedJobHandler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	savedJobGroup := router.Group("/saved-jobs")

	savedJobGroup.Use(middleware.JwtMiddleware())
	savedJobGroup.Use(middleware.RequireRoleMiddleware(string(enum.RoleSeeker)))

	savedJobGroup.Post("/", middleware.RequireRoleMiddleware(string(enum.RoleSeeker)), h.Create)
	savedJobGroup.Delete("/:id", middleware.RequireRoleMiddleware(string(enum.RoleSeeker)), h.Delete)
	savedJobGroup.Get("/", middleware.RequireRoleMiddleware(string(enum.RoleSeeker)), h.FindAllByJobSeeker)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var dto SavedJobCreateDto
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	jobSeekerID := c.Locals("user_id").(string)

	savedJob, err := h.service.Create(&dto, jobSeekerID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if savedJob == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "job already saved",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(savedJob)
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	jobSeekerID := c.Locals("user_id").(string)

	err := h.service.Delete(id, jobSeekerID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}

func (h *Handler) FindAllByJobSeeker(c *fiber.Ctx) error {
	opts := helper.NewFindAllOptionsFromQuery(c)
	jobSeekerID := c.Locals("user_id").(string)

	savedJobs, err := h.service.FindAllByJobSeeker(opts, jobSeekerID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(savedJobs)
}
