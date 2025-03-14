package main

import (
	"bazar/config"
	"bazar/internal/app/services"
	"bazar/internal/ethereum"
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

	ethClient, err := ethereum.NewClient(cfg.EthereumNodeUrl, cfg.ContractAddress)
	if err != nil {
		log.Fatalf("Failed to create Ethereum client: %v", err)
	}

	eventHandler := ethereum.NewEventHandler()

	logs := make(chan ethereum.LogEvent)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sub, err := ethClient.SubscribeToEvents(ctx, logs)
	if err != nil {
		log.Fatalf("Failed to subscribe to Ethereum events: %v", err)
	}
	defer sub.Unsubscribe()

	go eventHandler.HandleEvents(ctx, logs)

	newRouter := router.NewRouter(userHandler, nftHandler)
	newRouter.RegisterRoutes()

	if err := newRouter.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
