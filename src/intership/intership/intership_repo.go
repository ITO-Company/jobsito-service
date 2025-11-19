package intership

import (
	"github.com/ito-company/jobsito-service/helper"
	"github.com/ito-company/jobsito-service/src/model"
	"gorm.io/gorm"
)

type IntershipRepo interface {
	Create(m model.Intership) error
	FindById(id string) (*model.Intership, error)
	FindCompanyById(id string) (*model.CompanyProfile, error)
	FindJobSeekerById(id string) (*model.JobSeekerProfile, error)
	FindJobPostingById(id string) (*model.JobPosting, error)
	FindAll(companyID string, jobSeekerID string, opts *helper.FindAllOptions) ([]model.Intership, int64, error)
}

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) IntershipRepo {
	return &Repo{db: db}
}

func (r *Repo) Create(m model.Intership) error {
	return r.db.Create(&m).Error
}

func (r *Repo) FindById(id string) (*model.Intership, error) {
	var intership model.Intership
	if err := r.db.
		Preload("JobPosting").
		Preload("JobSeekerProfile").
		Preload("CompanyProfile").
		First(&intership, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &intership, nil
}

func (r *Repo) FindCompanyById(id string) (*model.CompanyProfile, error) {
	var company model.CompanyProfile
	if err := r.db.First(&company, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &company, nil
}

func (r *Repo) FindJobSeekerById(id string) (*model.JobSeekerProfile, error) {
	var jobSeeker model.JobSeekerProfile
	if err := r.db.First(&jobSeeker, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &jobSeeker, nil
}

func (r *Repo) FindJobPostingById(id string) (*model.JobPosting, error) {
	var jobPosting model.JobPosting
	if err := r.db.First(&jobPosting, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &jobPosting, nil
}

func (r *Repo) FindAll(companyID string, jobSeekerID string, opts *helper.FindAllOptions) ([]model.Intership, int64, error) {
	var interships []model.Intership
	query := r.db.Model(&model.Intership{}).
		Preload("JobPosting").
		Preload("JobSeekerProfile").
		Preload("CompanyProfile")

	if companyID != "" {
		query = query.Where("company_profile_id = ?", companyID)
	}

	if jobSeekerID != "" {
		query = query.Where("job_seeker_profile_id = ?", jobSeekerID)
	}

	var total int64
	query, total = helper.ApplyFindAllOptions(query, opts)

	err := query.Find(&interships).Error

	return interships, total, err
}
