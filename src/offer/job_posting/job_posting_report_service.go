package jobposting

import (
	"bytes"
	"fmt"
	"time"

	"github.com/go-pdf/fpdf"
	"github.com/ito-company/jobsito-service/src/model"
)

type JobPostingReportService interface {
	GenerateJobPostingListPDF(jobPostings []model.JobPosting) ([]byte, error)
	GenerateJobPostingDetailPDF(jobPosting *model.JobPosting) ([]byte, error)
}

type jobPostingReportService struct{}

func NewJobPostingReportService() JobPostingReportService {
	return &jobPostingReportService{}
}

func (rs *jobPostingReportService) GenerateJobPostingListPDF(jobPostings []model.JobPosting) ([]byte, error) {
	buf := new(bytes.Buffer)
	pdf := fpdf.New("L", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetMargins(10, 10, 10)

	// Header
	pdf.SetFont("Arial", "B", 16)
	pdf.SetTextColor(33, 37, 41)
	pdf.MultiCell(0, 10, "REPORTE DE OFERTAS DE TRABAJO", "", "C", false)

	// Date
	pdf.SetFont("Arial", "", 9)
	pdf.SetTextColor(100, 100, 100)
	pdf.MultiCell(0, 4, fmt.Sprintf("Fecha de generacion: %s", time.Now().Format("02/01/2006 15:04")), "", "C", false)
	pdf.MultiCell(0, 4, fmt.Sprintf("Total de ofertas: %d", len(jobPostings)), "", "C", false)
	pdf.Ln(3)

	// Table headers
	colWidths := []float64{50, 30, 25, 25, 20, 20}
	headers := []string{"Titulo", "Tipo de Contrato", "Nivel Experiencia", "Aplicaciones", "Aceptadas", "Tasa Exito"}

	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFillColor(52, 73, 94)

	for i, header := range headers {
		pdf.CellFormat(colWidths[i], 8, header, "1", 0, "C", true, 0, "")
	}
	pdf.Ln(8)

	// Table rows
	pdf.SetFont("Arial", "", 8)
	pdf.SetTextColor(0, 0, 0)

	alternateRow := false
	for _, jobPosting := range jobPostings {
		title := jobPosting.Title
		if len(title) > 25 {
			title = title[:22] + "..."
		}

		contractType := jobPosting.ContractType
		if len(contractType) > 15 {
			contractType = contractType[:12] + "..."
		}

		experience := jobPosting.ExperienceLevel
		if len(experience) > 15 {
			experience = experience[:12] + "..."
		}

		// Count applications
		totalApps := len(jobPosting.Applications)
		acceptedApps := 0
		for _, app := range jobPosting.Applications {
			if app.IsAccepted {
				acceptedApps++
			}
		}

		// Calculate success rate
		successRate := 0.0
		if totalApps > 0 {
			successRate = (float64(acceptedApps) / float64(totalApps)) * 100
		}

		// Alternate row colors
		if alternateRow {
			pdf.SetFillColor(245, 245, 245)
		} else {
			pdf.SetFillColor(255, 255, 255)
		}
		alternateRow = !alternateRow

		// Print row
		pdf.CellFormat(colWidths[0], 7, title, "1", 0, "L", true, 0, "")
		pdf.CellFormat(colWidths[1], 7, contractType, "1", 0, "L", true, 0, "")
		pdf.CellFormat(colWidths[2], 7, experience, "1", 0, "C", true, 0, "")
		pdf.CellFormat(colWidths[3], 7, fmt.Sprintf("%d", totalApps), "1", 0, "C", true, 0, "")
		pdf.CellFormat(colWidths[4], 7, fmt.Sprintf("%d", acceptedApps), "1", 0, "C", true, 0, "")
		pdf.CellFormat(colWidths[5], 7, fmt.Sprintf("%.1f%%", successRate), "1", 1, "C", true, 0, "")

		// Check if we need a new page
		if pdf.GetY() > 180 {
			pdf.AddPage()
			// Re-add headers on new page
			pdf.SetFont("Arial", "B", 9)
			pdf.SetTextColor(255, 255, 255)
			pdf.SetFillColor(52, 73, 94)
			for i, header := range headers {
				pdf.CellFormat(colWidths[i], 8, header, "1", 0, "C", true, 0, "")
			}
			pdf.Ln(8)
			pdf.SetFont("Arial", "", 8)
			pdf.SetTextColor(0, 0, 0)
		}
	}

	err := pdf.Output(buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (rs *jobPostingReportService) GenerateJobPostingDetailPDF(jobPosting *model.JobPosting) ([]byte, error) {
	buf := new(bytes.Buffer)
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetMargins(15, 15, 15)

	// Header
	pdf.SetFont("Arial", "B", 18)
	pdf.SetTextColor(33, 37, 41)
	pdf.MultiCell(0, 10, "REPORTE DETALLADO DE OFERTA DE TRABAJO", "", "C", false)

	// Separator line
	pdf.SetDrawColor(52, 152, 219)
	pdf.SetLineWidth(1)
	pdf.Line(15, pdf.GetY(), 195, pdf.GetY())
	pdf.Ln(5)

	// Generation date
	pdf.SetFont("Arial", "", 9)
	pdf.SetTextColor(100, 100, 100)
	pdf.MultiCell(0, 4, fmt.Sprintf("Fecha de generacion: %s", time.Now().Format("02/01/2006 15:04")), "", "R", false)
	pdf.Ln(3)

	// SECTION 1: Basic Information
	pdf.SetFont("Arial", "B", 11)
	pdf.SetTextColor(52, 152, 219)
	pdf.MultiCell(0, 6, "INFORMACION BASICA", "", "L", false)

	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(33, 37, 41)

	// Title
	pdf.SetFont("Arial", "B", 10)
	pdf.SetTextColor(52, 73, 94)
	pdf.MultiCell(0, 5, "Titulo:", "", "L", false)
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(50, 50, 50)
	pdf.SetX(20)
	pdf.MultiCell(0, 5, jobPosting.Title, "", "L", false)

	// Status
	pdf.SetFont("Arial", "B", 10)
	pdf.SetTextColor(52, 73, 94)
	pdf.MultiCell(0, 5, "Estado:", "", "L", false)
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(50, 50, 50)
	pdf.SetX(20)
	statusText := "Activa"
	if jobPosting.IsClosed {
		statusText = "Cerrada"
	}
	pdf.MultiCell(0, 5, statusText, "", "L", false)

	// Location
	pdf.SetFont("Arial", "B", 10)
	pdf.SetTextColor(52, 73, 94)
	pdf.MultiCell(0, 5, "Ubicacion:", "", "L", false)
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(50, 50, 50)
	pdf.SetX(20)
	pdf.MultiCell(0, 5, jobPosting.Location, "", "L", false)

	pdf.Ln(2)

	// SECTION 2: Work Details
	pdf.SetFont("Arial", "B", 11)
	pdf.SetTextColor(52, 152, 219)
	pdf.MultiCell(0, 6, "DETALLES DEL TRABAJO", "", "L", false)

	// Work Type
	pdf.SetFont("Arial", "B", 10)
	pdf.SetTextColor(52, 73, 94)
	pdf.MultiCell(0, 5, "Tipo de Trabajo:", "", "L", false)
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(50, 50, 50)
	pdf.SetX(20)
	pdf.MultiCell(0, 5, jobPosting.WorkType, "", "L", false)

	// Experience Level
	pdf.SetFont("Arial", "B", 10)
	pdf.SetTextColor(52, 73, 94)
	pdf.MultiCell(0, 5, "Nivel de Experiencia:", "", "L", false)
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(50, 50, 50)
	pdf.SetX(20)
	pdf.MultiCell(0, 5, jobPosting.ExperienceLevel, "", "L", false)

	// Contract Type
	pdf.SetFont("Arial", "B", 10)
	pdf.SetTextColor(52, 73, 94)
	pdf.MultiCell(0, 5, "Tipo de Contrato:", "", "L", false)
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(50, 50, 50)
	pdf.SetX(20)
	pdf.MultiCell(0, 5, jobPosting.ContractType, "", "L", false)

	// Remote
	pdf.SetFont("Arial", "B", 10)
	pdf.SetTextColor(52, 73, 94)
	pdf.MultiCell(0, 5, "Remoto:", "", "L", false)
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(50, 50, 50)
	pdf.SetX(20)
	remoteText := "No"
	if jobPosting.IsRemote {
		remoteText = "Si"
	}
	pdf.MultiCell(0, 5, remoteText, "", "L", false)

	// Hybrid
	pdf.SetFont("Arial", "B", 10)
	pdf.SetTextColor(52, 73, 94)
	pdf.MultiCell(0, 5, "Hibrido:", "", "L", false)
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(50, 50, 50)
	pdf.SetX(20)
	hybridText := "No"
	if jobPosting.IsHibrid {
		hybridText = "Si"
	}
	pdf.MultiCell(0, 5, hybridText, "", "L", false)

	pdf.Ln(2)

	// SECTION 3: Salary
	pdf.SetFont("Arial", "B", 11)
	pdf.SetTextColor(52, 152, 219)
	pdf.MultiCell(0, 6, "RANGO SALARIAL", "", "L", false)

	// Salary Min
	pdf.SetFont("Arial", "B", 10)
	pdf.SetTextColor(52, 73, 94)
	pdf.MultiCell(0, 5, "Minimo:", "", "L", false)
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(50, 50, 50)
	pdf.SetX(20)
	pdf.MultiCell(0, 5, jobPosting.SalaryMin, "", "L", false)

	// Salary Max
	pdf.SetFont("Arial", "B", 10)
	pdf.SetTextColor(52, 73, 94)
	pdf.MultiCell(0, 5, "Maximo:", "", "L", false)
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(50, 50, 50)
	pdf.SetX(20)
	pdf.MultiCell(0, 5, jobPosting.SalaryMax, "", "L", false)

	pdf.Ln(2)

	// SECTION 4: Description
	pdf.SetFont("Arial", "B", 11)
	pdf.SetTextColor(52, 152, 219)
	pdf.MultiCell(0, 6, "DESCRIPCION", "", "L", false)
	pdf.SetFont("Arial", "", 9)
	pdf.SetTextColor(50, 50, 50)
	pdf.SetX(15)
	pdf.MultiCell(0, 4, jobPosting.Description, "", "L", false)

	pdf.Ln(2)

	// SECTION 5: Requirements
	pdf.SetFont("Arial", "B", 11)
	pdf.SetTextColor(52, 152, 219)
	pdf.MultiCell(0, 6, "REQUISITOS", "", "L", false)
	pdf.SetFont("Arial", "", 9)
	pdf.SetTextColor(50, 50, 50)
	pdf.SetX(15)
	pdf.MultiCell(0, 4, jobPosting.Requirement, "", "L", false)

	pdf.Ln(2)

	// SECTION 6: Benefits
	pdf.SetFont("Arial", "B", 11)
	pdf.SetTextColor(52, 152, 219)
	pdf.MultiCell(0, 6, "BENEFICIOS", "", "L", false)
	pdf.SetFont("Arial", "", 9)
	pdf.SetTextColor(50, 50, 50)
	pdf.SetX(15)
	pdf.MultiCell(0, 4, jobPosting.Benefit, "", "L", false)

	pdf.Ln(5)

	// SECTION 7: Application Statistics
	pdf.SetFont("Arial", "B", 11)
	pdf.SetTextColor(52, 152, 219)
	pdf.MultiCell(0, 6, "ESTADISTICAS DE APLICACIONES", "", "L", false)

	// Count applications
	totalApps := len(jobPosting.Applications)
	acceptedApps := 0
	rejectedApps := 0
	for _, app := range jobPosting.Applications {
		if app.IsAccepted {
			acceptedApps++
		} else {
			rejectedApps++
		}
	}

	successRate := 0.0
	if totalApps > 0 {
		successRate = (float64(acceptedApps) / float64(totalApps)) * 100
	}

	// Summary boxes
	pdf.SetFont("Arial", "B", 10)
	pdf.SetTextColor(255, 255, 255)

	// Total Applications
	pdf.SetFillColor(52, 152, 219)
	pdf.SetX(15)
	pdf.MultiCell(50, 10, fmt.Sprintf("TOTAL APLICACIONES\n%d", totalApps), "", "C", true)

	// Accepted
	pdf.SetX(80)
	pdf.SetFillColor(46, 204, 113)
	pdf.MultiCell(50, 10, fmt.Sprintf("ACEPTADAS\n%d", acceptedApps), "", "C", true)

	// Rejected
	pdf.SetX(145)
	pdf.SetFillColor(231, 76, 60)
	pdf.MultiCell(0, 10, fmt.Sprintf("RECHAZADAS\n%d", rejectedApps), "", "C", true)

	pdf.Ln(3)

	// Success Rate
	pdf.SetFont("Arial", "B", 11)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFillColor(155, 89, 182)
	pdf.SetX(15)
	pdf.MultiCell(0, 8, fmt.Sprintf("TASA DE EXITO: %.1f%%", successRate), "", "C", true)

	pdf.Ln(5)

	// SECTION 8: Application Details
	if totalApps > 0 {
		pdf.SetFont("Arial", "B", 11)
		pdf.SetTextColor(52, 152, 219)
		pdf.MultiCell(0, 6, "DETALLE DE APLICACIONES", "", "L", false)

		// Table headers
		pdf.SetFont("Arial", "B", 9)
		pdf.SetTextColor(255, 255, 255)
		pdf.SetFillColor(52, 73, 94)
		pdf.SetX(15)
		pdf.CellFormat(60, 7, "Nombre del Aplicante", "1", 0, "L", true, 0, "")
		pdf.CellFormat(40, 7, "Email", "1", 0, "L", true, 0, "")
		pdf.CellFormat(30, 7, "Estado", "1", 0, "C", true, 0, "")
		pdf.CellFormat(25, 7, "Aceptado", "1", 1, "C", true, 0, "")

		// Table rows
		pdf.SetFont("Arial", "", 8)
		pdf.SetTextColor(0, 0, 0)

		for _, app := range jobPosting.Applications {
			applicantName := app.JobSeeker.Name
			if len(applicantName) > 25 {
				applicantName = applicantName[:22] + "..."
			}

			applicantEmail := app.JobSeeker.Email
			if len(applicantEmail) > 20 {
				applicantEmail = applicantEmail[:17] + "..."
			}

			acceptedText := "No"
			if app.IsAccepted {
				acceptedText = "Si"
			}

			pdf.SetX(15)
			pdf.CellFormat(60, 6, applicantName, "1", 0, "L", false, 0, "")
			pdf.CellFormat(40, 6, applicantEmail, "1", 0, "L", false, 0, "")
			pdf.CellFormat(30, 6, app.Status, "1", 0, "C", false, 0, "")
			pdf.CellFormat(25, 6, acceptedText, "1", 1, "C", false, 0, "")

			// Check if we need a new page
			if pdf.GetY() > 250 {
				pdf.AddPage()
				pdf.SetFont("Arial", "B", 9)
				pdf.SetTextColor(255, 255, 255)
				pdf.SetFillColor(52, 73, 94)
				pdf.SetX(15)
				pdf.CellFormat(60, 7, "Nombre del Aplicante", "1", 0, "L", true, 0, "")
				pdf.CellFormat(40, 7, "Email", "1", 0, "L", true, 0, "")
				pdf.CellFormat(30, 7, "Estado", "1", 0, "C", true, 0, "")
				pdf.CellFormat(25, 7, "Aceptado", "1", 1, "C", true, 0, "")
				pdf.SetFont("Arial", "", 8)
				pdf.SetTextColor(0, 0, 0)
			}
		}
	}

	err := pdf.Output(buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
