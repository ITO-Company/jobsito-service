package company

import (
	"github.com/ito-company/jobsito-service/src/model"
	"gorm.io/gorm"
)

type CompanyRepo interface {
	Create(m model.CompanyProfile) error
	FindByEmail(email string) (*model.CompanyProfile, error)
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
