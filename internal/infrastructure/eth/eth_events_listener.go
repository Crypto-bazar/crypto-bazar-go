package eth

import (
	"bazar/internal/contract"
	"bazar/internal/domain"
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type EthEventListener struct {
	contract *contract.Mycontract
}

func NewEthEventListener(contract *contract.Mycontract) *EthEventListener {
	return &EthEventListener{contract: contract}
}

func (l *EthEventListener) ListenNFTProposed(ctx context.Context) (<-chan domain.NFTProposedEvent, error) {
	eventCh := make(chan domain.NFTProposedEvent)
	rawCh := make(chan *contract.MycontractNFTProposed)

	sub, err := l.contract.WatchNFTProposed(&bind.WatchOpts{Context: ctx}, rawCh, nil)
	if err != nil {
		return nil, err
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
