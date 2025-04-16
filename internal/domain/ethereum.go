package domain

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type TokenMintedEvent struct {
	TokenId  *big.Int
	Owner    common.Address
	TokenURI string
}

type TokenListedForSaleEvent struct {
	TokenId uint64
	Price   *big.Int
	Seller  common.Address
}

type TokenSold struct {
	TokenId uint64
	Buyer   common.Address
	Price   *big.Int
}

type NFTProposedEvent struct {
	ProposalID string
	Proposer   string
	TokenURI   string
}

type NFTVotedEvent struct {
	ProposalID string
	Voter      string
	TokenURI   string
	Amount     big.Int
}

type NFTEventListener interface {
	ListenNFTProposed(ctx context.Context) (<-chan NFTProposedEvent, error)
	ListenNFTVoted(ctx context.Context) (<-chan NFTVotedEvent, error)
}
