package issue

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ito-company/jobsito-service/helper"
	"github.com/ito-company/jobsito-service/middleware"
	"github.com/ito-company/jobsito-service/src/enum"
)

type IssueHandler interface {
	RegisterRoutes(router fiber.Router)
	Create(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
}

type Handler struct {
	service IssueService
}

func NewHandler(service IssueService) IssueHandler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	issueGroup := router.Group("/issues")

	issueGroup.Use(middleware.JwtMiddleware())

	issueGroup.Post("/", middleware.RequireRoleMiddleware(string(enum.RoleCompany)), h.Create)
	issueGroup.Get("/milestone/:milestone_id", h.FindAll)
	issueGroup.Get("/:id", h.FindById)
	issueGroup.Patch("/:id", middleware.RequireRoleMiddleware(string(enum.RoleCompany)), h.Update)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var dto IssueCreateDto
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	issue, err := h.service.Create(dto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(issue)
}

func (h *Handler) FindById(c *fiber.Ctx) error {
	id := c.Params("id")

	issue, err := h.service.FindById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(issue)
}

func (h *Handler) FindAll(c *fiber.Ctx) error {
	milestoneId := c.Params("milestone_id")
	opts := helper.NewFindAllOptionsFromQuery(c)

	issues, err := h.service.FindAll(milestoneId, opts)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(issues)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var dto IssueUpdateDto
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	issue, err := h.service.Update(id, dto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(issue)
}
