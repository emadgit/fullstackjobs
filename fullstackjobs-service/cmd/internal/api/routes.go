package api

import (
	"fullstackjobs-service/cmd/internal/models"
	"fullstackjobs-service/cmd/internal/storage"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes registers all API endpoints
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/jobs", getJobs)
	api.Get("/jobs/:id", getJobByID)
}

// GET /api/jobs - Returns all jobs
func getJobs(c *fiber.Ctx) error {
	var jobs []models.Job
	if err := storage.DB.Order("posted_at desc").Find(&jobs).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch jobs",
		})
	}
	return c.JSON(jobs)
}

// GET /api/jobs/:id - Returns a single job by ID
func getJobByID(c *fiber.Ctx) error {
	jobID := c.Params("id")
	var job models.Job

	if err := storage.DB.First(&job, "id = ?", jobID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Job not found",
		})
	}
	return c.JSON(job)
}
