package routes

import (
	"go-product-api/pkg/handlers"

	"github.com/gofiber/fiber/v2"
)

func AppRoutes(app *fiber.App) {
    // User routes
    app.Post("/api/users", handlers.CreateUser)
	app.Get("/api/users", handlers.GetAllUsers)
    app.Get("/api/users/:id", handlers.GetUser)
    app.Patch("/api/users/:id", handlers.UpdateUser)
    app.Delete("/api/users/:id", handlers.DeleteUser)

    // Product routes
    app.Post("/api/products", handlers.CreateProduct)
    app.Get("/api/products/:id", handlers.GetProduct)
    app.Patch("/api/products/:id", handlers.UpdateProduct)
    app.Delete("/api/products/:id", handlers.DeleteProduct)

    // Category routes
    app.Post("/api/categories", handlers.CreateCategory)
    app.Get("/api/categories/:id", handlers.GetCategory)
    app.Patch("/api/categories/:id", handlers.UpdateCategory)
    app.Delete("/api/categories/:id", handlers.DeleteCategory)

}
