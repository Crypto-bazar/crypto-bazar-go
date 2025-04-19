package main

import (
	_ "bazar/docs"
	"bazar/internal/app/application"
	"log"
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
	app, err := application.NewApplication()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}
	defer app.Shutdown()

	if err := app.InitDependencies(); err != nil {
		log.Fatalf("Failed to initialize dependencies: %v", err)
	}

	if err := app.Run(); err != nil {
		log.Fatalf("Failed to run application: %v", err)
	}
}
