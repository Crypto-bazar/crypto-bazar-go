package database

import (
	"bazar/internal/domain"
	"bazar/internal/domain/entities"
	"bazar/internal/domain/requests"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type NFTRepository struct {
	db *sqlx.DB
}

func (n *NFTRepository) SetTokenId(updateTokenReq *requests.UpdateTokenIdReq) (*entities.NFT, error) {
	var nft entities.NFT
	updateQuery := "UPDATE nfts SET token_id = :token_id WHERE token_uri = :token_uri"
	_, err := n.db.NamedExec(updateQuery, map[string]any{
		"token_id":  updateTokenReq.TokenId,
		"token_uri": updateTokenReq.TokenURI,
	})

	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error updating NFT: %w", err)
	}

	err = n.db.Get(&nft, "SELECT * FROM nfts WHERE token_id = $1", updateTokenReq.TokenId)

	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error getting NFT: %w", err)
	}

	return &nft, nil
}

func (n *NFTRepository) CreateNFT(nft *entities.NFT) error {
	query := "INSERT INTO nfts (token_id, token_uri, name, description, price, owner_id, image_path) VALUES (:token_id, :token_uri, :name, :description, :price, :owner_id, :image_path)"
	_, err := n.db.NamedExec(query, &nft)

	if err != nil {
		log.Printf("DB error: %v", err)
		return fmt.Errorf("error creating NFT: %w", err)
	}

	return nil
}

func (n *NFTRepository) GetAllNFTs() (*[]entities.NFT, error) {
	var nfts []entities.NFT
	query := "SELECT * FROM nfts"
	err := n.db.Select(&nfts, query)

	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error getting all NFTs: %w", err)
	}

	return &nfts, nil
}

func (n *NFTRepository) GetNFTById(id string) (*entities.NFT, error) {
	var nft entities.NFT
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
