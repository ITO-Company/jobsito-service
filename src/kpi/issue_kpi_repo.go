package kpi

import (
	"time"

	"github.com/ito-company/jobsito-service/src/model"
	"gorm.io/gorm"
)

// IssueKPIRepo handles all queries for issue KPIs
type IssueKPIRepo interface {
	// For a specific company (through milestones)
	FindIssuesByCompany(companyID string) ([]model.FollowupIssue, error)
	FindIssuesByIntership(intershipID string) ([]model.FollowupIssue, error)

	// Count operations
	CountIssuesByCompanyAndStatus(companyID string, status string) (int64, error)
	CountIssuesByIntershipAndStatus(intershipID string, status string) (int64, error)

	// Aggregate operations
	CountOverdueIssuesByCompany(companyID string) (int64, error)
	CountOverdueIssuesByIntership(intershipID string) (int64, error)

	// Detail operations
	FindIssueCompletionTimeByCompany(companyID string) ([]time.Duration, error)
	FindIssueCompletionTimeByIntership(intershipID string) ([]time.Duration, error)

	// Issues with requests
	CountIssuesWithRequestsByCompany(companyID string) (int64, error)
	CountIssuesWithRequestsByIntership(intershipID string) (int64, error)
}

type IssueRepo struct {
	db *gorm.DB
}

func NewIssueKPIRepo(db *gorm.DB) IssueKPIRepo {
	return &IssueRepo{db: db}
}

// FindIssuesByCompany returns all issues for a company (through milestones)
func (r *IssueRepo) FindIssuesByCompany(companyID string) ([]model.FollowupIssue, error) {
	var issues []model.FollowupIssue
	err := r.db.
		Joins("JOIN followup_milestones ON followup_milestones.id = followup_issues.followup_milestone_id").
		Where("followup_milestones.company_profile_id = ?", companyID).
		Find(&issues).Error
	return issues, err
}

// FindIssuesByIntership returns all issues for a specific intership
func (r *IssueRepo) FindIssuesByIntership(intershipID string) ([]model.FollowupIssue, error) {
	var issues []model.FollowupIssue
	err := r.db.
		Joins("JOIN followup_milestones ON followup_milestones.id = followup_issues.followup_milestone_id").
		Where("followup_milestones.intership_id = ?", intershipID).
		Find(&issues).Error
	return issues, err
}

// CountIssuesByCompanyAndStatus counts issues by company and status
func (r *IssueRepo) CountIssuesByCompanyAndStatus(companyID string, status string) (int64, error) {
	var count int64
	err := r.db.
		Joins("JOIN followup_milestones ON followup_milestones.id = followup_issues.followup_milestone_id").
		Where("followup_milestones.company_profile_id = ? AND followup_issues.status = ?", companyID, status).
		Model(&model.FollowupIssue{}).
		Count(&count).Error
	return count, err
}

// CountIssuesByIntershipAndStatus counts issues by intership and status
func (r *IssueRepo) CountIssuesByIntershipAndStatus(intershipID string, status string) (int64, error) {
	var count int64
	err := r.db.
		Joins("JOIN followup_milestones ON followup_milestones.id = followup_issues.followup_milestone_id").
		Where("followup_milestones.intership_id = ? AND followup_issues.status = ?", intershipID, status).
		Model(&model.FollowupIssue{}).
		Count(&count).Error
	return count, err
}

// CountOverdueIssuesByCompany counts issues that are overdue
func (r *IssueRepo) CountOverdueIssuesByCompany(companyID string) (int64, error) {
	var count int64
	err := r.db.
		Joins("JOIN followup_milestones ON followup_milestones.id = followup_issues.followup_milestone_id").
		Where("followup_milestones.company_profile_id = ? AND followup_issues.due_date < ? AND followup_issues.status != ?", companyID, time.Now(), "approved").
		Model(&model.FollowupIssue{}).
		Count(&count).Error
	return count, err
}

// CountOverdueIssuesByIntership counts issues that are overdue for a specific intership
func (r *IssueRepo) CountOverdueIssuesByIntership(intershipID string) (int64, error) {
	var count int64
	err := r.db.
		Joins("JOIN followup_milestones ON followup_milestones.id = followup_issues.followup_milestone_id").
		Where("followup_milestones.intership_id = ? AND followup_issues.due_date < ? AND followup_issues.status != ?", intershipID, time.Now(), "approved").
		Model(&model.FollowupIssue{}).
		Count(&count).Error
	return count, err
}

// FindIssueCompletionTimeByCompany returns completion times for completed issues by company
func (r *IssueRepo) FindIssueCompletionTimeByCompany(companyID string) ([]time.Duration, error) {
	var issues []model.FollowupIssue
	err := r.db.
		Joins("JOIN followup_milestones ON followup_milestones.id = followup_issues.followup_milestone_id").
		Where("followup_milestones.company_profile_id = ? AND followup_issues.status = ?", companyID, "approved").
		Find(&issues).Error

	if err != nil {
		return nil, err
	}

	var durations []time.Duration
	for _, i := range issues {
		duration := i.UpdatedAt.Sub(i.CreatedAt)
		durations = append(durations, duration)
	}
	return durations, nil
}

// FindIssueCompletionTimeByIntership returns completion times for completed issues by intership
func (r *IssueRepo) FindIssueCompletionTimeByIntership(intershipID string) ([]time.Duration, error) {
	var issues []model.FollowupIssue
	err := r.db.
		Joins("JOIN followup_milestones ON followup_milestones.id = followup_issues.followup_milestone_id").
		Where("followup_milestones.intership_id = ? AND followup_issues.status = ?", intershipID, "approved").
		Find(&issues).Error

	if err != nil {
		return nil, err
	}

	var durations []time.Duration
	for _, i := range issues {
		duration := i.UpdatedAt.Sub(i.CreatedAt)
		durations = append(durations, duration)
	}
	return durations, nil
}

// CountIssuesWithRequestsByCompany counts issues that have requests
func (r *IssueRepo) CountIssuesWithRequestsByCompany(companyID string) (int64, error) {
	var count int64
	err := r.db.
		Joins("JOIN followup_milestones ON followup_milestones.id = followup_issues.followup_milestone_id").
		Joins("JOIN requests ON requests.followup_issue_id = followup_issues.id").
		Where("followup_milestones.company_profile_id = ?", companyID).
		Distinct("followup_issues.id").
		Model(&model.FollowupIssue{}).
		Count(&count).Error
	return count, err
}

// CountIssuesWithRequestsByIntership counts issues with requests for a specific intership
func (r *IssueRepo) CountIssuesWithRequestsByIntership(intershipID string) (int64, error) {
	var count int64
	err := r.db.
		Joins("JOIN followup_milestones ON followup_milestones.id = followup_issues.followup_milestone_id").
		Joins("JOIN requests ON requests.followup_issue_id = followup_issues.id").
		Where("followup_milestones.intership_id = ?", intershipID).
		Distinct("followup_issues.id").
		Model(&model.FollowupIssue{}).
		Count(&count).Error
	return count, err
}
