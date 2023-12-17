package main

import (
	"log"

	"github.com/santoadji21/santoadji21-go-fiber-product-api/config"
	"github.com/santoadji21/santoadji21-go-fiber-product-api/internal/db"
	"github.com/santoadji21/santoadji21-go-fiber-product-api/pkg/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
    // Initialize Fiber app
    app := fiber.New()

	// Load your configuration 
    cfg := config.LoadConfig()

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

    // Start the Fiber app on a specified port
    app.Listen(":3000")
}


