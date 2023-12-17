package routes

import (
	"github.com/santoadji21/santoadji21-go-fiber-product-api/pkg/handlers"
	"github.com/santoadji21/santoadji21-go-fiber-product-api/pkg/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AppRoutes(app *fiber.App) {
	// User routes
	app.Post("/api/users", handlers.CreateUser)
	app.Get("/api/users", handlers.GetAllUsers)
	app.Get("/api/users/:id", handlers.GetUser)
	app.Patch("/api/users/:id", handlers.UpdateUser)
	app.Delete("/api/users/:id", handlers.DeleteUser)

	// Auth routes
	app.Post("/api/login", handlers.Login)

	// Product routes
	app.Post("/api/product", middlewares.Protected(), handlers.CreateProduct)
	app.Get("/api/products", handlers.GetAllProducts)
	app.Get("/api/product/:id", handlers.GetProduct)
	app.Patch("/api/product/:id", middlewares.Protected(), handlers.UpdateProduct)
	app.Delete("/api/product/:id", middlewares.Protected(), handlers.DeleteProduct)

	// Category routes
	app.Post("/api/categories", handlers.CreateCategory)
	app.Get("/api/categories/:id", handlers.GetCategory)
	app.Patch("/api/categories/:id", handlers.UpdateCategory)
	app.Delete("/api/categories/:id", handlers.DeleteCategory)

}
