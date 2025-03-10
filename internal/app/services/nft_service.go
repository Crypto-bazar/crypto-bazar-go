package services

import (
	"bazar/internal/domain"
	"bazar/internal/domain/dto"
	"bazar/internal/domain/models"
	"fmt"
)

type NFTService struct {
	nftRepo  domain.NFTRepository
	userRepo domain.UserRepository
}

// CreateNFT implements domain.NFTService.
func (n *NFTService) CreateNFT(imagePath string, nft *dto.CreateNFTRequest) error {
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

// GetAllNFTs implements domain.NFTService.
func (n *NFTService) GetAllNFTs() (*[]models.NFT, error) {
	return n.nftRepo.GetAllNFTs()
}

// GetNFTById implements domain.NFTService.
func (n *NFTService) GetNFTById(id string) (*models.NFT, error) {
	return n.nftRepo.GetNFTById(id)
}

func NewNFTService(nftRepo domain.NFTRepository, userRepo domain.UserRepository) domain.NFTService {
	return &NFTService{nftRepo: nftRepo, userRepo: userRepo}
}
