package handlers

import (
	"bazar/internal/domain"
	"bazar/internal/domain/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service domain.UserService
}

// CreateUser implements domain.UserHandler.
func (u *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := u.service.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetAllUsers implements domain.UserHandler.
func (u *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := u.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting all users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUserById implements domain.UserHandler.
func (u *UserHandler) GetUserById(c *gin.Context) {
	id := c.Param("id")
	user, err := u.service.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func NewUserHandler(service domain.UserService) *UserHandler {
	return &UserHandler{service: service}
}
