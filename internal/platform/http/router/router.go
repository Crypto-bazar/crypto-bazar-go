package router

import (
	"bazar/internal/domain"
	"bazar/internal/platform/http/handlers"
	"bazar/internal/platform/http/middleware"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine      *gin.Engine
	userHandler domain.UserHandler
	nftHandler  *handlers.NFTHandler
}

func NewRouter(userHandler domain.UserHandler, nftHandler *handlers.NFTHandler) *Router {
	return &Router{engine: gin.Default(), userHandler: userHandler, nftHandler: nftHandler}
}

func (r *Router) RegisterRoutes() {
	r.engine.Use(middleware.Logger())
	r.engine.Use(middleware.Recovery())
	r.engine.Use(middleware.Cors())
	r.engine.Static("/uploads", "./uploads")

	api := r.engine.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.POST("/", r.userHandler.CreateUser)
			users.GET("/", r.userHandler.GetAllUsers)
			users.GET("/:id", r.userHandler.GetUserById)
			users.POST("/check", r.userHandler.CheckUserExists)
		}

		nft := api.Group("/nfts")
		{
			nft.POST("/", r.nftHandler.CreateNFT)
			nft.GET("/", r.nftHandler.GetAllNFTs)
			nft.GET("/:id", r.nftHandler.GetNFTById)
		}
	}
}

func (r *Router) Run(addr string) error {
	return r.engine.Run(addr)
}
