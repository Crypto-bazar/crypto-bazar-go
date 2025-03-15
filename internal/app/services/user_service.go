package services

import (
	"bazar/internal/domain"
	"bazar/internal/domain/requests"
	"bazar/internal/domain/models"
	"bazar/internal/platform/utils"
	"fmt"
)

type UserService struct {
	repo domain.UserRepository
}

func (u *UserService) CheckUserExists(address *requests.CheckUserRequest) (*bool, error) {
	return u.repo.CheckUserExists(address.Address)
}

func (u *UserService) CreateUser(user *requests.CreateUserRequest) error {
	verifyResult, err := utils.VerifySignature(user.Message, user.Signature, user.EthAddress)

	if err != nil {
		return fmt.Errorf("signature error %w", err)
	}

	if !verifyResult {
		return fmt.Errorf("signature verification failed")
	}

	newUser := models.User{
		EthAddress: user.EthAddress,
	}

	return u.repo.CreateUser(&newUser)
}

func (u *UserService) GetAllUsers() (*[]models.User, error) {
	return u.repo.GetAllUsers()
}

func (u *UserService) GetUserById(id string) (*models.User, error) {
	return u.repo.GetUserById(id)
}

func NewUserService(repo domain.UserRepository) domain.UserService {
	return &UserService{repo: repo}
}
