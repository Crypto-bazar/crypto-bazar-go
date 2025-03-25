package entities

import "time"

type Comment struct {
	ID        uint      `db:"id" json:"id"`
	NFTID     uint      `db:"nft_id" json:"nft_id"`
	OwnerId   uint      `db:"owner_id" json:"owner_id"`
	Content   string    `db:"content" json:"content"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
