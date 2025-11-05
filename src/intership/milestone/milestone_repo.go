package milestone

import (
	"github.com/ito-company/jobsito-service/helper"
	"github.com/ito-company/jobsito-service/src/model"
	"gorm.io/gorm"
)

type MilestoneRepo interface {
	Create(milestone *model.FollowupMilestone) error
	FindById(id string) (*model.FollowupMilestone, error)
	FindCompanyById(id string) (*model.CompanyProfile, error)
	FindJobSeekerById(id string) (*model.JobSeekerProfile, error)
	FindIntershipById(id string) (*model.Intership, error)
	FindAll(intershipId string, opts *helper.FindAllOptions) ([]model.FollowupMilestone, int64, error)
	Update(milestone *model.FollowupMilestone) error
}

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) MilestoneRepo {
	return &Repo{db: db}
}

func (r *Repo) Create(milestone *model.FollowupMilestone) error {
	return r.db.Create(milestone).Error
}

func (r *Repo) FindById(id string) (*model.FollowupMilestone, error) {
	var milestone model.FollowupMilestone
	if err := r.db.First(&milestone, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &milestone, nil
}

func (r *Repo) FindCompanyById(id string) (*model.CompanyProfile, error) {
	var company model.CompanyProfile
	if err := r.db.First(&company, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &company, nil
}

func (r *Repo) FindJobSeekerById(id string) (*model.JobSeekerProfile, error) {
	var jobSeeker model.JobSeekerProfile
	if err := r.db.First(&jobSeeker, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &jobSeeker, nil
}

func (r *Repo) FindIntershipById(id string) (*model.Intership, error) {
	var intership model.Intership
	if err := r.db.First(&intership, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &intership, nil
}

func (r *Repo) FindAll(intershipId string, opts *helper.FindAllOptions) ([]model.FollowupMilestone, int64, error) {
	var milestones []model.FollowupMilestone
	query := r.db.Model(&model.FollowupMilestone{}).
		Where("intership_id = ?", intershipId)

	var total int64
	query, total = helper.ApplyFindAllOptions(query, opts)

	err := query.Find(&milestones).Error

	return milestones, total, err
}

func (r *Repo) Update(milestone *model.FollowupMilestone) error {
	return r.db.Save(milestone).Error
}
