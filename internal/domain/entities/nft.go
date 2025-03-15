package entities

type NFT struct {
	ID          uint   `db:"id" json:"id"`
	TokenID     uint   `db:"token_id" json:"token_id"`
	TokenURI    string `db:"token_uri" json:"token_uri"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	Price       string `db:"price" json:"price"`
	Owner       uint   `db:"owner_id" json:"owner_id"`
	ImagePath   string `db:"image_path" json:"image_path"`
}
