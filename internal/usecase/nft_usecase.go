package usecase

import (
	"bazar/internal/domain"
	"bazar/internal/domain/entities"
	"bazar/internal/domain/requests"
	"bazar/pkg/utils"
	"fmt"
)

type NFTService struct {
	nftRepo  domain.NFTRepository
	userRepo domain.UserRepository
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
	tokenUri := utils.GenerateTokenURI("http://localhost:8080", imagePath)

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

func (n *NFTService) GetAllNFTs() (*[]entities.NFT, error) {
	return n.nftRepo.GetAllNFTs()
}

func (n *NFTService) GetNFTById(id string) (*entities.NFT, error) {
	return n.nftRepo.GetNFTById(id)
}

func NewNFTService(nftRepo domain.NFTRepository, userRepo domain.UserRepository) domain.NFTService {
	return &NFTService{nftRepo: nftRepo, userRepo: userRepo}
}
