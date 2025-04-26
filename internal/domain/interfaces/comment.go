package interfaces

import (
	"bazar/internal/domain/entities"
	"bazar/internal/domain/requests"

	"github.com/gin-gonic/gin"
)

type CommentRepository interface {
	CreateComment(comment *requests.CreateCommentReq) (*entities.Comment, error)
	GetCommentsByTokenId(tokenId string) (*[]entities.Comment, error)
}

type CommentService interface {
	CreateComment(comment *requests.CreateCommentReq) (*entities.Comment, error)
	GetCommentsByTokenId(tokenId string) (*[]entities.Comment, error)
}

type CommentHandler interface {
	CreateComment(c *gin.Context)
	GetCommentsByTokenId(c *gin.Context)
}
