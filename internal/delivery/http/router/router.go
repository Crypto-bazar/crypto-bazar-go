package router

import (
	"bazar/internal/delivery/http/middleware"
	"bazar/internal/domain/interfaces"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine         *gin.Engine
	userHandler    interfaces.UserHandler
	nftHandler     interfaces.NFTHandler
	commentHandler interfaces.CommentHandler
	wsHandler      *gin.HandlerFunc
}

func NewRouter(userHandler interfaces.UserHandler, nftHandler interfaces.NFTHandler, commentHandler interfaces.CommentHandler, wsHandler *gin.HandlerFunc) *Router {
	return &Router{engine: gin.Default(), userHandler: userHandler, nftHandler: nftHandler, commentHandler: commentHandler, wsHandler: wsHandler}
}

func (r *Router) RegisterRoutes() {
	r.engine.Use(middleware.Logger())
	r.engine.Use(middleware.Recovery())
	r.engine.Use(middleware.Cors())
	r.engine.Static("/uploads", "./uploads")
	r.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.engine.GET("/ws", *r.wsHandler)

	api := r.engine.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.POST("/", r.userHandler.CreateUser)
			users.GET("/", r.userHandler.GetAllUsers)
			users.GET("/:address", r.userHandler.GetUserByAddress)
			users.POST("/check", r.userHandler.CheckUserExists)
			users.POST("/:eth_address/avatar", r.userHandler.UploadAvatarHandler)

		}

		nft := api.Group("/nfts")
		{
			nft.POST("/", r.nftHandler.CreateNFT)
			nft.GET("/", r.nftHandler.GetAllNFTs)
			nft.GET("/:id", r.nftHandler.GetNFTById)
			nft.PUT("/", r.nftHandler.SetTokenAddress)
			nft.GET("/sales", r.nftHandler.GetSalesNFT)
			nft.GET("/user", r.nftHandler.GetUserNFT)
			nft.GET("/proposed", r.nftHandler.GetProposedNFTs)

			nft.GET("/favourites:address", r.nftHandler.GetFavouriteNFTS)
			nft.POST("/favourites", r.nftHandler.AddFavouriteNFT)
			nft.DELETE("/favourites", r.nftHandler.RemoveFavouriteNFT)
		}

		comment := api.Group("comments")
		{
			comment.POST("/", r.commentHandler.CreateComment)
			comment.GET("/:tokenId", r.commentHandler.GetCommentsByTokenId)
		}
	}
}

func (r *Router) Run(addr string) error {
	return r.engine.Run(addr)
}
