package entities

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// NFTProposal представляет предложение для NFT.
type NFTProposal struct {
	// TokenURI URI для токена
	// swagger:model
	TokenURI string `json:"token_uri"`

	// Proposer адрес пользователя, предложившего NFT
	// swagger:model
	Proposer common.Address `json:"proposer"`

	// Votes количество голосов за предложение
	// swagger:model
	Votes *big.Int `json:"votes"`

	// Minted флаг, указывающий, был ли токен скомпилирован
	// swagger:model
	Minted bool `json:"minted"`
}
