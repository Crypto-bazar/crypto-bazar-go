package usecase

import (
	"bazar/internal/domain/entities"
	"bazar/internal/domain/interfaces"
	"bazar/internal/domain/requests"
	"bazar/pkg/utils"
	"fmt"
)

type UserService struct {
	repo interfaces.UserRepository
}

func (u *UserService) UpdateAvatarURL(ethAddress, avatarURL string) (*entities.User, error) {
	return u.repo.UpdateAvatarURL(ethAddress, avatarURL)
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

	newUser := entities.User{
		EthAddress: user.EthAddress,
	}

	return u.repo.CreateUser(&newUser)
}

func (u *UserService) GetAllUsers() (*[]entities.User, error) {
	return u.repo.GetAllUsers()
}

func (u *UserService) GetUserByAddress(id string) (*entities.User, error) {
	return u.repo.GetUserByAddress(id)
}

func NewUserService(repo interfaces.UserRepository) interfaces.UserService {
	return &UserService{repo: repo}
}
