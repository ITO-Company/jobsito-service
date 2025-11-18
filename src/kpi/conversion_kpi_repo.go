package kpi

import (
	"time"

	"github.com/ito-company/jobsito-service/src/model"
	"gorm.io/gorm"
)

// ConversionKPIRepo handles all queries for conversion KPIs by company
type ConversionKPIRepo interface {
	// Company level conversions
	CountApplicationsByCompany(companyID string) (int64, error)
	CountAcceptedApplicationsByCompany(companyID string) (int64, error)
	CountInitiatedIntershipsByCompany(companyID string) (int64, error)
	CountCompletedIntershipsByCompany(companyID string) (int64, error)

	// Per job posting conversions
	CountApplicationsByJobPosting(jobPostingID string) (int64, error)
	CountAcceptedApplicationsByJobPosting(jobPostingID string) (int64, error)
	CountInitiatedIntershipsByJobPosting(jobPostingID string) (int64, error)
	CountCompletedIntershipsByJobPosting(jobPostingID string) (int64, error)

	// Timing analysis
	FindAverageTimeApplicationToAcceptanceByCompany(companyID string) (time.Duration, error)
	FindAverageTimeAcceptanceToIntershipByCompany(companyID string) (time.Duration, error)
	FindAverageTimeApplicationToAcceptanceByJobPosting(jobPostingID string) (time.Duration, error)
	FindAverageTimeAcceptanceToIntershipByJobPosting(jobPostingID string) (time.Duration, error)

	// Salary analysis
	FindAverageSalaryProposedVsOfferedByCompany(companyID string) (proposedAvg string, offeredMin string, offeredMax string, err error)
	FindAverageSalaryProposedVsOfferedByJobPosting(jobPostingID string) (proposedAvg string, offeredMin string, offeredMax string, err error)
}

type ConversionRepo struct {
	db *gorm.DB
}

func NewConversionKPIRepo(db *gorm.DB) ConversionKPIRepo {
	return &ConversionRepo{db: db}
}

// CountApplicationsByCompany counts total applications for a company's job postings
func (r *ConversionRepo) CountApplicationsByCompany(companyID string) (int64, error) {
	var count int64
	err := r.db.
		Joins("JOIN job_postings ON job_postings.id = applications.job_posting_id").
		Where("job_postings.company_profile_id = ?", companyID).
		Model(&model.Application{}).
		Count(&count).Error
	return count, err
}

// CountAcceptedApplicationsByCompany counts accepted applications for a company
func (r *ConversionRepo) CountAcceptedApplicationsByCompany(companyID string) (int64, error) {
	var count int64
	err := r.db.
		Joins("JOIN job_postings ON job_postings.id = applications.job_posting_id").
		Where("job_postings.company_profile_id = ? AND applications.is_accepted = ?", companyID, true).
		Model(&model.Application{}).
		Count(&count).Error
	return count, err
}

// CountInitiatedIntershipsByCompany counts interships that have been initiated for a company
func (r *ConversionRepo) CountInitiatedIntershipsByCompany(companyID string) (int64, error) {
	var count int64
	err := r.db.
		Where("company_profile_id = ?", companyID).
		Model(&model.Intership{}).
		Count(&count).Error
	return count, err
}

// CountCompletedIntershipsByCompany counts completed interships for a company
func (r *ConversionRepo) CountCompletedIntershipsByCompany(companyID string) (int64, error) {
	var count int64
	err := r.db.
		Where("company_profile_id = ? AND status = ?", companyID, "approved").
		Model(&model.Intership{}).
		Count(&count).Error
	return count, err
}

// CountApplicationsByJobPosting counts total applications for a specific job posting
func (r *ConversionRepo) CountApplicationsByJobPosting(jobPostingID string) (int64, error) {
	var count int64
	err := r.db.
		Where("job_posting_id = ?", jobPostingID).
		Model(&model.Application{}).
		Count(&count).Error
	return count, err
}

// CountAcceptedApplicationsByJobPosting counts accepted applications for a specific job posting
func (r *ConversionRepo) CountAcceptedApplicationsByJobPosting(jobPostingID string) (int64, error) {
	var count int64
	err := r.db.
		Where("job_posting_id = ? AND is_accepted = ?", jobPostingID, true).
		Model(&model.Application{}).
		Count(&count).Error
	return count, err
}

// CountInitiatedIntershipsByJobPosting counts interships initiated from a specific job posting
func (r *ConversionRepo) CountInitiatedIntershipsByJobPosting(jobPostingID string) (int64, error) {
	var count int64
	err := r.db.
		Where("job_posting_id = ?", jobPostingID).
		Model(&model.Intership{}).
		Count(&count).Error
	return count, err
}

// CountCompletedIntershipsByJobPosting counts completed interships from a specific job posting
func (r *ConversionRepo) CountCompletedIntershipsByJobPosting(jobPostingID string) (int64, error) {
	var count int64
	err := r.db.
		Where("job_posting_id = ? AND status = ?", jobPostingID, "approved").
		Model(&model.Intership{}).
		Count(&count).Error
	return count, err
}

// FindAverageTimeApplicationToAcceptanceByCompany calculates average time from application creation to acceptance
func (r *ConversionRepo) FindAverageTimeApplicationToAcceptanceByCompany(companyID string) (time.Duration, error) {
	var result struct {
		AvgSeconds float64
	}

	err := r.db.
		Joins("JOIN job_postings ON job_postings.id = applications.job_posting_id").
		Where("job_postings.company_profile_id = ? AND applications.is_accepted = ?", companyID, true).
		Model(&model.Application{}).
		Select("AVG(EXTRACT(EPOCH FROM (applications.updated_at - applications.created_at))) as avg_seconds").
		Scan(&result).Error

	if err != nil {
		return 0, err
	}

	return time.Duration(int64(result.AvgSeconds)) * time.Second, nil
}

// FindAverageTimeAcceptanceToIntershipByCompany calculates average time from accepted application to intership initiation
func (r *ConversionRepo) FindAverageTimeAcceptanceToIntershipByCompany(companyID string) (time.Duration, error) {
	var result struct {
		AvgSeconds float64
	}

	err := r.db.
		Joins("JOIN applications ON applications.id = interships.job_posting_id").
		Joins("JOIN job_postings ON job_postings.id = applications.job_posting_id").
		Where("job_postings.company_profile_id = ? AND applications.is_accepted = ?", companyID, true).
		Model(&model.Intership{}).
		Select("AVG(EXTRACT(EPOCH FROM (interships.created_at - applications.updated_at))) as avg_seconds").
		Scan(&result).Error

	if err != nil {
		return 0, err
	}

	return time.Duration(int64(result.AvgSeconds)) * time.Second, nil
}

// FindAverageTimeApplicationToAcceptanceByJobPosting calculates average time for a specific job posting
func (r *ConversionRepo) FindAverageTimeApplicationToAcceptanceByJobPosting(jobPostingID string) (time.Duration, error) {
	var result struct {
		AvgSeconds float64
	}

	err := r.db.
		Where("job_posting_id = ? AND is_accepted = ?", jobPostingID, true).
		Model(&model.Application{}).
		Select("AVG(EXTRACT(EPOCH FROM (updated_at - created_at))) as avg_seconds").
		Scan(&result).Error

	if err != nil {
		return 0, err
	}

	return time.Duration(int64(result.AvgSeconds)) * time.Second, nil
}

// FindAverageTimeAcceptanceToIntershipByJobPosting calculates average time from acceptance to intership for a job posting
func (r *ConversionRepo) FindAverageTimeAcceptanceToIntershipByJobPosting(jobPostingID string) (time.Duration, error) {
	var result struct {
		AvgSeconds float64
	}

	err := r.db.
		Joins("JOIN applications ON applications.job_posting_id = ?", jobPostingID).
		Where("applications.is_accepted = ?", true).
		Model(&model.Intership{}).
		Select("AVG(EXTRACT(EPOCH FROM (interships.created_at - applications.updated_at))) as avg_seconds").
		Scan(&result).Error

	if err != nil {
		return 0, err
	}

	return time.Duration(int64(result.AvgSeconds)) * time.Second, nil
}

// FindAverageSalaryProposedVsOfferedByCompany compares average proposed salary vs offered
func (r *ConversionRepo) FindAverageSalaryProposedVsOfferedByCompany(companyID string) (proposedAvg string, offeredMin string, offeredMax string, err error) {
	var result struct {
		ProposedAvg string
		OfferedMin  string
		OfferedMax  string
	}

	err = r.db.
		Joins("JOIN job_postings ON job_postings.id = applications.job_posting_id").
		Where("job_postings.company_profile_id = ?", companyID).
		Model(&model.Application{}).
		Select(
			"AVG(CAST(applications.proposed_salary AS NUMERIC)) as proposed_avg",
			"AVG(CAST(job_postings.salary_min AS NUMERIC)) as offered_min",
			"AVG(CAST(job_postings.salary_max AS NUMERIC)) as offered_max",
		).
		Scan(&result).Error

	return result.ProposedAvg, result.OfferedMin, result.OfferedMax, err
}

// FindAverageSalaryProposedVsOfferedByJobPosting compares salary for a specific job posting
func (r *ConversionRepo) FindAverageSalaryProposedVsOfferedByJobPosting(jobPostingID string) (proposedAvg string, offeredMin string, offeredMax string, err error) {
	var result struct {
		ProposedAvg string
		OfferedMin  string
		OfferedMax  string
	}

	err = r.db.
		Where("job_posting_id = ?", jobPostingID).
		Model(&model.Application{}).
		Select(
			"AVG(CAST(proposed_salary AS NUMERIC)) as proposed_avg",
			"(SELECT salary_min FROM job_postings WHERE id = ?) as offered_min",
			"(SELECT salary_max FROM job_postings WHERE id = ?) as offered_max",
			jobPostingID,
			jobPostingID,
		).
		Scan(&result).Error

	return result.ProposedAvg, result.OfferedMin, result.OfferedMax, err
}
