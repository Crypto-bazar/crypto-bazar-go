package usecase

import (
	"bazar/internal/domain"
	"bazar/internal/domain/entities"
)

type CommentService struct {
	CommentRepo domain.CommentRepository
}

func (c *CommentService) CreateComment(comment *entities.Comment) (*entities.Comment, error) {
	return c.CommentRepo.CreateComment(comment)
}

// GetCommentById implements domain.CommentService.
func (c *CommentService) GetCommentById(id string) (*entities.Comment, error) {
	return c.CommentRepo.GetCommentById(id)
}

func NewCommentService(repo domain.CommentRepository) domain.CommentService {
	return &CommentService{CommentRepo: repo}
}
