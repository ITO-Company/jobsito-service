package intership

import (
	"bytes"
	"fmt"
	"time"

	"github.com/go-pdf/fpdf"
	"github.com/ito-company/jobsito-service/src/dto"
)

type ReportService interface {
	GenerateOverviewPDF(overview *dto.IntershipOverviewDto) ([]byte, error)
	GenerateOverviewListPDF(overviews []dto.IntershipOverviewDto, total int64) ([]byte, error)
}

type reportService struct{}

func NewReportService() ReportService {
	return &reportService{}
}

func (rs *reportService) GenerateOverviewPDF(overview *dto.IntershipOverviewDto) ([]byte, error) {
	buf := new(bytes.Buffer)
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetMargins(15, 15, 15)

	// Header
	pdf.SetFont("Arial", "B", 18)
	pdf.SetTextColor(33, 37, 41)
	pdf.MultiCell(0, 10, "REPORTE DETALLADO DE PASANTIA", "", "C", false)

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

	// SECTION 1: Company and Internship Info
	pdf.SetFont("Arial", "B", 11)
	pdf.SetTextColor(52, 152, 219)
	pdf.MultiCell(0, 6, "INFORMACION DE LA PASANTIA", "", "L", false)

	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(33, 37, 41)

	// Company
	pdf.SetFont("Arial", "B", 10)
	pdf.SetTextColor(52, 73, 94)
	pdf.MultiCell(0, 5, "Empresa:", "", "L", false)
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(50, 50, 50)
	pdf.SetX(20)
	pdf.MultiCell(0, 5, overview.CompanyProfile.Name, "", "L", false)

	// Position
	pdf.SetFont("Arial", "B", 10)
	pdf.SetTextColor(52, 73, 94)
	pdf.MultiCell(0, 5, "Posicion:", "", "L", false)
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(50, 50, 50)
	pdf.SetX(20)
	pdf.MultiCell(0, 5, overview.JobPosting.Title, "", "L", false)

	// Status
	pdf.SetFont("Arial", "B", 10)
	pdf.SetTextColor(52, 73, 94)
	pdf.MultiCell(0, 5, "Estado:", "", "L", false)
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(50, 50, 50)
	pdf.SetX(20)
	pdf.MultiCell(0, 5, overview.Status, "", "L", false)

	pdf.Ln(3)

	// SECTION 2: Intern Info
	pdf.SetFont("Arial", "B", 11)
	pdf.SetTextColor(52, 152, 219)
	pdf.MultiCell(0, 6, "INFORMACION DEL PASANTE", "", "L", false)

	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(33, 37, 41)

	// Name
	pdf.SetFont("Arial", "B", 10)
	pdf.SetTextColor(52, 73, 94)
	pdf.MultiCell(0, 5, "Nombre:", "", "L", false)
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(50, 50, 50)
	pdf.SetX(20)
	pdf.MultiCell(0, 5, overview.JobSeeker.Name, "", "L", false)

	// Email
	pdf.SetFont("Arial", "B", 10)
	pdf.SetTextColor(52, 73, 94)
	pdf.MultiCell(0, 5, "Email:", "", "L", false)
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(50, 50, 50)
	pdf.SetX(20)
	pdf.MultiCell(0, 5, overview.JobSeeker.Email, "", "L", false)

	pdf.Ln(3)

	// SECTION 3: Dates
	pdf.SetFont("Arial", "B", 11)
	pdf.SetTextColor(52, 152, 219)
	pdf.MultiCell(0, 6, "FECHAS", "", "L", false)

	// Start Date
	pdf.SetFont("Arial", "B", 10)
	pdf.SetTextColor(52, 73, 94)
	pdf.MultiCell(0, 5, "Inicio:", "", "L", false)
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(50, 50, 50)
	pdf.SetX(20)
	pdf.MultiCell(0, 5, overview.StartDate.Format("02/01/2006"), "", "L", false)

	// End Date
	pdf.SetFont("Arial", "B", 10)
	pdf.SetTextColor(52, 73, 94)
	pdf.MultiCell(0, 5, "Termino:", "", "L", false)
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(50, 50, 50)
	pdf.SetX(20)
	pdf.MultiCell(0, 5, overview.EndDate.Format("02/01/2006"), "", "L", false)

	pdf.Ln(5)

	// SECTION 4: Summary boxes
	pdf.SetFont("Arial", "B", 11)
	pdf.SetTextColor(52, 152, 219)
	pdf.MultiCell(0, 6, "RESUMEN", "", "L", false)

	// Count totals
	milestoneTotalCount := 0
	for _, count := range overview.MilestoneCounts {
		milestoneTotalCount += count.Count
	}

	issueTotalCount := 0
	for _, count := range overview.IssueCounts {
		issueTotalCount += count.Count
	}

	requestTotalCount := 0
	for _, count := range overview.RequestCounts {
		requestTotalCount += count.Count
	}

	// Milestones
	pdf.SetFont("Arial", "B", 10)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFillColor(52, 152, 219)
	pdf.SetX(15)
	pdf.MultiCell(50, 8, fmt.Sprintf("HITOS\n%d", milestoneTotalCount), "", "C", true)

	// Issues
	pdf.SetX(80)
	pdf.SetFillColor(155, 89, 182)
	pdf.MultiCell(50, 8, fmt.Sprintf("ISSUES\n%d", issueTotalCount), "", "C", true)

	// Requests
	pdf.SetX(145)
	pdf.SetFillColor(46, 204, 113)
	pdf.MultiCell(0, 8, fmt.Sprintf("REQUESTS\n%d", requestTotalCount), "", "C", true)

	pdf.Ln(5)

	// SECTION 5: Details by status
	pdf.SetFont("Arial", "B", 11)
	pdf.SetTextColor(52, 152, 219)
	pdf.MultiCell(0, 6, "DETALLES POR ESTADO", "", "L", false)

	if len(overview.MilestoneCounts) > 0 {
		pdf.SetFont("Arial", "B", 10)
		pdf.SetTextColor(52, 152, 219)
		pdf.SetX(15)
		pdf.MultiCell(0, 5, "Hitos:", "", "L", false)
		pdf.SetFont("Arial", "", 9)
		pdf.SetTextColor(50, 50, 50)
		for _, count := range overview.MilestoneCounts {
			pdf.SetX(25)
			pdf.MultiCell(0, 4, fmt.Sprintf("- %s: %d", count.Status, count.Count), "", "L", false)
		}
		pdf.Ln(2)
	}

	if len(overview.IssueCounts) > 0 {
		pdf.SetFont("Arial", "B", 10)
		pdf.SetTextColor(155, 89, 182)
		pdf.SetX(15)
		pdf.MultiCell(0, 5, "Issues:", "", "L", false)
		pdf.SetFont("Arial", "", 9)
		pdf.SetTextColor(50, 50, 50)
		for _, count := range overview.IssueCounts {
			pdf.SetX(25)
			pdf.MultiCell(0, 4, fmt.Sprintf("- %s: %d", count.Status, count.Count), "", "L", false)
		}
		pdf.Ln(2)
	}

	if len(overview.RequestCounts) > 0 {
		pdf.SetFont("Arial", "B", 10)
		pdf.SetTextColor(46, 204, 113)
		pdf.SetX(15)
		pdf.MultiCell(0, 5, "Requests:", "", "L", false)
		pdf.SetFont("Arial", "", 9)
		pdf.SetTextColor(50, 50, 50)
		for _, count := range overview.RequestCounts {
			pdf.SetX(25)
			pdf.MultiCell(0, 4, fmt.Sprintf("- %s: %d", count.Status, count.Count), "", "L", false)
		}
	}

	err := pdf.Output(buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (rs *reportService) GenerateOverviewListPDF(overviews []dto.IntershipOverviewDto, total int64) ([]byte, error) {
	buf := new(bytes.Buffer)
	pdf := fpdf.New("L", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetMargins(10, 10, 10)

	// Header
	pdf.SetFont("Arial", "B", 16)
	pdf.SetTextColor(33, 37, 41)
	pdf.MultiCell(0, 10, "REPORTE DE TODAS LAS PASANTIAS", "", "C", false)

	// Date and total
	pdf.SetFont("Arial", "", 9)
	pdf.SetTextColor(100, 100, 100)
	pdf.MultiCell(0, 4, fmt.Sprintf("Fecha de generacion: %s", time.Now().Format("02/01/2006 15:04")), "", "C", false)
	pdf.MultiCell(0, 4, fmt.Sprintf("Total de registros: %d", total), "", "C", false)
	pdf.Ln(3)

	// Table headers
	colWidths := []float64{35, 30, 22, 22, 18, 18, 18}
	headers := []string{"Empresa", "Pasante", "Inicio", "Termino", "Hitos", "Issues", "Requests"}

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
	for _, overview := range overviews {
		// Get company name
		companyName := overview.CompanyProfile.Name
		if len(companyName) > 18 {
			companyName = companyName[:15] + "..."
		}

		// Get pasante name
		pasanteName := overview.JobSeeker.Name
		if len(pasanteName) > 16 {
			pasanteName = pasanteName[:13] + "..."
		}

		// Calculate totals
		milestoneTotalCount := 0
		for _, count := range overview.MilestoneCounts {
			milestoneTotalCount += count.Count
		}

		issueTotalCount := 0
		for _, count := range overview.IssueCounts {
			issueTotalCount += count.Count
		}

		requestTotalCount := 0
		for _, count := range overview.RequestCounts {
			requestTotalCount += count.Count
		}

		// Alternate row colors
		if alternateRow {
			pdf.SetFillColor(245, 245, 245)
		} else {
			pdf.SetFillColor(255, 255, 255)
		}
		alternateRow = !alternateRow

		// Print row
		pdf.CellFormat(colWidths[0], 7, companyName, "1", 0, "L", true, 0, "")
		pdf.CellFormat(colWidths[1], 7, pasanteName, "1", 0, "L", true, 0, "")
		pdf.CellFormat(colWidths[2], 7, overview.StartDate.Format("02/01/06"), "1", 0, "C", true, 0, "")
		pdf.CellFormat(colWidths[3], 7, overview.EndDate.Format("02/01/06"), "1", 0, "C", true, 0, "")
		pdf.CellFormat(colWidths[4], 7, fmt.Sprintf("%d", milestoneTotalCount), "1", 0, "C", true, 0, "")
		pdf.CellFormat(colWidths[5], 7, fmt.Sprintf("%d", issueTotalCount), "1", 0, "C", true, 0, "")
		pdf.CellFormat(colWidths[6], 7, fmt.Sprintf("%d", requestTotalCount), "1", 1, "C", true, 0, "")

		// Check if we need a new page
		if pdf.GetY() > 185 {
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
