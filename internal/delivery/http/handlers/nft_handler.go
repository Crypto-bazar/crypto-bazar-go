package handlers

import (
	"bazar/internal/domain/interfaces"
	"bazar/internal/domain/requests"
	"bazar/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NFTHandler struct {
	service interfaces.NFTService
}

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

func (u *NFTHandler) GetSalesNFT(c *gin.Context) {
	nfts, err := u.service.GetSalesNFT()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, nfts)
}

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

func NewNFTHandler(service interfaces.NFTService) interfaces.NFTHandler {
	return &NFTHandler{service: service}
}
