package seed

import (
	"log"

	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) {
	if err := SeedGlobalTags(db); err != nil {
		log.Fatalf("Error al seedear tags globales: %v", err)
	}
}
