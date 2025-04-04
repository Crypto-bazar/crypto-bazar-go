package main

import (
	"bazar/config"
	"bazar/internal/contract"
	"bazar/internal/delivery/http/handlers"
	"bazar/internal/delivery/http/router"
	"bazar/internal/infrastructure/database"
	"bazar/internal/infrastructure/eth"
	"bazar/internal/usecase"
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

	userRepo := database.NewUserRepository(db)
	userService := usecase.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	nftRepo := database.NewNFTRepository(db)
	nftService := usecase.NewNFTService(nftRepo, userRepo)

	commentRepo := database.NewCommentRepo(db)
	commentService := usecase.NewCommentService(commentRepo)
	commentHandler := handlers.NewCommentHandler(commentService)

	client, err := eth.NewClient(cfg.EthereumNodeUrl, cfg.ContractAddress)
	if err != nil {
		fmt.Printf("Error create client: %v", err)
	}

	transcations := eth.NewTransaction(client, contract.Abi)
	transcationsService := usecase.NewTransactionUseCase(transcations)
	nftHandler := handlers.NewNFTHandler(nftService, *transcationsService)
	newRouter := router.NewRouter(userHandler, nftHandler, commentHandler)
	newRouter.RegisterRoutes()

	eventHandler := usecase.NewNFTEventHandler(nftService)

	listener, err := eth.NewListener(client, contract.Abi, eventHandler)
	if err != nil {
		fmt.Printf("Error create listener: %v", err)
	}

	go listener.StartListening()

	if err := newRouter.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
