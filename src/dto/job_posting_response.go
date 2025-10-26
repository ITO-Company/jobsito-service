package dto

import (
	"time"

	"github.com/ito-company/jobsito-service/src/model"
	"github.com/jinzhu/copier"
)

type JobPostingResponse struct {
	ID              string              `json:"id"`
	Title           string              `json:"title"`
	Description     string              `json:"description"`
	Requirement     string              `json:"requirement"`
	SalaryMin       string              `json:"salary_min"`
	SalaryMax       string              `json:"salary_max"`
	WorkType        string              `json:"work_type"`
	ExperienceLevel string              `json:"experience_level"`
	Location        string              `json:"location"`
	IsRemote        string              `json:"is_remote"`
	IsHibrid        string              `json:"is_hibrid"`
	ContractType    string              `json:"contract_type"`
	Benefit         string              `json:"benefit"`
	Status          string              `json:"status"`
	ExpiresAt       time.Time           `json:"expires_at"`
	Tags            []GlobalTagResponse `json:"tags,omitempty"`
}

func JobPostingToDto(m *model.JobPosting) JobPostingResponse {
	var dto JobPostingResponse
	copier.Copy(&dto, m)

	tags := make([]GlobalTagResponse, 0, len(m.JobPostingTags))
	for _, jpt := range m.JobPostingTags {
		tag := jpt.GlobalTag
		tags = append(tags, GlobalTagResponse{
			ID:       tag.ID.String(),
			Name:     tag.Name,
			Category: tag.Category,
			Color:    tag.Color,
		})
	}
	dto.Tags = tags

	return dto
}

func JobPostingToListDto(m []model.JobPosting) []JobPostingResponse {
	out := make([]JobPostingResponse, len(m))
	for i := range m {
		out[i] = JobPostingToDto(&m[i])
	}
	return out
}
