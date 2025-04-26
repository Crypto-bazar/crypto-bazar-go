package database

import (
	"bazar/internal/domain/interfaces"
	"bazar/internal/domain/requests"
	"bazar/internal/domain/responses"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type CommentRepo struct {
	db *sqlx.DB
}

func (c *CommentRepo) CreateComment(comment *requests.CreateCommentReq) (*responses.CommentResponse, error) {
	if comment == nil {
		return nil, fmt.Errorf("nil comment request")
	}

	query := `
        WITH owner AS (
            SELECT id, eth_address
            FROM users
            WHERE eth_address = $1
        )
        INSERT INTO comments (nft_id, owner_id, content)
        VALUES ($2, (SELECT id FROM owner), $3)
        RETURNING 
            comments.id,
            comments.nft_id,
            (SELECT eth_address FROM owner) as owner_address,
            comments.content,
            comments.created_at;
    `

	var response responses.CommentResponse

	err := c.db.QueryRowx(query, comment.OwnerAddress, comment.TokenId, comment.Content).StructScan(&response)
	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error creating comment: %w", err)
	}

	return &response, nil
}

func (c CommentRepo) GetCommentsByTokenId(tokenId string) (*[]responses.CommentResponse, error) {
	var comments []responses.CommentResponse
	query := `
        SELECT 
            c.id,
            c.nft_id,
            u.eth_address as eth_address,
            c.content,
            c.created_at
        FROM comments c
        JOIN users u ON c.owner_id = u.id
        WHERE c.nft_id = $1
        ORDER BY c.created_at DESC`

	err := c.db.Select(&comments, query, tokenId)
	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error getting comments by token ID: %w", err)
	}

	if comments == nil {
		return &[]responses.CommentResponse{}, nil
	}

	return &comments, nil
}

func NewCommentRepo(db *sqlx.DB) interfaces.CommentRepository {
	return &CommentRepo{db: db}
}
