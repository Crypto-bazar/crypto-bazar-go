package domain

import (
	"bazar/internal/domain/entities"
	"bazar/internal/domain/requests"
	"context"

	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	CreateUser(user *entities.User) error
	GetAllUsers() (*[]entities.User, error)
	GetUserById(id string) (*entities.User, error)
	GetUserByAddress(address string) (*entities.User, error)
	CheckUserExists(address string) (*bool, error)
}

type UserService interface {
	CreateUser(user *requests.CreateUserRequest) error
	GetAllUsers() (*[]entities.User, error)
	GetUserById(id string) (*entities.User, error)
	CheckUserExists(address *requests.CheckUserRequest) (*bool, error)
}

type UserHandler interface {
	CreateUser(c *gin.Context)
	GetAllUsers(c *gin.Context)
	GetUserById(c *gin.Context)
	CheckUserExists(c *gin.Context)
}

type NFTRepository interface {
	CreateNFT(nft *entities.NFT) (*entities.NFT, error) 
	GetAllNFTs() (*[]entities.NFT, error)
	GetNFTById(id string) (*entities.NFT, error)
	SetTokenId(updateTokenReq *requests.UpdateTokenIdReq) (*entities.NFT, error)
}

type NFTService interface {
	CreateNFT(imagePath string, nft *requests.CreateNFTRequest) (*entities.NFT, error)
	GetAllNFTs() (*[]entities.NFT, error)
	GetNFTById(id string) (*entities.NFT, error)
	SetTokenId(updateTokenReq *requests.UpdateTokenIdReq) (*entities.NFT, error)
}

type NFTHandler interface {
	CreateNFT(c *gin.Context)
	GetAllNFTs(c *gin.Context)
	GetNFTById(c *gin.Context)
	SetTokenAddress(c *gin.Context)
}

type EventHandler interface {
	OnTokenMinted(ctx context.Context, event TokenMintedEvent) error
	OnTokenListedForSale(ctx context.Context, event TokenListedForSaleEvent) error
	OnTokenSold(ctx context.Context, event TokenSold) error
}
