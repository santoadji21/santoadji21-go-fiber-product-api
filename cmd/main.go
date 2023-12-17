package main

import (
	"log"

	"github.com/santoadji21/santoadji21-go-fiber-product-api/config"
	"github.com/santoadji21/santoadji21-go-fiber-product-api/internal/db"
	"github.com/santoadji21/santoadji21-go-fiber-product-api/pkg/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"                                     // swagger middleware
	_ "github.com/santoadji21/santoadji21-go-fiber-product-api/docs" // swagger docs
	_ "github.com/santoadji21/santoadji21-go-fiber-product-api/pkg/handlers"
)

// @title Fiber Product API
// @version 1.0
// @description This is a sample server for Fiber Product API.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Load your configuration
	cfg := config.DbCfg()

	// Initialize the database
	db.ConnectDB(cfg)

	// Check if the database connection was successful
	if db.GetDB() != nil {
		log.Println("Successfully connected to the database")
	} else {
		log.Fatalln("Failed to connect to the database")
	}

	// Setup routes
	routes.AppRoutes(app)

	// Setup swagger middleware
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Start the Fiber app on a specified port
	app.Listen(":3000")
}
