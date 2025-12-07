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
	GetOverview(c *fiber.Ctx) error
	GetOverviewList(c *fiber.Ctx) error
	GetOverviewPDF(c *fiber.Ctx) error
	GetOverviewListPDF(c *fiber.Ctx) error
}

type Handler struct {
	service       IntershipService
	reportService ReportService
}

func NewHandler(service IntershipService) IntershipHandler {
	return &Handler{
		service:       service,
		reportService: NewReportService(),
	}
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	intershipGroup := router.Group("/interships")

	intershipGroup.Get("/overview/list/pdf/cor", h.GetOverviewListPDF)
	intershipGroup.Get("/overview/list/cor", h.GetOverviewList)

	intershipGroup.Use(middleware.JwtMiddleware())

	intershipGroup.Post("/", middleware.RequireRoleMiddleware(string(enum.RoleCompany)), h.Create)
	intershipGroup.Get("/job-seeker", middleware.RequireRoleMiddleware(string(enum.RoleSeeker)), h.FindAllForJobSeeker)
	intershipGroup.Get("/company", middleware.RequireRoleMiddleware(string(enum.RoleCompany)), h.FindAllForCompany)
	intershipGroup.Get("/overview/list/pdf", h.GetOverviewListPDF)
	intershipGroup.Get("/overview/list", h.GetOverviewList)
	intershipGroup.Get("/:id/overview/pdf", h.GetOverviewPDF)
	intershipGroup.Get("/:id", h.FindById)
	intershipGroup.Get("/:id/overview", h.GetOverview)
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

func (h *Handler) GetOverview(c *fiber.Ctx) error {
	id := c.Params("id")

	overview, err := h.service.FindByIdWithOverview(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(overview)
}

func (h *Handler) GetOverviewList(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	opts := helper.NewFindAllOptionsFromQuery(c)

	var userIDStr string
	if userID != nil {
		userIDStr = userID.(string)
	}

	// Intentar obtener como company primero
	response, err := h.service.FindAllWithOverview(userIDStr, "", opts)
	if err == nil && response.Total > 0 {
		return c.Status(fiber.StatusOK).JSON(response)
	}

	// Si no hay user_id (coordinador), retornar todos sin filtrar
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *Handler) GetOverviewPDF(c *fiber.Ctx) error {
	id := c.Params("id")

	overview, err := h.service.FindByIdWithOverview(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	pdfBytes, err := h.reportService.GenerateOverviewPDF(overview)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate PDF",
		})
	}

	c.Set("Content-Type", "application/pdf")
	c.Set("Content-Disposition", "attachment; filename=intership-"+id+".pdf")
	return c.Send(pdfBytes)
}

func (h *Handler) GetOverviewListPDF(c *fiber.Ctx) error {
	userID := c.Locals("user_id")

	var userIDStr string
	if userID != nil {
		userIDStr = userID.(string)
	}

	// Obtener todas las pasantías sin límite
	opts := &helper.FindAllOptions{
		Limit:  100000,
		Offset: 0,
	}

	// Obtener datos con overview
	response, err := h.service.FindAllWithOverview(userIDStr, "", opts)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	pdfBytes, err := h.reportService.GenerateOverviewListPDF(response.Data, response.Total)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate PDF",
		})
	}

	c.Set("Content-Type", "application/pdf")
	c.Set("Content-Disposition", "attachment; filename=interships-list.pdf")
	return c.Send(pdfBytes)
}
