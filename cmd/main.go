package main

import (
	"bazar/config"
	handlers2 "bazar/internal/http/handlers"
	"bazar/internal/http/router"
	"bazar/internal/infrastructure/database"
	"bazar/internal/infrastructure/eth"
	"bazar/internal/usecase"
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

	userRepo := database.NewUserRepository(db)
	userService := usecase.NewUserService(userRepo)
	userHandler := handlers2.NewUserHandler(userService)

	nftRepo := database.NewNFTRepository(db)
	nftService := usecase.NewNFTService(nftRepo, userRepo)

	commentRepo := database.NewCommentRepo(db)
	commentService := usecase.NewCommentService(commentRepo)
	commentHandler := handlers2.NewCommentHandler(commentService)

	ctx, cancel := context.WithCancel(context.Background())
	client := eth.NewClient(cfg.EthereumNodeUrl)
	defer client.Close()

	instance := eth.LoadContract(cfg.ContractAddress, client)
	defer cancel()

	go eth.ListenNFTProposed(ctx, instance)

	transcations := eth.NewTransaction(instance)
	transcationsService := usecase.NewTransactionUseCase(transcations, ctx)
	nftHandler := handlers2.NewNFTHandler(nftService, *transcationsService)
	newRouter := router.NewRouter(userHandler, nftHandler, commentHandler)
	newRouter.RegisterRoutes()

	if err := newRouter.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
