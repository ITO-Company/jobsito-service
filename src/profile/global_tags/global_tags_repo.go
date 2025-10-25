package globaltags

import (
	"github.com/ito-company/jobsito-service/helper"
	"github.com/ito-company/jobsito-service/src/model"
	"gorm.io/gorm"
)

type GlobalTagsRepo interface {
	FindAll(opts *helper.FindAllOptions) ([]model.GlobalTag, int64, error)
}

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) GlobalTagsRepo {
	return &Repo{db: db}
}

func (r *Repo) FindAll(opts *helper.FindAllOptions) ([]model.GlobalTag, int64, error) {
	var finded []model.GlobalTag
	query := r.db.Model(model.GlobalTag{})
	var total int64
	query, total = helper.ApplyFindAllOptions(query, opts)

	err := query.Find(&finded).Error
	return finded, total, err
}
