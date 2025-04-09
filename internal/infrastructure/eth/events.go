package eth

import (
	"bazar/internal/contract"
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func ListenNFTProposed(ctx context.Context, instance *contract.Mycontract) {
	eventCh := make(chan *contract.MycontractNFTProposed)

	sub, err := instance.WatchNFTProposed(&bind.WatchOpts{
		Context: ctx,
	}, eventCh, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe: %v", err)
	}
	defer sub.Unsubscribe()

	fmt.Println("Listening for NFTProposed events...")
	for {
		select {
		case event := <-eventCh:
			fmt.Printf("Proposal: %s\n Proposer: %s\n URI: %s\n", event.ProposalId, event.Proposer, event.TokenURI)
		case err := <-sub.Err():
			log.Fatalf("Subscription error: %v", err)
		}
	}
}
