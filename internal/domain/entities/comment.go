package entities

import "time"

// Comment представляет комментарий к NFT.
type Comment struct {
    // ID уникальный идентификатор комментария
    // swagger:model
    ID uint `db:"id" json:"id"`

    // NFTID идентификатор NFT, к которому относится комментарий
    // swagger:model
    NFTID uint `db:"nft_id" json:"nft_id"`

    // OwnerId идентификатор пользователя, создавшего комментарий
    // swagger:model
    OwnerId uint `db:"owner_id" json:"owner_id"`

    // Content текст комментария
    // swagger:model
    Content string `db:"content" json:"content"`

    // CreatedAt дата создания комментария
    // swagger:model
    CreatedAt time.Time `db:"created_at" json:"created_at"`
}
