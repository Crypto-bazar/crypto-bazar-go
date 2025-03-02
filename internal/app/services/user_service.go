package services

import (
	"bazar/internal/domain"
	"bazar/internal/domain/models"
)

type UserService struct {
	repo domain.UserRepository
}

// CreateUser implements domain.UserService.
func (u *UserService) CreateUser(user *models.User) error {
	return u.repo.CreateUser(user)
}

// GetAllUsers implements domain.UserService.
func (u *UserService) GetAllUsers() (*[]models.User, error) {
	return u.repo.GetAllUsers()
}

// GetUserById implements domain.UserService.
func (u *UserService) GetUserById(id uint) (*models.User, error) {
	return u.repo.GetUserById(id)
}

func NewUserRepository(repo domain.UserRepository) domain.UserService {
	return &UserService{repo: repo}
}
