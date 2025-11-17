package issue

import (
	"github.com/ito-company/jobsito-service/helper"
	"github.com/ito-company/jobsito-service/src/model"
	"gorm.io/gorm"
)

type IssueRepo interface {
	Create(issue *model.FollowupIssue) error
	FindById(id string) (*model.FollowupIssue, error)
	FindMilestoneById(id string) (*model.FollowupMilestone, error)
	FindAll(milestoneId string, opts *helper.FindAllOptions) ([]model.FollowupIssue, int64, error)
	Update(issue *model.FollowupIssue) error
}

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) IssueRepo {
	return &Repo{db: db}
}

func (r *Repo) Create(issue *model.FollowupIssue) error {
	return r.db.Create(issue).Error
}

func (r *Repo) FindById(id string) (*model.FollowupIssue, error) {
	var issue model.FollowupIssue
	if err := r.db.First(&issue, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &issue, nil
}

func (r *Repo) FindMilestoneById(id string) (*model.FollowupMilestone, error) {
	var milestone model.FollowupMilestone
	if err := r.db.First(&milestone, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &milestone, nil
}

func (r *Repo) FindAll(milestoneId string, opts *helper.FindAllOptions) ([]model.FollowupIssue, int64, error) {
	var issues []model.FollowupIssue
	query := r.db.Model(&model.FollowupIssue{}).Where("followup_milestone_id = ?", milestoneId)

	var total int64
	query, total = helper.ApplyFindAllOptions(query, opts)

	err := query.Find(&issues).Error

	return issues, total, err
}

func (r *Repo) Update(issue *model.FollowupIssue) error {
	return r.db.Save(issue).Error
}
