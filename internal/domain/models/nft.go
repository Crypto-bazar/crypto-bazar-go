package models

type NFT struct {
	ID        uint   `db:"id" json:"id"`
	TokenID   string `db:"token_id" form:"token_id"`
	Owner     uint   `db:"owner_id" form:"owner_id"`
	ImagePath string `db:"image_path" json:"image_path"`
}
