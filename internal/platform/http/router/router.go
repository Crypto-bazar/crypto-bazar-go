package router

import (
	"bazar/internal/platform/http/handlers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine      *gin.Engine
	userHandler *handlers.UserHandler
}

func NewRouter(userHandler *handlers.UserHandler) *Router {
	return &Router{engine: gin.Default(), userHandler: userHandler}
}

func (r *Router) RegisterRoutes() {
	api := r.engine.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.POST("/", r.userHandler.CreateUser)
			users.GET("/", r.userHandler.GetAllUsers)
			users.GET("/:id", r.userHandler.GetUserById)
		}
	}
}

func (r *Router) Run(addr string) error {
	return r.engine.Run(addr)
}
