package config

import (
	"log"

	"github.com/ito-company/jobsito-service/src/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.GlobalTag{},
		&model.CompanyProfile{},
		&model.JobSeekerProfile{},
		&model.JobPosting{},
		&model.Application{},
		&model.JobPostingTags{},
		&model.JobSeekerTags{},
		&model.SavedJob{},
		&model.Intership{},
		&model.FollowupMilestone{},
		&model.FollowupIssue{},
		&model.Request{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database", err)
	}
}
