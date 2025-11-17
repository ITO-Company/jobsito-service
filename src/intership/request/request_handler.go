package request

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ito-company/jobsito-service/helper"
	"github.com/ito-company/jobsito-service/middleware"
	"github.com/ito-company/jobsito-service/src/enum"
)

type RequestHandler interface {
	RegisterRoutes(router fiber.Router)
	Create(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	FindByIssueId(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Review(c *fiber.Ctx) error
}

type Handler struct {
	service RequestService
}

func NewHandler(service RequestService) RequestHandler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	requestGroup := router.Group("/requests")

	requestGroup.Use(middleware.JwtMiddleware())

	requestGroup.Post("/", middleware.RequireRoleMiddleware(string(enum.RoleSeeker)), h.Create)
	requestGroup.Get("/issue/:issue_id", h.FindByIssueId)
	requestGroup.Get("/:id", h.FindById)
	requestGroup.Patch("/:id", middleware.RequireRoleMiddleware(string(enum.RoleSeeker)), h.Update)
	requestGroup.Patch("/:id/review", middleware.RequireRoleMiddleware(string(enum.RoleCompany)), h.Review)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var dto RequestCreateDto
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	request, err := h.service.Create(dto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(request)
}

func (h *Handler) FindById(c *fiber.Ctx) error {
	id := c.Params("id")

	request, err := h.service.FindById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(request)
}

func (h *Handler) FindByIssueId(c *fiber.Ctx) error {
	issueId := c.Params("issue_id")
	opts := helper.NewFindAllOptionsFromQuery(c)

	requests, err := h.service.FindByIssueId(issueId, opts)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(requests)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var dto RequestUpdateDto
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	request, err := h.service.Update(id, dto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(request)
}

func (h *Handler) Review(c *fiber.Ctx) error {
	id := c.Params("id")

	var dto RequestReviewDto
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	request, err := h.service.Review(id, dto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(request)
}
