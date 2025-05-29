package api

import (
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes registers all API endpoints
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/jobs", getJobs)
	api.Get("/jobs/:id", getJobByID)
}

// Dummy handler for GET /jobs
func getJobs(c *fiber.Ctx) error {
	// Return a placeholder response
	return c.JSON(fiber.Map{
		"jobs": []string{"Google - Fullstack Engineer", "Netflix - Senior Fullstack Dev"},
	})
}

// Dummy handler for GET /jobs/:id
func getJobByID(c *fiber.Ctx) error {
	jobID := c.Params("id")
	return c.JSON(fiber.Map{
		"id":   jobID,
		"role": "Fullstack Engineer",
	})
}
