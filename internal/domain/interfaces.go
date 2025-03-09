package domain

import (
	"bazar/internal/domain/dto"
	"bazar/internal/domain/models"

	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetAllUsers() (*[]models.User, error)
	GetUserById(id string) (*models.User, error)
	GetUserByAddress(address string) (*models.User, error)
	CheckUserExists(address string) (*bool, error)
}

type UserService interface {
	CreateUser(user *dto.CreateUserRequest) error
	GetAllUsers() (*[]models.User, error)
	GetUserById(id string) (*models.User, error)
	CheckUserExists(address *dto.CheckUserRequest) (*bool, error)
}

type UserHandler interface {
	CreateUser(c *gin.Context)
	GetAllUsers(c *gin.Context)
	GetUserById(c *gin.Context)
	CheckUserExists(c *gin.Context)
}

type NFTRepository interface {
	CreateNFT(nft *models.NFT) error
	GetAllNFTs() (*[]models.NFT, error)
	GetNFTById(id string) (*models.NFT, error)
}

type NFTService interface {
	CreateNFT(imagePath string, nft *dto.CreateNFTRequest) error
	GetAllNFTs() (*[]models.NFT, error)
	GetNFTById(id string) (*models.NFT, error)
}

type NFTHandler interface {
	CreateNFT(c *gin.Context)
	GetAllNFTs(c *gin.Context)
	GetNFTById(c *gin.Context)
}
