package handlers

import (
	"bazar/internal/domain"
	"bazar/internal/domain/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NFTHandler struct {
	service domain.NFTService
}

func (u *NFTHandler) CreateNFT(c *gin.Context) {
	var nft models.NFT

	if err := c.BindJSON(&nft); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := u.service.CreateNFT(&nft); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, nft)
}

func (u *NFTHandler) GetAllNFTs(c *gin.Context) {
	nfts, err := u.service.GetAllNFTs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting all NFTs"})
		return
	}

	c.JSON(http.StatusOK, nfts)
}

func (u *NFTHandler) GetNFTById(c *gin.Context) {
	id := c.Param("id")
	nft, err := u.service.GetNFTById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "NFT not found"})
		return
	}
	c.JSON(http.StatusOK, nft)
}

func NewNFTHandler(service domain.NFTService) *NFTHandler {
	return &NFTHandler{service: service}
}
