package application

import (
	"bazar/internal/app/config"
	"bazar/internal/delivery/http/handlers"
	"bazar/internal/delivery/http/router"
	"bazar/internal/delivery/websocket"
	"bazar/internal/infrastructure/database"
	"bazar/internal/infrastructure/eth"
	"bazar/internal/usecase"
	"context"
	"fmt"
	"log"
	"runtime/debug"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jmoiron/sqlx"
)

type Application struct {
	Config    *config.Config
	DB        *sqlx.DB
	Router    *router.Router
	EthClient *ethclient.Client
	WSHub     *websocket.Hub
}

func NewApplication() (*Application, error) {
	cfg := config.LoadConfig()

	db, err := config.ConnectDB(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	ethClient, err := config.NewClient(cfg.EthereumNodeUrl, 10, time.Second)
	if err != nil {
		return nil, fmt.Errorf("failed to connect ethereum client: %w", err)
	}

	wsHub := websocket.NewHub()

	return &Application{
		Config:    cfg,
		DB:        db,
		EthClient: ethClient,
		WSHub:     wsHub,
	}, nil
}

func (a *Application) InitDependencies() error {
	userRepo := database.NewUserRepository(a.DB)
	userService := usecase.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	nftRepo := database.NewNFTRepository(a.DB)
	nftService := usecase.NewNFTService(nftRepo, userRepo)

	commentRepo := database.NewCommentRepo(a.DB)
	commentService := usecase.NewCommentService(commentRepo)
	commentHandler := handlers.NewCommentHandler(commentService)

	ctx := context.Background()
	instance := eth.LoadContract(a.Config.ContractAddress, a.EthClient)

	eventListener := eth.NewEthEventListener(instance)
	processor := usecase.NewNFTProcessor(eventListener, nftRepo, a.WSHub)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("[panic recovered in processor.Run]: %v", r)
				debug.PrintStack()
			}
		}()

		if err := processor.Run(ctx); err != nil {
			log.Printf("processor error: %v", err)
		}
	}()

	transactions := eth.NewTransaction(instance)
	transactionsService := usecase.NewTransactionUseCase(transactions, ctx)
	nftHandler := handlers.NewNFTHandler(nftService, *transactionsService)
	wsHandler := websocket.Handler(a.WSHub)

	a.Router = router.NewRouter(userHandler, nftHandler, commentHandler, &wsHandler)

	return nil
}

func (a *Application) Run() error {
	a.Router.RegisterRoutes()
	return a.Router.Run(":" + a.Config.ServerPort)
}

func (a *Application) Shutdown() {
	if err := a.DB.Close(); err != nil {
		log.Printf("error closing database: %v", err)
	}
	a.EthClient.Close()
}
