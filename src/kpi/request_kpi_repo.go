package kpi

import (
	"time"

	"github.com/ito-company/jobsito-service/src/model"
	"gorm.io/gorm"
)

// RequestKPIRepo handles all queries for request KPIs
type RequestKPIRepo interface {
	// For a specific company (through issues/milestones)
	FindRequestsByCompany(companyID string) ([]model.Request, error)
	FindRequestsByIntership(intershipID string) ([]model.Request, error)

	// Count operations
	CountRequestsByCompanyAndStatus(companyID string, status string) (int64, error)
	CountRequestsByIntershipAndStatus(intershipID string, status string) (int64, error)

	// Aggregate operations
	CountPendingRequestsByCompany(companyID string) (int64, error)
	CountPendingRequestsByIntership(intershipID string) (int64, error)

	// Detail operations
	FindRequestReviewTimeByCompany(companyID string) ([]time.Duration, error)
	FindRequestReviewTimeByIntership(intershipID string) ([]time.Duration, error)

	// Issues with pending requests
	CountIssuesWithPendingRequestsByCompany(companyID string) (int64, error)
	CountIssuesWithPendingRequestsByIntership(intershipID string) (int64, error)
}

type RequestRepo struct {
	db *gorm.DB
}

func NewRequestKPIRepo(db *gorm.DB) RequestKPIRepo {
	return &RequestRepo{db: db}
}

// FindRequestsByCompany returns all requests for a company (through milestones/issues)
func (r *RequestRepo) FindRequestsByCompany(companyID string) ([]model.Request, error) {
	var requests []model.Request
	err := r.db.
		Joins("JOIN followup_issues ON followup_issues.id = requests.followup_issue_id").
		Joins("JOIN followup_milestones ON followup_milestones.id = followup_issues.followup_milestone_id").
		Where("followup_milestones.company_profile_id = ?", companyID).
		Find(&requests).Error
	return requests, err
}

// FindRequestsByIntership returns all requests for a specific intership
func (r *RequestRepo) FindRequestsByIntership(intershipID string) ([]model.Request, error) {
	var requests []model.Request
	err := r.db.
		Joins("JOIN followup_issues ON followup_issues.id = requests.followup_issue_id").
		Joins("JOIN followup_milestones ON followup_milestones.id = followup_issues.followup_milestone_id").
		Where("followup_milestones.intership_id = ?", intershipID).
		Find(&requests).Error
	return requests, err
}

// CountRequestsByCompanyAndStatus counts requests by company and status
func (r *RequestRepo) CountRequestsByCompanyAndStatus(companyID string, status string) (int64, error) {
	var count int64
	err := r.db.
		Joins("JOIN followup_issues ON followup_issues.id = requests.followup_issue_id").
		Joins("JOIN followup_milestones ON followup_milestones.id = followup_issues.followup_milestone_id").
		Where("followup_milestones.company_profile_id = ? AND requests.status = ?", companyID, status).
		Model(&model.Request{}).
		Count(&count).Error
	return count, err
}

// CountRequestsByIntershipAndStatus counts requests by intership and status
func (r *RequestRepo) CountRequestsByIntershipAndStatus(intershipID string, status string) (int64, error) {
	var count int64
	err := r.db.
		Joins("JOIN followup_issues ON followup_issues.id = requests.followup_issue_id").
		Joins("JOIN followup_milestones ON followup_milestones.id = followup_issues.followup_milestone_id").
		Where("followup_milestones.intership_id = ? AND requests.status = ?", intershipID, status).
		Model(&model.Request{}).
		Count(&count).Error
	return count, err
}

// CountPendingRequestsByCompany counts pending requests for a company
func (r *RequestRepo) CountPendingRequestsByCompany(companyID string) (int64, error) {
	var count int64
	err := r.db.
		Joins("JOIN followup_issues ON followup_issues.id = requests.followup_issue_id").
		Joins("JOIN followup_milestones ON followup_milestones.id = followup_issues.followup_milestone_id").
		Where("followup_milestones.company_profile_id = ? AND requests.status = ?", companyID, "pending").
		Model(&model.Request{}).
		Count(&count).Error
	return count, err
}

// CountPendingRequestsByIntership counts pending requests for a specific intership
func (r *RequestRepo) CountPendingRequestsByIntership(intershipID string) (int64, error) {
	var count int64
	err := r.db.
		Joins("JOIN followup_issues ON followup_issues.id = requests.followup_issue_id").
		Joins("JOIN followup_milestones ON followup_milestones.id = followup_issues.followup_milestone_id").
		Where("followup_milestones.intership_id = ? AND requests.status = ?", intershipID, "pending").
		Model(&model.Request{}).
		Count(&count).Error
	return count, err
}

// FindRequestReviewTimeByCompany returns review times for reviewed requests by company
func (r *RequestRepo) FindRequestReviewTimeByCompany(companyID string) ([]time.Duration, error) {
	var requests []model.Request
	err := r.db.
		Joins("JOIN followup_issues ON followup_issues.id = requests.followup_issue_id").
		Joins("JOIN followup_milestones ON followup_milestones.id = followup_issues.followup_milestone_id").
		Where("followup_milestones.company_profile_id = ? AND requests.status != ?", companyID, "pending").
		Find(&requests).Error

	if err != nil {
		return nil, err
	}

	var durations []time.Duration
	for _, req := range requests {
		duration := req.UpdatedAt.Sub(req.CreatedAt)
		durations = append(durations, duration)
	}
	return durations, nil
}

// FindRequestReviewTimeByIntership returns review times for reviewed requests by intership
func (r *RequestRepo) FindRequestReviewTimeByIntership(intershipID string) ([]time.Duration, error) {
	var requests []model.Request
	err := r.db.
		Joins("JOIN followup_issues ON followup_issues.id = requests.followup_issue_id").
		Joins("JOIN followup_milestones ON followup_milestones.id = followup_issues.followup_milestone_id").
		Where("followup_milestones.intership_id = ? AND requests.status != ?", intershipID, "pending").
		Find(&requests).Error

	if err != nil {
		return nil, err
	}

	var durations []time.Duration
	for _, req := range requests {
		duration := req.UpdatedAt.Sub(req.CreatedAt)
		durations = append(durations, duration)
	}
	return durations, nil
}

// CountIssuesWithPendingRequestsByCompany counts issues that have pending requests for a company
func (r *RequestRepo) CountIssuesWithPendingRequestsByCompany(companyID string) (int64, error) {
	var count int64
	err := r.db.
		Joins("JOIN followup_issues ON followup_issues.id = requests.followup_issue_id").
		Joins("JOIN followup_milestones ON followup_milestones.id = followup_issues.followup_milestone_id").
		Where("followup_milestones.company_profile_id = ? AND requests.status = ?", companyID, "pending").
		Distinct("followup_issues.id").
		Model(&model.Request{}).
		Count(&count).Error
	return count, err
}

// CountIssuesWithPendingRequestsByIntership counts issues with pending requests for a specific intership
func (r *RequestRepo) CountIssuesWithPendingRequestsByIntership(intershipID string) (int64, error) {
	var count int64
	err := r.db.
		Joins("JOIN followup_issues ON followup_issues.id = requests.followup_issue_id").
		Joins("JOIN followup_milestones ON followup_milestones.id = followup_issues.followup_milestone_id").
		Where("followup_milestones.intership_id = ? AND requests.status = ?", intershipID, "pending").
		Distinct("followup_issues.id").
		Model(&model.Request{}).
		Count(&count).Error
	return count, err
}
