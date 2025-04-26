package usecase

import (
	"bazar/internal/domain/entities"
	"bazar/internal/domain/interfaces"
	"bazar/internal/domain/requests"
)

type CommentService struct {
	CommentRepo interfaces.CommentRepository
}

// GetCommentsByTokenId implements interfaces.CommentService.
func (c *CommentService) GetCommentsByTokenId(tokenId string) (*[]entities.Comment, error) {
	return c.CommentRepo.GetCommentsByTokenId(tokenId)
}

func (c *CommentService) CreateComment(req *requests.CreateCommentReq) (*entities.Comment, error) {
	return c.CommentRepo.CreateComment(req)
}

func NewCommentService(repo interfaces.CommentRepository) interfaces.CommentService {
	return &CommentService{CommentRepo: repo}
}
