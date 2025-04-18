package interfaces

import (
	"bazar/internal/domain/entities"
	"bazar/internal/domain/requests"

	"github.com/gin-gonic/gin"
)

type CommentRepository interface {
	CreateComment(comment *entities.Comment) (*entities.Comment, error)
	GetCommentById(id string) (*entities.Comment, error)
}

type CommentService interface {
	CreateComment(comment *requests.CreateCommentReq) (*entities.Comment, error)
	GetCommentById(id string) (*entities.Comment, error)
}

type CommentHandler interface {
	CreateComment(c *gin.Context)
	GetCommentById(c *gin.Context)
}
