package handlers

import (
	"bazar/internal/domain"
	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	service domain.CommentService
}

func (c2 *CommentHandler) CreateComment(c *gin.Context) {
	// TODO implement me
//   var req 
}

func (c2 *CommentHandler) GetCommentById(c *gin.Context) {
	// TODO implement me
	panic("implement me")
}

func NewCommentHandler(service domain.CommentService) domain.CommentHandler {
	return &CommentHandler{service: service}
}
