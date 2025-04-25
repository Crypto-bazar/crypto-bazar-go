package eth

import (
	"bazar/internal/contract"
	"bazar/internal/domain/entities"
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type Transaction struct {
	contract *contract.Contract
}

func NewTransaction(contractInstance *contract.Contract) *Transaction {
	return &Transaction{contract: contractInstance}
}

func (t *Transaction) GetProposedNFTs(ctx context.Context) (*[]entities.NFTProposal, error) {
	proposals, err := t.contract.GetProposeNFT(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	var result []entities.NFTProposal
	for _, p := range proposals {
		result = append(result, entities.NFTProposal{
			Minted:   p.Minted,
			Votes:    p.Votes,
			Proposer: p.Proposer,
			TokenURI: p.TokenURI,
		})
	}

	return &result, nil
}
