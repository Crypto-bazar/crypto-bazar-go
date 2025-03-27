package database

import (
	"bazar/internal/domain/entities"
	"bazar/internal/domain/interfaces"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type CommentRepo struct {
	db *sqlx.DB
}

func (c CommentRepo) CreateComment(comment *entities.Comment) (*entities.Comment, error) {
	query := `
		INSERT INTO comments (nft_id, owner_id, content)
		VALUES (:nft_id, :owner_id, :content)
		RETURNING *`

	rows, err := c.db.NamedQuery(query, comment)
	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error creating comment: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.StructScan(comment)
		if err != nil {
			log.Printf("Error scanning comment: %v", err)
			return nil, fmt.Errorf("error scanning comment: %w", err)
		}
		return comment, nil
	}

	return nil, fmt.Errorf("no rows returned after insert")
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
