package ethereum

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	ethClient       *ethclient.Client
	contractAddress common.Address
}

func NewClient(nodeUrl, contractAddress string) (*Client, error) {
	client, err := ethclient.DialContext(context.Background(), nodeUrl)
	if err != nil {
		return nil, err
	}

	return &Client{
		ethClient:       client,
		contractAddress: common.HexToAddress(contractAddress),
	}, nil
}

func (c *Client) SubscribeToEvents(ctx context.Context, logsChan chan LogEvent) (ethereum.Subscription, error) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{c.contractAddress},
	}

	// Преобразуем наш кастомный тип LogEvent в go-ethereum types.Log
	internalChan := make(chan types.Log)

	sub, err := c.ethClient.SubscribeFilterLogs(ctx, query, internalChan)
	if err != nil {
		return nil, err
	}

	// Запускаем горутину для конвертации логов
	go func() {
		for log := range internalChan {
			logsChan <- LogEvent(log)
		}
	}()

	log.Println("Subscribed to Ethereum events")
	return sub, nil
}
