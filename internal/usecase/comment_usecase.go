package usecase

import (
	"bazar/internal/domain/interfaces"
	"bazar/internal/domain/requests"
	"bazar/internal/domain/responses"
)

type CommentService struct {
	CommentRepo interfaces.CommentRepository
}

func (c *CommentService) GetCommentsByTokenId(tokenId string) (*[]responses.CommentResponse, error) {
	return c.CommentRepo.GetCommentsByTokenId(tokenId)
}

func (c *CommentService) CreateComment(req *requests.CreateCommentReq) (*responses.CommentResponse, error) {
	return c.CommentRepo.CreateComment(req)
}

func NewCommentService(repo interfaces.CommentRepository) interfaces.CommentService {
	return &CommentService{CommentRepo: repo}
}
