package config

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	Host            string
	User            string
	Password        string
	Name            string
	DBPort          string
	ServerPort      string
	EthereumNodeUrl string
	ContractAddress string
}

func NewClient(nodeURL string, maxRetries int, initialDelay time.Duration) (*ethclient.Client, error) {
	var client *ethclient.Client
	var err error

	for i := 0; i < maxRetries; i++ {
		client, err = ethclient.Dial(nodeURL)
		if err == nil {
			return client, nil
		}

		delay := calculateBackoff(i, initialDelay)
		log.Printf("Attempt %d: failed to connect to Ethereum node: %v (retrying in %v)", i+1, err, delay)
		time.Sleep(delay)
	}

	return nil, fmt.Errorf("failed to connect to Ethereum node after %d attempts: %v", maxRetries, err)
}

func ConnectDBWithRetry(cfg *Config, maxRetries int, initialDelay time.Duration) (*sqlx.DB, error) {
	var db *sqlx.DB
	var err error

	for i := 0; i < maxRetries; i++ {
		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.Host, cfg.DBPort, cfg.User, cfg.Password, cfg.Name,
		)

		db, err = sqlx.Open("postgres", dsn)
		if err != nil {
			log.Printf("Attempt %d: failed to open database connection: %v", i+1, err)
			time.Sleep(calculateBackoff(i, initialDelay))
			continue
		}

		err = db.Ping()
		if err != nil {
			log.Printf("Attempt %d: failed to ping database: %v", i+1, err)
			_ = db.Close() // Закрываем неудачное соединение
			time.Sleep(calculateBackoff(i, initialDelay))
			continue
		}

		log.Printf("Successfully connected to database after %d attempts", i+1)
		return db, nil
	}

	return nil, fmt.Errorf("failed to connect to database after %d attempts: %v", maxRetries, err)
}

func calculateBackoff(attempt int, initialDelay time.Duration) time.Duration {
	return time.Duration(math.Pow(2, float64(attempt))) * initialDelay
}

func ConnectDB(cfg *Config) (*sqlx.DB, error) {
	return ConnectDBWithRetry(cfg, 10, time.Second)
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
		DBPort:          getEnv("DB_PORT", "7432"),
		EthereumNodeUrl: getEnv("ETHEREUM_NODE_URL", ""),
		ContractAddress: getEnv("CONTRACT_ADDRESS", ""),
		ServerPort:      getEnv("SERVER_PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
