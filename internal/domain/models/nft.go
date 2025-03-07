package models

type NFT struct {
	ID          uint   `db:"id" json:"id"`
	TokenID     string `db:"token_id" form:"token_id"`
	Name        string `db:"name" form:"name"`
	Description string `db:"description" form:"description"`
	Price       string `db:"price" form:"price"`
	Owner       uint   `db:"owner_id" form:"owner_id"`
	ImagePath   string `db:"image_path" json:"image_path"`
}
