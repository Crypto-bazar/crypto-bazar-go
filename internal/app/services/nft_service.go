package services

import (
	"bazar/internal/domain"
	"bazar/internal/domain/requests"
	"bazar/internal/domain/models"
	"fmt"
)

type NFTService struct {
	nftRepo  domain.NFTRepository
	userRepo domain.UserRepository
}

func (n *NFTService) SetTokenAddress(updateTokenReq *requests.UpdateTokenAddressReq) (*models.NFT, error) {
	return n.nftRepo.SetTokenAddress(updateTokenReq)
}

func (n *NFTService) CreateNFT(imagePath string, nft *requests.CreateNFTRequest) error {
	owner, err := n.userRepo.GetUserByAddress(nft.OwnerAddress)

	if err != nil {
		return fmt.Errorf("error, user not found %w", err)
	}

	nftModel := models.NFT{
		TokenID:     nft.TokenID,
		Name:        nft.Name,
		Description: nft.Description,
		Price:       nft.Price,
		Owner:       owner.ID,
		ID:          0,
		ImagePath:   imagePath,
	}
	return n.nftRepo.CreateNFT(&nftModel)
}

func (n *NFTService) GetAllNFTs() (*[]models.NFT, error) {
	return n.nftRepo.GetAllNFTs()
}

func (n *NFTService) GetNFTById(id string) (*models.NFT, error) {
	return n.nftRepo.GetNFTById(id)
}

func NewNFTService(nftRepo domain.NFTRepository, userRepo domain.UserRepository) domain.NFTService {
	return &NFTService{nftRepo: nftRepo, userRepo: userRepo}
}
