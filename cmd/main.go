package main

import (
	"bazar/config"
	"bazar/internal/app/services"
	"bazar/internal/platform/database/repositories"
	"bazar/internal/platform/http/handlers"
	"bazar/internal/platform/http/router"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

func main() {
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
	nftService := services.NewNFTService(nftRepo, userRepo)
	nftHandler := handlers.NewNFTHandler(nftService)

	newRouter := router.NewRouter(userHandler, nftHandler)
	newRouter.RegisterRoutes()

	if err := newRouter.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
