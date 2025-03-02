package services

import (
	"bazar/internal/domain"
	"bazar/internal/domain/models"
)

type NFTService struct {
	repo domain.NFTRepository
}

// CreateNFT implements domain.NFTService.
func (n *NFTService) CreateNFT(nft *models.NFT) error {
	return n.repo.CreateNFT(nft)
}

// GetAllNFTs implements domain.NFTService.
func (n *NFTService) GetAllNFTs() (*[]models.NFT, error) {
	return n.repo.GetAllNFTs()
}

// GetNFTById implements domain.NFTService.
func (n *NFTService) GetNFTById(id string) (*models.NFT, error) {
	return n.repo.GetNFTById(id)
}

func NewUserService(repo domain.NFTRepository) domain.NFTService {
	return &NFTService{repo: repo}
}
