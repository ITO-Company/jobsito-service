package intership

import (
	"bytes"
	"fmt"
	"time"

	"github.com/go-pdf/fpdf"
	"github.com/ito-company/jobsito-service/src/dto"
	"github.com/ito-company/jobsito-service/src/model"
)

type ReportService interface {
	GenerateOverviewPDF(overview *dto.IntershipOverviewDto) ([]byte, error)
	GenerateOverviewListPDF(overviews []dto.IntershipOverviewDto, total int64) ([]byte, error)
	GenerateDetailedPDF(intership *model.Intership) ([]byte, error)
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

	// Modern Header with Background
	pdf.SetFillColor(41, 128, 185)
	pdf.Rect(0, 0, 210, 35, "F")

	pdf.SetFont("Arial", "B", 22)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetY(12)
	pdf.MultiCell(0, 8, "REPORTE DE PASANTIA", "", "C", false)

	pdf.SetFont("Arial", "", 10)
	pdf.SetY(22)
	pdf.MultiCell(0, 5, "Documento Confidencial", "", "C", false)

	// Info bar
	pdf.SetY(38)
	pdf.SetFont("Arial", "I", 8)
	pdf.SetTextColor(120, 120, 120)
	pdf.MultiCell(0, 4, fmt.Sprintf("Generado: %s | ID: INS-%s", time.Now().Format("02/01/2006 15:04"), overview.ID), "", "R", false)

	pdf.SetTextColor(0, 0, 0)
	pdf.Ln(5)

	// SECTION 1: Company and Internship Info
	pdf.SetDrawColor(220, 220, 220)
	pdf.SetFillColor(248, 249, 250)
	pdf.RoundedRect(15, pdf.GetY(), 180, 35, 2, "1234", "FD")

	yStart := pdf.GetY()
	pdf.SetY(yStart + 3)

	pdf.SetFont("Arial", "B", 12)
	pdf.SetTextColor(41, 128, 185)
	pdf.SetX(20)
	pdf.Cell(0, 6, "INFORMACION DE LA PASANTIA")
	pdf.Ln(8)

	// Company
	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(90, 90, 90)
	pdf.SetX(25)
	pdf.Cell(40, 5, "Empresa:")
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(50, 50, 50)
	pdf.Cell(0, 5, overview.CompanyProfile.Name)
	pdf.Ln(6)

	// Position
	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(90, 90, 90)
	pdf.SetX(25)
	pdf.Cell(40, 5, "Posicion:")
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(50, 50, 50)
	pdf.Cell(0, 5, overview.JobPosting.Title)
	pdf.Ln(6)

	// Status with colored badge
	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(90, 90, 90)
	pdf.SetX(25)
	pdf.Cell(40, 5, "Estado:")

	// Status badge color
	statusColor := map[string][3]int{
		"ACTIVE":    {46, 204, 113},
		"COMPLETED": {52, 152, 219},
		"SUSPENDED": {241, 196, 15},
		"CANCELLED": {231, 76, 60},
	}
	color, exists := statusColor[overview.Status]
	if !exists {
		color = [3]int{149, 165, 166}
	}

	pdf.SetFillColor(color[0], color[1], color[2])
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFont("Arial", "B", 9)
	pdf.SetX(65)
	pdf.CellFormat(35, 5, " "+overview.Status+" ", "", 0, "C", true, 0, "")

	pdf.SetY(yStart + 35)
	pdf.Ln(3)

	// SECTION 2: Intern Info
	pdf.SetDrawColor(220, 220, 220)
	pdf.SetFillColor(248, 249, 250)
	pdf.RoundedRect(15, pdf.GetY(), 180, 25, 2, "1234", "FD")

	yStart = pdf.GetY()
	pdf.SetY(yStart + 3)

	pdf.SetFont("Arial", "B", 12)
	pdf.SetTextColor(41, 128, 185)
	pdf.SetX(20)
	pdf.Cell(0, 6, "INFORMACION DEL PASANTE")
	pdf.Ln(8)

	// Name
	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(90, 90, 90)
	pdf.SetX(25)
	pdf.Cell(40, 5, "Nombre:")
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(50, 50, 50)
	pdf.Cell(0, 5, overview.JobSeeker.Name)
	pdf.Ln(6)

	// Email
	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(90, 90, 90)
	pdf.SetX(25)
	pdf.Cell(40, 5, "Email:")
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(50, 50, 50)
	pdf.Cell(0, 5, overview.JobSeeker.Email)

	pdf.SetY(yStart + 25)
	pdf.Ln(3)

	// SECTION 3: Dates - Timeline
	pdf.SetDrawColor(220, 220, 220)
	pdf.SetFillColor(248, 249, 250)
	pdf.RoundedRect(15, pdf.GetY(), 180, 25, 2, "1234", "FD")

	yStart = pdf.GetY()
	pdf.SetY(yStart + 3)

	pdf.SetFont("Arial", "B", 12)
	pdf.SetTextColor(41, 128, 185)
	pdf.SetX(20)
	pdf.Cell(0, 6, "LINEA DE TIEMPO")
	pdf.Ln(8)

	// Start Date
	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(90, 90, 90)
	pdf.SetX(25)
	pdf.Cell(20, 5, "Inicio:")
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(46, 204, 113)
	pdf.Cell(40, 5, overview.StartDate.Format("02/01/2006"))

	// Arrow
	pdf.SetX(85)
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(150, 150, 150)
	pdf.Cell(20, 5, "  =>  ")

	// End Date
	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(90, 90, 90)
	pdf.Cell(20, 5, "Termino:")
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(231, 76, 60)
	pdf.Cell(0, 5, overview.EndDate.Format("02/01/2006"))

	pdf.SetY(yStart + 25)
	pdf.Ln(5)

	// SECTION 4: Summary boxes - Dashboard
	pdf.SetFont("Arial", "B", 12)
	pdf.SetTextColor(41, 128, 185)
	pdf.SetX(20)
	pdf.Cell(0, 6, "DASHBOARD DE ACTIVIDADES")
	pdf.Ln(10)

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

	yStart = pdf.GetY()

	// Milestones Card
	pdf.SetDrawColor(52, 152, 219)
	pdf.SetFillColor(52, 152, 219)
	pdf.RoundedRect(20, yStart, 50, 30, 3, "1234", "FD")
	pdf.SetY(yStart + 5)
	pdf.SetX(20)
	pdf.SetFont("Arial", "B", 22)
	pdf.SetTextColor(255, 255, 255)
	pdf.CellFormat(50, 10, fmt.Sprintf("%d", milestoneTotalCount), "", 0, "C", false, 0, "")
	pdf.SetY(yStart + 17)
	pdf.SetX(20)
	pdf.SetFont("Arial", "B", 9)
	pdf.CellFormat(50, 5, "HITOS", "", 0, "C", false, 0, "")
	pdf.SetY(yStart + 23)
	pdf.SetX(20)
	pdf.SetFont("Arial", "", 7)
	pdf.CellFormat(50, 4, "Total Registrados", "", 0, "C", false, 0, "")

	// Issues Card
	pdf.SetDrawColor(155, 89, 182)
	pdf.SetFillColor(155, 89, 182)
	pdf.RoundedRect(75, yStart, 50, 30, 3, "1234", "FD")
	pdf.SetY(yStart + 5)
	pdf.SetX(75)
	pdf.SetFont("Arial", "B", 22)
	pdf.CellFormat(50, 10, fmt.Sprintf("%d", issueTotalCount), "", 0, "C", false, 0, "")
	pdf.SetY(yStart + 17)
	pdf.SetX(75)
	pdf.SetFont("Arial", "B", 9)
	pdf.CellFormat(50, 5, "ISSUES", "", 0, "C", false, 0, "")
	pdf.SetY(yStart + 23)
	pdf.SetX(75)
	pdf.SetFont("Arial", "", 7)
	pdf.CellFormat(50, 4, "Total Reportados", "", 0, "C", false, 0, "")

	// Requests Card
	pdf.SetDrawColor(46, 204, 113)
	pdf.SetFillColor(46, 204, 113)
	pdf.RoundedRect(130, yStart, 50, 30, 3, "1234", "FD")
	pdf.SetY(yStart + 5)
	pdf.SetX(130)
	pdf.SetFont("Arial", "B", 22)
	pdf.CellFormat(50, 10, fmt.Sprintf("%d", requestTotalCount), "", 0, "C", false, 0, "")
	pdf.SetY(yStart + 17)
	pdf.SetX(130)
	pdf.SetFont("Arial", "B", 9)
	pdf.CellFormat(50, 5, "REQUESTS", "", 0, "C", false, 0, "")
	pdf.SetY(yStart + 23)
	pdf.SetX(130)
	pdf.SetFont("Arial", "", 7)
	pdf.CellFormat(50, 4, "Total Solicitados", "", 0, "C", false, 0, "")

	pdf.SetY(yStart + 30)
	pdf.Ln(8)

	// SECTION 5: Details by status
	pdf.SetFont("Arial", "B", 12)
	pdf.SetTextColor(41, 128, 185)
	pdf.SetX(20)
	pdf.Cell(0, 6, "DESGLOSE POR ESTADO")
	pdf.Ln(8)

	if len(overview.MilestoneCounts) > 0 {
		// Milestones Table
		pdf.SetFont("Arial", "B", 10)
		pdf.SetTextColor(255, 255, 255)
		pdf.SetFillColor(52, 152, 219)
		pdf.SetX(20)
		pdf.CellFormat(90, 7, " HITOS", "1", 0, "L", true, 0, "")
		pdf.CellFormat(40, 7, "CANTIDAD", "1", 1, "C", true, 0, "")

		pdf.SetFont("Arial", "", 9)
		pdf.SetTextColor(50, 50, 50)
		alternate := false
		for _, count := range overview.MilestoneCounts {
			if alternate {
				pdf.SetFillColor(245, 245, 245)
			} else {
				pdf.SetFillColor(255, 255, 255)
			}
			alternate = !alternate

			pdf.SetX(20)
			pdf.CellFormat(90, 6, "  "+count.Status, "1", 0, "L", true, 0, "")
			pdf.SetFont("Arial", "B", 9)
			pdf.CellFormat(40, 6, fmt.Sprintf("%d", count.Count), "1", 1, "C", true, 0, "")
			pdf.SetFont("Arial", "", 9)
		}
		pdf.Ln(3)
	}

	if len(overview.IssueCounts) > 0 {
		// Issues Table
		pdf.SetFont("Arial", "B", 10)
		pdf.SetTextColor(255, 255, 255)
		pdf.SetFillColor(155, 89, 182)
		pdf.SetX(20)
		pdf.CellFormat(90, 7, " ISSUES", "1", 0, "L", true, 0, "")
		pdf.CellFormat(40, 7, "CANTIDAD", "1", 1, "C", true, 0, "")

		pdf.SetFont("Arial", "", 9)
		pdf.SetTextColor(50, 50, 50)
		alternate := false
		for _, count := range overview.IssueCounts {
			if alternate {
				pdf.SetFillColor(245, 245, 245)
			} else {
				pdf.SetFillColor(255, 255, 255)
			}
			alternate = !alternate

			pdf.SetX(20)
			pdf.CellFormat(90, 6, "  "+count.Status, "1", 0, "L", true, 0, "")
			pdf.SetFont("Arial", "B", 9)
			pdf.CellFormat(40, 6, fmt.Sprintf("%d", count.Count), "1", 1, "C", true, 0, "")
			pdf.SetFont("Arial", "", 9)
		}
		pdf.Ln(3)
	}

	if len(overview.RequestCounts) > 0 {
		// Requests Table
		pdf.SetFont("Arial", "B", 10)
		pdf.SetTextColor(255, 255, 255)
		pdf.SetFillColor(46, 204, 113)
		pdf.SetX(20)
		pdf.CellFormat(90, 7, " REQUESTS", "1", 0, "L", true, 0, "")
		pdf.CellFormat(40, 7, "CANTIDAD", "1", 1, "C", true, 0, "")

		pdf.SetFont("Arial", "", 9)
		pdf.SetTextColor(50, 50, 50)
		alternate := false
		for _, count := range overview.RequestCounts {
			if alternate {
				pdf.SetFillColor(245, 245, 245)
			} else {
				pdf.SetFillColor(255, 255, 255)
			}
			alternate = !alternate

			pdf.SetX(20)
			pdf.CellFormat(90, 6, "  "+count.Status, "1", 0, "L", true, 0, "")
			pdf.SetFont("Arial", "B", 9)
			pdf.CellFormat(40, 6, fmt.Sprintf("%d", count.Count), "1", 1, "C", true, 0, "")
			pdf.SetFont("Arial", "", 9)
		}
	}

	// Footer
	pdf.Ln(5)
	pdf.SetY(280)
	pdf.SetFont("Arial", "I", 7)
	pdf.SetTextColor(150, 150, 150)
	pdf.SetDrawColor(200, 200, 200)
	pdf.Line(15, 278, 195, 278)
	pdf.MultiCell(0, 3, "Este documento es confidencial y solo debe ser utilizado para fines autorizados.", "", "C", false)

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

	// Modern Header
	pdf.SetFillColor(41, 128, 185)
	pdf.Rect(0, 0, 297, 30, "F")

	pdf.SetFont("Arial", "B", 20)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetY(10)
	pdf.MultiCell(0, 8, "REPORTE GENERAL DE PASANTIAS", "", "C", false)

	// Stats bar
	pdf.SetY(35)
	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFillColor(52, 152, 219)
	pdf.SetX(10)
	pdf.CellFormat(60, 7, fmt.Sprintf(" Total: %d registros", total), "", 0, "L", true, 0, "")
	pdf.SetFillColor(46, 204, 113)
	pdf.SetX(75)
	pdf.CellFormat(70, 7, fmt.Sprintf(" Generado: %s", time.Now().Format("02/01/2006 15:04")), "", 0, "L", true, 0, "")

	pdf.SetTextColor(0, 0, 0)
	pdf.Ln(10)

	// Modern Table headers
	colWidths := []float64{40, 35, 25, 25, 20, 20, 20}
	headers := []string{"Empresa", "Pasante", "Inicio", "Termino", "Hitos", "Issues", "Requests"}
	headerColors := [][3]int{
		{52, 73, 94},
		{52, 73, 94},
		{46, 204, 113},
		{231, 76, 60},
		{52, 152, 219},
		{155, 89, 182},
		{46, 204, 113},
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
	for _, overview := range overviews {
		companyName := overview.CompanyProfile.Name
		if len(companyName) > 15 {
			companyName = companyName[:15] + "..."
		}

		pasanteName := overview.JobSeeker.Name
		if len(pasanteName) > 13 {
			pasanteName = pasanteName[:13] + "..."
		}

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

		if alternateRow {
			pdf.SetFillColor(248, 249, 250)
		} else {
			pdf.SetFillColor(255, 255, 255)
		}
		alternateRow = !alternateRow

		pdf.SetFont("Arial", "", 7)
		pdf.CellFormat(colWidths[0], 7, companyName, "1", 0, "L", true, 0, "")
		pdf.CellFormat(colWidths[1], 7, pasanteName, "1", 0, "L", true, 0, "")

		pdf.SetTextColor(46, 204, 113)
		pdf.CellFormat(colWidths[2], 7, overview.StartDate.Format("02/01/06"), "1", 0, "C", true, 0, "")
		pdf.SetTextColor(231, 76, 60)
		pdf.CellFormat(colWidths[3], 7, overview.EndDate.Format("02/01/06"), "1", 0, "C", true, 0, "")

		pdf.SetFont("Arial", "B", 7)
		pdf.SetTextColor(52, 152, 219)
		pdf.CellFormat(colWidths[4], 7, fmt.Sprintf("%d", milestoneTotalCount), "1", 0, "C", true, 0, "")
		pdf.SetTextColor(155, 89, 182)
		pdf.CellFormat(colWidths[5], 7, fmt.Sprintf("%d", issueTotalCount), "1", 0, "C", true, 0, "")
		pdf.SetTextColor(46, 204, 113)
		pdf.CellFormat(colWidths[6], 7, fmt.Sprintf("%d", requestTotalCount), "1", 1, "C", true, 0, "")

		pdf.SetTextColor(50, 50, 50)

		if pdf.GetY() > 185 {
			pdf.AddPage()

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
	pdf.MultiCell(0, 3, fmt.Sprintf("Documento confidencial - Pagina 1 - Total de registros: %d", total), "", "C", false)

	err := pdf.Output(buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// GenerateDetailedPDF generates a comprehensive PDF report for a single internship with full details
func (rs *reportService) GenerateDetailedPDF(intership *model.Intership) ([]byte, error) {
	buf := new(bytes.Buffer)
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetMargins(15, 15, 15)

	// Modern Header
	pdf.SetFillColor(41, 128, 185)
	pdf.Rect(0, 0, 210, 35, "F")

	pdf.SetFont("Arial", "B", 22)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetY(12)
	pdf.MultiCell(0, 8, "REPORTE DETALLADO DE PASANTIA", "", "C", false)

	pdf.SetFont("Arial", "", 10)
	pdf.SetY(22)
	pdf.MultiCell(0, 5, "Documento Confidencial", "", "C", false)

	// Info bar
	pdf.SetY(38)
	pdf.SetFont("Arial", "I", 8)
	pdf.SetTextColor(120, 120, 120)
	pdf.MultiCell(0, 4, fmt.Sprintf("Generado: %s | ID: %s", time.Now().Format("02/01/2006 15:04"), intership.ID), "", "R", false)

	pdf.SetTextColor(0, 0, 0)
	pdf.Ln(5)

	// Basic Information Section
	pdf.SetDrawColor(220, 220, 220)
	pdf.SetFillColor(248, 249, 250)
	pdf.RoundedRect(15, pdf.GetY(), 180, 45, 2, "1234", "FD")

	yStart := pdf.GetY()
	pdf.SetY(yStart + 3)

	pdf.SetFont("Arial", "B", 12)
	pdf.SetTextColor(41, 128, 185)
	pdf.SetX(20)
	pdf.Cell(0, 6, "INFORMACION BASICA")
	pdf.Ln(8)

	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(90, 90, 90)
	pdf.SetX(25)
	pdf.Cell(40, 5, "Empresa:")
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(50, 50, 50)
	pdf.Cell(0, 5, intership.CompanyProfile.CompanyName)
	pdf.Ln(6)

	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(90, 90, 90)
	pdf.SetX(25)
	pdf.Cell(40, 5, "Posicion:")
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(50, 50, 50)
	pdf.Cell(0, 5, intership.JobPosting.Title)
	pdf.Ln(6)

	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(90, 90, 90)
	pdf.SetX(25)
	pdf.Cell(40, 5, "Pasante:")
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(50, 50, 50)
	pdf.Cell(0, 5, intership.JobSeekerProfile.Name)
	pdf.Ln(6)

	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(90, 90, 90)
	pdf.SetX(25)
	pdf.Cell(40, 5, "Estado:")

	statusColor := map[string][3]int{
		"ACTIVE":    {46, 204, 113},
		"COMPLETED": {52, 152, 219},
		"SUSPENDED": {241, 196, 15},
		"CANCELLED": {231, 76, 60},
	}
	color, exists := statusColor[string(intership.Status)]
	if !exists {
		color = [3]int{149, 165, 166}
	}

	pdf.SetFillColor(color[0], color[1], color[2])
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFont("Arial", "B", 9)
	pdf.CellFormat(35, 5, " "+string(intership.Status)+" ", "", 0, "C", true, 0, "")

	pdf.SetY(yStart + 45)
	pdf.Ln(3)

	// Timeline Section
	pdf.SetDrawColor(220, 220, 220)
	pdf.SetFillColor(248, 249, 250)
	pdf.RoundedRect(15, pdf.GetY(), 180, 20, 2, "1234", "FD")

	yStart = pdf.GetY()
	pdf.SetY(yStart + 3)

	pdf.SetFont("Arial", "B", 12)
	pdf.SetTextColor(41, 128, 185)
	pdf.SetX(20)
	pdf.Cell(0, 6, "PERIODO")
	pdf.Ln(8)

	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(90, 90, 90)
	pdf.SetX(25)
	pdf.Cell(20, 5, "Inicio:")
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(46, 204, 113)
	pdf.Cell(40, 5, intership.StartDate.Format("02/01/2006"))

	pdf.SetX(85)
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(150, 150, 150)
	pdf.Cell(20, 5, "  =>  ")

	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(90, 90, 90)
	pdf.Cell(20, 5, "Termino:")
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(231, 76, 60)
	pdf.Cell(0, 5, intership.EndDate.Format("02/01/2006"))

	pdf.SetY(yStart + 20)
	pdf.Ln(5)

	// DASHBOARD DE ACTIVIDADES - RESUMEN
	pdf.SetFont("Arial", "B", 12)
	pdf.SetTextColor(41, 128, 185)
	pdf.SetX(20)
	pdf.Cell(0, 6, "DASHBOARD DE ACTIVIDADES")
	pdf.Ln(10)

	// Contar totales
	milestoneTotalCount := len(intership.Milestones)
	issueTotalCount := 0
	requestTotalCount := 0

	for _, milestone := range intership.Milestones {
		issueTotalCount += len(milestone.FollowupIssues)
		for _, issue := range milestone.FollowupIssues {
			requestTotalCount += len(issue.Requests)
		}
	}

	yStart = pdf.GetY()

	// Milestones Card
	pdf.SetDrawColor(52, 152, 219)
	pdf.SetFillColor(52, 152, 219)
	pdf.RoundedRect(20, yStart, 50, 30, 3, "1234", "FD")
	pdf.SetY(yStart + 5)
	pdf.SetX(20)
	pdf.SetFont("Arial", "B", 22)
	pdf.SetTextColor(255, 255, 255)
	pdf.CellFormat(50, 10, fmt.Sprintf("%d", milestoneTotalCount), "", 0, "C", false, 0, "")
	pdf.SetY(yStart + 17)
	pdf.SetX(20)
	pdf.SetFont("Arial", "B", 9)
	pdf.CellFormat(50, 5, "HITOS", "", 0, "C", false, 0, "")
	pdf.SetY(yStart + 23)
	pdf.SetX(20)
	pdf.SetFont("Arial", "", 7)
	pdf.CellFormat(50, 4, "Total Registrados", "", 0, "C", false, 0, "")

	// Issues Card
	pdf.SetDrawColor(155, 89, 182)
	pdf.SetFillColor(155, 89, 182)
	pdf.RoundedRect(75, yStart, 50, 30, 3, "1234", "FD")
	pdf.SetY(yStart + 5)
	pdf.SetX(75)
	pdf.SetFont("Arial", "B", 22)
	pdf.CellFormat(50, 10, fmt.Sprintf("%d", issueTotalCount), "", 0, "C", false, 0, "")
	pdf.SetY(yStart + 17)
	pdf.SetX(75)
	pdf.SetFont("Arial", "B", 9)
	pdf.CellFormat(50, 5, "ISSUES", "", 0, "C", false, 0, "")
	pdf.SetY(yStart + 23)
	pdf.SetX(75)
	pdf.SetFont("Arial", "", 7)
	pdf.CellFormat(50, 4, "Total Reportados", "", 0, "C", false, 0, "")

	// Requests Card
	pdf.SetDrawColor(46, 204, 113)
	pdf.SetFillColor(46, 204, 113)
	pdf.RoundedRect(130, yStart, 50, 30, 3, "1234", "FD")
	pdf.SetY(yStart + 5)
	pdf.SetX(130)
	pdf.SetFont("Arial", "B", 22)
	pdf.CellFormat(50, 10, fmt.Sprintf("%d", requestTotalCount), "", 0, "C", false, 0, "")
	pdf.SetY(yStart + 17)
	pdf.SetX(130)
	pdf.SetFont("Arial", "B", 9)
	pdf.CellFormat(50, 5, "REQUESTS", "", 0, "C", false, 0, "")
	pdf.SetY(yStart + 23)
	pdf.SetX(130)
	pdf.SetFont("Arial", "", 7)
	pdf.CellFormat(50, 4, "Total Solicitados", "", 0, "C", false, 0, "")

	pdf.SetY(yStart + 30)
	pdf.Ln(10)

	// Milestones Section - DETALLES COMPLETOS
	pdf.SetFont("Arial", "B", 12)
	pdf.SetTextColor(41, 128, 185)
	pdf.SetX(20)
	pdf.Cell(0, 6, "HITOS (MILESTONES) - DETALLE")
	pdf.Ln(8)

	if len(intership.Milestones) > 0 {
		for i, milestone := range intership.Milestones {
			// Check if we need a new page
			if pdf.GetY() > 240 {
				pdf.AddPage()
				pdf.Ln(10)
			}

			// Milestone Card
			pdf.SetDrawColor(52, 152, 219)
			pdf.SetFillColor(248, 249, 250)
			pdf.RoundedRect(20, pdf.GetY(), 160, 30, 2, "1234", "FD")

			yStart = pdf.GetY()
			pdf.SetY(yStart + 3)

			pdf.SetFont("Arial", "B", 10)
			pdf.SetTextColor(52, 152, 219)
			pdf.SetX(25)
			pdf.Cell(0, 5, fmt.Sprintf("Hito #%d: %s", i+1, milestone.Title))
			pdf.Ln(6)

			pdf.SetFont("Arial", "", 8)
			pdf.SetTextColor(60, 60, 60)
			pdf.SetX(25)
			description := milestone.Description
			if len(description) > 77 {
				description = description[:77] + "..."
			}
			pdf.Cell(0, 4, description)
			pdf.Ln(5)

			pdf.SetFont("Arial", "B", 8)
			pdf.SetTextColor(90, 90, 90)
			pdf.SetX(25)
			pdf.Cell(30, 4, "Estado:")
			pdf.SetFont("Arial", "", 8)
			pdf.SetTextColor(50, 50, 50)
			pdf.Cell(40, 4, string(milestone.Status))

			pdf.SetFont("Arial", "B", 8)
			pdf.SetTextColor(90, 90, 90)
			pdf.Cell(30, 4, "Fecha Limite:")
			pdf.SetFont("Arial", "", 8)
			pdf.SetTextColor(50, 50, 50)
			pdf.Cell(0, 4, milestone.DueDate.Format("02/01/2006"))

			pdf.SetY(yStart + 30)
			pdf.Ln(3)
		}
	} else {
		pdf.SetFont("Arial", "I", 9)
		pdf.SetTextColor(150, 150, 150)
		pdf.SetX(25)
		pdf.Cell(0, 5, "No hay hitos registrados")
		pdf.Ln(8)
	}

	// Issues Section - DETALLES COMPLETOS
	pdf.SetFont("Arial", "B", 12)
	pdf.SetTextColor(155, 89, 182)
	pdf.SetX(20)
	pdf.Cell(0, 6, "PROBLEMAS (ISSUES) - DETALLE")
	pdf.Ln(8)

	issuesFound := false
	for _, milestone := range intership.Milestones {
		if len(milestone.FollowupIssues) > 0 {
			issuesFound = true
			for i, issue := range milestone.FollowupIssues {
				if pdf.GetY() > 240 {
					pdf.AddPage()
					pdf.Ln(10)
				}

				pdf.SetDrawColor(155, 89, 182)
				pdf.SetFillColor(248, 249, 250)
				pdf.RoundedRect(20, pdf.GetY(), 160, 30, 2, "1234", "FD")

				yStart = pdf.GetY()
				pdf.SetY(yStart + 3)

				pdf.SetFont("Arial", "B", 10)
				pdf.SetTextColor(155, 89, 182)
				pdf.SetX(25)
				pdf.Cell(0, 5, fmt.Sprintf("Issue #%d: %s", i+1, issue.Title))
				pdf.Ln(6)

				pdf.SetFont("Arial", "", 8)
				pdf.SetTextColor(60, 60, 60)
				pdf.SetX(25)
				description := issue.Description
				if len(description) > 77 {
					description = description[:77] + "..."
				}
				pdf.Cell(0, 4, description)
				pdf.Ln(5)
				pdf.SetFont("Arial", "B", 8)
				pdf.SetTextColor(90, 90, 90)
				pdf.SetX(25)
				pdf.Cell(30, 4, "Estado:")
				pdf.SetFont("Arial", "", 8)
				pdf.SetTextColor(50, 50, 50)
				pdf.Cell(40, 4, string(issue.Status))

				pdf.SetFont("Arial", "B", 8)
				pdf.SetTextColor(90, 90, 90)
				pdf.Cell(30, 4, "Fecha Limite:")
				pdf.SetFont("Arial", "", 8)
				pdf.SetTextColor(50, 50, 50)
				pdf.Cell(0, 4, issue.DueDate.Format("02/01/2006"))

				pdf.SetY(yStart + 30)
				pdf.Ln(3)
			}
		}
	}

	if !issuesFound {
		pdf.SetFont("Arial", "I", 9)
		pdf.SetTextColor(150, 150, 150)
		pdf.SetX(25)
		pdf.Cell(0, 5, "No hay issues registrados")
		pdf.Ln(8)
	}

	// Requests Section - DETALLES COMPLETOS
	pdf.SetFont("Arial", "B", 12)
	pdf.SetTextColor(46, 204, 113)
	pdf.SetX(20)
	pdf.Cell(0, 6, "SOLICITUDES (REQUESTS) - DETALLE")
	pdf.Ln(8)

	requestsFound := false
	for _, milestone := range intership.Milestones {
		for _, issue := range milestone.FollowupIssues {
			if len(issue.Requests) > 0 {
				requestsFound = true
				for i, request := range issue.Requests {
					if pdf.GetY() > 240 {
						pdf.AddPage()
						pdf.Ln(10)
					}

					pdf.SetDrawColor(46, 204, 113)
					pdf.SetFillColor(248, 249, 250)
					pdf.RoundedRect(20, pdf.GetY(), 160, 30, 2, "1234", "FD")

					yStart = pdf.GetY()
					pdf.SetY(yStart + 3)

					pdf.SetFont("Arial", "B", 10)
					pdf.SetTextColor(46, 204, 113)
					pdf.SetX(25)
					pdf.Cell(0, 5, fmt.Sprintf("Request #%d: %s", i+1, request.Title))
					pdf.Ln(6)

					pdf.SetFont("Arial", "", 8)
					pdf.SetTextColor(60, 60, 60)
					pdf.SetX(25)
					description := request.Description
					if len(description) > 80 {
						description = description[:77] + "..."
					}
					pdf.Cell(0, 4, description)
					pdf.Ln(5)

					pdf.SetFont("Arial", "B", 8)
					pdf.SetTextColor(90, 90, 90)
					pdf.SetX(25)
					pdf.Cell(30, 4, "Estado:")
					pdf.SetFont("Arial", "", 8)
					pdf.SetTextColor(50, 50, 50)
					pdf.Cell(40, 4, string(request.Status))

					pdf.SetFont("Arial", "B", 8)
					pdf.SetTextColor(90, 90, 90)
					pdf.Cell(30, 4, "Comentario:")
					pdf.SetFont("Arial", "", 8)
					pdf.SetTextColor(50, 50, 50)
					if request.CompanyComment != "" {
						comment := request.CompanyComment
						if len(comment) > 50 {
							comment = comment[:50] + "..."
						}
						pdf.Cell(0, 4, comment)
					} else {
						pdf.Cell(0, 4, "Sin comentarios")
					}

					pdf.SetY(yStart + 30)
					pdf.Ln(3)
				}
			}
		}
	}

	if !requestsFound {
		pdf.SetFont("Arial", "I", 9)
		pdf.SetTextColor(150, 150, 150)
		pdf.SetX(25)
		pdf.Cell(0, 5, "No hay requests registradas")
		pdf.Ln(8)
	}

	// Footer
	pdf.SetY(280)
	pdf.SetFont("Arial", "I", 7)
	pdf.SetTextColor(150, 150, 150)
	pdf.SetDrawColor(200, 200, 200)
	pdf.Line(15, 278, 195, 278)
	pdf.MultiCell(0, 3, "Este documento es confidencial y solo debe ser utilizado para fines autorizados.", "", "C", false)

	err := pdf.Output(buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
