package main

import (
	"bazar/config"
	"bazar/internal/app/services"
	"bazar/internal/contract"
	"bazar/internal/event"
	"bazar/internal/infrastructure"
	"bazar/internal/platform/database/repositories"
	"bazar/internal/platform/http/handlers"
	"bazar/internal/platform/http/router"
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
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
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	nftRepo := repositories.NewNFTRepository(db)
	nftService := services.NewNFTService(nftRepo, userRepo)
	nftHandler := handlers.NewNFTHandler(nftService)

	newRouter := router.NewRouter(userHandler, nftHandler)
	newRouter.RegisterRoutes()

	client, err := infrastructure.NewClient(cfg.EthereumNodeUrl, cfg.ContractAddress)
	if err != nil {
		fmt.Printf("Error create client: %v", err)
	}

	listener, err := event.NewListener(client, contract.ContractAbi)
	if err != nil {
		fmt.Printf("Error create listener: %v", err)
	}

	go listener.StartListening(context.Background())

	if err := newRouter.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
