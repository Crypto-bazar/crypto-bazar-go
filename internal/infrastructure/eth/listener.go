package eth

import (
	"bazar/internal/domain"
	"bazar/internal/domain/interfaces"
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
	eventHandler interfaces.EventHandler
}

func NewListener(client *Client, abi string, eventHandler interfaces.EventHandler) (*Listener, error) {
	return &Listener{client: client, abi: abi, eventHandler: eventHandler}, nil

}

func (l *Listener) parseAbi() *abi.ABI {
	parsedABI, err := abi.JSON(strings.NewReader(l.abi))
	if err != nil {
		log.Fatal(err)
	}

	return &parsedABI
}

func (l *Listener) subscribeToEvents() (ethereum.Subscription, chan types.Log) {
	logs := make(chan types.Log)

	if l.client == nil {
		log.Printf("Ethereum client is nil, skipping event subscription")
		return nil, logs
	}

	sub, err := l.client.SubscribeToEvents(logs)
	if err != nil {
		log.Printf("Subscription error: %v.", err)

	}

	return sub, logs
}

func (l *Listener) StartListening() {
	sub, logs := l.subscribeToEvents()

	if sub == nil {
		log.Println("Subscription is nil, skipping listening loop")
		time.Sleep(5 * time.Second)
	}

	parsedABI := l.parseAbi()

	for {
		select {
		case err := <-sub.Err():
			log.Printf("Subscription error: %v. Reconnecting...", err)
			sub.Unsubscribe()
			time.Sleep(5 * time.Second)

		case vLog := <-logs:
			l.handleLog(parsedABI, vLog)
		}
	}
}

func (l *Listener) handleLog(parsedABI *abi.ABI, vLog types.Log) {
	events := NewEvents(parsedABI, &vLog)

	switch vLog.Topics[0].Hex() {
	case parsedABI.Events["TokenMinted"].ID.Hex():
		eventData, err := events.MintedEvent()

		if err != nil {
			fmt.Printf("Error unpacking TokenMinted event: %v", err)
			return
		}

		if err := l.eventHandler.OnTokenMinted(eventData); err != nil {
			log.Printf("Error processing TokenMinted event: %v", err)
		}

	case parsedABI.Events["TokenListedForSale"].ID.Hex():
		eventData, err := events.SalesEvent()

		if err != nil {
			fmt.Printf("Error unpacking TokenListedForSale event: %v", err)
			return
		}

		if err := l.eventHandler.OnTokenListedForSale(eventData); err != nil {
			fmt.Printf("Error unpacking TokenListedForSale event: %v", err)
			return
		}

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
