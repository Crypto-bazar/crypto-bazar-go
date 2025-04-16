package usecase

import (
	"bazar/internal/domain"
	"context"
	"fmt"
)

type NFTProcessor struct {
	eventListener domain.NFTEventListener
}

func NewNFTProcessor(listener domain.NFTEventListener) *NFTProcessor {
	return &NFTProcessor{eventListener: listener}
}

func (p *NFTProcessor) Run(ctx context.Context) error {
	events, err := p.eventListener.ListenNFTProposed(ctx)
	if err != nil {
		return err
	}

	for event := range events {
		if err := p.processProposal(event); err != nil {
			return err
		}
	}
	return nil
}

func (p *NFTProcessor) processProposal(event domain.NFTProposedEvent) error {
	fmt.Printf("Результат: %s", event)

	return nil
}
