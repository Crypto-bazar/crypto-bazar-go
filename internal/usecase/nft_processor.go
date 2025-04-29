package usecase

import (
	"bazar/internal/delivery/websocket"
	"bazar/internal/domain"
	"bazar/internal/domain/interfaces"
	"context"
	"fmt"
	"log"
)

type NFTProcessor struct {
	eventListener domain.NFTEventListener
	nftRepo       interfaces.NFTRepository
	hub           *websocket.Hub
}

func NewNFTProcessor(listener domain.NFTEventListener, nftRepo interfaces.NFTRepository, hub *websocket.Hub) *NFTProcessor {
	return &NFTProcessor{eventListener: listener, nftRepo: nftRepo, hub: hub}
}

func (p *NFTProcessor) Run(ctx context.Context) error {
	proposedCh, err := p.eventListener.ListenNFTProposed(ctx)
	if err != nil {
		log.Printf("error processing proposal: %v", err)
	}

	voteCh, err := p.eventListener.ListenNFTVoted(ctx)
	if err != nil {
		return err
	}

	mintedCh, err := p.eventListener.ListendNFTMinted(ctx)
	if err != nil {
		return err
	}

	salesCh, err := p.eventListener.ListenNFTInSale(ctx)
	if err != nil {
		return err
	}

	soldCh, err := p.eventListener.ListenNFTSold(ctx)
	if err != nil {
		return err
	}

	for {
		select {
		case event := <-proposedCh:
			if err := p.processProposal(event); err != nil {
				log.Printf("error processing proposal: %v", err)
			}
		case event := <-voteCh:
			if err := p.processVote(event); err != nil {
				log.Printf("error processing vote: %v", err)
			}
		case event := <-mintedCh:
			if err := p.processMinted(event); err != nil {
				log.Printf("error processing minted: %v", err)
			}
		case event := <-salesCh:
			if err := p.processSale(event); err != nil {
				log.Printf("error processing sales: %v", err)
			}
		case event := <-soldCh:
			if err := p.processSold(event); err != nil {
				log.Printf("error processing sold: %v", err)
			}
		case <-ctx.Done():
			return nil

		}
	}
}

func (p *NFTProcessor) processSold(event domain.NFTSoldEvent) error {
	nft, err := p.nftRepo.UpdateSoldByTokenId(event.TokenId, event.Owner)
	if err != nil {
		return fmt.Errorf("error with updating sold status: %w", err)
	}
	p.hub.Broadcast(struct {
		Type string `json:"type"`
		NFT  any    `json:"nft"`
	}{
		Type: "sold",
		NFT:  nft,
	})
	return nil
}

func (p *NFTProcessor) processSale(event domain.NFTInSaleEvent) error {
	nft, err := p.nftRepo.UpdateSaleByTokenId(event.TokenId, event.Price.String())
	if err != nil {
		return fmt.Errorf("error with updating sale status: %w", err)
	}
	p.hub.Broadcast(struct {
		Type string `json:"type"`
		NFT  any    `json:"nft"`
	}{
		Type: "sale",
		NFT:  nft,
	})
	return nil
}

func (p *NFTProcessor) processProposal(event domain.NFTProposedEvent) error {
	nft, err := p.nftRepo.UpdateProposedByTokenURI(event.ProposalID, event.TokenURI)
	log.Print(event)
	if err != nil {
		return fmt.Errorf("error with updating proposed status: %w", err)
	}

	p.hub.Broadcast(struct {
		Type string `json:"type"`
		NFT  any    `json:"nft"`
	}{
		Type: "proposed",
		NFT:  nft,
	})

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
	nft, err := p.nftRepo.UpdateId(event.TokenURI, event.TokenId)
	if err != nil {
		return fmt.Errorf("error with updating proposed status: %w", err)
	}

	p.hub.Broadcast((struct {
		Type string `json:"type"`
		NFT  any    `json:"nft"`
	}{
		Type: "minted",
		NFT:  nft,
	}))

	return nil
}
