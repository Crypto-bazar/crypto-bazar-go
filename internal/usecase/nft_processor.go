package usecase

import (
	"bazar/internal/domain"
	"bazar/internal/domain/interfaces"
	"context"
	"fmt"
)

type NFTProcessor struct {
	eventListener domain.NFTEventListener
	nftRepo       interfaces.NFTRepository
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
	_, err := p.nftRepo.UpdateProposedByTokenURI(true, event.TokenURI)
	if err != nil {
		return fmt.Errorf("error with updating proposed status: %w", err)
	}

	return nil
}
