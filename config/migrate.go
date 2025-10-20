package config

import (
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	createEnum := `
	DO $$
	BEGIN
	    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'state') THEN
	        CREATE TYPE state AS ENUM ('PENDING','ACTIVE','FINISHED','REGENERATED','ERROR');
	    END IF;
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'provider') THEN
	        CREATE TYPE provider AS ENUM ('OPENAI','GEMINI','ELEVENLAB');
	    END IF;
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'audio_line') THEN
	        CREATE TYPE audio_line AS ENUM ('TTS','SFX');
	    END IF;
	END$$;`
	if err := db.Exec(createEnum).Error; err != nil {
		log.Fatal("Failed to create enums", err)
	}

	err := db.AutoMigrate()

	if err != nil {
		log.Fatal("Failed to migrate database", err)
	}
}
