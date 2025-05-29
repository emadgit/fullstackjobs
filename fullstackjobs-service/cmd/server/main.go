package main

import (
	"fullstackjobs-service/cmd/internal/api"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Register routes
	api.SetupRoutes(app)

	// Start server
	log.Println("🚀 Server running on http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
