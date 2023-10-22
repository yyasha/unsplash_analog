package http_api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// setup all routes
func StartHttpServer() {
	app := fiber.New()
	// Allow CORS
	app.Use(cors.New())
	// setup routes
	setupRoutes(app)
	// start listen
	log.Println("HTTP server started")
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}

// SetupRoutes func for describe group of api routes
func setupRoutes(app *fiber.App) {

}
