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

	// Modern Header
	pdf.SetFillColor(22, 160, 133)
	pdf.Rect(0, 0, 297, 30, "F")

	pdf.SetFont("Arial", "B", 20)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetY(10)
	pdf.MultiCell(0, 8, "REPORTE DE OFERTAS DE TRABAJO", "", "C", false)

	// Stats bar
	pdf.SetY(35)
	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFillColor(52, 152, 219)
	pdf.SetX(10)
	pdf.CellFormat(60, 7, fmt.Sprintf(" Total: %d ofertas", len(jobPostings)), "", 0, "L", true, 0, "")
	pdf.SetFillColor(155, 89, 182)
	pdf.SetX(75)
	pdf.CellFormat(70, 7, fmt.Sprintf(" %s", time.Now().Format("02/01/2006 15:04")), "", 0, "L", true, 0, "")

	pdf.SetTextColor(0, 0, 0)
	pdf.Ln(10)

	// Modern Table headers
	colWidths := []float64{55, 35, 30, 28, 25, 27}
	headers := []string{"Titulo", "Tipo Contrato", "Experiencia", "Aplicaciones", "Aceptadas", "Tasa Exito"}
	headerColors := [][3]int{
		{52, 73, 94},
		{41, 128, 185},
		{142, 68, 173},
		{52, 152, 219},
		{46, 204, 113},
		{243, 156, 18},
	}

	pdf.SetFont("Arial", "B", 8)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetDrawColor(200, 200, 200)

	for i, header := range headers {
		pdf.SetFillColor(headerColors[i][0], headerColors[i][1], headerColors[i][2])
		pdf.CellFormat(colWidths[i], 8, header, "1", 0, "C", true, 0, "")
	}
	pdf.Ln(8)

	// Table rows
	pdf.SetFont("Arial", "", 7)
	pdf.SetTextColor(50, 50, 50)

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
			pdf.SetFillColor(248, 249, 250)
		} else {
			pdf.SetFillColor(255, 255, 255)
		}
		alternateRow = !alternateRow

		// Print row
		pdf.SetFont("Arial", "", 7)
		pdf.SetTextColor(50, 50, 50)
		pdf.CellFormat(colWidths[0], 7, title, "1", 0, "L", true, 0, "")
		pdf.CellFormat(colWidths[1], 7, contractType, "1", 0, "L", true, 0, "")
		pdf.CellFormat(colWidths[2], 7, experience, "1", 0, "C", true, 0, "")

		// Metrics with color coding
		pdf.SetFont("Arial", "B", 7)
		pdf.SetTextColor(52, 152, 219)
		pdf.CellFormat(colWidths[3], 7, fmt.Sprintf("%d", totalApps), "1", 0, "C", true, 0, "")
		pdf.SetTextColor(46, 204, 113)
		pdf.CellFormat(colWidths[4], 7, fmt.Sprintf("%d", acceptedApps), "1", 0, "C", true, 0, "")

		// Success rate with color based on value
		if successRate >= 70 {
			pdf.SetTextColor(46, 204, 113) // Green
		} else if successRate >= 40 {
			pdf.SetTextColor(243, 156, 18) // Orange
		} else {
			pdf.SetTextColor(231, 76, 60) // Red
		}
		pdf.CellFormat(colWidths[5], 7, fmt.Sprintf("%.1f%%", successRate), "1", 1, "C", true, 0, "")

		pdf.SetTextColor(50, 50, 50)

		// Check if we need a new page
		if pdf.GetY() > 180 {
			pdf.AddPage()

			// Re-add headers on new page
			pdf.SetFont("Arial", "B", 8)
			pdf.SetTextColor(255, 255, 255)
			for i, header := range headers {
				pdf.SetFillColor(headerColors[i][0], headerColors[i][1], headerColors[i][2])
				pdf.CellFormat(colWidths[i], 8, header, "1", 0, "C", true, 0, "")
			}
			pdf.Ln(8)
			pdf.SetFont("Arial", "", 7)
			pdf.SetTextColor(50, 50, 50)
		}
	}

	// Footer
	pdf.SetY(200)
	pdf.SetFont("Arial", "I", 7)
	pdf.SetTextColor(150, 150, 150)
	pdf.SetDrawColor(200, 200, 200)
	pdf.Line(10, 198, 287, 198)
	pdf.MultiCell(0, 3, fmt.Sprintf("Documento confidencial - Total de ofertas: %d", len(jobPostings)), "", "C", false)

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

	// Modern Header
	pdf.SetFillColor(22, 160, 133)
	pdf.Rect(0, 0, 210, 35, "F")

	pdf.SetFont("Arial", "B", 20)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetY(12)
	pdf.MultiCell(0, 8, "OFERTA DE TRABAJO", "", "C", false)

	pdf.SetFont("Arial", "", 9)
	pdf.SetY(22)
	pdf.MultiCell(0, 5, "Reporte Detallado", "", "C", false)

	// Info bar
	pdf.SetY(38)
	pdf.SetFont("Arial", "I", 8)
	pdf.SetTextColor(120, 120, 120)
	pdf.MultiCell(0, 4, fmt.Sprintf("Generado: %s | ID: JP-%d", time.Now().Format("02/01/2006 15:04"), jobPosting.ID), "", "R", false)

	pdf.SetTextColor(0, 0, 0)
	pdf.Ln(5)

	// SECTION 1: Basic Information
	pdf.SetDrawColor(220, 220, 220)
	pdf.SetFillColor(248, 249, 250)
	pdf.RoundedRect(15, pdf.GetY(), 180, 40, 2, "1234", "FD")

	yStart := pdf.GetY()
	pdf.SetY(yStart + 3)

	pdf.SetFont("Arial", "B", 12)
	pdf.SetTextColor(22, 160, 133)
	pdf.SetX(20)
	pdf.Cell(0, 6, "INFORMACION BASICA")
	pdf.Ln(8)

	// Title
	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(90, 90, 90)
	pdf.SetX(25)
	pdf.Cell(35, 5, "Titulo:")
	pdf.SetFont("Arial", "B", 10)
	pdf.SetTextColor(22, 160, 133)
	pdf.MultiCell(0, 5, jobPosting.Title, "", "L", false)
	pdf.Ln(2)

	// Status with badge
	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(90, 90, 90)
	pdf.SetX(25)
	pdf.Cell(35, 5, "Estado:")

	statusText := "ACTIVA"
	statusColor := [3]int{46, 204, 113}
	if jobPosting.IsClosed {
		statusText = "CERRADA"
		statusColor = [3]int{231, 76, 60}
	}

	pdf.SetFillColor(statusColor[0], statusColor[1], statusColor[2])
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFont("Arial", "B", 9)
	pdf.CellFormat(35, 5, " "+statusText+" ", "", 0, "C", true, 0, "")
	pdf.Ln(7)

	// Location
	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(90, 90, 90)
	pdf.SetX(25)
	pdf.Cell(35, 5, "Ubicacion:")
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(50, 50, 50)
	pdf.Cell(0, 5, jobPosting.Location)

	pdf.SetY(yStart + 40)
	pdf.Ln(3)

	// SECTION 2: Work Details
	pdf.SetDrawColor(220, 220, 220)
	pdf.SetFillColor(248, 249, 250)
	pdf.RoundedRect(15, pdf.GetY(), 180, 45, 2, "1234", "FD")

	yStart = pdf.GetY()
	pdf.SetY(yStart + 3)

	pdf.SetFont("Arial", "B", 12)
	pdf.SetTextColor(22, 160, 133)
	pdf.SetX(20)
	pdf.Cell(0, 6, "DETALLES DEL TRABAJO")
	pdf.Ln(8)

	// Work Type & Experience in same line
	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(90, 90, 90)
	pdf.SetX(25)
	pdf.Cell(35, 5, "Tipo:")
	pdf.SetFont("Arial", "", 9)
	pdf.SetTextColor(50, 50, 50)
	pdf.Cell(50, 5, jobPosting.WorkType)

	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(90, 90, 90)
	pdf.Cell(25, 5, "Experiencia:")
	pdf.SetFont("Arial", "", 9)
	pdf.SetTextColor(50, 50, 50)
	pdf.Cell(0, 5, jobPosting.ExperienceLevel)
	pdf.Ln(6)

	// Contract Type
	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(90, 90, 90)
	pdf.SetX(25)
	pdf.Cell(35, 5, "Contrato:")
	pdf.SetFont("Arial", "", 9)
	pdf.SetTextColor(50, 50, 50)
	pdf.Cell(0, 5, jobPosting.ContractType)
	pdf.Ln(6)

	// Remote & Hybrid as badges
	pdf.SetX(25)
	pdf.SetFont("Arial", "B", 8)

	if jobPosting.IsRemote {
		pdf.SetFillColor(52, 152, 219)
		pdf.SetTextColor(255, 255, 255)
		pdf.CellFormat(30, 5, " REMOTO ", "", 0, "C", true, 0, "")
	}

	pdf.SetX(60)
	if jobPosting.IsHibrid {
		pdf.SetFillColor(155, 89, 182)
		pdf.SetTextColor(255, 255, 255)
		pdf.CellFormat(30, 5, " HIBRIDO ", "", 0, "C", true, 0, "")
	}

	pdf.SetY(yStart + 45)
	pdf.Ln(3)

	// SECTION 3: Salary
	pdf.SetFont("Arial", "B", 12)
	pdf.SetTextColor(22, 160, 133)
	pdf.SetX(20)
	pdf.Cell(0, 6, "RANGO SALARIAL")
	pdf.Ln(8)

	yStart = pdf.GetY()

	// Min Salary Card
	pdf.SetDrawColor(46, 204, 113)
	pdf.SetFillColor(46, 204, 113)
	pdf.RoundedRect(25, yStart, 70, 20, 2, "1234", "FD")
	pdf.SetY(yStart + 3)
	pdf.SetX(25)
	pdf.SetFont("Arial", "B", 8)
	pdf.SetTextColor(255, 255, 255)
	pdf.CellFormat(70, 5, "MINIMO", "", 0, "C", false, 0, "")
	pdf.SetY(yStart + 10)
	pdf.SetX(25)
	pdf.SetFont("Arial", "B", 11)
	pdf.CellFormat(70, 6, jobPosting.SalaryMin, "", 0, "C", false, 0, "")

	// Max Salary Card
	pdf.SetDrawColor(231, 76, 60)
	pdf.SetFillColor(231, 76, 60)
	pdf.RoundedRect(105, yStart, 70, 20, 2, "1234", "FD")
	pdf.SetY(yStart + 3)
	pdf.SetX(105)
	pdf.SetFont("Arial", "B", 8)
	pdf.CellFormat(70, 5, "MAXIMO", "", 0, "C", false, 0, "")
	pdf.SetY(yStart + 10)
	pdf.SetX(105)
	pdf.SetFont("Arial", "B", 11)
	pdf.CellFormat(70, 6, jobPosting.SalaryMax, "", 0, "C", false, 0, "")

	pdf.SetY(yStart + 20)
	pdf.Ln(5)

	// SECTION 4: Description
	pdf.SetFont("Arial", "B", 12)
	pdf.SetTextColor(22, 160, 133)
	pdf.SetX(20)
	pdf.Cell(0, 6, "DESCRIPCION")
	pdf.Ln(6)

	pdf.SetDrawColor(220, 220, 220)
	pdf.SetFillColor(255, 255, 255)
	pdf.RoundedRect(20, pdf.GetY(), 160, 0, 2, "1234", "D")
	pdf.SetFont("Arial", "", 9)
	pdf.SetTextColor(60, 60, 60)
	pdf.SetX(25)
	pdf.MultiCell(150, 4, jobPosting.Description, "", "J", false)
	pdf.Ln(3)

	// SECTION 5: Requirements
	pdf.SetFont("Arial", "B", 12)
	pdf.SetTextColor(22, 160, 133)
	pdf.SetX(20)
	pdf.Cell(0, 6, "REQUISITOS")
	pdf.Ln(6)

	pdf.SetDrawColor(220, 220, 220)
	pdf.RoundedRect(20, pdf.GetY(), 160, 0, 2, "1234", "D")
	pdf.SetFont("Arial", "", 9)
	pdf.SetTextColor(60, 60, 60)
	pdf.SetX(25)
	pdf.MultiCell(150, 4, jobPosting.Requirement, "", "J", false)
	pdf.Ln(3)

	// SECTION 6: Benefits
	pdf.SetFont("Arial", "B", 12)
	pdf.SetTextColor(22, 160, 133)
	pdf.SetX(20)
	pdf.Cell(0, 6, "BENEFICIOS")
	pdf.Ln(6)

	pdf.SetDrawColor(220, 220, 220)
	pdf.RoundedRect(20, pdf.GetY(), 160, 0, 2, "1234", "D")
	pdf.SetFont("Arial", "", 9)
	pdf.SetTextColor(60, 60, 60)
	pdf.SetX(25)
	pdf.MultiCell(150, 4, jobPosting.Benefit, "", "J", false)
	pdf.Ln(5)

	// SECTION 7: Application Statistics
	pdf.SetFont("Arial", "B", 12)
	pdf.SetTextColor(22, 160, 133)
	pdf.SetX(20)
	pdf.Cell(0, 6, "ESTADISTICAS DE APLICACIONES")
	pdf.Ln(10)

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

	yStart = pdf.GetY()

	// Total Applications Card
	pdf.SetDrawColor(52, 152, 219)
	pdf.SetFillColor(52, 152, 219)
	pdf.RoundedRect(25, yStart, 50, 25, 3, "1234", "FD")
	pdf.SetY(yStart + 4)
	pdf.SetX(25)
	pdf.SetFont("Arial", "B", 18)
	pdf.SetTextColor(255, 255, 255)
	pdf.CellFormat(50, 8, fmt.Sprintf("%d", totalApps), "", 0, "C", false, 0, "")
	pdf.SetY(yStart + 14)
	pdf.SetX(25)
	pdf.SetFont("Arial", "B", 8)
	pdf.CellFormat(50, 5, "APLICACIONES", "", 0, "C", false, 0, "")

	// Accepted Card
	pdf.SetDrawColor(46, 204, 113)
	pdf.SetFillColor(46, 204, 113)
	pdf.RoundedRect(80, yStart, 50, 25, 3, "1234", "FD")
	pdf.SetY(yStart + 4)
	pdf.SetX(80)
	pdf.SetFont("Arial", "B", 18)
	pdf.CellFormat(50, 8, fmt.Sprintf("%d", acceptedApps), "", 0, "C", false, 0, "")
	pdf.SetY(yStart + 14)
	pdf.SetX(80)
	pdf.SetFont("Arial", "B", 8)
	pdf.CellFormat(50, 5, "ACEPTADAS", "", 0, "C", false, 0, "")

	// Rejected Card
	pdf.SetDrawColor(231, 76, 60)
	pdf.SetFillColor(231, 76, 60)
	pdf.RoundedRect(135, yStart, 50, 25, 3, "1234", "FD")
	pdf.SetY(yStart + 4)
	pdf.SetX(135)
	pdf.SetFont("Arial", "B", 18)
	pdf.CellFormat(50, 8, fmt.Sprintf("%d", rejectedApps), "", 0, "C", false, 0, "")
	pdf.SetY(yStart + 14)
	pdf.SetX(135)
	pdf.SetFont("Arial", "B", 8)
	pdf.CellFormat(50, 5, "RECHAZADAS", "", 0, "C", false, 0, "")

	pdf.SetY(yStart + 25)
	pdf.Ln(8)

	// Success Rate with progress bar
	pdf.SetFont("Arial", "B", 10)
	pdf.SetTextColor(50, 50, 50)
	pdf.SetX(25)
	pdf.Cell(0, 6, "TASA DE EXITO")
	pdf.Ln(8)

	// Progress bar background
	pdf.SetDrawColor(220, 220, 220)
	pdf.SetFillColor(240, 240, 240)
	pdf.RoundedRect(25, pdf.GetY(), 160, 8, 2, "1234", "FD")

	// Progress bar fill
	barColor := [3]int{46, 204, 113}
	if successRate < 40 {
		barColor = [3]int{231, 76, 60}
	} else if successRate < 70 {
		barColor = [3]int{243, 156, 18}
	}

	barWidth := (successRate / 100) * 160
	pdf.SetFillColor(barColor[0], barColor[1], barColor[2])
	if barWidth > 0 {
		pdf.RoundedRect(25, pdf.GetY(), barWidth, 8, 2, "1234", "F")
	}

	// Percentage text
	pdf.SetFont("Arial", "B", 10)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetY(pdf.GetY() + 1)
	pdf.SetX(25)
	pdf.CellFormat(160, 6, fmt.Sprintf("%.1f%%", successRate), "", 0, "C", false, 0, "")

	pdf.Ln(10)

	// SECTION 8: Application Details
	if totalApps > 0 {
		pdf.SetFont("Arial", "B", 12)
		pdf.SetTextColor(22, 160, 133)
		pdf.SetX(20)
		pdf.Cell(0, 6, "DETALLE DE APLICACIONES")
		pdf.Ln(8)

		// Table headers
		colWidths := []float64{60, 45, 30, 25}
		headers := []string{"Nombre", "Email", "Estado", "Aceptado"}
		headerColors := [][3]int{
			{52, 73, 94},
			{41, 128, 185},
			{155, 89, 182},
			{46, 204, 113},
		}

		pdf.SetFont("Arial", "B", 8)
		pdf.SetTextColor(255, 255, 255)
		pdf.SetDrawColor(200, 200, 200)
		pdf.SetX(20)

		for i, header := range headers {
			pdf.SetFillColor(headerColors[i][0], headerColors[i][1], headerColors[i][2])
			pdf.CellFormat(colWidths[i], 7, header, "1", 0, "C", true, 0, "")
		}
		pdf.Ln(7)

		// Table rows
		pdf.SetFont("Arial", "", 7)
		pdf.SetTextColor(50, 50, 50)

		alternate := false
		for _, app := range jobPosting.Applications {
			applicantName := app.JobSeeker.Name
			if len(applicantName) > 28 {
				applicantName = applicantName[:25] + "..."
			}

			applicantEmail := app.JobSeeker.Email
			if len(applicantEmail) > 22 {
				applicantEmail = applicantEmail[:19] + "..."
			}

			// Alternate row colors
			if alternate {
				pdf.SetFillColor(248, 249, 250)
			} else {
				pdf.SetFillColor(255, 255, 255)
			}
			alternate = !alternate

			pdf.SetX(20)
			pdf.CellFormat(colWidths[0], 6, applicantName, "1", 0, "L", true, 0, "")
			pdf.CellFormat(colWidths[1], 6, applicantEmail, "1", 0, "L", true, 0, "")
			pdf.CellFormat(colWidths[2], 6, app.Status, "1", 0, "C", true, 0, "")

			// Accepted badge
			if app.IsAccepted {
				pdf.SetTextColor(46, 204, 113)
				pdf.SetFont("Arial", "B", 7)
				pdf.CellFormat(colWidths[3], 6, "SI", "1", 1, "C", true, 0, "")
			} else {
				pdf.SetTextColor(231, 76, 60)
				pdf.SetFont("Arial", "B", 7)
				pdf.CellFormat(colWidths[3], 6, "NO", "1", 1, "C", true, 0, "")
			}
			pdf.SetFont("Arial", "", 7)
			pdf.SetTextColor(50, 50, 50)

			// Check if we need a new page
			if pdf.GetY() > 250 {
				pdf.AddPage()

				// Re-add headers
				pdf.SetFont("Arial", "B", 8)
				pdf.SetTextColor(255, 255, 255)
				pdf.SetX(20)
				for i, header := range headers {
					pdf.SetFillColor(headerColors[i][0], headerColors[i][1], headerColors[i][2])
					pdf.CellFormat(colWidths[i], 7, header, "1", 0, "C", true, 0, "")
				}
				pdf.Ln(7)
				pdf.SetFont("Arial", "", 7)
				pdf.SetTextColor(50, 50, 50)
			}
		}
	}

	// Footer
	pdf.SetY(280)
	pdf.SetFont("Arial", "I", 7)
	pdf.SetTextColor(150, 150, 150)
	pdf.SetDrawColor(200, 200, 200)
	pdf.Line(15, 278, 195, 278)
	pdf.MultiCell(0, 3, "Documento confidencial - Solo para uso autorizado", "", "C", false)

	err := pdf.Output(buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
