package dto

import (
	"time"

	"github.com/ito-company/jobsito-service/src/model"
	"github.com/jinzhu/copier"
)

type IntershipResponseDto struct {
	ID             string          `json:"id"`
	StartDate      time.Time       `json:"start_date"`
	EndDate        time.Time       `json:"end_date"`
	Status         string          `json:"status"`
	CompanyProfile CompanyResponse `json:"company_profile,omitempty"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}

func IntershipToResponse(m *model.Intership) IntershipResponseDto {
	var dto IntershipResponseDto
	copier.Copy(&dto, m)

	dto.ID = m.ID.String()
	dto.CompanyProfile = CompanyToDto(&m.CompanyProfile)

	return dto
}

func IntershipToListDto(m []model.Intership) []IntershipResponseDto {
	out := make([]IntershipResponseDto, len(m))
	for i := range m {
		out[i] = IntershipToResponse(&m[i])
	}
	return out
}
