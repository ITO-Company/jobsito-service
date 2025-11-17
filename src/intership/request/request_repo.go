package request

import (
	"github.com/ito-company/jobsito-service/helper"
	"github.com/ito-company/jobsito-service/src/model"
	"gorm.io/gorm"
)

type RequestRepo interface {
	Create(request *model.Request) error
	FindById(id string) (*model.Request, error)
	FindByIssueId(issueId string, opts *helper.FindAllOptions) ([]model.Request, int64, error)
	FindIssueById(id string) (*model.FollowupIssue, error)
	Update(request *model.Request) error
}

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) RequestRepo {
	return &Repo{db: db}
}

func (r *Repo) Create(request *model.Request) error {
	return r.db.Create(request).Error
}

func (r *Repo) FindById(id string) (*model.Request, error) {
	var request model.Request
	if err := r.db.First(&request, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &request, nil
}

func (r *Repo) FindByIssueId(issueId string, opts *helper.FindAllOptions) ([]model.Request, int64, error) {
	var requests []model.Request
	query := r.db.Model(&model.Request{}).Where("followup_issue_id = ?", issueId)

	var total int64
	query, total = helper.ApplyFindAllOptions(query, opts)

	err := query.Find(&requests).Error

	return requests, total, err
}

func (r *Repo) FindIssueById(id string) (*model.FollowupIssue, error) {
	var issue model.FollowupIssue
	if err := r.db.First(&issue, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &issue, nil
}

func (r *Repo) Update(request *model.Request) error {
	return r.db.Save(request).Error
}
