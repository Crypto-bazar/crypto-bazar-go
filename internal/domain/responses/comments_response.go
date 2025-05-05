package responses

import "time"

// CommentResponse Comment представляет комментарий к NFT.
type CommentResponse struct {
	// ID уникальный идентификатор комментария
	// swagger:model
	ID uint `db:"id" json:"id"`

	// NFTID идентификатор NFT, к которому относится комментарий
	// swagger:model
	NFTID uint `db:"nft_id" json:"nft_id"`

	// OwnerId идентификатор пользователя, создавшего комментарий
	// swagger:model
	OwnerAddress string `db:"owner_address" json:"owner_address"`

	// Content текст комментария
	// swagger:model
	Content string `db:"content" json:"content"`

	// CreatedAt дата создания комментария
	// swagger:model
	CreatedAt time.Time `db:"created_at" json:"created_at"`

	// AvatarUrl URL аватара пользователя
	// swagger:model
	AvatarUrl string `db:"avatar_url" json:"avatar_url"`
}
