package jobposting

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/ito-company/jobsito-service/helper"
	"github.com/ito-company/jobsito-service/src/model"
	"gorm.io/gorm"
)

type JobPostingRepo interface {
	Create(m model.JobPosting) error
	CreateWithTags(job model.JobPosting, tagIDs []string) error
	FindById(id string) (*model.JobPosting, error)
	FindCompanyById(id string) (*model.CompanyProfile, error)
	Update(m model.JobPosting) error
	SoftDelete(id string) error
	FindAll(opts *helper.FindAllOptions, tagIDs []string, companyID string) ([]model.JobPosting, int64, error)
	AddTagToJobPosting(jobPostingId, tagId string) error
	RemoveTagFromJobPosting(jobPostingId, tagId string) error
}

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) JobPostingRepo {
	return &Repo{db: db}
}

func (r *Repo) Create(m model.JobPosting) error {
	return r.db.Create(&m).Error
}

func (r *Repo) CreateWithTags(job model.JobPosting, tagIDs []string) error {
	tx := r.db.Begin()

	if err := tx.Create(&job).Error; err != nil {
		tx.Rollback()
		return err
	}

	var tags []model.GlobalTag
	if err := tx.Where("id IN ?", tagIDs).Find(&tags).Error; err != nil {
		tx.Rollback()
		return err
	}

	if len(tags) != len(tagIDs) {
		tx.Rollback()
		return fmt.Errorf("one or more tags not found")
	}

	for _, tag := range tags {
		jpt := model.JobPostingTags{
			ID:           uuid.New(),
			JobPostingId: job.ID,
			GlobalTagID:  tag.ID,
		}
		if err := tx.Create(&jpt).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (r *Repo) FindById(id string) (*model.JobPosting, error) {
	var job model.JobPosting
	err := r.db.
		Preload("JobPostingTags.GlobalTag").
		Preload("CompanyProfile").
		Where("id = ?", id).
		First(&job).Error
	return &job, err
}

func (r *Repo) FindCompanyById(id string) (*model.CompanyProfile, error) {
	var company model.CompanyProfile
	err := r.db.Where("id = ?", id).First(&company).Error
	return &company, err
}

func (r *Repo) Update(m model.JobPosting) error {
	return r.db.Save(&m).Error
}

func (r *Repo) SoftDelete(id string) error {
	return r.db.Where("id = ?", id).Delete(&model.JobPosting{}).Error
}

func (r *Repo) FindAll(opts *helper.FindAllOptions, tagIDs []string, companyID string) ([]model.JobPosting, int64, error) {
	var finded []model.JobPosting
	query := r.db.Model(&model.JobPosting{}).
		Preload("JobPostingTags.GlobalTag")

	if len(tagIDs) > 0 {
		query = query.Joins("JOIN job_posting_tags ON job_postings.id = job_posting_tags.job_posting_id").
			Where("job_posting_tags.global_tag_id IN ?", tagIDs)
	}

	if companyID != "" {
		query = query.Where("company_profile_id = ?", companyID)
	}

	if opts != nil && opts.OrderBy == "" {
		opts.OrderBy = "job_postings.created_at"
	}

	var total int64
	query, total = helper.ApplyFindAllOptions(query, opts)

	err := query.Find(&finded).Error
	return finded, total, err
}

func (r *Repo) AddTagToJobPosting(jobPostingId, tagId string) error {
	var existing model.JobPostingTags
	err := r.db.Where("job_posting_id = ? AND global_tag_id = ?", jobPostingId, tagId).First(&existing).Error
	if err == nil {
		return nil
	}
	if err != gorm.ErrRecordNotFound {
		return err
	}

	var tag model.GlobalTag
	if err := r.db.Where("id = ?", tagId).First(&tag).Error; err != nil {
		return fmt.Errorf("tag not found")
	}

	jpt := model.JobPostingTags{
		ID:           uuid.New(),
		JobPostingId: uuid.MustParse(jobPostingId),
		GlobalTagID:  uuid.MustParse(tagId),
		GlobalTag:    tag,
	}
	return r.db.Create(&jpt).Error
}

func (r *Repo) RemoveTagFromJobPosting(jobPostingId, tagId string) error {
	return r.db.Where("job_posting_id = ? AND global_tag_id = ?", jobPostingId, tagId).
		Delete(&model.JobPostingTags{}).Error
}
