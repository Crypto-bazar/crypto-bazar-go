package eth

import (
	"bazar/internal/domain"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func MintedEvent(parsedABI *abi.ABI, vLog types.Log) (*domain.TokenMintedEvent, error) {
	var eventData domain.TokenMintedEvent
	eventData.TokenId = new(big.Int).SetBytes(vLog.Topics[1].Bytes())
	eventData.Owner = common.HexToAddress(vLog.Topics[2].Hex())

	if err := parsedABI.UnpackIntoInterface(&eventData, "TokenMinted", vLog.Data); err != nil {
		return nil, fmt.Errorf("error unpacking TokenMinted event: %v", err)
	}
	
	return &eventData, nil
}
