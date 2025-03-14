package ethereum

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
)

type LogEvent struct {
	Data string
}

type EventHandler struct{}

func NewEventHandler() *EventHandler {
	return &EventHandler{}
}

func (h *EventHandler) HandleEvents(ctx context.Context, logs chan types.Log) {
	for {
		select {
		case <-ctx.Done():
			log.Println("Stopping event handler")
			return

		case logEvent := <-logs:
			log.Printf("New Ethereum event: %+v", logEvent)
		}
	}
}
