package interfaces

import (
	"bazar/internal/domain/entities"
	"bazar/internal/domain/requests"

	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	CreateUser(user *entities.User) error
	GetAllUsers() (*[]entities.User, error)
	GetUserById(id string) (*entities.User, error)
	GetUserByAddress(address string) (*entities.User, error)
	CheckUserExists(address string) (*bool, error)
	UpdateAvatarURL(ethAddress, avatarURL string) (*entities.User, error)
}

type UserService interface {
	CreateUser(user *requests.CreateUserRequest) error
	GetAllUsers() (*[]entities.User, error)
	GetUserById(id string) (*entities.User, error)
	CheckUserExists(address *requests.CheckUserRequest) (*bool, error)
	UpdateAvatarURL(ethAddress, avatarURL string) (*entities.User, error)
}

type UserHandler interface {
	CreateUser(c *gin.Context)
	GetAllUsers(c *gin.Context)
	GetUserById(c *gin.Context)
	CheckUserExists(c *gin.Context)
	UploadAvatarHandler(c *gin.Context)
}
