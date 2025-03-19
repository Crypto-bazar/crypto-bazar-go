package eth

import (
	"bazar/internal/domain"
	"fmt"
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
	eventData.TokenId = new(big.Int).SetBytes(e.vLog.Topics[1].Bytes())
	eventData.Owner = common.HexToAddress(e.vLog.Topics[2].Hex())

	if err := e.parsedABI.UnpackIntoInterface(&eventData, "TokenMinted", e.vLog.Data); err != nil {
		return nil, fmt.Errorf("error unpacking TokenMinted event: %v", err)
	}

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
