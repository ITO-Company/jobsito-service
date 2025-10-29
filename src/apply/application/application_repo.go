package application

import (
	"github.com/ito-company/jobsito-service/helper"
	"github.com/ito-company/jobsito-service/src/model"
	"gorm.io/gorm"
)

type ApplicationRepo interface {
	Create(m *model.Application) error
	Update(m *model.Application) error
	FindById(id string) (*model.Application, error)
	FindByIDAndJobSeeker(id string, jobSeekerID string) (*model.Application, error)
	FindJobPostingById(id string) (*model.JobPosting, error)
	FindAllByJobSeeker(jobSeekerID string, opts *helper.FindAllOptions) ([]model.Application, int64, error)
	FindAllByJobPostingAndCompany(jobPostingID string, companyID string, opts *helper.FindAllOptions) ([]model.Application, int64, error)
}

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) ApplicationRepo {
	return &Repo{db: db}
}

func (r *Repo) Create(m *model.Application) error {
	return r.db.Create(m).Error
}

func (r *Repo) Update(m *model.Application) error {
	return r.db.Save(m).Error
}

func (r *Repo) FindById(id string) (*model.Application, error) {
	var application model.Application
	if err := r.db.
		Preload("JobSeeker").
		Where("id = ?", id).
		First(&application).Error; err != nil {
		return nil, err
	}
	return &application, nil
}

func (r *Repo) FindByIDAndJobSeeker(id string, jobSeekerID string) (*model.Application, error) {
	var application model.Application
	if err := r.db.
		Where("id = ? AND job_seeker_id = ?", id, jobSeekerID).
		First(&application).Error; err != nil {
		return nil, err
	}
	return &application, nil
}

func (r *Repo) FindJobPostingById(id string) (*model.JobPosting, error) {
	var jobPosting model.JobPosting
	if err := r.db.
		Preload("CompanyProfile").
		Where("id = ?", id).
		First(&jobPosting).Error; err != nil {
		return nil, err
	}
	return &jobPosting, nil
}

func (r *Repo) FindAllByJobSeeker(jobSeekerID string, opts *helper.FindAllOptions) ([]model.Application, int64, error) {
	var applications []model.Application
	query := r.db.Model(&model.Application{}).Where("job_seeker_id = ?", jobSeekerID)
	query, total := helper.ApplyFindAllOptions(query, opts)
	err := query.Find(&applications).Error
	return applications, total, err
}

func (r *Repo) FindAllByJobPostingAndCompany(jobPostingID string, companyID string, opts *helper.FindAllOptions) ([]model.Application, int64, error) {
	var applications []model.Application
	query := r.db.Model(&model.Application{}).
		Joins("JOIN job_postings ON job_postings.id = applications.job_posting_id").
		Where("applications.job_posting_id = ? AND job_postings.company_profile_id = ?", jobPostingID, companyID).
		Preload("JobSeeker")
	query, total := helper.ApplyFindAllOptions(query, opts)
	err := query.Find(&applications).Error
	return applications, total, err
}
