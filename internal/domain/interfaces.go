package domain

import (
	"bazar/internal/domain/entities"
	"bazar/internal/domain/requests"
	"github.com/gin-gonic/gin"
)

type CommentRepository interface {
	CreateComment(comment *entities.Comment) (*entities.Comment, error)
	GetCommentById(id string) (*entities.Comment, error)
}

type CommentService interface {
	CreateComment(comment *entities.Comment) (*entities.Comment, error)
	GetCommentById(id string) (*entities.Comment, error)
}

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
	GetSalesNFT() (*[]entities.NFT, error)
	CreateNFT(nft *entities.NFT) (*entities.NFT, error)
	GetAllNFTs() (*[]entities.NFT, error)
	GetNFTById(id string) (*entities.NFT, error)
	SetTokenId(updateTokenReq *requests.UpdateTokenIdReq) (*entities.NFT, error)
	SetTokenPrice(updateTokenReq *requests.UpdateTokenPriceReq) (*entities.NFT, error)
	GetUserNFT(address string) (*[]entities.NFT, error)
	SetTokenSales(req *requests.UpdateTokenStatusReq) (*entities.NFT, error)
}

type NFTService interface {
	GetSalesNFT() (*[]entities.NFT, error)
	CreateNFT(imagePath string, nft *requests.CreateNFTRequest) (*entities.NFT, error)
	GetAllNFTs() (*[]entities.NFT, error)
	GetNFTById(id string) (*entities.NFT, error)
	SetTokenId(updateTokenReq *requests.UpdateTokenIdReq) (*entities.NFT, error)
	SetTokenPrice(updateTokenReq *requests.UpdateTokenPriceReq) (*entities.NFT, error)
	GetUserNFT(address string) (*[]entities.NFT, error)
	SetTokenSales(req *requests.UpdateTokenStatusReq) (*entities.NFT, error)
}
type NFTHandler interface {
	CreateNFT(c *gin.Context)
	GetAllNFTs(c *gin.Context)
	GetNFTById(c *gin.Context)
	SetTokenAddress(c *gin.Context)
	GetSalesNFT(c *gin.Context)
	GetUserNFT(c *gin.Context)
}

type CommentHandler interface {
	CreateComment(c *gin.Context)
	GetCommentById(c *gin.Context)
}

type EventHandler interface {
	OnTokenMinted(event *TokenMintedEvent) error
	OnTokenListedForSale(event *TokenListedForSaleEvent) error
	OnTokenSold(event TokenSold) error
}
