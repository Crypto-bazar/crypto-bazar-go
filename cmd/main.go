package main

import (
	"bazar/config"
	_ "bazar/docs"
	"bazar/internal/delivery/http/handlers"
	"bazar/internal/delivery/http/router"
	"bazar/internal/delivery/websocket"
	"bazar/internal/infrastructure/database"
	"bazar/internal/infrastructure/eth"
	"bazar/internal/usecase"
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

// @title           Bazar NFT Marketplace API
// @version         1.0
// @description     Backend for Bazar - NFT marketplace on Ethereum
// @termsOfService  http://swagger.io/terms/

// @contact.name   Egor Gladkikh
// @contact.url    https://github.com/egorgladkikh
// @contact.email  egor@example.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey  BearerAuth
// @in                          header
// @name                        Authorization
// @description                 Type "Bearer" followed by a space and JWT token.

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

	userRepo := database.NewUserRepository(db)
	userService := usecase.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	nftRepo := database.NewNFTRepository(db)
	nftService := usecase.NewNFTService(nftRepo, userRepo)
	hub := websocket.NewHub()

	commentRepo := database.NewCommentRepo(db)
	commentService := usecase.NewCommentService(commentRepo)
	commentHandler := handlers.NewCommentHandler(commentService)

	ctx, cancel := context.WithCancel(context.Background())
	client := eth.NewClient(cfg.EthereumNodeUrl)
	defer client.Close()

	instance := eth.LoadContract(cfg.ContractAddress, client)
	defer cancel()

	eventListener := eth.NewEthEventListener(instance)

	processor := usecase.NewNFTProcessor(eventListener, nftRepo, hub)

	go func() {
		if err := processor.Run(ctx); err != nil {
			log.Printf("processor error: %v", err)
		}
	}()

	transactions := eth.NewTransaction(instance)
	transactionsService := usecase.NewTransactionUseCase(transactions, ctx)
	nftHandler := handlers.NewNFTHandler(nftService, *transactionsService)
	wsHandler := websocket.Handler(hub)

	newRouter := router.NewRouter(userHandler, nftHandler, commentHandler, &wsHandler)
	newRouter.RegisterRoutes()

	if err := newRouter.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
