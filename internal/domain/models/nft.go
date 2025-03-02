package models

type NFT struct {
	ID      uint   `db:"id"`
	TokenID string `db:"token_id"`
	Owner   uint `db:"owner_id"`
}
