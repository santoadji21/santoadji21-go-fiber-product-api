package models

type Category struct {
	Model
	Name string `json:"name" gorm:"unique"`
}
