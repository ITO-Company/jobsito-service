package kpi

import (
	"time"
)

// MilestoneKPIDto holds all KPI metrics for a company's milestones
type MilestoneKPIDto struct {
	CompanyID                 string  `json:"company_id"`
	TotalMilestones           int64   `json:"total_milestones"`
	CompletedMilestones       int64   `json:"completed_milestones"`
	PendingMilestones         int64   `json:"pending_milestones"`
	ActiveMilestones          int64   `json:"active_milestones"`
	OverdueMilestones         int64   `json:"overdue_milestones"`
	CompletionRate            float64 `json:"completion_rate"`
	AverageCompletionTimeDays float64 `json:"average_completion_time_days"`
	OverduePercentage         float64 `json:"overdue_percentage"`
}

// MilestoneKPIService handles business logic for milestone KPIs
type MilestoneKPIService interface {
	GetCompanyMilestoneKPI(companyID string) (*MilestoneKPIDto, error)
	GetIntershipMilestoneKPI(intershipID string) (*MilestoneKPIDto, error)
}

type MilestoneKPIServiceImpl struct {
	repo MilestoneKPIRepo
}

func NewMilestoneKPIService(repo MilestoneKPIRepo) MilestoneKPIService {
	return &MilestoneKPIServiceImpl{repo: repo}
}

func (s *MilestoneKPIServiceImpl) GetCompanyMilestoneKPI(companyID string) (*MilestoneKPIDto, error) {
	completed, err := s.repo.CountMilestonesByCompanyAndStatus(companyID, "approved")
	if err != nil {
		return nil, err
	}

	pending, err := s.repo.CountMilestonesByCompanyAndStatus(companyID, "pending")
	if err != nil {
		return nil, err
	}

	active, err := s.repo.CountMilestonesByCompanyAndStatus(companyID, "active")
	if err != nil {
		return nil, err
	}

	overdue, err := s.repo.CountOverdueMilestonesByCompany(companyID)
	if err != nil {
		return nil, err
	}

	completionTimes, err := s.repo.FindMilestoneCompletionTimeByCompany(companyID)
	if err != nil {
		return nil, err
	}

	total := completed + pending + active + overdue

	completionRate := 0.0
	if total > 0 {
		completionRate = float64(completed) / float64(total) * 100
	}

	overduePercentage := 0.0
	if total > 0 {
		overduePercentage = float64(overdue) / float64(total) * 100
	}

	avgCompletionTime := 0.0
	if len(completionTimes) > 0 {
		var totalTime time.Duration
		for _, d := range completionTimes {
			totalTime += d
		}
		avgCompletionTime = totalTime.Hours() / float64(len(completionTimes)) / 24 // Convert to days
	}

	return &MilestoneKPIDto{
		CompanyID:                 companyID,
		TotalMilestones:           total,
		CompletedMilestones:       completed,
		PendingMilestones:         pending,
		ActiveMilestones:          active,
		OverdueMilestones:         overdue,
		CompletionRate:            completionRate,
		AverageCompletionTimeDays: avgCompletionTime,
		OverduePercentage:         overduePercentage,
	}, nil
}

func (s *MilestoneKPIServiceImpl) GetIntershipMilestoneKPI(intershipID string) (*MilestoneKPIDto, error) {
	completed, err := s.repo.CountMilestonesByIntershipAndStatus(intershipID, "approved")
	if err != nil {
		return nil, err
	}

	pending, err := s.repo.CountMilestonesByIntershipAndStatus(intershipID, "pending")
	if err != nil {
		return nil, err
	}

	active, err := s.repo.CountMilestonesByIntershipAndStatus(intershipID, "active")
	if err != nil {
		return nil, err
	}

	overdue, err := s.repo.CountOverdueMilestonesByIntership(intershipID)
	if err != nil {
		return nil, err
	}

	completionTimes, err := s.repo.FindMilestoneCompletionTimeByIntership(intershipID)
	if err != nil {
		return nil, err
	}

	total := completed + pending + active + overdue

	completionRate := 0.0
	if total > 0 {
		completionRate = float64(completed) / float64(total) * 100
	}

	overduePercentage := 0.0
	if total > 0 {
		overduePercentage = float64(overdue) / float64(total) * 100
	}

	avgCompletionTime := 0.0
	if len(completionTimes) > 0 {
		var totalTime time.Duration
		for _, d := range completionTimes {
			totalTime += d
		}
		avgCompletionTime = totalTime.Hours() / float64(len(completionTimes)) / 24 // Convert to days
	}

	return &MilestoneKPIDto{
		CompanyID:                 intershipID,
		TotalMilestones:           total,
		CompletedMilestones:       completed,
		PendingMilestones:         pending,
		ActiveMilestones:          active,
		OverdueMilestones:         overdue,
		CompletionRate:            completionRate,
		AverageCompletionTimeDays: avgCompletionTime,
		OverduePercentage:         overduePercentage,
	}, nil
}
