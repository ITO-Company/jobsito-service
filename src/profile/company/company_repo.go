package company

import (
	"github.com/ito-company/jobsito-service/src/model"
	"gorm.io/gorm"
)

type CompanyRepo interface {
	Create(m model.CompanyProfile) error
	FindByEmail(email string) (*model.CompanyProfile, error)
	Update(m model.CompanyProfile) error
	SoftDelete(id string) error
}

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) CompanyRepo {
	return &Repo{db: db}
}

func (r *Repo) Create(m model.CompanyProfile) error {
	return r.db.Create(&m).Error
}

func (r *Repo) FindByEmail(email string) (*model.CompanyProfile, error) {
	var company model.CompanyProfile
	err := r.db.Where("email = ?", email).First(&company).Error
	return &company, err
}

func (r *Repo) Update(m model.CompanyProfile) error {
	return r.db.Save(&m).Error
}

func (r *Repo) SoftDelete(id string) error {
	return r.db.Where("id = ?", id).Delete(&model.CompanyProfile{}).Error
}
