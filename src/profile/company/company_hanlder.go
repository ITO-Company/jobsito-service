package company

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ito-company/jobsito-service/src/dto"
)

type CompanyHandler interface {
	RegisterRoutes(router fiber.Router)
	Signup(c *fiber.Ctx) error
	Signin(c *fiber.Ctx) error
}

type Handler struct {
	service CompanyService
}

func NewHandler(service CompanyService) CompanyHandler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	router.Post("/company/signup", h.Signup)
	router.Post("/company/signin", h.Signin)
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
