package utils

import (
	"github.com/santoadji21/santoadji21-go-fiber-product-api/pkg/models"
)

func UserResponse(user models.User) models.User {
	return models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}

func UsersResponse(users []models.User) []models.User {
	var usersResponse []models.User
	for _, user := range users {
		userResponse := UserResponse(user)
		usersResponse = append(usersResponse, userResponse)
	}
	return usersResponse
}
