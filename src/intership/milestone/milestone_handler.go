package milestone

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ito-company/jobsito-service/helper"
	"github.com/ito-company/jobsito-service/middleware"
	"github.com/ito-company/jobsito-service/src/enum"
)

type MilestoneHandler interface {
	RegisterRoutes(router fiber.Router)
	Create(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
}

type Handler struct {
	service MilestoneService
}

func NewHandler(service MilestoneService) MilestoneHandler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	milestoneGroup := router.Group("/milestones")

	milestoneGroup.Use(middleware.JwtMiddleware())

	milestoneGroup.Post("/", middleware.RequireRoleMiddleware((string(enum.RoleCompany))), h.Create)
	milestoneGroup.Get("/intern/:intership_id", h.FindAll)
	milestoneGroup.Get("/:id", h.FindById)
	milestoneGroup.Patch("/:id", middleware.RequireRoleMiddleware((string(enum.RoleCompany))), h.Update)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var dto MilestoneCreateDto
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	id := c.Locals("user_id").(string)

	milestone, err := h.service.Create(dto, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(milestone)
}

func (h *Handler) FindById(c *fiber.Ctx) error {
	id := c.Params("id")

	milestone, err := h.service.FindById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(milestone)
}

func (h *Handler) FindAll(c *fiber.Ctx) error {
	opts := helper.NewFindAllOptionsFromQuery(c)
	intershipId := c.Params("intership_id", "")

	milestones, err := h.service.FindAll(intershipId, opts)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(milestones)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	var dto MilestoneUpdateDto
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	id := c.Params("id")
	userId := c.Locals("user_id").(string)
	milestone, err := h.service.Update(id, dto, userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(milestone)
}
