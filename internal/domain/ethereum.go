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

type NFTMintedEvent struct {
	TokenId  string
	TokenURI string
	Owner    string
}

type NFTInSaleEvent struct {
	TokenId string
	Owner   string
	Price   *big.Int
}

type NFTSoldEvent struct {
	TokenId string
	Owner   string
	Price   *big.Int
}

type NFTEventListener interface {
	ListenNFTProposed(ctx context.Context) (eventCh <-chan NFTProposedEvent, errCh <-chan error, err error)
	ListenNFTVoted(ctx context.Context) (eventCh <-chan NFTVotedEvent, errCh <-chan error, err error)
	ListenedNFTMinted(ctx context.Context) (eventCh <-chan NFTMintedEvent, errCh <-chan error, err error)
	ListenNFTInSale(ctx context.Context) (eventCh <-chan NFTInSaleEvent, errCh <-chan error, err error)
	ListenNFTSold(ctx context.Context) (eventCh <-chan NFTSoldEvent, errCh <-chan error, err error)
}
