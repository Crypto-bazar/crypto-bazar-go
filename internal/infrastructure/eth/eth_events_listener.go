package eth

import (
	"bazar/internal/contract"
	"bazar/internal/domain"
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/event"
)

type EventListener struct {
	contract *contract.Contract
}

func NewEthEventListener(contract *contract.Contract) *EventListener {
	return &EventListener{contract: contract}
}

func subscribeWithErrorHandling[T any, R any](
	ctx context.Context,
	subscribeFn func(*bind.WatchOpts, chan R) (event.Subscription, error),
	convertFn func(R) T,
) (<-chan T, <-chan error, error) {
	eventCh := make(chan T)
	errCh := make(chan error, 1)
	rawCh := make(chan R)

	sub, err := subscribeFn(&bind.WatchOpts{Context: ctx}, rawCh)
	if err != nil {
		return nil, nil, fmt.Errorf("subscribe error: %w", err)
	}

	go func() {
		defer close(eventCh)
		defer sub.Unsubscribe()

		for {
			select {
			case raw := <-rawCh:
				eventCh <- convertFn(raw)
			case err := <-sub.Err():
				errCh <- err
				return
			case <-ctx.Done():
				return
			}
		}
	}()

	return eventCh, errCh, nil
}

func (l *EventListener) ListenNFTProposed(ctx context.Context) (<-chan domain.NFTProposedEvent, <-chan error, error) {
	return subscribeWithErrorHandling(
		ctx,
		func(opts *bind.WatchOpts, ch chan *contract.ContractNFTProposed) (event.Subscription, error) {
			return l.contract.WatchNFTProposed(opts, ch, nil)
		},
		func(ev *contract.ContractNFTProposed) domain.NFTProposedEvent {
			return domain.NFTProposedEvent{
				ProposalID: ev.ProposalId.String(),
				Proposer:   ev.Proposer.Hex(),
				TokenURI:   ev.TokenURI,
			}
		},
	)
}

func (l *EventListener) ListenNFTSold(ctx context.Context) (<-chan domain.NFTSoldEvent, <-chan error, error) {
	return subscribeWithErrorHandling(
		ctx,
		func(opts *bind.WatchOpts, ch chan *contract.ContractNFTSold) (event.Subscription, error) {
			return l.contract.WatchNFTSold(opts, ch, nil, nil)
		},
		func(ev *contract.ContractNFTSold) domain.NFTSoldEvent {
			return domain.NFTSoldEvent{
				TokenId: ev.TokenId.String(),
				Owner:   ev.Buyer.Hex(),
				Price:   ev.Price,
			}
		},
	)
}

func (l *EventListener) ListenNFTVoted(ctx context.Context) (<-chan domain.NFTVotedEvent, <-chan error, error) {
	return subscribeWithErrorHandling(
		ctx,
		func(opts *bind.WatchOpts, ch chan *contract.ContractVoted) (event.Subscription, error) {
			return l.contract.WatchVoted(opts, ch, nil, nil)
		},
		func(ev *contract.ContractVoted) domain.NFTVotedEvent {
			return domain.NFTVotedEvent{
				ProposalID: ev.ProposalId.String(),
				Voter:      ev.Voter.Hex(),
				TokenURI:   ev.TokenURI,
				Amount:     *ev.Amount,
			}
		},
	)
}

func (l *EventListener) ListenedNFTMinted(ctx context.Context) (<-chan domain.NFTMintedEvent, <-chan error, error) {
	return subscribeWithErrorHandling(
		ctx,
		func(opts *bind.WatchOpts, ch chan *contract.ContractNFTMinted) (event.Subscription, error) {
			return l.contract.WatchNFTMinted(opts, ch, nil)
		},
		func(ev *contract.ContractNFTMinted) domain.NFTMintedEvent {
			return domain.NFTMintedEvent{
				TokenId:  ev.TokenId.String(),
				TokenURI: ev.TokenURI,
				Owner:    ev.Owner.Hex(),
			}
		},
	)
}

func (l *EventListener) ListenNFTInSale(ctx context.Context) (<-chan domain.NFTInSaleEvent, <-chan error, error) {
	return subscribeWithErrorHandling(
		ctx,
		func(opts *bind.WatchOpts, ch chan *contract.ContractNFTInSale) (event.Subscription, error) {
			return l.contract.WatchNFTInSale(opts, ch, nil, nil)
		},
		func(ev *contract.ContractNFTInSale) domain.NFTInSaleEvent {
			return domain.NFTInSaleEvent{
				TokenId: ev.TokenId.String(),
				Owner:   ev.Owner.Hex(),
				Price:   ev.Price,
			}
		},
	)
}
