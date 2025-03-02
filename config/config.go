package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type DB struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

func ConnectDB(cfg *DB) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(

		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name,
	)

	db, err := sqlx.Open("postgres", dsn)

	if err != nil {
		return nil, fmt.Errorf("error connection database, %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("lost connection databse, %v", err)
	}

	log.Println("Success connection!")
	return db, nil
}

func LoadConfig() *DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	return &DB{
		Host:     getEnv("DB_HOST", "localhost"),
		User:     getEnv("DB_USER", ""),
		Password: getEnv("DB_PASSWORD", ""),
		Name:     getEnv("DB_NAME", ""),
		Port:     getEnv("DB_PORT", "5432"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
