package jobseeker

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ito-company/jobsito-service/helper"
	"github.com/ito-company/jobsito-service/middleware"
	"github.com/ito-company/jobsito-service/src/dto"
	"github.com/ito-company/jobsito-service/src/enum"
)

type JobSeekerHandler interface {
	RegisterRoutes(router fiber.Router)
	Signup(c *fiber.Ctx) error
	Signin(c *fiber.Ctx) error
	InternSignin(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	FindByEmail(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	SoftDelete(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	AddTagToJobSeeker(c *fiber.Ctx) error
	RemoveTagFromJobSeeker(c *fiber.Ctx) error
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
	jobSeekerGroup.Post("/intern-signin", h.InternSignin)

	jobSeekerGroup.Use(middleware.JwtMiddleware())
	jobSeekerGroup.Post("/:tag_id", middleware.RequireRoleMiddleware(string(enum.RoleSeeker)), h.AddTagToJobSeeker)
	jobSeekerGroup.Get("/", middleware.RequireRoleMiddleware(string(enum.RoleCompany)), h.FindAll)
	jobSeekerGroup.Get("/me", middleware.RequireRoleMiddleware(string(enum.RoleSeeker)), h.FindByEmail)
	jobSeekerGroup.Get("/:id", middleware.RequireRoleMiddleware(string(enum.RoleCompany)), h.FindById)
	jobSeekerGroup.Patch("/me", middleware.RequireRoleMiddleware(string(enum.RoleSeeker)), h.Update)
	jobSeekerGroup.Delete("/me", middleware.RequireRoleMiddleware(string(enum.RoleSeeker)), h.SoftDelete)
	jobSeekerGroup.Delete("/:tag_id", middleware.RequireRoleMiddleware(string(enum.RoleSeeker)), h.RemoveTagFromJobSeeker)
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

func (h *Handler) InternSignin(c *fiber.Ctx) error {
	var dto dto.InternSigninDto
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	token, err := h.service.InternSignin(dto)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

func (h *Handler) Update(c *fiber.Ctx) error {
	var input JobSeekerUpdateDto
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	email := c.Locals("email").(string)

	updatedJobSeeker, err := h.service.Update(email, input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(updatedJobSeeker)
}

func (h *Handler) FindByEmail(c *fiber.Ctx) error {
	email := c.Locals("email").(string)
	jobSeeker, err := h.service.FindByEmail(email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(jobSeeker)
}

func (h *Handler) FindById(c *fiber.Ctx) error {
	id := c.Params("id")
	jobSeeker, err := h.service.FindById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(jobSeeker)
}

func (h *Handler) SoftDelete(c *fiber.Ctx) error {
	id := c.Locals("user_id").(string)
	if err := h.service.SoftDelete(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(nil)
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

func (h *Handler) AddTagToJobSeeker(c *fiber.Ctx) error {
	jobseekerId := c.Locals("user_id").(string)
	tagId := c.Params("tag_id")
	proficiency := c.Query("proficiency", "")

	if err := h.service.AddTagToJobSeeker(jobseekerId, tagId, proficiency); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}

func (h *Handler) RemoveTagFromJobSeeker(c *fiber.Ctx) error {
	jobseekerId := c.Locals("user_id").(string)
	tagId := c.Params("tag_id")

	if err := h.service.RemoveTagFromJobSeeker(jobseekerId, tagId); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}
