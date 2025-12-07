package dto

import (
	"time"

	"github.com/ito-company/jobsito-service/src/model"
)

type IntershipJobPostingResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	SalaryMin   string `json:"salary_min"`
	SalaryMax   string `json:"salary_max"`
}

type IntershipJobSeekerResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type IntershipCompanyResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type IntershipResponseDto struct {
	ID             string                      `json:"id"`
	StartDate      time.Time                   `json:"start_date"`
	EndDate        time.Time                   `json:"end_date"`
	Status         string                      `json:"status"`
	JobPosting     IntershipJobPostingResponse `json:"job_posting,omitempty"`
	JobSeeker      IntershipJobSeekerResponse  `json:"job_seeker,omitempty"`
	CompanyProfile IntershipCompanyResponse    `json:"company_profile,omitempty"`
	CreatedAt      time.Time                   `json:"created_at"`
	UpdatedAt      time.Time                   `json:"updated_at"`
}

func IntershipToResponse(m *model.Intership) IntershipResponseDto {
	dto := IntershipResponseDto{
		ID:        m.ID.String(),
		StartDate: m.StartDate,
		EndDate:   m.EndDate,
		Status:    string(m.Status),
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}

	if m.JobPosting.ID != [16]byte{} {
		dto.JobPosting = IntershipJobPostingResponse{
			ID:          m.JobPosting.ID.String(),
			Title:       m.JobPosting.Title,
			Description: m.JobPosting.Description,
			SalaryMin:   m.JobPosting.SalaryMin,
			SalaryMax:   m.JobPosting.SalaryMax,
		}
	}

	if m.JobSeekerProfile.ID != [16]byte{} {
		dto.JobSeeker = IntershipJobSeekerResponse{
			ID:    m.JobSeekerProfile.ID.String(),
			Name:  m.JobSeekerProfile.Name,
			Email: m.JobSeekerProfile.Email,
		}
	}

	if m.CompanyProfile.ID != [16]byte{} {
		dto.CompanyProfile = IntershipCompanyResponse{
			ID:   m.CompanyProfile.ID.String(),
			Name: m.CompanyProfile.CompanyName,
		}
	}

	return dto
}

func IntershipToListDto(m []model.Intership) []IntershipResponseDto {
	out := make([]IntershipResponseDto, len(m))
	for i := range m {
		out[i] = IntershipToResponse(&m[i])
	}
	return out
}

// Overview DTOs
type StatusCount struct {
	Status string `json:"status"`
	Count  int    `json:"count"`
}

type IntershipOverviewDto struct {
	ID             string                      `json:"id"`
	StartDate      time.Time                   `json:"start_date"`
	EndDate        time.Time                   `json:"end_date"`
	Status         string                      `json:"status"`
	JobPosting     IntershipJobPostingResponse `json:"job_posting,omitempty"`
	JobSeeker      IntershipJobSeekerResponse  `json:"job_seeker,omitempty"`
	CompanyProfile IntershipCompanyResponse    `json:"company_profile,omitempty"`
	CreatedAt      time.Time                   `json:"created_at"`
	UpdatedAt      time.Time                   `json:"updated_at"`

	// Overview information
	MilestoneCounts []StatusCount `json:"milestone_counts"`
	IssueCounts     []StatusCount `json:"issue_counts"`
	RequestCounts   []StatusCount `json:"request_counts"`
}

func IntershipToOverview(m *model.Intership) IntershipOverviewDto {
	base := IntershipToResponse(m)

	// Count milestones by status
	milestoneCounts := make(map[string]int)
	for _, milestone := range m.Milestones {
		milestoneCounts[string(milestone.Status)]++
	}

	// Count issues and requests by status
	issueCounts := make(map[string]int)
	requestCounts := make(map[string]int)

	for _, milestone := range m.Milestones {
		for _, issue := range milestone.FollowupIssues {
			issueCounts[string(issue.Status)]++
			for _, request := range issue.Requests {
				requestCounts[string(request.Status)]++
			}
		}
	}

	// Convert maps to slices
	milestoneSlice := make([]StatusCount, 0)
	for status, count := range milestoneCounts {
		milestoneSlice = append(milestoneSlice, StatusCount{Status: status, Count: count})
	}

	issueSlice := make([]StatusCount, 0)
	for status, count := range issueCounts {
		issueSlice = append(issueSlice, StatusCount{Status: status, Count: count})
	}

	requestSlice := make([]StatusCount, 0)
	for status, count := range requestCounts {
		requestSlice = append(requestSlice, StatusCount{Status: status, Count: count})
	}

	overview := IntershipOverviewDto{
		ID:              base.ID,
		StartDate:       base.StartDate,
		EndDate:         base.EndDate,
		Status:          base.Status,
		JobPosting:      base.JobPosting,
		JobSeeker:       base.JobSeeker,
		CompanyProfile:  base.CompanyProfile,
		CreatedAt:       base.CreatedAt,
		UpdatedAt:       base.UpdatedAt,
		MilestoneCounts: milestoneSlice,
		IssueCounts:     issueSlice,
		RequestCounts:   requestSlice,
	}

	return overview
}

func IntershipToOverviewList(m []model.Intership) []IntershipOverviewDto {
	out := make([]IntershipOverviewDto, len(m))
	for i := range m {
		out[i] = IntershipToOverview(&m[i])
	}
	return out
}
