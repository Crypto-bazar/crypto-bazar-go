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

func (n *NFTService) SetTokenId(updateTokenReq *requests.UpdateTokenIdReq) (*entities.NFT, error) {
	return n.nftRepo.SetTokenId(updateTokenReq)
}

func (n *NFTService) CreateNFT(imagePath string, nft *requests.CreateNFTRequest) error {
	owner, err := n.userRepo.GetUserByAddress(nft.OwnerAddress)
	tokenUri := utils.GenerateTokenURI("http://localhost:8080", imagePath)

	if err != nil {
		return fmt.Errorf("error, user not found %w", err)
	}

	nftModel := entities.NFT{
		TokenID:     0,
		Name:        nft.Name,
		Description: nft.Description,
		Price:       nft.Price,
		Owner:       owner.ID,
		ID:          0,
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
