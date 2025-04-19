package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	Host            string
	User            string
	Password        string
	Name            string
	Port            string
	EthereumNodeUrl string
	ContractAddress string
}

func ConnectDB(cfg *Config) (*sqlx.DB, error) {
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

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Host:            getEnv("DB_HOST", "localhost"),
		User:            getEnv("DB_USER", ""),
		Password:        getEnv("DB_PASSWORD", ""),
		Name:            getEnv("DB_NAME", ""),
		Port:            getEnv("DB_PORT", "5432"),
		EthereumNodeUrl: getEnv("ETHEREUM_NODE_URL", ""),
		ContractAddress: getEnv("CONTRACT_ADDRESS", ""),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
