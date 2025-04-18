package handlers

import (
	"bazar/internal/domain/interfaces"
	"bazar/internal/domain/requests"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CommentHandler отвечает за обработку запросов, связанных с комментариями.
type CommentHandler struct {
	service interfaces.CommentService
}

// CreateComment godoc
// @Summary Создать комментарий
// @Tags Comments
// @Accept json
// @Produce json
// @Param request body requests.CreateCommentReq true "Комментарий"
// @Success 201 {object} entities.Comment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/comments [post]
func (ch *CommentHandler) CreateComment(c *gin.Context) {
	var req requests.CreateCommentReq

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	comment, err := ch.service.CreateComment(&req)
	if err != nil {
		log.Printf("Error creating comment: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating comment"})
		return
	}

	c.JSON(http.StatusCreated, &comment)
}

// GetCommentById godoc
// @Summary Получить комментарий по ID
// @Tags Comments
// @Produce json
// @Param id path string true "ID комментария"
// @Success 200 {object} entities.Comment
// @Failure 404 {object} map[string]string
// @Router /api/v1/comments/{id} [get]
func (ch *CommentHandler) GetCommentById(c *gin.Context) {
	// TODO implement me
	panic("implement me")
}

// NewCommentHandler создает новый экземпляр CommentHandler.
func NewCommentHandler(service interfaces.CommentService) interfaces.CommentHandler {
	return &CommentHandler{service: service}
}
