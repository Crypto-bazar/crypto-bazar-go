package services

import (
	"bazar/internal/domain"
	"bazar/internal/domain/dto"
	"bazar/internal/domain/models"
	crypto "bazar/internal/pkg"
	"fmt"
)

type UserService struct {
	repo domain.UserRepository
}

// CheckUserExists implements domain.UserService.
func (u *UserService) CheckUserExists(address *dto.CheckUserRequest) (*bool, error) {
	return u.repo.CheckUserExists(address.Address)
}

// CreateUser implements domain.UserService.
func (u *UserService) CreateUser(user *dto.CreateUserRequest) error {
	verifyResult := crypto.VerifySignature(user.EthAddress, user.Message, user.Signature)

	if !verifyResult {
		return fmt.Errorf("signature verification failed")
	}

	newUser := models.User{
		ID:         0,
		EthAddress: user.EthAddress,
	}

	return u.repo.CreateUser(&newUser)
}

// GetAllUsers implements domain.UserService.
func (u *UserService) GetAllUsers() (*[]models.User, error) {
	return u.repo.GetAllUsers()
}

// GetUserById implements domain.UserService.
func (u *UserService) GetUserById(id string) (*models.User, error) {
	return u.repo.GetUserById(id)
}

func NewUserService(repo domain.UserRepository) domain.UserService {
	return &UserService{repo: repo}
}
