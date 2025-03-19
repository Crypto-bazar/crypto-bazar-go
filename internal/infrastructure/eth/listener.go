package eth

import (
	"bazar/internal/domain"
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Listener struct {
	client       *Client
	abi          string
	eventHandler domain.EventHandler
}

func NewListener(client *Client, abi string, eventHandler domain.EventHandler) (*Listener, error) {
	return &Listener{client: client, abi: abi, eventHandler: eventHandler}, nil

}

func (l *Listener) parseAbi() *abi.ABI {
	parsedABI, err := abi.JSON(strings.NewReader(l.abi))
	if err != nil {
		log.Fatal(err)
	}

	return &parsedABI
}

func (l *Listener) subscribeToEvents(ctx context.Context) (ethereum.Subscription, chan types.Log) {
	logs := make(chan types.Log)

	if l.client == nil {
		log.Printf("Ethereum client is nil, skipping event subscription")
		return nil, logs
	}

	sub, err := l.client.SubscribeToEvents(ctx, logs)
	if err != nil {
		log.Printf("Subscription error: %v.", err)

	}

	return sub, logs
}

func (l *Listener) StartListening(ctx context.Context) {
	for {
		sub, logs := l.subscribeToEvents(ctx)

		if sub == nil {
			log.Println("Subscription is nil, skipping listening loop")
			time.Sleep(5 * time.Second)
			continue
		}

		parsedABI := l.parseAbi()

		for {
			select {
			case err := <-sub.Err():
				log.Printf("Subscription error: %v. Reconnecting...", err)
				sub.Unsubscribe()
				time.Sleep(5 * time.Second)
				continue

			case vLog := <-logs:
				l.handleLog(ctx, parsedABI, vLog)
			}
		}
	}
}

func (l *Listener) handleLog(ctx context.Context, parsedABI *abi.ABI, vLog types.Log) {
	switch vLog.Topics[0].Hex() {
	case parsedABI.Events["TokenMinted"].ID.Hex():
		eventData, err := MintedEvent(parsedABI, vLog)
		
		if err != nil {
			fmt.Printf("Error unpacking TokenMinted event: %v", err)
			return
		}

		if err := l.eventHandler.OnTokenMinted(ctx, eventData); err != nil {
			log.Printf("Error processing TokenMinted event: %v", err)
		}

	case parsedABI.Events["TokenListedForSale"].ID.Hex():
		var eventData domain.TokenListedForSaleEvent
		eventData.TokenId = new(big.Int).SetBytes(vLog.Topics[1].Bytes()).Uint64()
		eventData.Seller = common.HexToAddress(vLog.Topics[2].Hex())

		if err := parsedABI.UnpackIntoInterface(&eventData, "TokenListedForSale", vLog.Data); err != nil {
			log.Printf("Error unpacking TokenListedForSale event: %v", err)
			return
		}

		fmt.Printf("Token Listed For Sale: TokenId=%d, Price=%s, Seller=%s\n", eventData.TokenId, eventData.Price.String(), eventData.Seller.Hex())

	case parsedABI.Events["TokenSold"].ID.Hex():
		var eventData domain.TokenSold
		eventData.TokenId = new(big.Int).SetBytes(vLog.Topics[1].Bytes()).Uint64()
		eventData.Buyer = common.HexToAddress(vLog.Topics[2].Hex())

		if err := parsedABI.UnpackIntoInterface(&eventData, "TokenSold", vLog.Data); err != nil {
			log.Printf("Error unpacking TokenSold event: %v", err)
			return
		}

		fmt.Printf("Token Sold: TokenId=%d, Buyer=%s, Price=%s\n", eventData.TokenId, eventData.Buyer.Hex(), eventData.Price.String())
	}
}
