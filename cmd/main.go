package main

import (
	"bazar/internal/app"
	"fmt"
	"log"
)

func main() {
	application, err := app.BootStrap()
	if err != nil {
		log.Fatalf("Failed to bootstrap application: %v", err)
	}

	defer func() {
		if err := application.DB.Close(); err != nil {
			fmt.Printf("Error closing database connection: %v", err)
		}
	}()

	if err := application.Router.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
