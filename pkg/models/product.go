package models

// Product represents the product model
type Product struct {
	Model
	Name        string  `json:"name" gorm:"unique;column:name"`
	Description string  `json:"description"`
	Qty         int     `json:"qty"`
	Price       float64 `json:"price"`
	Discount    float64 `json:"discount"`
	CategoryID  uint    `json:"category_id"`
}
