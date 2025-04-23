package entities

// NFT представляет сам NFT.
type NFT struct {
	// ID уникальный идентификатор NFT
	// swagger:model
	ID uint `db:"id" json:"id"`

	// TokenID идентификатор токена
	// swagger:model
	TokenID uint `db:"token_id" json:"token_id"`

	// ProposalId идентификатор предложения для этого NFT
	// swagger:model
	ProposalId uint `db:"proposal_id" json:"proposal_id"`

	// TokenURI URI для токена
	// swagger:model
	TokenURI string `db:"token_uri" json:"token_uri"`

	// Name название NFT
	// swagger:model
	Name string `db:"name" json:"name"`

	// Description описание NFT
	// swagger:model
	Description string `db:"description" json:"description"`

	// Price цена NFT
	// swagger:model
	Price string `db:"price" json:"price"`

	// Owner идентификатор владельца
	// swagger:model
	Owner uint `db:"owner_id" json:"owner_id"`

	// ImagePath путь к изображению для NFT
	// swagger:model
	ImagePath string `db:"image_path" json:"image_path"`

	// Sales флаг, указывающий, продается ли NFT
	// swagger:model
	Sales bool `db:"in_sales" json:"in_sales"`

	// Proposed флаг, указывающий, предложен ли NFT
	// swagger:model
	Proposed bool `db:"proposed" json:"proposed"`

	// Votes количество голосов за этот NFT
	// swagger:model
	Votes string `db:"votes_amount" json:"votes_amount"`
}
