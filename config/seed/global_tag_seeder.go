package seed

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/ito-company/jobsito-service/src/model"
	"gorm.io/gorm"
)

func SeedGlobalTags(db *gorm.DB) error {
	tags := []model.GlobalTag{
		{
			ID:         uuid.New(),
			Name:       "Arquitecto de Software",
			Category:   "Programación",
			Color:      "#1abc9c",
			IsApproved: "true",
			UsageCount: "0",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			ID:         uuid.New(),
			Name:       "Diseñador UI/UX",
			Category:   "Programación",
			Color:      "#e67e22",
			IsApproved: "true",
			UsageCount: "0",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			ID:         uuid.New(),
			Name:       "Go",
			Category:   "Programación",
			Color:      "#00ADD8",
			IsApproved: "true",
			UsageCount: "0",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			ID:         uuid.New(),
			Name:       "Backend",
			Category:   "Programación",
			Color:      "#34495e",
			IsApproved: "true",
			UsageCount: "0",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			ID:         uuid.New(),
			Name:       "Frontend",
			Category:   "Programación",
			Color:      "#2980b9",
			IsApproved: "true",
			UsageCount: "0",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			ID:         uuid.New(),
			Name:       "DevOps",
			Category:   "Programación",
			Color:      "#27ae60",
			IsApproved: "true",
			UsageCount: "0",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			ID:         uuid.New(),
			Name:       "QA",
			Category:   "Programación",
			Color:      "#8e44ad",
			IsApproved: "true",
			UsageCount: "0",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	for _, tag := range tags {
		var existing model.GlobalTag
		if err := db.Where("name = ?", tag.Name).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&tag).Error; err != nil {
				return fmt.Errorf("no se pudo crear el tag global %s: %v", tag.Name, err)
			} else {
				log.Printf("Tag global creado: %s", tag.Name)
			}
		}
	}

	return nil
}
