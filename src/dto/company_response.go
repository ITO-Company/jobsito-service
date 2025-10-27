package dto

import (
	"github.com/ito-company/jobsito-service/src/model"
	"github.com/jinzhu/copier"
)

type CompanyResponse struct {
	ID          string               `json:"id"`
	CompanyName string               `json:"company_name"`
	Email       string               `json:"email"`
	Description string               `json:"description"`
	Website     string               `json:"website"`
	Phone       string               `json:"phone"`
	Address     string               `json:"address"`
	Industry    string               `json:"industry"`
	CompanySize string               `json:"company_size"`
	LogoUrl     string               `json:"logo_url"`
	IsVerified  bool                 `json:"is_verified"`
	JobPosting  []JobPostingResponse `json:"job_postings,omitempty"`
}

func CompanyToDto(m *model.CompanyProfile) CompanyResponse {
	var dto CompanyResponse
	copier.Copy(&dto, m)

	if len(m.JobPostings) > 0 {
		dto.JobPosting = JobPostingToListDto(m.JobPostings)
	}

	return dto
}

func CompanyToListDto(m []model.CompanyProfile) []CompanyResponse {
	out := make([]CompanyResponse, len(m))
	for i := range m {
		out[i] = CompanyToDto(&m[i])
	}
	return out
}
