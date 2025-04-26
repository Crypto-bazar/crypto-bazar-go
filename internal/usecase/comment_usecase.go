package usecase

import (
	"bazar/internal/domain/entities"
	"bazar/internal/domain/interfaces"
	"bazar/internal/domain/requests"
)

type CommentService struct {
	CommentRepo interfaces.CommentRepository
}

func (c *CommentService) CreateComment(req *requests.CreateCommentReq) (*entities.Comment, error) {
	return c.CommentRepo.CreateComment(req)
}

func (c *CommentService) GetCommentById(id string) (*entities.Comment, error) {
	return c.CommentRepo.GetCommentById(id)
}

func NewCommentService(repo interfaces.CommentRepository) interfaces.CommentService {
	return &CommentService{CommentRepo: repo}
}
