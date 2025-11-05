package intership

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ito-company/jobsito-service/helper"
	"github.com/ito-company/jobsito-service/middleware"
	"github.com/ito-company/jobsito-service/src/enum"
)

type IntershipHandler interface {
	RegisterRoutes(router fiber.Router)
	Create(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	FindAllForJobSeeker(c *fiber.Ctx) error
	FindAllForCompany(c *fiber.Ctx) error
}

type Handler struct {
	service IntershipService
}

func NewHandler(service IntershipService) IntershipHandler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	intershipGroup := router.Group("/interships")

	intershipGroup.Use(middleware.JwtMiddleware())

	intershipGroup.Post("/", middleware.RequireRoleMiddleware(string(enum.RoleCompany)), h.Create)
	intershipGroup.Get("/job-seeker", middleware.RequireRoleMiddleware(string(enum.RoleSeeker)), h.FindAllForJobSeeker)
	intershipGroup.Get("/company", middleware.RequireRoleMiddleware(string(enum.RoleCompany)), h.FindAllForCompany)
	intershipGroup.Get("/:id", h.FindById)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var dto IntershipCreateDto
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	id := c.Locals("user_id").(string)

	intership, err := h.service.Create(id, dto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(intership)
}

func (h *Handler) FindById(c *fiber.Ctx) error {
	id := c.Params("id")

	intership, err := h.service.FindById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(intership)
}

func (h *Handler) FindAllForJobSeeker(c *fiber.Ctx) error {
	jobSeekerID := c.Locals("user_id").(string)
	opts := helper.NewFindAllOptionsFromQuery(c)

	response, err := h.service.FindAll("", jobSeekerID, opts)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *Handler) FindAllForCompany(c *fiber.Ctx) error {
	companyID := c.Locals("user_id").(string)
	opts := helper.NewFindAllOptionsFromQuery(c)

	response, err := h.service.FindAll(companyID, "", opts)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
