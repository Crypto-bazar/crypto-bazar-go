package interfaces

import (
	"bazar/internal/domain/entities"
	"bazar/internal/domain/requests"
	"bazar/internal/domain/responses"

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
	UpdateSaleByTokenId(tokenId string, price string) (*entities.NFT, error)
	UpdateSoldByTokenId(tokenId, buyer string) (*entities.NFT, error)
	AddFavouriteNFT(nftId, ethAddress string) (*responses.FavouriteNFTsResponse, error)
	RemoveFavouriteNFT(nftId, ethAddress string) (*responses.FavouriteNFTsResponse, error)
	GetFavouriteNFTS(ethAddress string) (*responses.FavouriteNFTsResponse, error)
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
	AddFavouriteNFT(nftId, ethAddress string) (*responses.FavouriteNFTsResponse, error)
	RemoveFavouriteNFT(nftId, ethAddress string) (*responses.FavouriteNFTsResponse, error)
	GetFavouriteNFTS(ethAddress string) (*responses.FavouriteNFTsResponse, error)
}
type NFTHandler interface {
	CreateNFT(c *gin.Context)
	GetAllNFTs(c *gin.Context)
	GetNFTById(c *gin.Context)
	SetTokenAddress(c *gin.Context)
	GetSalesNFT(c *gin.Context)
	GetUserNFT(c *gin.Context)
	GetProposedNFTs(c *gin.Context)
	AddFavouriteNFT(c *gin.Context)
	RemoveFavouriteNFT(c *gin.Context)
	GetFavouriteNFTS(c *gin.Context)
}
