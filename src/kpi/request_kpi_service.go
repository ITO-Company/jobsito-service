package kpi

import (
	"time"
)

// RequestKPIDto holds all KPI metrics for a company's requests
type RequestKPIDto struct {
	ContextID                 string  `json:"context_id"` // Company ID or Intership ID
	TotalRequests             int64   `json:"total_requests"`
	ApprovedRequests          int64   `json:"approved_requests"`
	RejectedRequests          int64   `json:"rejected_requests"`
	PendingRequests           int64   `json:"pending_requests"`
	IssuesWithPendingRequests int64   `json:"issues_with_pending_requests"`
	ApprovalRate              float64 `json:"approval_rate"`
	RejectionRate             float64 `json:"rejection_rate"`
	AverageReviewTimeDays     float64 `json:"average_review_time_days"`
	PendingPercentage         float64 `json:"pending_percentage"`
}

// RequestKPIService handles business logic for request KPIs
type RequestKPIService interface {
	GetCompanyRequestKPI(companyID string) (*RequestKPIDto, error)
	GetIntershipRequestKPI(intershipID string) (*RequestKPIDto, error)
}

type RequestKPIServiceImpl struct {
	repo RequestKPIRepo
}

func NewRequestKPIService(repo RequestKPIRepo) RequestKPIService {
	return &RequestKPIServiceImpl{repo: repo}
}

func (s *RequestKPIServiceImpl) GetCompanyRequestKPI(companyID string) (*RequestKPIDto, error) {
	approved, err := s.repo.CountRequestsByCompanyAndStatus(companyID, "approved")
	if err != nil {
		return nil, err
	}

	rejected, err := s.repo.CountRequestsByCompanyAndStatus(companyID, "rejected")
	if err != nil {
		return nil, err
	}

	pending, err := s.repo.CountPendingRequestsByCompany(companyID)
	if err != nil {
		return nil, err
	}

	issuesWithPendingRequests, err := s.repo.CountIssuesWithPendingRequestsByCompany(companyID)
	if err != nil {
		return nil, err
	}

	reviewTimes, err := s.repo.FindRequestReviewTimeByCompany(companyID)
	if err != nil {
		return nil, err
	}

	total := approved + rejected + pending

	approvalRate := 0.0
	if total > 0 {
		approvalRate = float64(approved) / float64(total) * 100
	}

	rejectionRate := 0.0
	if total > 0 {
		rejectionRate = float64(rejected) / float64(total) * 100
	}

	pendingPercentage := 0.0
	if total > 0 {
		pendingPercentage = float64(pending) / float64(total) * 100
	}

	avgReviewTime := 0.0
	if len(reviewTimes) > 0 {
		var totalTime time.Duration
		for _, d := range reviewTimes {
			totalTime += d
		}
		avgReviewTime = totalTime.Hours() / float64(len(reviewTimes)) / 24 // Convert to days
	}

	return &RequestKPIDto{
		ContextID:                 companyID,
		TotalRequests:             total,
		ApprovedRequests:          approved,
		RejectedRequests:          rejected,
		PendingRequests:           pending,
		IssuesWithPendingRequests: issuesWithPendingRequests,
		ApprovalRate:              approvalRate,
		RejectionRate:             rejectionRate,
		AverageReviewTimeDays:     avgReviewTime,
		PendingPercentage:         pendingPercentage,
	}, nil
}

func (s *RequestKPIServiceImpl) GetIntershipRequestKPI(intershipID string) (*RequestKPIDto, error) {
	approved, err := s.repo.CountRequestsByIntershipAndStatus(intershipID, "approved")
	if err != nil {
		return nil, err
	}

	rejected, err := s.repo.CountRequestsByIntershipAndStatus(intershipID, "rejected")
	if err != nil {
		return nil, err
	}

	pending, err := s.repo.CountPendingRequestsByIntership(intershipID)
	if err != nil {
		return nil, err
	}

	issuesWithPendingRequests, err := s.repo.CountIssuesWithPendingRequestsByIntership(intershipID)
	if err != nil {
		return nil, err
	}

	reviewTimes, err := s.repo.FindRequestReviewTimeByIntership(intershipID)
	if err != nil {
		return nil, err
	}

	total := approved + rejected + pending

	approvalRate := 0.0
	if total > 0 {
		approvalRate = float64(approved) / float64(total) * 100
	}

	rejectionRate := 0.0
	if total > 0 {
		rejectionRate = float64(rejected) / float64(total) * 100
	}

	pendingPercentage := 0.0
	if total > 0 {
		pendingPercentage = float64(pending) / float64(total) * 100
	}

	avgReviewTime := 0.0
	if len(reviewTimes) > 0 {
		var totalTime time.Duration
		for _, d := range reviewTimes {
			totalTime += d
		}
		avgReviewTime = totalTime.Hours() / float64(len(reviewTimes)) / 24 // Convert to days
	}

	return &RequestKPIDto{
		ContextID:                 intershipID,
		TotalRequests:             total,
		ApprovedRequests:          approved,
		RejectedRequests:          rejected,
		PendingRequests:           pending,
		IssuesWithPendingRequests: issuesWithPendingRequests,
		ApprovalRate:              approvalRate,
		RejectionRate:             rejectionRate,
		AverageReviewTimeDays:     avgReviewTime,
		PendingPercentage:         pendingPercentage,
	}, nil
}
