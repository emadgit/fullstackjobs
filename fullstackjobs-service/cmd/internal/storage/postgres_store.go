package storage

import (
	"errors"
	"fullstackjobs-service/cmd/internal/models"
	"log"

	"gorm.io/gorm"
)

func SaveJobs(jobs []models.Job) error {
	for _, job := range jobs {
		// Skip if duplicate
		var existing models.Job
		result := DB.First(&existing, "link = ? OR (title = ? AND company = ?)", job.Link, job.Title, job.Company)
		if result.Error == nil {
			continue // job already exists
		}
		if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("⚠️ DB query error for job: %v\n", result.Error)
		}

		err := DB.Create(&job).Error
		if err != nil {
			log.Println("❌ Failed to insert job:", err)
		}
	}
	return nil
}
