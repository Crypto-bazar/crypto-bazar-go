package database

import (
	"bazar/internal/domain/entities"
	"bazar/internal/domain/interfaces"
	"bazar/internal/domain/requests"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type NFTRepository struct {
	db *sqlx.DB
}

func NewNFTRepository(db *sqlx.DB) interfaces.NFTRepository {
	return &NFTRepository{db: db}
}

func (n *NFTRepository) UpdateId(tokenUri string, id string) (*entities.NFT, error) {
	query := "UPDATE nfts SET token_id = $1 WHERE token_uri = $2"
	_, err := n.db.Exec(query, id, tokenUri)
	if err != nil {
		log.Printf("Db error: %v", err)
		return nil, fmt.Errorf("error updating token id: %w", err)
	}

	var nft entities.NFT
	query = "SELECT * FROM nfts WHERE token_uri = $1"
	err = n.db.Get(&nft, query, tokenUri)
	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("err	or getting NFT: %w", err)
	}

	return &nft, nil
}

func (n *NFTRepository) UpdateVotesByTokenURI(tokenURI string, amount string) (*entities.NFT, error) {
	query := "UPDATE nfts SET votes_amount = votes_amount + $1 WHERE token_uri = $2"
	_, err := n.db.Exec(query, amount, tokenURI)
	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error updating votes: %w", err)
	}

	var nft entities.NFT
	query = "SELECT * FROM nfts WHERE token_uri = $1"
	err = n.db.Get(&nft, query, tokenURI)
	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("err	or getting NFT: %w", err)
	}

	return &nft, nil
}

func (n *NFTRepository) UpdateProposedByTokenURI(id string, tokenURI string) (*entities.NFT, error) {
	query := "UPDATE nfts SET proposed = true, proposal_id = $1 WHERE token_uri = $2"
	_, err := n.db.Exec(query, id, tokenURI)
	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error updating proposed status: %w", err)
	}

	query = "SELECT * FROM nfts WHERE token_uri = $1"
	var nft entities.NFT
	err = n.db.Get(&nft, query, tokenURI)
	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error getting NFT: %w", err)
	}

	return &nft, nil
}

func (n *NFTRepository) SetTokenSales(req *requests.UpdateTokenStatusReq) (*entities.NFT, error) {
	var nft entities.NFT
	query := `
		UPDATE nfts SET in_sales = :status
		WHERE token_id = :token_id`

	res, err := n.db.NamedExec(query, req)

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error updating NFT: %w", err)
	}

	query = `
		SELECT * FROM nfts
		WHERE token_id = $1`

	err = n.db.Get(&nft, query, req.TokenID)
	if err != nil {
		return nil, fmt.Errorf("error getting nft: %w", err)
	}

	return &nft, nil
}

func (n *NFTRepository) GetUserNFT(address string) (*[]entities.NFT, error) {
	var nft []entities.NFT
	query := `
		SELECT nfts.* FROM nfts 
		JOIN "users" ON nfts.owner_id = "users".id 
		WHERE "users".eth_address = $1
		ORDER BY nfts.id`

	err := n.db.Select(&nft, query, address)
	if err != nil {
		return nil, fmt.Errorf("error getting nft: %w", err)
	}

	return &nft, nil
}

func (n *NFTRepository) SetTokenPrice(updateTokenReq *requests.UpdateTokenPriceReq) (*entities.NFT, error) {
	query := `
		UPDATE nfts SET price = :price 
		WHERE token_id = :token_id`

	res, err := n.db.NamedExec(query, updateTokenReq)
	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error updating NFT: %w", err)
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return nil, fmt.Errorf("no rows updated")
	}

	var nft entities.NFT
	query = `
		SELECT * FROM nfts 
		WHERE token_id = $1`

	err = n.db.Get(&nft, query, updateTokenReq.TokenId)
	if err != nil {
		log.Printf("Error fetching updated NFT: %v", err)
		return nil, fmt.Errorf("error fetching updated NFT: %w", err)
	}

	return &nft, nil
}

func (n *NFTRepository) SetTokenId(updateTokenReq *requests.UpdateTokenIdReq) (*entities.NFT, error) {
	var nft entities.NFT
	updateQuery := `
		UPDATE nfts 
		SET token_id = :token_id 
		WHERE token_uri = :token_uri 
		RETURNING id, token_id, token_uri, name, description, price, owner_id, image_path`

	_, err := n.db.NamedExec(updateQuery, map[string]any{
		"token_id":  updateTokenReq.TokenId.Int64(),
		"token_uri": updateTokenReq.TokenURI,
	})

	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error updating NFT: %w", err)
	}

	query := `
		SELECT * FROM nfts 
		WHERE token_id = $1`

	err = n.db.Get(&nft, query, updateTokenReq.TokenId.Int64())

	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error getting NFT: %w", err)
	}

	return &nft, nil
}

func (n *NFTRepository) CreateNFT(nft *entities.NFT) (*entities.NFT, error) {
	var createdNFT entities.NFT
	query := `
        INSERT INTO nfts (token_id, token_uri, name, description, price, owner_id, image_path)
        VALUES (:token_id, :token_uri, :name, :description, :price, :owner_id, :image_path)
        RETURNING id, token_id, token_uri, name, description, price, owner_id, image_path
    `

	rows, err := n.db.NamedQuery(query, nft)
	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error creating NFT: %w", err)
	}

	defer func(rows *sqlx.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows: ", err)
		}
	}(rows)

	if rows.Next() {
		err = rows.StructScan(nft)
		if err != nil {
			log.Printf("Error scanning NFT: %v", err)
			return nil, fmt.Errorf("error scanning NFT: %w", err)
		}
	} else {
		return nil, fmt.Errorf("no rows returned after insert")
	}

	query = `SELECT * FROM nfts WHERE token_id = $1`

	err = n.db.Get(&createdNFT, query, nft.TokenID)
	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error getting NFT: %w", err)
	}

	return &createdNFT, nil
}

func (n *NFTRepository) GetAllNFTs() (*[]entities.NFT, error) {
	var nfts []entities.NFT
	query := `
			SELECT * FROM nfts
			ORDER BY id`
	err := n.db.Select(&nfts, query)

	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error getting all NFTs: %w", err)
	}

	return &nfts, nil
}

func (n *NFTRepository) GetSalesNFT() (*[]entities.NFT, error) {
	var nfts []entities.NFT
	query := `
		SELECT * FROM nfts 
		WHERE price > 0`

	err := n.db.Select(&nfts, query)
	if err != nil {
		return nil, fmt.Errorf("error getting sales NFTs: %w", err)
	}

	return &nfts, nil
}

func (n *NFTRepository) GetNFTById(id string) (*entities.NFT, error) {
	var nft entities.NFT
	query := `
		SELECT * FROM nfts 
		WHERE id = $1`

	err := n.db.Get(&nft, query, id)

	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error getting NFT by ID: %w", err)
	}

	return &nft, nil
}
