package jobseeker

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ito-company/jobsito-service/src/dto"
)

type JobSeekerHandler interface {
	RegisterRoutes(router fiber.Router)
	Signup(c *fiber.Ctx) error
	Signin(c *fiber.Ctx) error
}

type Handler struct {
	service JobSeekerService
}

func NewHandler(service JobSeekerService) JobSeekerHandler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	jobSeekerGroup := router.Group("/job-seekers")
	jobSeekerGroup.Post("/signup", h.Signup)
	jobSeekerGroup.Post("/signin", h.Signin)
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
