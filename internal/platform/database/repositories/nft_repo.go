package repositories

import (
	"bazar/internal/domain"
	"bazar/internal/domain/dto"
	"bazar/internal/domain/models"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type NFTRepository struct {
	db *sqlx.DB
}

func (n *NFTRepository) SetTokenAddress(updateTokenReq *dto.UpdateTokenAddressReq) (*models.NFT, error) {
	var nft models.NFT

	updateQuery := "UPDATE nfts SET token_id = :token_id WHERE id = :id"
	_, err := n.db.NamedExec(updateQuery, map[string]interface{}{
		"token_id": updateTokenReq.ContractAddress,
		"id":       updateTokenReq.Id,
	})

	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error updating NFT: %w", err)
	}

	err = n.db.Get(&nft, "SELECT * FROM nfts WHERE id = $1", updateTokenReq.Id)

	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error updating NFT: %w", err)
	}

	return &nft, nil
}

func (n *NFTRepository) CreateNFT(nft *models.NFT) error {
	query := "INSERT INTO nfts (token_id, name, symbol, description, price, owner_id, image_path) VALUES (:token_id, :name, :symbol, :description, :price, :owner_id, :image_path)"
	_, err := n.db.NamedExec(query, &nft)

	if err != nil {
		log.Printf("DB error: %v", err)
		return fmt.Errorf("error creating NFT: %w", err)
	}

	return nil
}

func (n *NFTRepository) GetAllNFTs() (*[]models.NFT, error) {
	var nfts []models.NFT
	query := "SELECT * FROM nfts"
	err := n.db.Select(&nfts, query)

	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error getting all NFTs: %w", err)
	}

	return &nfts, nil
}

func (n *NFTRepository) GetNFTById(id string) (*models.NFT, error) {
	var nft models.NFT
	query := "SELECT * FROM nfts WHERE id = ?"
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
