package storage

import (
	"fullstackjobs-service/cmd/internal/models"
	"log"
)

func SaveJobs(jobs []models.Job) error {
	for _, job := range jobs {
		// Skip if duplicate
		var existing models.Job
		result := DB.First(&existing, "link = ?", job.Link)
		if result.Error == nil {
			continue // job already exists
		}

		err := DB.Create(&job).Error
		if err != nil {
			log.Println("❌ Failed to insert job:", err)
		}
	}
	return nil
}
