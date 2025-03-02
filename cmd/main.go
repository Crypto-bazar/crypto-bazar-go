package main

import (
	"bazar/config"
	"bazar/internal/app/services"
	"bazar/internal/platform/database/repositories"
	"bazar/internal/platform/http/handlers"
	"bazar/internal/platform/http/router"
	"log"
)

func main() {
	cfg := config.LoadConfig()
	db, err := config.ConnectDB(cfg)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	defer db.Close()
	log.Println("Success connection!")

	userRepo := repositories.NewUserRepository(db)

	userService := services.NewUserRepository(userRepo)

	userHandler := handlers.NewUserHandler(userService)

	router := router.NewRouter(userHandler)
	router.RegisterRoutes()

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
