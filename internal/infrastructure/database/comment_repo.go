package database

import (
	"bazar/internal/domain"
	"bazar/internal/domain/entities"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type CommentRepo struct {
	db *sqlx.DB
}

func (c CommentRepo) CreateComment(comment *entities.Comment) (*entities.Comment, error) {
	query := `
		INSERT INTO comments ( nft_id, owner_id, content)
		VALUES (:nft_id, :owner_id, :content)`

	rows, err := c.db.NamedQuery(query, comment)
	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("error create comment: %w", err)
	}

	defer func(rows *sqlx.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows: ", err)
		}
	}(rows)

	if rows.Next() {
		err := rows.StructScan(comment)
		if err != nil {
			log.Printf("Error scanning comment: %v", err)
			return nil, fmt.Errorf("error scanning Comment: %w", err)
		}
	} else {
		return nil, fmt.Errorf("no rows returned after insert")
	}

	return comment, nil
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

func NewCommentRepo(db *sqlx.DB) domain.CommentRepository {
	return &CommentRepo{db: db}
}
