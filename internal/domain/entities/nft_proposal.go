package entities

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type NFTProposal struct {
	TokenURI string
	Proposer common.Address
	Votes    *big.Int
	Minted   bool
}
