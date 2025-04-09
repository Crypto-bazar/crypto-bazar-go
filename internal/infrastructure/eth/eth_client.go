package eth

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func NewClient(url string) *ethclient.Client {
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	return client
}
