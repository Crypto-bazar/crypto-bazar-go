package services

import (
	"bazar/internal/domain"
	"bazar/internal/domain/models"
)

type NFTService struct {
	repo domain.UserRepository
}

// CreateNFT implements domain.NFTService.
func (n *NFTService) CreateNFT(nft *models.NFT) error {
	panic("unimplemented")
}

// GetAllNFTs implements domain.NFTService.
func (n *NFTService) GetAllNFTs() (*[]models.NFT, error) {
	panic("unimplemented")
}

// GetNFTById implements domain.NFTService.
func (n *NFTService) GetNFTById(id string) (*models.NFT, error) {
	panic("unimplemented")
}

func NewUserService(repo domain.UserRepository) domain.NFTService {
	return &NFTService{repo: repo}
}
