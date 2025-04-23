package handlers

import (
	"bazar/internal/domain/interfaces"
	"bazar/internal/domain/requests"
	"bazar/pkg/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

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

// GetUserByAddress godoc
// @Summary Получить пользователя по Address
// @Tags Users
// @Produce json
// @Param address path string true "User Address"
// @Success 200 {object} entities.User
// @Failure 500 {object} map[string]string
// @Router /api/v1/users/{address} [get]
func (u *UserHandler) GetUserByAddress(c *gin.Context) {
	id := c.Param("address")
	user, err := u.service.GetUserByAddress(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UploadAvatarHandler godoc
// @Summary Загрузить аватар пользователя
// @Tags Users
// @Accept multipart/form-data
// @Produce json
// @Param eth_address path string true "Ethereum Address пользователя"
// @Param avatar formData file true "Файл аватара (JPG/PNG)"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/users/{eth_address}/avatar [post]
func (u *UserHandler) UploadAvatarHandler(c *gin.Context) {
	ethAddress := c.Param("eth_address")
	if ethAddress == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ETH address is required"})
		return
	}

	file, _ := c.FormFile("avatar")
	if file == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	fileName := fmt.Sprintf("%s-%s", ethAddress, utils.GenerateRandomString(8)) + filepath.Ext(file.Filename)

	uploadDir := "./uploads/avatars/"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.MkdirAll(uploadDir, os.ModePerm)
	}

	filePath := filepath.Join(uploadDir, fileName)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	avatarURL := "uploads/avatars/" + fileName
	if _, err := u.service.UpdateAvatarURL(ethAddress, avatarURL); err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update avatar URL in the database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Avatar uploaded successfully", "avatar_url": avatarURL})
}

func NewUserHandler(service interfaces.UserService) interfaces.UserHandler {
	return &UserHandler{service: service}
}
