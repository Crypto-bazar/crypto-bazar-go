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

func (c CommentRepo) GetCommentById(id string) (*entities.Comment, error) {
	var comment entities.Comment
	query := `
		SELECT * FROM comments
		WHERE id = $1`

	err := c.db.Get(&comment, query, id)

	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error getting comment by ID: %w", err)
	}

	return &comment, nil
}

func NewCommentRepo(db *sqlx.DB) interfaces.CommentRepository {
	return &CommentRepo{db: db}
}
