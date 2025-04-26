package interfaces

import (
	"bazar/internal/domain/requests"
	"bazar/internal/domain/responses"

	"github.com/gin-gonic/gin"
)

type CommentRepository interface {
	CreateComment(comment *requests.CreateCommentReq) (*responses.CommentResponse, error)
	GetCommentsByTokenId(tokenId string) (*[]responses.CommentResponse, error)
}

type CommentService interface {
	CreateComment(comment *requests.CreateCommentReq) (*responses.CommentResponse, error)
	GetCommentsByTokenId(tokenId string) (*[]responses.CommentResponse, error)
}

type CommentHandler interface {
	CreateComment(c *gin.Context)
	GetCommentsByTokenId(c *gin.Context)
}
