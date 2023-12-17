# Product API

## Description
The Product API is a Go-based RESTful API developed using the Fiber framework. It provides a comprehensive suite of functionalities for managing products, users, categories, and authentication. This API is designed to be fast, efficient, and easy to integrate with modern web applications.

## Setup Project

To set up the project on your local machine, follow these steps:

1. **Clone the repository:**
   ```bash
   git clone https://github.com/santoadji21/santoadji21-go-fiber-product-api
   ```

2. **Navigate to the project directory:**
   ```bash
   cd santoadji21-go-fiber-product-api
   ```

3. **Run the application:**
   ```bash
   go run ./cmd
   ```

## Available Routes

### User Routes
- `POST /api/users`: Create a new user
- `GET /api/users`: Retrieve all users
- `GET /api/users/:id`: Retrieve a user by ID
- `PATCH /api/users/:id`: Update a user by ID
- `DELETE /api/users/:id`: Delete a user by ID

### Auth Routes
- `POST /api/login`: User login

### Product Routes
- `POST /api/product`: Create a new product (Protected)
- `GET /api/products`: Retrieve all products
- `GET /api/product/:id`: Retrieve a product by ID
- `PATCH /api/product/:id`: Update a product by ID (Protected)
- `DELETE /api/product/:id`: Delete a product by ID (Protected)

### Category Routes
- `POST /api/category`: Create a new category (Protected)
- `GET /api/categories`: Retrieve all categories
- `GET /api/category/:id`: Retrieve a category by ID
- `PATCH /api/category/:id`: Update a category by ID (Protected)
- `DELETE /api/category/:id`: Delete a category by ID (Protected)

## Documentation

For a detailed description of the API endpoints, including request and response formats, visit the Swagger documentation:

- [Swagger API Documentation](http://localhost:3000/swagger/index.html) (Accessible when running the API locally)