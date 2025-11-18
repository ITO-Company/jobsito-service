package kpi

import (
	"time"

	"github.com/ito-company/jobsito-service/src/model"
	"gorm.io/gorm"
)

// MilestoneKPIRepo handles all queries for milestone KPIs
type MilestoneKPIRepo interface {
	// For a specific company
	FindMilestonesByCompany(companyID string) ([]model.FollowupMilestone, error)
	FindMilestonesByIntership(intershipID string) ([]model.FollowupMilestone, error)

	// Count operations
	CountMilestonesByCompanyAndStatus(companyID string, status string) (int64, error)
	CountMilestonesByIntershipAndStatus(intershipID string, status string) (int64, error)

	// Aggregate operations
	CountOverdueMilestonesByCompany(companyID string) (int64, error)
	CountOverdueMilestonesByIntership(intershipID string) (int64, error)

	// Detail operations
	FindMilestoneCompletionTimeByCompany(companyID string) ([]time.Duration, error)
	FindMilestoneCompletionTimeByIntership(intershipID string) ([]time.Duration, error)
}

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) MilestoneKPIRepo {
	return &Repo{db: db}
}

// FindMilestonesByCompany returns all milestones created by a specific company
func (r *Repo) FindMilestonesByCompany(companyID string) ([]model.FollowupMilestone, error) {
	var milestones []model.FollowupMilestone
	err := r.db.
		Where("company_profile_id = ?", companyID).
		Find(&milestones).Error
	return milestones, err
}

// FindMilestonesByIntership returns all milestones for a specific intership
func (r *Repo) FindMilestonesByIntership(intershipID string) ([]model.FollowupMilestone, error) {
	var milestones []model.FollowupMilestone
	err := r.db.
		Where("intership_id = ?", intershipID).
		Find(&milestones).Error
	return milestones, err
}

// CountMilestonesByCompanyAndStatus counts milestones by company and status
func (r *Repo) CountMilestonesByCompanyAndStatus(companyID string, status string) (int64, error) {
	var count int64
	err := r.db.
		Where("company_profile_id = ? AND status = ?", companyID, status).
		Model(&model.FollowupMilestone{}).
		Count(&count).Error
	return count, err
}

// CountMilestonesByIntershipAndStatus counts milestones by intership and status
func (r *Repo) CountMilestonesByIntershipAndStatus(intershipID string, status string) (int64, error) {
	var count int64
	err := r.db.
		Where("intership_id = ? AND status = ?", intershipID, status).
		Model(&model.FollowupMilestone{}).
		Count(&count).Error
	return count, err
}

// CountOverdueMilestonesByCompany counts milestones that are overdue (DueDate < now and not completed)
func (r *Repo) CountOverdueMilestonesByCompany(companyID string) (int64, error) {
	var count int64
	err := r.db.
		Where("company_profile_id = ? AND due_date < ? AND status != ?", companyID, time.Now(), "approved").
		Model(&model.FollowupMilestone{}).
		Count(&count).Error
	return count, err
}

// CountOverdueMilestonesByIntership counts milestones that are overdue for a specific intership
func (r *Repo) CountOverdueMilestonesByIntership(intershipID string) (int64, error) {
	var count int64
	err := r.db.
		Where("intership_id = ? AND due_date < ? AND status != ?", intershipID, time.Now(), "approved").
		Model(&model.FollowupMilestone{}).
		Count(&count).Error
	return count, err
}

// FindMilestoneCompletionTimeByCompany returns completion times for completed milestones by company
func (r *Repo) FindMilestoneCompletionTimeByCompany(companyID string) ([]time.Duration, error) {
	var milestones []model.FollowupMilestone
	err := r.db.
		Where("company_profile_id = ? AND status = ?", companyID, "approved").
		Find(&milestones).Error

	if err != nil {
		return nil, err
	}

	var durations []time.Duration
	for _, m := range milestones {
		duration := m.UpdatedAt.Sub(m.CreatedAt)
		durations = append(durations, duration)
	}
	return durations, nil
}

// FindMilestoneCompletionTimeByIntership returns completion times for completed milestones by intership
func (r *Repo) FindMilestoneCompletionTimeByIntership(intershipID string) ([]time.Duration, error) {
	var milestones []model.FollowupMilestone
	err := r.db.
		Where("intership_id = ? AND status = ?", intershipID, "approved").
		Find(&milestones).Error

	if err != nil {
		return nil, err
	}

	var durations []time.Duration
	for _, m := range milestones {
		duration := m.UpdatedAt.Sub(m.CreatedAt)
		durations = append(durations, duration)
	}
	return durations, nil
}
