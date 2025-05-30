package main

import (
	"fullstackjobs-service/cmd/internal/api"
	scrapper "fullstackjobs-service/cmd/internal/jobs"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/robfig/cron/v3"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	// Register routes
	api.SetupRoutes(app)

	// Start job scraping in a separate goroutine
	c := cron.New()
	_, err := c.AddFunc("@every 1h", func() {
		log.Println("🕒 Starting job scraping...")
		scrapper.ScrapeJobs()
	})
	if err != nil {
		log.Fatalf("Failed to schedule job scraping: %v", err)
	}
	c.Start()
	// Start server
	log.Println("🚀 Server running on http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
