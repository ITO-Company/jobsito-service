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
		&model.ApplicationStatusHistory{},
		&model.JobPostingTags{},
		&model.JobSeekerTags{},
		&model.DailyJobStats{},
		&model.JobView{},
		&model.ProfileView{},
		&model.SavedJob{},
		&model.Intership{},
		&model.FollowupMilestone{},
		&model.FollowupIssue{},
		&model.WeeklyCompanyMetrics{},
		&model.Request{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database", err)
	}
}
