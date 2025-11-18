package kpi

import (
	"time"
)

// IssueKPIDto holds all KPI metrics for a company's issues
type IssueKPIDto struct {
	ContextID                 string  `json:"context_id"` // Company ID or Intership ID
	TotalIssues               int64   `json:"total_issues"`
	ResolvedIssues            int64   `json:"resolved_issues"`
	PendingIssues             int64   `json:"pending_issues"`
	ActiveIssues              int64   `json:"active_issues"`
	OverdueIssues             int64   `json:"overdue_issues"`
	IssuesWithRequests        int64   `json:"issues_with_requests"`
	ResolutionRate            float64 `json:"resolution_rate"`
	AverageResolutionTimeDays float64 `json:"average_resolution_time_days"`
	OverduePercentage         float64 `json:"overdue_percentage"`
}

// IssueKPIService handles business logic for issue KPIs
type IssueKPIService interface {
	GetCompanyIssueKPI(companyID string) (*IssueKPIDto, error)
	GetIntershipIssueKPI(intershipID string) (*IssueKPIDto, error)
}

type IssueKPIServiceImpl struct {
	repo IssueKPIRepo
}

func NewIssueKPIService(repo IssueKPIRepo) IssueKPIService {
	return &IssueKPIServiceImpl{repo: repo}
}

func (s *IssueKPIServiceImpl) GetCompanyIssueKPI(companyID string) (*IssueKPIDto, error) {
	resolved, err := s.repo.CountIssuesByCompanyAndStatus(companyID, "approved")
	if err != nil {
		return nil, err
	}

	pending, err := s.repo.CountIssuesByCompanyAndStatus(companyID, "pending")
	if err != nil {
		return nil, err
	}

	active, err := s.repo.CountIssuesByCompanyAndStatus(companyID, "active")
	if err != nil {
		return nil, err
	}

	overdue, err := s.repo.CountOverdueIssuesByCompany(companyID)
	if err != nil {
		return nil, err
	}

	issuesWithRequests, err := s.repo.CountIssuesWithRequestsByCompany(companyID)
	if err != nil {
		return nil, err
	}

	resolutionTimes, err := s.repo.FindIssueCompletionTimeByCompany(companyID)
	if err != nil {
		return nil, err
	}

	total := resolved + pending + active + overdue

	resolutionRate := 0.0
	if total > 0 {
		resolutionRate = float64(resolved) / float64(total) * 100
	}

	overduePercentage := 0.0
	if total > 0 {
		overduePercentage = float64(overdue) / float64(total) * 100
	}

	avgResolutionTime := 0.0
	if len(resolutionTimes) > 0 {
		var totalTime time.Duration
		for _, d := range resolutionTimes {
			totalTime += d
		}
		avgResolutionTime = totalTime.Hours() / float64(len(resolutionTimes)) / 24 // Convert to days
	}

	return &IssueKPIDto{
		ContextID:                 companyID,
		TotalIssues:               total,
		ResolvedIssues:            resolved,
		PendingIssues:             pending,
		ActiveIssues:              active,
		OverdueIssues:             overdue,
		IssuesWithRequests:        issuesWithRequests,
		ResolutionRate:            resolutionRate,
		AverageResolutionTimeDays: avgResolutionTime,
		OverduePercentage:         overduePercentage,
	}, nil
}

func (s *IssueKPIServiceImpl) GetIntershipIssueKPI(intershipID string) (*IssueKPIDto, error) {
	resolved, err := s.repo.CountIssuesByIntershipAndStatus(intershipID, "approved")
	if err != nil {
		return nil, err
	}

	pending, err := s.repo.CountIssuesByIntershipAndStatus(intershipID, "pending")
	if err != nil {
		return nil, err
	}

	active, err := s.repo.CountIssuesByIntershipAndStatus(intershipID, "active")
	if err != nil {
		return nil, err
	}

	overdue, err := s.repo.CountOverdueIssuesByIntership(intershipID)
	if err != nil {
		return nil, err
	}

	issuesWithRequests, err := s.repo.CountIssuesWithRequestsByIntership(intershipID)
	if err != nil {
		return nil, err
	}

	resolutionTimes, err := s.repo.FindIssueCompletionTimeByIntership(intershipID)
	if err != nil {
		return nil, err
	}

	total := resolved + pending + active + overdue

	resolutionRate := 0.0
	if total > 0 {
		resolutionRate = float64(resolved) / float64(total) * 100
	}

	overduePercentage := 0.0
	if total > 0 {
		overduePercentage = float64(overdue) / float64(total) * 100
	}

	avgResolutionTime := 0.0
	if len(resolutionTimes) > 0 {
		var totalTime time.Duration
		for _, d := range resolutionTimes {
			totalTime += d
		}
		avgResolutionTime = totalTime.Hours() / float64(len(resolutionTimes)) / 24 // Convert to days
	}

	return &IssueKPIDto{
		ContextID:                 intershipID,
		TotalIssues:               total,
		ResolvedIssues:            resolved,
		PendingIssues:             pending,
		ActiveIssues:              active,
		OverdueIssues:             overdue,
		IssuesWithRequests:        issuesWithRequests,
		ResolutionRate:            resolutionRate,
		AverageResolutionTimeDays: avgResolutionTime,
		OverduePercentage:         overduePercentage,
	}, nil
}
