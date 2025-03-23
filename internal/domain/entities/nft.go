package entities

import "time"

type NFT struct {
	ID          uint       `db:"id" json:"id"`
	TokenID     uint       `db:"token_id" json:"token_id"`
	TokenURI    string     `db:"token_uri" json:"token_uri"`
	Name        string     `db:"name" json:"name"`
	Description string     `db:"description" json:"description"`
	Price       string     `db:"price" json:"price"`
	Owner       uint       `db:"owner_id" json:"owner_id"`
	ImagePath   string     `db:"image_path" json:"image_path"`
	Sales       bool       `db:"in_sales" json:"in_sales"`
	CreatedAt   *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at"`
}
