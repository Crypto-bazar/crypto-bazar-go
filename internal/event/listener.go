package event

import (
	"bazar/internal/infrastructure/blockchain"
	"context"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Listener struct {
	client       *ethclient.Client
	contract     *blockchain.GeneratedMetaData // ABI контракта
	contractAddr common.Address
	eventRepo    repository.EventRepository
}

func NewListener(client *ethclient.Client, contractAddr common.Address, eventRepo repository.EventRepository) (*Listener, error) {
	contract, err := generated.NewNFTMarketplace(contractAddr, client)
	if err != nil {
		return nil, err
	}

	return &Listener{
		client:       client,
		contract:     contract,
		contractAddr: contractAddr,
		eventRepo:    eventRepo,
	}, nil
}

func (l *Listener) StartListening(ctx context.Context) {
	log.Println("Запуск прослушивания событий...")

	query := ethereum.FilterQuery{
		Addresses: []common.Address{l.contractAddr},
	}

	logs := make(chan types.Log)
	sub, err := l.client.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		log.Fatalf("Ошибка подписки на логи: %v", err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case err := <-sub.Err():
			log.Println("Ошибка в подписке:", err)
		case vLog := <-logs:
			event, err := l.contract.ParseNFTCreated(vLog)
			if err != nil {
				log.Println("Ошибка парсинга события:", err)
				continue
			}

			nftEvent := domain.NFTCreatedEvent{
				Creator: event.Creator.Hex(),
				TokenID: event.TokenId.Uint64(),
				URI:     event.Uri,
			}

			err = l.eventRepo.SaveEvent(nftEvent)
			if err != nil {
				log.Println("Ошибка сохранения события:", err)
			}

			log.Printf("Получено событие: %+v\n", nftEvent)
		}
	}
}
