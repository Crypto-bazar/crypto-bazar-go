package app

import (
	"bazar/config"
	"bazar/internal/app/services"
	"bazar/internal/platform/database/repositories"
	"bazar/internal/platform/http/handlers"
	"bazar/internal/platform/http/router"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type App struct {
	Router *router.Router
	DB     *sqlx.DB
}

func BootStrap() (*App, error) {
	cfg := config.LoadConfig()
	db, err := config.ConnectDB(cfg)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("Error closing database connection: %v", err)
		}
	}(db)

	log.Println("Success connection!")

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserRepository(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	nftRepo := repositories.NewNFTRepository(db)
	nftService := services.NewNFTService(nftRepo)
	nftHandler := handlers.NewNFTHandler(nftService)

	newRouter := router.NewRouter(userHandler, nftHandler)
	newRouter.RegisterRoutes()

	return &App{
		Router: newRouter,
		DB:     db,
	}, nil
}
