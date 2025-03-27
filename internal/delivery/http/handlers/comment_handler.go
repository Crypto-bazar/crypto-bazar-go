package handlers

import (
	"bazar/internal/domain/interfaces"
	"bazar/internal/domain/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	service interfaces.CommentService
}

func (ch *CommentHandler) CreateComment(c *gin.Context) {
	var req requests.CreateCommentReq

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	comment, err := ch.service.CreateComment(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating comment"})
		return
	}

	c.JSON(http.StatusCreated, &comment)

}

func (ch *CommentHandler) GetCommentById(c *gin.Context) {
	// TODO implement me
	panic("implement me")
}

func NewCommentHandler(service interfaces.CommentService) interfaces.CommentHandler {
	return &CommentHandler{service: service}
}
