package models

type NFT struct {
	ID      uint   `db:"id"`
	TokenID string `db:"token_id"`
	Owner   string `db:"owner_id"`
}
