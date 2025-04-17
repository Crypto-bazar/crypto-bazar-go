package handlers

import (
	"bazar/internal/domain/interfaces"
	"bazar/internal/domain/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler отвечает за работу с пользователями.
type UserHandler struct {
	service interfaces.UserService
}

// CheckUserExists godoc
// @Summary Проверить существование пользователя по адресу
// @Tags Users
// @Accept json
// @Produce json
// @Param request body requests.CheckUserRequest true "Ethereum Address"
// @Success 200 {object} bool
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/users/check [post]
func (u *UserHandler) CheckUserExists(c *gin.Context) {
	var address requests.CheckUserRequest
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

// CreateUser godoc
// @Summary Создать нового пользователя
// @Tags Users
// @Accept json
// @Produce json
// @Param user body requests.CreateUserRequest true "User payload"
// @Success 201 {object} requests.CreateUserRequest
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/users [post]
func (u *UserHandler) CreateUser(c *gin.Context) {
	var user requests.CreateUserRequest

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

// GetAllUsers godoc
// @Summary Получить список всех пользователей
// @Tags Users
// @Produce json
// @Success 200 {array} entities.User
// @Failure 500 {object} map[string]string
// @Router /api/v1/users [get]
func (u *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := u.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting all users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUserById godoc
// @Summary Получить пользователя по ID
// @Tags Users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} entities.User
// @Failure 500 {object} map[string]string
// @Router /api/v1/users/{id} [get]
func (u *UserHandler) GetUserById(c *gin.Context) {
	id := c.Param("id")
	user, err := u.service.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// NewUserHandler создаёт новый обработчик UserHandler.
func NewUserHandler(service interfaces.UserService) interfaces.UserHandler {
	return &UserHandler{service: service}
}
