package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"

	"fullstackjobs-service/cmd/internal/api"
	scrapper "fullstackjobs-service/cmd/internal/jobs"
	"fullstackjobs-service/cmd/internal/storage"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  No .env file found, continuing without it")
	}

	app := fiber.New()
	app.Use(cors.New())

	storage.InitDB() // Initialize DB connection
	if storage.DB == nil {
		log.Fatal("❌ Database connection failed")
	}

	api.SetupRoutes(app)

	// Schedule job scraping every hour
	c := cron.New()
	_, err = c.AddFunc("@every 1h", func() {
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
