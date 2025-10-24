package jobseeker

import (
	"github.com/ito-company/jobsito-service/src/model"
	"gorm.io/gorm"
)

type JobSeekerRepo interface {
	Create(m model.JobSeekerProfile) error
	FindByEmail(email string) (model.JobSeekerProfile, error)
	Update(m model.JobSeekerProfile) error
	SoftDelete(id string) error
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

func (r *Repo) FindByEmail(email string) (model.JobSeekerProfile, error) {
	var jobSeeker model.JobSeekerProfile
	err := r.db.Where("email = ?", email).First(&jobSeeker).Error
	return jobSeeker, err
}

func (r *Repo) Update(m model.JobSeekerProfile) error {
	return r.db.Save(&m).Error
}

func (r *Repo) SoftDelete(id string) error {
	return r.db.Where("id = ?", id).Delete(&model.JobSeekerProfile{}).Error
}
