package eth

import (
	"bazar/internal/domain"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Events struct {
	parsedABI *abi.ABI
	vLog      *types.Log
}

func NewEvents(abi *abi.ABI, vLog *types.Log) *Events {
	return &Events{parsedABI: abi, vLog: vLog}
}

func (e *Events) MintedEvent() (*domain.TokenMintedEvent, error) {
	var eventData domain.TokenMintedEvent
	eventData.TokenId = e.vLog.Topics[1].Big()

	if err := e.parsedABI.UnpackIntoInterface(&eventData, "NFTMinted", e.vLog.Data); err != nil {
		return nil, fmt.Errorf("error unpacking TokenMinted event: %v", err)
	}
	log.Printf("Topics: %v", e.vLog.Topics)
	log.Printf("tokenId: %s", eventData.TokenId.String())
	log.Printf("%s\n", eventData)

	return &eventData, nil
}

func (e *Events) SalesEvent() (*domain.TokenListedForSaleEvent, error) {
	var eventData domain.TokenListedForSaleEvent
	eventData.TokenId = new(big.Int).SetBytes(e.vLog.Topics[1].Bytes()).Uint64()
	eventData.Seller = common.HexToAddress(e.vLog.Topics[2].Hex())
	fmt.Printf("Token Listed For Sale: TokenId=%d, Price=%s, Seller=%s\n", eventData.TokenId, eventData.Price.String(), eventData.Seller.Hex())

	if err := e.parsedABI.UnpackIntoInterface(&eventData, "TokenListedForSale", e.vLog.Data); err != nil {
		return nil, fmt.Errorf("error unpacking TokenListedForSale event: %v", err)
	}

	return &eventData, nil
}

func (e *Events) SoldEvent() (*domain.TokenSold, error) {
	var eventData domain.TokenSold
	eventData.Buyer = common.HexToAddress(e.vLog.Topics[2].Hex())
	eventData.TokenId = new(big.Int).SetBytes(e.vLog.Topics[1].Bytes()).Uint64()

	if err := e.parsedABI.UnpackIntoInterface(&eventData, "TokenSold", e.vLog.Data); err != nil {
		return nil, fmt.Errorf("error unpacking TokenSold event")
	}

	return &eventData, nil
}
