package repositories

import (
	"bazar/internal/domain"
	"bazar/internal/domain/models"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type NFTRepository struct {
	db *sqlx.DB
}

// CreateNFT implements domain.NFTRepository.
func (n *NFTRepository) CreateNFT(nft *models.NFT) error {
	fmt.Printf("%+v\n", nft)
	query := "INSERT INTO nft (token_id, owner_id) VALUES (:token_id, :owner_id)"
	_, err := n.db.NamedExec(query, &nft)

	if err != nil {
		log.Printf("DB error: %v", err)
		return fmt.Errorf("error creating NFT: %w", err)
	}

	return nil
}

// GetAllNFTs implements domain.NFTRepository.
func (n *NFTRepository) GetAllNFTs() (*[]models.NFT, error) {
	var nfts []models.NFT
	query := "SELECT * FROM nft"
	err := n.db.Get(&nfts, query)

	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error getting all NFTs: %w", err)
	}

	return &nfts, nil
}

// GetNFTById implements domain.NFTRepository.
func (n *NFTRepository) GetNFTById(id string) (*models.NFT, error) {
	var nft models.NFT
	query := "SELECT * FROM nft WHERE id = ?"
	err := n.db.Get(&nft, query, id)

	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error getting NFT by ID: %w", err)
	}

	return &nft, nil
}

func NewNFTRepository(db *sqlx.DB) domain.NFTRepository {
	return &NFTRepository{db: db}
}
