package database

import (
	"bazar/internal/domain/entities"
	"bazar/internal/domain/interfaces"
	"bazar/internal/domain/requests"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type CommentRepo struct {
	db *sqlx.DB
}

func (c *CommentRepo) CreateComment(comment *requests.CreateCommentReq) (*entities.Comment, error) {
	if comment == nil {
		return nil, fmt.Errorf("nil comment request")
	}

	query := `
		WITH owner AS (
			SELECT id
			FROM users
			WHERE eth_address = $1
		)
		INSERT INTO comments (nft_id, owner_id, content)
		VALUES ($2, (SELECT id FROM owner), $3)
		RETURNING *;
	`

	var createdComment entities.Comment

	err := c.db.QueryRowx(query, comment.OwnerAddress, comment.TokenId, comment.Content).StructScan(&createdComment)
	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error creating comment: %w", err)
	}

	return &createdComment, nil
}

func (c CommentRepo) GetCommentsByTokenId(tokenId string) (*[]entities.Comment, error) {
    var comments []entities.Comment
    query := `
        SELECT * FROM comments
        WHERE nft_id = $1
        ORDER BY created_at DESC`

    err := c.db.Select(&comments, query, tokenId)

    if err != nil {
        log.Printf("DB error: %v", err)
        return nil, fmt.Errorf("error getting comments by token ID: %w", err)
    }

    if comments == nil {
        return &[]entities.Comment{}, nil
    }

    return &comments, nil
}

func NewCommentRepo(db *sqlx.DB) interfaces.CommentRepository {
	return &CommentRepo{db: db}
}
