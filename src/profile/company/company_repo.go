package company

import (
	"github.com/ito-company/jobsito-service/helper"
	"github.com/ito-company/jobsito-service/src/model"
	"gorm.io/gorm"
)

type CompanyRepo interface {
	Create(m model.CompanyProfile) error
	FindByEmail(email string) (*model.CompanyProfile, error)
	FindById(id string) (*model.CompanyProfile, error)
	Update(m model.CompanyProfile) error
	SoftDelete(id string) error
	FindAll(opts *helper.FindAllOptions) ([]model.CompanyProfile, int64, error)
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

// Ver si es eficiente cargar los jobpostings,
// creo que es mejor un findall de jobposting que se encargue
func (r *Repo) FindByEmail(email string) (*model.CompanyProfile, error) {
	var company model.CompanyProfile
	err := r.db.
		Preload("JobPostings.JobPostingTags.GlobalTag").
		Where("email = ?", email).
		First(&company).Error
	return &company, err
}

func (r *Repo) FindById(id string) (*model.CompanyProfile, error) {
	var company model.CompanyProfile
	err := r.db.
		Where("id = ?", id).
		First(&company).Error
	return &company, err
}

func (r *Repo) Update(m model.CompanyProfile) error {
	return r.db.Save(&m).Error
}

func (r *Repo) SoftDelete(id string) error {
	return r.db.Where("id = ?", id).Delete(&model.CompanyProfile{}).Error
}

func (r *Repo) FindAll(opts *helper.FindAllOptions) ([]model.CompanyProfile, int64, error) {
	var finded []model.CompanyProfile
	query := r.db.Model(model.CompanyProfile{})
	var total int64
	query, total = helper.ApplyFindAllOptions(query, opts)

	err := query.Find(&finded).Error
	return finded, total, err
}
