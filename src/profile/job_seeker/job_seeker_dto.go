package jobseeker

import (
	"github.com/ito-company/jobsito-service/src/model"
	"github.com/jinzhu/copier"
)

type JobSeekerUpdateDto struct {
	Name              *string `json:"name"`
	Bio               *string `json:"bio"`
	Phone             *string `json:"phone"`
	Location          *string `json:"location"`
	CvUrl             *string `json:"cv_url"`
	PortfolioUrl      *string `json:"portfolio_url"`
	ExpectedSalaryMin *string `json:"expected_salary_min"`
	ExpectedSalaryMax *string `json:"expected_salary_max"`
	Availability      *string `json:"availability"`
	Skills            *string `json:"skills"`
	Experience        *string `json:"experience"`
}

type JobSeekerResponse struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	Bio               string `json:"bio"`
	Phone             string `json:"phone"`
	Location          string `json:"location"`
	CvUrl             string `json:"cv_url"`
	PortfolioUrl      string `json:"portfolio_url"`
	ExpectedSalaryMin string `json:"expected_salary_min"`
	ExpectedSalaryMax string `json:"expected_salary_max"`
	Availability      string `json:"availability"`
	Skills            string `json:"skills"`
	Experience        string `json:"experience"`
	IsActive          bool   `json:"is_active"`
}

func JobSeekerToDto(m *model.JobSeekerProfile) JobSeekerResponse {
	var dto JobSeekerResponse
	copier.Copy(&dto, m)
	return dto
}

func JobSeekerToListDto(m []model.JobSeekerProfile) []JobSeekerResponse {
	out := make([]JobSeekerResponse, len(m))
	for i := range m {
		out[i] = JobSeekerToDto(&m[i])
	}
	return out
}
