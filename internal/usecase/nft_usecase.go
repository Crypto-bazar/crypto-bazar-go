package usecase

import (
	"bazar/internal/domain"
	"bazar/internal/domain/entities"
	"bazar/internal/domain/interfaces"
	"bazar/internal/domain/requests"
	"bazar/internal/domain/responses"
	"bazar/pkg/utils"
	"fmt"
)

type NFTService struct {
	nftRepo  interfaces.NFTRepository
	userRepo interfaces.UserRepository
}

func (n *NFTService) AddFavouriteNFT(nftId string, ethAddress string) (*responses.FavouriteNFTsResponse, error) {
	return n.nftRepo.AddFavouriteNFT(nftId, ethAddress)
}

func (n *NFTService) GetFavouriteNFTS(ethAddress string) (*responses.FavouriteNFTsResponse, error) {
	return n.nftRepo.GetFavouriteNFTS(ethAddress)
}

func (n *NFTService) RemoveFavouriteNFT(nftId string, ethAddress string) (*responses.FavouriteNFTsResponse, error) {
	return n.nftRepo.RemoveFavouriteNFT(nftId, ethAddress)
}

func NewNFTService(nftRepo interfaces.NFTRepository, userRepo interfaces.UserRepository) interfaces.NFTService {
	return &NFTService{nftRepo: nftRepo, userRepo: userRepo}
}

func (n *NFTService) SetNFTProposed(event *domain.NFTProposedEvent) (*entities.NFT, error) {
	nft, err := n.nftRepo.UpdateProposedByTokenURI(event.ProposalID, event.TokenURI)
	if err != nil {
		return nil, fmt.Errorf("error with updating proposed status: %w", err)
	}

	return nft, nil
}

func (n *NFTService) SetTokenSales(req *requests.UpdateTokenStatusReq) (*entities.NFT, error) {
	return n.nftRepo.SetTokenSales(req)
}

func (n *NFTService) GetUserNFT(address string) (*[]entities.NFT, error) {
	return n.nftRepo.GetUserNFT(address)
}

func (n *NFTService) SetTokenPrice(updateTokenReq *requests.UpdateTokenPriceReq) (*entities.NFT, error) {
	return n.nftRepo.SetTokenPrice(updateTokenReq)
}

func (n *NFTService) GetSalesNFT() (*[]entities.NFT, error) {
	return n.nftRepo.GetSalesNFT()
}

func (n *NFTService) SetTokenId(updateTokenReq *requests.UpdateTokenIdReq) (*entities.NFT, error) {
	return n.nftRepo.SetTokenId(updateTokenReq)
}

func (n *NFTService) CreateNFT(imagePath string, nft *requests.CreateNFTRequest) (*entities.NFT, error) {
	owner, err := n.userRepo.GetUserByAddress(nft.OwnerAddress)
	tokenUri := utils.GenerateTokenURI("http://localhost:8080")

	if err != nil {
		return nil, fmt.Errorf("error, user not found %w", err)
	}

	nftModel := entities.NFT{
		Price:       "0",
		TokenID:     0,
		Name:        nft.Name,
		Description: nft.Description,
		Owner:       owner.ID,
		ImagePath:   imagePath,
		TokenURI:    tokenUri,
	}
	return n.nftRepo.CreateNFT(&nftModel)
}

func (n *NFTService) GetAllNFTs() (*[]entities.NFTResponse, error) {
	return n.nftRepo.GetAllNFTs()
}

func (n *NFTService) GetNFTById(id string) (*entities.NFTResponse, error) {
	return n.nftRepo.GetNFTById(id)
}
