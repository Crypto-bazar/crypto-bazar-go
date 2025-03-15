package handlers

import (
	"bazar/internal/domain"
	"bazar/internal/domain/requests"
	"bazar/pkg/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NFTHandler struct {
	service domain.NFTService
}

func (u *NFTHandler) SetTokenAddress(c *gin.Context) {
	var req requests.UpdateTokenIdReq

	if err := c.ShouldBind(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	nft, err := u.service.SetTokenAddress(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, nft)

}

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

	fmt.Println("NFT created successfully:", req)

	if err := u.service.CreateNFT(imagePath, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, req)
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

func NewNFTHandler(service domain.NFTService) domain.NFTHandler {
	return &NFTHandler{service: service}
}
