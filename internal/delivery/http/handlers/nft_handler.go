package handlers

import (
	"bazar/internal/domain/interfaces"
	"bazar/internal/domain/requests"
	"bazar/internal/usecase"
	"bazar/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NFTHandler обрабатывает запросы, связанные с NFT.
type NFTHandler struct {
	service      interfaces.NFTService
	transcations usecase.TransactionUseCase
}

// GetFavouriteNFTS godoc
// @Summary Получить избранные NFT пользователя
// @Description Возвращает список NFT, добавленных в избранное для указанного Ethereum-адреса
// @Tags NFT
// @Produce json
// @Param address query string true "Ethereum Address пользователя"
// @Success 200 {array} entities.NFT
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/nfts/favourites [get]
func (u *NFTHandler) GetFavouriteNFTS(c *gin.Context) {
	ethAddress := c.Query("address")
	if ethAddress == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ethereum address is required"})
		return
	}

	nfts, err := u.service.GetFavouriteNFTS(ethAddress)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, nfts)
}

// AddFavouriteNFT godoc
// @Summary Добавить NFT в избранные
// @Description Добавляет указанный NFT в список избранных для пользователя
// @Tags NFT
// @Accept json
// @Produce json
// @Param data body requests.AddFavouriteNFT true "Данные для добавления в избранное"
// @Success 200 {object} entities.NFT
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/nfts/favourites [post]
func (u *NFTHandler) AddFavouriteNFT(c *gin.Context) {
	var req requests.AddFavouriteNFT
	log.Print(req)

	if err := c.ShouldBind(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	nft, err := u.service.AddFavouriteNFT(req.NftId, req.EthAddress)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, nft)
}

// RemoveFavouriteNFT godoc
// @Summary Удалить NFT из избранных
// @Description Удаляет указанный NFT из списка избранных пользователя
// @Tags NFT
// @Accept json
// @Produce json
// @Param data body requests.RemoveFavouriteNFT true "Данные для удаления из избранного"
// @Success 200 {object} entities.NFT
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/nfts/favourites [delete]
func (u *NFTHandler) RemoveFavouriteNFT(c *gin.Context) {
	var req requests.RemoveFavouriteNFT

	if err := c.ShouldBind(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	nft, err := u.service.RemoveFavouriteNFT(req.NftId, req.EthAddress)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, nft)
}

// GetProposedNFTs godoc
// @Summary Получить предложенные NFT
// @Tags NFT
// @Produce json
// @Success 200 {array} entities.NFT
// @Failure 500 {object} map[string]string
// @Router /api/v1/nfts/proposed [get]
func (u *NFTHandler) GetProposedNFTs(c *gin.Context) {
	nfts, err := u.transcations.GetProposedNFTs()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, nfts)
}

// GetUserNFT godoc
// @Summary Получить NFT пользователя по Ethereum-адресу
// @Tags NFT
// @Produce json
// @Param address query string true "Ethereum Address"
// @Success 200 {array} entities.NFT
// @Failure 500 {object} map[string]string
// @Router /api/v1/nfts/user [get]
func (u *NFTHandler) GetUserNFT(c *gin.Context) {
	ethAddress := c.Query("address")

	nft, err := u.service.GetUserNFT(ethAddress)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, nft)
}

// GetSalesNFT godoc
// @Summary Получить NFT на продаже
// @Tags NFT
// @Produce json
// @Success 200 {array} entities.NFT
// @Failure 500 {object} map[string]string
// @Router /api/v1/nfts/sales [get]
func (u *NFTHandler) GetSalesNFT(c *gin.Context) {
	nfts, err := u.service.GetSalesNFT()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, nfts)
}

// SetTokenAddress godoc
// @Summary Обновить Token ID для NFT
// @Tags NFT
// @Accept json
// @Produce json
// @Param data body requests.UpdateTokenIdReq true "Update Token ID Payload"
// @Success 200 {object} entities.NFT
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/nfts [put]
func (u *NFTHandler) SetTokenAddress(c *gin.Context) {
	var req requests.UpdateTokenIdReq

	if err := c.ShouldBind(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	nft, err := u.service.SetTokenId(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, nft)
}

// CreateNFT godoc
// @Summary Создать новый NFT
// @Tags NFT
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "Название NFT"
// @Param description formData string true "Описание NFT"
// @Param price formData number true "Цена"
// @Param image formData file true "Файл изображения"
// @Param eth_address formData string true "Адрес владельца"
// @Success 201 {object} entities.NFT
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/nfts [post]
func (u *NFTHandler) CreateNFT(c *gin.Context) {
	var req requests.CreateNFTRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	imagePath, err := utils.FileSave(c, "image", "uploads")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	nft, err := u.service.CreateNFT(imagePath, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, nft)
}

// GetAllNFTs godoc
// @Summary Получить все NFT
// @Tags NFT
// @Produce json
// @Success 200 {array} entities.NFT
// @Failure 500 {object} map[string]string
// @Router /api/v1/nfts [get]
func (u *NFTHandler) GetAllNFTs(c *gin.Context) {
	nfts, err := u.service.GetAllNFTs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting all NFTs"})
		return
	}

	c.JSON(http.StatusOK, nfts)
}

// GetNFTById godoc
// @Summary Получить NFT по ID
// @Tags NFT
// @Produce json
// @Param id path string true "NFT ID"
// @Success 200 {object} entities.NFT
// @Failure 500 {object} map[string]string
// @Router /api/v1/nfts/{id} [get]
func (u *NFTHandler) GetNFTById(c *gin.Context) {
	id := c.Param("id")
	nft, err := u.service.GetNFTById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "NFT not found"})
		return
	}
	c.JSON(http.StatusOK, nft)
}

// NewNFTHandler создаёт новый экземпляр NFTHandler.
func NewNFTHandler(service interfaces.NFTService, transactions usecase.TransactionUseCase) interfaces.NFTHandler {
	return &NFTHandler{service: service, transcations: transactions}
}
