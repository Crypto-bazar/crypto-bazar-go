package handlers

import (
	"bazar/internal/domain"
	"bazar/internal/domain/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service domain.UserService
}

// CheckUserExists implements domain.UserHandler.
func (u *UserHandler) CheckUserExists(c *gin.Context) {
	var address dto.CheckUserRequest
	if err := c.BindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	isExists, err := u.service.CheckUserExists(&address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, isExists)
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	var user dto.CreateUserRequest

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := u.service.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

func NewUserHandler(service domain.UserService) domain.UserHandler {
	return &UserHandler{service: service}
}
