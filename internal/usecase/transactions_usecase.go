package usecase

import (
	"bazar/internal/domain/entities"
	"bazar/internal/infrastructure/eth"
	"context"
)

type TransactionUseCase struct {
	transaction *eth.Transaction
	ctx         context.Context
}

func NewTransactionUseCase(transaction *eth.Transaction, ctx context.Context) *TransactionUseCase {
	return &TransactionUseCase{
		transaction: transaction,
		ctx:         ctx,
	}
}

func (t *TransactionUseCase) GetProposedNFTs() (*[]entities.NFTProposal, error) {
	proposals, err := t.transaction.GetProposedNFTs(t.ctx)
	if err != nil {
		return nil, err
	}

	return proposals, nil
}
