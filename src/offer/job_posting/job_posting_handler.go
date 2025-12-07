package jobposting

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/ito-company/jobsito-service/helper"
	"github.com/ito-company/jobsito-service/middleware"
	"github.com/ito-company/jobsito-service/src/enum"
)

type JobPostingHandler interface {
	RegisterRoutes(router fiber.Router)
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	SoftDelete(c *fiber.Ctx) error
	AddTagToJobPosting(c *fiber.Ctx) error
	RemoveTagFromJobPosting(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	GenerateJobPostingListPDF(c *fiber.Ctx) error
	GenerateJobPostingDetailPDF(c *fiber.Ctx) error
}

type Handler struct {
	service       JobPostingService
	reportService JobPostingReportService
}

func NewHandler(service JobPostingService) JobPostingHandler {
	return &Handler{
		service:       service,
		reportService: NewJobPostingReportService(),
	}
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	jobPostingGroup := router.Group("/job-postings")

	jobPostingGroup.Use(middleware.JwtMiddleware())

	jobPostingGroup.Get("/", h.FindAll)
	jobPostingGroup.Get("/:id", h.FindById)
	jobPostingGroup.Get("/report/list", middleware.RequireRoleMiddleware(string(enum.RoleCompany)), h.GenerateJobPostingListPDF)
	jobPostingGroup.Get("/:id/report", middleware.RequireRoleMiddleware(string(enum.RoleCompany)), h.GenerateJobPostingDetailPDF)
	jobPostingGroup.Post("/", middleware.RequireRoleMiddleware(string(enum.RoleCompany)), h.Create)
	jobPostingGroup.Patch("/:id", middleware.RequireRoleMiddleware(string(enum.RoleCompany)), h.Update)
	jobPostingGroup.Delete("/:id", middleware.RequireRoleMiddleware(string(enum.RoleCompany)), h.SoftDelete)
	jobPostingGroup.Post("/:id/tags/:tag_id", middleware.RequireRoleMiddleware(string(enum.RoleCompany)), h.AddTagToJobPosting)
	jobPostingGroup.Delete("/:id/tags/:tag_id", middleware.RequireRoleMiddleware(string(enum.RoleCompany)), h.RemoveTagFromJobPosting)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var dto JobPostingCreateDto
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	id := c.Locals("user_id").(string)

	job, err := h.service.Create(id, dto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(job)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	var input JobPostingUpdateDto
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	companyId := c.Locals("user_id").(string)
	jobId := c.Params("id")

	job, err := h.service.Update(companyId, jobId, input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(job)
}

func (h *Handler) SoftDelete(c *fiber.Ctx) error {
	companyId := c.Locals("user_id").(string)
	jobId := c.Params("id")

	if err := h.service.AuthorizeCompanyAction(companyId, jobId); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.service.SoftDelete(jobId); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}

func (h *Handler) AddTagToJobPosting(c *fiber.Ctx) error {
	companyId := c.Locals("user_id").(string)
	jobId := c.Params("id")
	tagId := c.Params("tag_id")

	if err := h.service.AuthorizeCompanyAction(companyId, jobId); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.service.AddTagToJobPosting(jobId, tagId); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}

func (h *Handler) RemoveTagFromJobPosting(c *fiber.Ctx) error {
	companyId := c.Locals("user_id").(string)
	jobId := c.Params("id")
	tagId := c.Params("tag_id")

	if err := h.service.AuthorizeCompanyAction(companyId, jobId); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.service.RemoveTagFromJobPosting(jobId, tagId); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}

func (h *Handler) FindAll(c *fiber.Ctx) error {
	opts := helper.NewFindAllOptionsFromQuery(c)
	tagIDs := c.Query("tag_ids")
	var tags []string
	if tagIDs != "" {
		tags = strings.Split(tagIDs, ",")
	}

	companyID := c.Query("company_id", "")

	finded, err := h.service.FindAll(opts, tags, companyID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(finded)
}

func (h *Handler) FindById(c *fiber.Ctx) error {
	id := c.Params("id")
	job, err := h.service.FindById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(job)
}

func (h *Handler) GenerateJobPostingListPDF(c *fiber.Ctx) error {
	companyID := c.Locals("user_id").(string)

	// Get all job postings with applications for this company
	jobs, err := h.service.FindAllWithApplications(companyID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	pdfBytes, err := h.reportService.GenerateJobPostingListPDF(jobs)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate PDF",
		})
	}

	c.Set("Content-Type", "application/pdf")
	c.Set("Content-Disposition", "attachment; filename=ofertas-trabajo.pdf")
	return c.Send(pdfBytes)
}

func (h *Handler) GenerateJobPostingDetailPDF(c *fiber.Ctx) error {
	jobPostingID := c.Params("id")
	companyID := c.Locals("user_id").(string)

	// Get full job posting with applications
	fullJob, err := h.service.FindByIdWithApplications(jobPostingID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Job posting not found",
		})
	}

	// Verify that the job posting belongs to the company
	if fullJob.CompanyProfileId.String() != companyID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	pdfBytes, err := h.reportService.GenerateJobPostingDetailPDF(fullJob)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate PDF",
		})
	}

	c.Set("Content-Type", "application/pdf")
	c.Set("Content-Disposition", "attachment; filename=oferta-"+jobPostingID+".pdf")
	return c.Send(pdfBytes)
}