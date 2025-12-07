package jobseeker

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/ito-company/jobsito-service/helper"
	"github.com/ito-company/jobsito-service/src/model"
	"gorm.io/gorm"
)

type JobSeekerRepo interface {
	Create(m model.JobSeekerProfile) error
	FindByEmail(email string) (model.JobSeekerProfile, error)
	FindById(id string) (model.JobSeekerProfile, error)
	Update(m model.JobSeekerProfile) error
	SoftDelete(id string) error
	FindAll(opts *helper.FindAllOptions) ([]model.JobSeekerProfile, int64, error)
	AddTagToJobSeeker(jobSeekerId, tagId, proficiency string) error
	RemoveTagFromJobSeeker(jobSeekerId, tagId string) error
	IsJobSeekerInInternship(jobSeekerId, internshipId string) (bool, error)
}

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) JobSeekerRepo {
	return &Repo{db: db}
}

func (r *Repo) Create(m model.JobSeekerProfile) error {
	return r.db.Create(&m).Error
}

func (r *Repo) FindById(id string) (model.JobSeekerProfile, error) {
	var jobSeeker model.JobSeekerProfile
	err := r.db.Where("id = ?", id).First(&jobSeeker).Error
	return jobSeeker, err
}

func (r *Repo) FindByEmail(email string) (model.JobSeekerProfile, error) {
	var jobSeeker model.JobSeekerProfile
	err := r.db.
		Preload("JobSeekerTags.GlobalTag").
		Where("email = ?", email).
		First(&jobSeeker).Error
	return jobSeeker, err
}

func (r *Repo) Update(m model.JobSeekerProfile) error {
	return r.db.Save(&m).Error
}

func (r *Repo) SoftDelete(id string) error {
	return r.db.Where("id = ?", id).Delete(&model.JobSeekerProfile{}).Error
}

func (r *Repo) FindAll(opts *helper.FindAllOptions) ([]model.JobSeekerProfile, int64, error) {
	var finded []model.JobSeekerProfile
	query := r.db.Model(model.JobSeekerProfile{})
	var total int64
	query, total = helper.ApplyFindAllOptions(query, opts)

	err := query.Find(&finded).Error
	return finded, total, err
}

func (r *Repo) AddTagToJobSeeker(jobSeekerId, tagId, proficiency string) error {
	var existing model.JobSeekerTags
	err := r.db.Where("job_seeker_profile_id = ? AND global_tag_id = ?", jobSeekerId, tagId).First(&existing).Error
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

	jst := model.JobSeekerTags{
		ID:                 uuid.New(),
		ProficiencyLevel:   proficiency,
		JobSeekerProfileID: uuid.MustParse(jobSeekerId),
		GlobalTagID:        uuid.MustParse(tagId),
		GlobalTag:          tag,
	}
	return r.db.Create(&jst).Error
}

func (r *Repo) RemoveTagFromJobSeeker(jobSeekerId, tagId string) error {
	return r.db.Where("job_seeker_profile_id = ? AND global_tag_id = ?", jobSeekerId, tagId).
		Delete(&model.JobSeekerTags{}).Error
}

func (r *Repo) IsJobSeekerInInternship(jobSeekerId, internshipId string) (bool, error) {
	var internship model.Intership
	err := r.db.
		Where("id = ? AND job_seeker_profile_id = ?", internshipId, jobSeekerId).
		First(&internship).
		Error

	if err == gorm.ErrRecordNotFound {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}
