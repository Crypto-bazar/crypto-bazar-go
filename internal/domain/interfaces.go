package domain

import (
	"bazar/internal/domain/models"

	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetAllUsers() (*[]models.User, error)
	GetUserById(id string) (*models.User, error)
}

type UserService interface {
	CreateUser(user *models.User) error
	GetAllUsers() (*[]models.User, error)
	GetUserById(id string) (*models.User, error)
}

type UserHandler interface {
	CreateUser(c *gin.Context)
	GetAllUsers(c *gin.Context)
	GetUserById(c *gin.Context)
}
