package domain

import (
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
