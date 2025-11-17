package savedjob

import (
	"github.com/ito-company/jobsito-service/helper"
	"github.com/ito-company/jobsito-service/src/model"
	"gorm.io/gorm"
)

type SavedJobRepo interface {
	Create(m *model.SavedJob) error
	Delete(id string) error
	FindById(id string) (*model.SavedJob, error)
	FindByJobSeekerAndJobPosting(jobSeekerID string, jobPostingID string) (*model.SavedJob, error)
	FindJobPostingById(id string) (*model.JobPosting, error)
	FindAllByJobSeeker(jobSeekerID string, opts *helper.FindAllOptions) ([]model.SavedJob, int64, error)
}

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) SavedJobRepo {
	return &Repo{db: db}
}

func (r *Repo) Create(m *model.SavedJob) error {
	return r.db.Create(m).Error
}

func (r *Repo) Delete(id string) error {
	return r.db.Delete(&model.SavedJob{}, "id = ?", id).Error
}

func (r *Repo) FindById(id string) (*model.SavedJob, error) {
	var savedJob model.SavedJob
	if err := r.db.
		Preload("JobPosting").
		Preload("JobPosting.CompanyProfile").
		Preload("JobPosting.JobPostingTags").
		Preload("JobPosting.JobPostingTags.GlobalTag").
		Where("id = ?", id).
		First(&savedJob).Error; err != nil {
		return nil, err
	}
	return &savedJob, nil
}

func (r *Repo) FindByJobSeekerAndJobPosting(jobSeekerID string, jobPostingID string) (*model.SavedJob, error) {
	var savedJob model.SavedJob
	if err := r.db.
		Where("job_seeker_profile_id = ? AND job_posting_id = ?", jobSeekerID, jobPostingID).
		First(&savedJob).Error; err != nil {
		return nil, err
	}
	return &savedJob, nil
}

func (r *Repo) FindJobPostingById(id string) (*model.JobPosting, error) {
	var jobPosting model.JobPosting
	if err := r.db.
		Preload("CompanyProfile").
		Preload("JobPostingTags").
		Preload("JobPostingTags.GlobalTag").
		Where("id = ?", id).
		First(&jobPosting).Error; err != nil {
		return nil, err
	}
	return &jobPosting, nil
}

func (r *Repo) FindAllByJobSeeker(jobSeekerID string, opts *helper.FindAllOptions) ([]model.SavedJob, int64, error) {
	var savedJobs []model.SavedJob
	query := r.db.Model(&model.SavedJob{}).
		Preload("JobPosting").
		Preload("JobPosting.CompanyProfile").
		Preload("JobPosting.JobPostingTags").
		Preload("JobPosting.JobPostingTags.GlobalTag").
		Where("job_seeker_profile_id = ?", jobSeekerID)

	var total int64
	query, total = helper.ApplyFindAllOptions(query, opts)

	err := query.Find(&savedJobs).Error
	return savedJobs, total, err
}
