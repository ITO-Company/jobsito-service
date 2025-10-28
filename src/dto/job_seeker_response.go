package dto

import (
	"github.com/ito-company/jobsito-service/src/model"
	"github.com/jinzhu/copier"
)

type JobSeekerTagResponse struct {
	ID               string            `json:"id"`
	ProficiencyLevel string            `json:"proficiency_level"`
	GlobalTag        GlobalTagResponse `json:"global_tag"`
}

type JobSeekerResponse struct {
	ID                string                 `json:"id"`
	Name              string                 `json:"name"`
	Email             string                 `json:"email"`
	Bio               string                 `json:"bio"`
	Phone             string                 `json:"phone"`
	Location          string                 `json:"location"`
	CvUrl             string                 `json:"cv_url"`
	PortfolioUrl      string                 `json:"portfolio_url"`
	ExpectedSalaryMin string                 `json:"expected_salary_min"`
	ExpectedSalaryMax string                 `json:"expected_salary_max"`
	Availability      string                 `json:"availability"`
	Skills            string                 `json:"skills"`
	Experience        string                 `json:"experience"`
	IsActive          bool                   `json:"is_active"`
	Tags              []JobSeekerTagResponse `json:"tags"`
}

func JobSeekerToDto(m *model.JobSeekerProfile) JobSeekerResponse {
	var dto JobSeekerResponse
	copier.Copy(&dto, m)

	tags := make([]JobSeekerTagResponse, 0, len(m.JobSeekerTags))
	for _, tag := range m.JobSeekerTags {
		tags = append(tags, JobSeekerTagResponse{
			ID:               tag.ID.String(),
			ProficiencyLevel: tag.ProficiencyLevel,
			GlobalTag: GlobalTagResponse{
				ID:       tag.GlobalTag.ID.String(),
				Name:     tag.GlobalTag.Name,
				Category: tag.GlobalTag.Category,
				Color:    tag.GlobalTag.Color,
			},
		})
	}
	dto.Tags = tags

	return dto
}

func JobSeekerToListDto(m []model.JobSeekerProfile) []JobSeekerResponse {
	out := make([]JobSeekerResponse, len(m))
	for i := range m {
		out[i] = JobSeekerToDto(&m[i])
	}
	return out
}
