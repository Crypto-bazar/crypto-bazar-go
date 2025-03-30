package usecase

import (
	"bazar/internal/domain/entities"
	"bazar/internal/infrastructure/eth"
)

type TransactionUseCase struct {
	transaction *eth.Transaction
}

func NewTransactionUseCase(transaction *eth.Transaction) *TransactionUseCase {
	return &TransactionUseCase{
		transaction: transaction,
	}
}

func (t *TransactionUseCase) GetProposedNFTs() (*[]entities.NFTProposal, error) {
	proposals, err := t.transaction.GetProposedNFTs()
	if err != nil {
		return nil, err
	}

	return proposals, nil
}
