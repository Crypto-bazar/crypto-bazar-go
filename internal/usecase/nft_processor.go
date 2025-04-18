package usecase

import (
	"bazar/internal/delivery/http/ws"
	"bazar/internal/domain"
	"bazar/internal/domain/interfaces"
	"context"
	"fmt"
)

type NFTProcessor struct {
	eventListener domain.NFTEventListener
	nftRepo       interfaces.NFTRepository
	hub           *ws.Hub
}

func NewNFTProcessor(listener domain.NFTEventListener, nftRepo interfaces.NFTRepository, hub *ws.Hub) *NFTProcessor {
	return &NFTProcessor{eventListener: listener, nftRepo: nftRepo, hub: hub}
}

func (p *NFTProcessor) Run(ctx context.Context) error {
	proposedCh, err := p.eventListener.ListenNFTProposed(ctx)
	if err != nil {
		return err
	}

	voteCh, err := p.eventListener.ListenNFTVoted(ctx)
	if err != nil {
		return err
	}

	mintedCh, err := p.eventListener.ListendNFTMinted(ctx)
	if err != nil {
		return err
	}

	for {
		select {
		case event := <-proposedCh:
			if err := p.processProposal(event); err != nil {
				return err
			}
		case event := <-voteCh:
			if err := p.processVote(event); err != nil {
				return err
			}
		case event := <-mintedCh:
			if err := p.processMinted(event); err != nil {
				return err
			}
		case <-ctx.Done():
			return nil

		}
	}
}

func (p *NFTProcessor) processProposal(event domain.NFTProposedEvent) error {
	_, err := p.nftRepo.UpdateProposedByTokenURI(event.ProposalID, event.TokenURI)
	if err != nil {
		return fmt.Errorf("error with updating proposed status: %w", err)
	}

	return nil
}

func (p *NFTProcessor) processVote(event domain.NFTVotedEvent) error {
	nft, err := p.nftRepo.UpdateVotesByTokenURI(event.TokenURI, event.Amount.String())
	if err != nil {
		return fmt.Errorf("error with updating votes: %w", err)
	}

	p.hub.Broadcast(struct {
		Type string `json:"type"`
		NFT  any    `json:"nft"`
	}{
		Type: "vote",
		NFT:  nft,
	})

	return nil
}

func (p *NFTProcessor) processMinted(event domain.NFTMintedEvent) error {
	_, err := p.nftRepo.UpdateId(event.TokenURI, event.TokenId)
	if err != nil {
		return fmt.Errorf("error with updating proposed status: %w", err)
	}
	return nil
}
