package eth

import (
	"bazar/internal/domain/entities"
	"context"
	"errors"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

type Transaction struct {
	client    *Client
	parsedABI *abi.ABI
}

func NewTransaction(client *Client, abiData string) *Transaction {
	parsedABI, err := abi.JSON(strings.NewReader(abiData))
	if err != nil {
		log.Fatal(err)
	}
	return &Transaction{client: client, parsedABI: &parsedABI}
}

func (t *Transaction) GetProposedNFTs() (*[]entities.NFTProposal, error) {
	methodID := t.parsedABI.Methods["getProposeNFT"].ID

	data := append(methodID[:])

	msg := ethereum.CallMsg{
		To:   &t.client.contractAddress,
		Data: data,
	}

	result, err := t.client.ethClient.CallContract(context.Background(), msg, nil)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, errors.New("empty response from contract")
	}

	var proposals []entities.NFTProposal
	err = t.parsedABI.UnpackIntoInterface(&proposals, "getProposeNFT", result)
	if err != nil {
		return nil, err
	}

	return &proposals, nil
}
