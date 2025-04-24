package interfaces

import (
	"bazar/internal/domain/entities"
	"bazar/internal/domain/requests"

	"github.com/gin-gonic/gin"
)

type NFTRepository interface {
	GetSalesNFT() (*[]entities.NFT, error)
	CreateNFT(nft *entities.NFT) (*entities.NFT, error)
	GetAllNFTs() (*[]entities.NFTResponse, error)
	GetNFTById(id string) (*entities.NFTResponse, error)
	SetTokenId(updateTokenReq *requests.UpdateTokenIdReq) (*entities.NFT, error)
	SetTokenPrice(updateTokenReq *requests.UpdateTokenPriceReq) (*entities.NFT, error)
	GetUserNFT(address string) (*[]entities.NFT, error)
	SetTokenSales(req *requests.UpdateTokenStatusReq) (*entities.NFT, error)
	UpdateProposedByTokenURI(id string, tokenURI string) (*entities.NFT, error)
	UpdateVotesByTokenURI(tokenURI string, amount string) (*entities.NFT, error)
	UpdateId(tokenUri string, id string) (*entities.NFT, error)
}

type NFTService interface {
	GetSalesNFT() (*[]entities.NFT, error)
	CreateNFT(imagePath string, nft *requests.CreateNFTRequest) (*entities.NFT, error)
	GetAllNFTs() (*[]entities.NFTResponse, error)
	GetNFTById(id string) (*entities.NFTResponse, error)
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
	GetProposedNFTs(c *gin.Context)
}
