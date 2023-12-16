package models

// Product represents the product model
type Product struct {
    Model
    Name        string
    Description string
    Qty         int
    Price       float64
    Discount    float64
    CategoryID  uint
}
