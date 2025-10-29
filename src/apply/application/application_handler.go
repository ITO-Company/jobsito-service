package application

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ito-company/jobsito-service/helper"
	"github.com/ito-company/jobsito-service/middleware"
	"github.com/ito-company/jobsito-service/src/enum"
)

type ApplicationHandler interface {
	RegisterRoutes(router fiber.Router)
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	FindAllByJobSeeker(c *fiber.Ctx) error
	FindAllByJobPostingAndCompany(c *fiber.Ctx) error
}

type Handler struct {
	service ApplicationService
}

func NewHandler(service ApplicationService) ApplicationHandler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	applicationGroup := router.Group("/applications")

	applicationGroup.Use(middleware.JwtMiddleware())

	applicationGroup.Post("/", middleware.RequireRoleMiddleware(string(enum.RoleSeeker)), h.Create)
	applicationGroup.Patch("/:id", middleware.RequireRoleMiddleware(string(enum.RoleCompany)), h.Update)
	applicationGroup.Get("/:id", h.FindById)
	applicationGroup.Get("/job-seeker", h.FindAllByJobSeeker)
	applicationGroup.Get("/job-posting/:job_posting_id/company", h.FindAllByJobPostingAndCompany)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var dto ApplicationCreateDto
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	id := c.Locals("user_id").(string)

	application, err := h.service.Create(&dto, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(application)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	var dto ApplicationUpdateDto
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	applicationID := c.Params("id")

	application, err := h.service.Update(&dto, applicationID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(application)
}

func (h *Handler) FindById(c *fiber.Ctx) error {
	id := c.Params("id")
	job, err := h.service.FindByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(job)
}

func (h *Handler) FindAllByJobSeeker(c *fiber.Ctx) error {
	opts := helper.NewFindAllOptionsFromQuery(c)
	jobSeekerID := c.Locals("user_id").(string)

	applications, err := h.service.FindAllByJobSeeker(opts, jobSeekerID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(applications)
}

func (h *Handler) FindAllByJobPostingAndCompany(c *fiber.Ctx) error {
	opts := helper.NewFindAllOptionsFromQuery(c)
	companyID := c.Locals("user_id").(string)
	jobPostingID := c.Params("job_posting_id")

	applications, err := h.service.FindAllByJobPostingAndCompany(opts, jobPostingID, companyID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(applications)
}
