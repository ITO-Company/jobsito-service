package kpi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ito-company/jobsito-service/middleware"
	"github.com/ito-company/jobsito-service/src/enum"
)

type KPIHandler interface {
	RegisterRoutes(router fiber.Router)
	GetCompanyMilestoneKPI(c *fiber.Ctx) error
	GetIntershipMilestoneKPI(c *fiber.Ctx) error
	GetCompanyIssueKPI(c *fiber.Ctx) error
	GetIntershipIssueKPI(c *fiber.Ctx) error
	GetCompanyRequestKPI(c *fiber.Ctx) error
	GetIntershipRequestKPI(c *fiber.Ctx) error
	GetCompanyConversionKPI(c *fiber.Ctx) error
	GetJobPostingConversionKPI(c *fiber.Ctx) error
}

type Handler struct {
	milestoneService  MilestoneKPIService
	issueService      IssueKPIService
	requestService    RequestKPIService
	conversionService ConversionKPIService
}

func NewHandler(
	milestoneService MilestoneKPIService,
	issueService IssueKPIService,
	requestService RequestKPIService,
	conversionService ConversionKPIService,
) KPIHandler {
	return &Handler{
		milestoneService:  milestoneService,
		issueService:      issueService,
		requestService:    requestService,
		conversionService: conversionService,
	}
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	kpiGroup := router.Group("/kpis")

	kpiGroup.Use(middleware.JwtMiddleware())

	// Milestone KPIs
	kpiGroup.Get("/milestones/company", middleware.RequireRoleMiddleware(string(enum.RoleCompany)), h.GetCompanyMilestoneKPI)
	kpiGroup.Get("/milestones/intership/:intership_id", h.GetIntershipMilestoneKPI)

	// Issue KPIs
	kpiGroup.Get("/issues/company", middleware.RequireRoleMiddleware(string(enum.RoleCompany)), h.GetCompanyIssueKPI)
	kpiGroup.Get("/issues/intership/:intership_id", h.GetIntershipIssueKPI)

	// Request KPIs
	kpiGroup.Get("/requests/company", middleware.RequireRoleMiddleware(string(enum.RoleCompany)), h.GetCompanyRequestKPI)
	kpiGroup.Get("/requests/intership/:intership_id", h.GetIntershipRequestKPI)

	// Conversion KPIs
	kpiGroup.Get("/conversions/company", middleware.RequireRoleMiddleware(string(enum.RoleCompany)), h.GetCompanyConversionKPI)
	kpiGroup.Get("/conversions/job-posting/:job_posting_id", h.GetJobPostingConversionKPI)
}

func (h *Handler) GetCompanyMilestoneKPI(c *fiber.Ctx) error {
	companyID := c.Locals("user_id").(string)

	kpi, err := h.milestoneService.GetCompanyMilestoneKPI(companyID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(kpi)
}

func (h *Handler) GetIntershipMilestoneKPI(c *fiber.Ctx) error {
	intershipID := c.Params("intership_id")

	kpi, err := h.milestoneService.GetIntershipMilestoneKPI(intershipID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(kpi)
}

func (h *Handler) GetCompanyIssueKPI(c *fiber.Ctx) error {
	companyID := c.Locals("user_id").(string)

	kpi, err := h.issueService.GetCompanyIssueKPI(companyID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(kpi)
}

func (h *Handler) GetIntershipIssueKPI(c *fiber.Ctx) error {
	intershipID := c.Params("intership_id")

	kpi, err := h.issueService.GetIntershipIssueKPI(intershipID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(kpi)
}

func (h *Handler) GetCompanyRequestKPI(c *fiber.Ctx) error {
	companyID := c.Locals("user_id").(string)

	kpi, err := h.requestService.GetCompanyRequestKPI(companyID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(kpi)
}

func (h *Handler) GetIntershipRequestKPI(c *fiber.Ctx) error {
	intershipID := c.Params("intership_id")

	kpi, err := h.requestService.GetIntershipRequestKPI(intershipID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(kpi)
}

func (h *Handler) GetCompanyConversionKPI(c *fiber.Ctx) error {
	companyID := c.Locals("user_id").(string)

	kpi, err := h.conversionService.GetCompanyConversionKPI(companyID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(kpi)
}

func (h *Handler) GetJobPostingConversionKPI(c *fiber.Ctx) error {
	jobPostingID := c.Params("job_posting_id")

	kpi, err := h.conversionService.GetJobPostingConversionKPI(jobPostingID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(kpi)
}
