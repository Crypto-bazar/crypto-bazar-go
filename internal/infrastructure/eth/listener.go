package eth

import (
	"bazar/internal/domain"
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Listener struct {
	client *Client
	abi    string
}

func NewListener(client *Client, abi string) (*Listener, error) {
	return &Listener{client: client, abi: abi}, nil

}

func (l *Listener) StartListening(ctx context.Context) {

	logs := make(chan types.Log)
	sub, err := l.client.SubscribeToEvents(ctx, logs)
	if err != nil {
		log.Fatal(err)
	}

	parsedABI, err := abi.JSON(strings.NewReader(l.abi))
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			if vLog.Topics[0].Hex() == parsedABI.Events["TokenMinted"].ID.Hex() {
				var eventData domain.TokenMintedEvent

				eventData.TokenId = new(big.Int).SetBytes(vLog.Topics[1].Bytes())
				eventData.Owner = common.HexToAddress(vLog.Topics[2].Hex())

				err := parsedABI.UnpackIntoInterface(&eventData, "TokenMinted", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("Token Minted: TokenId=%s, Owner=%s, TokenURI=%s\n", eventData.TokenId.String(), eventData.Owner.Hex(), eventData.TokenURI)
			}

			if vLog.Topics[0].Hex() == parsedABI.Events["TokenListedForSale"].ID.Hex() {
				var eventData domain.TokenListedForSaleEvent

				eventData.TokenId = new(big.Int).SetBytes(vLog.Topics[1].Bytes()).Uint64()
				eventData.Seller = common.HexToAddress(vLog.Topics[2].Hex())

				err := parsedABI.UnpackIntoInterface(&eventData, "TokenListedForSale", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("Token Listed For Sale: TokenId=%d, Price=%s, Seller=%s\n", eventData.TokenId, eventData.Price.String(), eventData.Seller.Hex())
			}

			if vLog.Topics[0].Hex() == parsedABI.Events["TokenSold"].ID.Hex() {
				var eventData domain.TokenSold

				eventData.TokenId = new(big.Int).SetBytes(vLog.Topics[1].Bytes()).Uint64()
				eventData.Buyer = common.HexToAddress(vLog.Topics[2].Hex())

				err := parsedABI.UnpackIntoInterface(&eventData, "TokenSold", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("Token Sold: TokenId=%d, Buyer=%s, Price=%s\n", eventData.TokenId, eventData.Buyer.Hex(), eventData.Price.String())
			}
		}
	}
}
