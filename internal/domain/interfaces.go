package domain

import "bazar/internal/domain/models"

type UserRepository interface {
	CreateUser(user *models.User) error
	GetAllUsers() (*[]models.User, error)
	GetUserById(id uint) (*models.User, error)
}

type UserService interface {
	CreateUser(user *models.User) error
	GetAllUsers() (*[]models.User, error)
	GetUserById(id uint) (*models.User, error)
}
