package models

// User represents the user model
type User struct {
    Model
    FirstName string `json:"firstName" gorm:"column:first_name"`
    LastName  string `json:"lastName" gorm:"column:last_name"`
    Email     string `json:"email" gorm:"unique;column:email"`
    Password  string `json:"-" gorm:"->;<-:create;<hidden>"`
}

