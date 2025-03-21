package eth

import (
	"context"

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

func (c *Client) SubscribeToEvents(logs chan types.Log) (ethereum.Subscription, error) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{c.contractAddress},
	}

	sub, err := c.ethClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		return nil, err
	}
	return sub, nil
}
