package eth

import (
	"bazar/internal/contract"
	"bazar/internal/domain"
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type EventListener struct {
	contract *contract.Contract
}

func NewEthEventListener(contract *contract.Contract) *EventListener {
	return &EventListener{contract: contract}
}

func (l *EventListener) ListenNFTProposed(ctx context.Context) (<-chan domain.NFTProposedEvent, error) {
	eventCh := make(chan domain.NFTProposedEvent)
	rawCh := make(chan *contract.ContractNFTProposed)

	sub, err := l.contract.WatchNFTProposed(&bind.WatchOpts{Context: ctx}, rawCh, nil)
	if err != nil {
		return nil, fmt.Errorf("error watching NFT proposed event: %w", err)
	}

	go func() {
		defer sub.Unsubscribe()
		defer close(eventCh)

		for {
			select {
			case event := <-rawCh:
				eventCh <- domain.NFTProposedEvent{
					ProposalID: event.ProposalId.String(),
					Proposer:   event.Proposer.Hex(),
					TokenURI:   event.TokenURI,
				}
			case <-sub.Err():
				return
			case <-ctx.Done():
				return
			}
		}
	}()
	return eventCh, nil
}

func (l *EventListener) ListenNFTVoted(ctx context.Context) (<-chan domain.NFTVotedEvent, error) {
	eventCh := make(chan domain.NFTVotedEvent)
	rawCh := make(chan *contract.ContractVoted)

	sub, err := l.contract.WatchVoted(&bind.WatchOpts{Context: ctx}, rawCh, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("error watching NFT voted event: %w", err)
	}

	go func() {
		defer sub.Unsubscribe()
		defer close(eventCh)

		for {
			select {
			case event := <-rawCh:
				eventCh <- domain.NFTVotedEvent{
					ProposalID: event.ProposalId.String(),
					Voter:      event.Voter.Hex(),
					TokenURI:   event.TokenURI,
					Amount:     *event.Amount,
				}
			case <-sub.Err():
				return
			case <-ctx.Done():
				return
			}
		}
	}()
	return eventCh, nil
}

func (l *EventListener) ListendNFTMinted(ctx context.Context) (<-chan domain.NFTMintedEvent, error) {
	eventCh := make(chan domain.NFTMintedEvent)
	rawCh := make(chan *contract.ContractNFTMinted)

	sub, err := l.contract.WatchNFTMinted(&bind.WatchOpts{Context: ctx}, rawCh, nil)
	if err != nil {
		return nil, fmt.Errorf("error watching NFT voted event: %w", err)
	}

	go func() {
		defer sub.Unsubscribe()
		defer close(eventCh)

		for {
			select {
			case event := <-rawCh:
				eventCh <- domain.NFTMintedEvent{
					TokenId:  event.TokenId.String(),
					TokenURI: event.TokenURI,
					Owner:    event.Owner.Hex(),
				}
			case <-sub.Err():
				return
			case <-ctx.Done():
				return
			}
		}
	}()
	return eventCh, nil
}
