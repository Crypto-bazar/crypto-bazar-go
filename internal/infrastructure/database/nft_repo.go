package database

import (
	"bazar/internal/domain/entities"
	"bazar/internal/domain/interfaces"
	"bazar/internal/domain/requests"
	"bazar/internal/domain/responses"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type NFTRepository struct {
	db *sqlx.DB
}

func (n *NFTRepository) AddFavouriteNFT(nftId string, ethAddress string) (*responses.FavouriteNFTsResponse, error) {
	tx, err := n.db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("error starting transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			if commitErr := tx.Commit(); commitErr != nil {
				err = fmt.Errorf("error committing transaction: %w", commitErr)
				tx.Rollback()
			}
		}
	}()

	// 1. Get user ID
	var userId int64
	if err = tx.Get(&userId, "SELECT id FROM users WHERE eth_address = $1", ethAddress); err != nil {
		return nil, fmt.Errorf("error getting user id: %w", err)
	}

	// 2. Check if NFT already in favourites
	var exists bool
	if err = tx.Get(&exists, `
        SELECT EXISTS(
            SELECT 1 FROM favourite_nfts 
            WHERE user_id = $1 AND nft_id = $2
        )`, userId, nftId); err != nil {
		return nil, fmt.Errorf("error checking if NFT exists in favorites: %w", err)
	}

	if exists {
		return nil, fmt.Errorf("NFT already in favourites")
	}

	// 3. Add to favourites (без проверки в таблице nfts)
	if _, err = tx.Exec(`
        INSERT INTO favourite_nfts (user_id, nft_id) 
        VALUES ($1, $2)`, userId, nftId); err != nil {
		return nil, fmt.Errorf("error adding to favourites: %w", err)
	}

	// 4. Get updated favourites list
	query := `
        SELECT f.nft_id
        FROM favourite_nfts f
        WHERE f.user_id = $1
        ORDER BY f.nft_id DESC
    `

	var nftItems []responses.FavouriteNFTResponse
	if err = tx.Select(&nftItems, query, userId); err != nil {
		return nil, fmt.Errorf("error getting updated favourites: %w", err)
	}

	nftIds := make([]int64, len(nftItems))
	for i, item := range nftItems {
		nftIds[i] = item.NftID
	}

	return &responses.FavouriteNFTsResponse{
		NftIds: nftIds,
	}, nil
}

func (n *NFTRepository) GetFavouriteNFTS(ethAddress string) (*responses.FavouriteNFTsResponse, error) {
	log.Println("GetFavouriteNFTS called with ethAddress:", ethAddress)

	query := `
        WITH owner AS (
            SELECT id 
            FROM users
            WHERE eth_address = $1
        )
        SELECT 
            f.nft_id
        FROM favourite_nfts f
        WHERE f.user_id = (SELECT id FROM owner)
        ORDER BY f.nft_id DESC
    `

	var nftItems []responses.FavouriteNFTResponse
	err := n.db.Select(&nftItems, query, ethAddress)
	if err != nil {
		return nil, fmt.Errorf("error querying favourite NFTs: %w", err)
	}

	// Преобразуем в массив ID
	nftIds := make([]int64, len(nftItems))
	for i, item := range nftItems {
		nftIds[i] = item.NftID
	}

	return &responses.FavouriteNFTsResponse{
		NftIds: nftIds,
	}, nil
}

func (n *NFTRepository) RemoveFavouriteNFT(nftId string, ethAddress string) (*responses.FavouriteNFTsResponse, error) {
	tx, err := n.db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("error starting transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			if commitErr := tx.Commit(); commitErr != nil {
				err = fmt.Errorf("error committing transaction: %w", commitErr)
				tx.Rollback()
			}
		}
	}()

	// 1. Get user ID
	var userId int64
	if err = tx.Get(&userId, "SELECT id FROM users WHERE eth_address = $1", ethAddress); err != nil {
		return nil, fmt.Errorf("error getting user id: %w", err)
	}

	// 2. Check if NFT exists in favorites
	var exists bool
	if err = tx.Get(&exists, `
        SELECT EXISTS(
            SELECT 1 FROM favourite_nfts 
            WHERE user_id = $1 AND nft_id = $2
        )`, userId, nftId); err != nil {
		return nil, fmt.Errorf("error checking if NFT exists in favorites: %w", err)
	}

	if !exists {
		return nil, fmt.Errorf("NFT was not in favourites")
	}

	// 3. Remove from favorites
	result, err := tx.Exec(`
        DELETE FROM favourite_nfts 
        WHERE user_id = $1 AND nft_id = $2`, userId, nftId)
	if err != nil {
		return nil, fmt.Errorf("error removing from favourites: %w", err)
	}

	// 4. Verify deletion
	_, err = result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("error checking rows affected: %w", err)
	}

	// 5. Get updated favorites list
	query := `
        SELECT f.nft_id
        FROM favourite_nfts f
        WHERE f.user_id = $1
        ORDER BY f.nft_id DESC
    `

	var nftItems []responses.FavouriteNFTResponse
	if err = tx.Select(&nftItems, query, userId); err != nil {
		return nil, fmt.Errorf("error getting updated favourites: %w", err)
	}

	// Convert to IDs array
	nftIds := make([]int64, len(nftItems))
	for i, item := range nftItems {
		nftIds[i] = item.NftID
	}

	return &responses.FavouriteNFTsResponse{
		NftIds: nftIds,
	}, nil
}

func NewNFTRepository(db *sqlx.DB) interfaces.NFTRepository {
	return &NFTRepository{db: db}
}

func (n *NFTRepository) UpdateSoldByTokenId(tokenId, buyer string) (*entities.NFT, error) {
	query := `
			WITH owner AS (
				SELECT id
				FROM users
				WHERE eth_address = $1
			)
			UPDATE nfts
			SET owner_id = (SELECT id FROM owner), in_sales = false
			WHERE token_id = $2;
			`
	log.Printf("TokenID: %s, Buyer: %s", tokenId, buyer)
	_, err := n.db.Exec(query, buyer, tokenId)
	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error updating NFT sale status: %w", err)
	}

	var nft entities.NFT
	query = "SELECT * FROM nfts WHERE token_id = $1"
	err = n.db.Get(&nft, query, tokenId)
	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error getting NFT: %w", err)
	}

	return &nft, nil
}

func (n *NFTRepository) UpdateSaleByTokenId(tokenId string, price string) (*entities.NFT, error) {
	query := "UPDATE nfts SET in_sales = true, price = $1 WHERE token_id = $2"
	_, err := n.db.Exec(query, price, tokenId)
	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error updating sale price: %w", err)
	}

	var nft entities.NFT
	query = "SELECT * FROM nfts WHERE token_id = $1"
	err = n.db.Get(&nft, query, tokenId)
	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error getting NFT: %w", err)
	}

	return &nft, nil
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

func (n *NFTRepository) GetAllNFTs() (*[]entities.NFTResponse, error) {
	var nfts []entities.NFTResponse

	query := `
		SELECT
			n.id,
			n.token_id,
			n.proposal_id,
			n.token_uri,
			n.name,
			n.description,
			n.price::TEXT,
			u.eth_address AS owner,
			n.image_path,
			n.in_sales,
			n.proposed,
			n.votes_amount::TEXT
		FROM nfts n
		JOIN users u ON n.owner_id = u.id
		ORDER BY n.id
	`

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
		WHERE in_sales = true
		ORDER BY id`

	err := n.db.Select(&nfts, query)
	if err != nil {
		return nil, fmt.Errorf("error getting sales NFTs: %w", err)
	}

	if nfts == nil {
		nfts = []entities.NFT{}
	}

	return &nfts, nil
}
func (n *NFTRepository) GetNFTById(id string) (*entities.NFTResponse, error) {
	var nft entities.NFTResponse

	query := `
        SELECT
            n.id,
            n.token_id,
            n.proposal_id,
            n.token_uri,
            n.name,
            n.description,
            n.price::TEXT,
            u.eth_address AS owner,
            n.image_path,
            n.in_sales,
            n.proposed,
            n.votes_amount::TEXT
        FROM nfts n
        JOIN users u ON n.owner_id = u.id
        WHERE n.id = $1
        LIMIT 1
    `

	err := n.db.Get(&nft, query, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("NFT with ID %s not found", id)
		}
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error getting NFT by ID: %w", err)
	}

	return &nft, nil
}
