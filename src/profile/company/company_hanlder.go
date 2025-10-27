package company

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ito-company/jobsito-service/helper"
	"github.com/ito-company/jobsito-service/middleware"
	"github.com/ito-company/jobsito-service/src/dto"
	"github.com/ito-company/jobsito-service/src/enum"
)

type CompanyHandler interface {
	RegisterRoutes(router fiber.Router)
	Signup(c *fiber.Ctx) error
	Signin(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	SoftDelete(c *fiber.Ctx) error
	FindByEmail(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
}

type Handler struct {
	service CompanyService
}

func NewHandler(service CompanyService) CompanyHandler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	companyGroup := router.Group("/company")
	companyGroup.Post("/signup", h.Signup)
	companyGroup.Post("/signin", h.Signin)

	companyGroup.Use(middleware.JwtMiddleware())
	companyGroup.Get("/", middleware.RequireRoleMiddleware(string(enum.RoleSeeker)), h.FindAll)
	companyGroup.Get("/me", middleware.RequireRoleMiddleware(string(enum.RoleCompany)), h.FindByEmail)
	companyGroup.Get("/:id", middleware.RequireRoleMiddleware(string(enum.RoleSeeker)), h.FindById)
	companyGroup.Patch("/me", middleware.RequireRoleMiddleware(string(enum.RoleCompany)), h.Update)
	companyGroup.Delete("/me", middleware.RequireRoleMiddleware(string(enum.RoleCompany)), h.SoftDelete)
}

func (h *Handler) Signup(c *fiber.Ctx) error {
	var dto dto.SignupDto
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	token, err := h.service.Signup(dto)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"token": token,
	})
}

func (h *Handler) Signin(c *fiber.Ctx) error {
	var dto dto.SigninDto
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	token, err := h.service.Signin(dto)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

func (h *Handler) Update(c *fiber.Ctx) error {
	var input CompanyUpdateDto

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	email := c.Locals("email").(string)

	updatedCompany, err := h.service.Update(email, input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(updatedCompany)
}

func (h *Handler) SoftDelete(c *fiber.Ctx) error {
	id := c.Locals("user_id").(string)

	if err := h.service.SoftDelete(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}

func (h *Handler) FindByEmail(c *fiber.Ctx) error {
	email := c.Locals("email").(string)
	company, err := h.service.FindByEmail(email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(company)
}

func (h *Handler) FindById(c *fiber.Ctx) error {
	id := c.Params("id")
	company, err := h.service.FindById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(company)
}

func (h *Handler) FindAll(c *fiber.Ctx) error {
	opts := helper.NewFindAllOptionsFromQuery(c)
	project, err := h.service.FindAll(opts)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(project)
}
