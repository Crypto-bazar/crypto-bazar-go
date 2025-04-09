package eth

import (
	"bazar/internal/contract"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func LoadContract(address string, client *ethclient.Client) *contract.Mycontract {
	contractAddress := common.HexToAddress(address)
	instance, err := contract.NewMycontract(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to load contract: %v", err)
	}
	return instance
}
