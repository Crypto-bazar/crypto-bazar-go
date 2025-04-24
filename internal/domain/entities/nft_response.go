package entities

type NFTResponse struct {
	ID          int    `json:"id" db:"id"`
	TokenID     int    `json:"token_id" db:"token_id"`
	ProposalID  int    `json:"proposal_id" db:"proposal_id"`
	TokenURI    string `json:"token_uri" db:"token_uri"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Price       string `json:"price" db:"price"`
	Owner       string `json:"owner" db:"owner"`
	ImagePath   string `json:"image_path" db:"image_path"`
	InSales     bool   `json:"in_sales" db:"in_sales"`
	Proposed    bool   `json:"proposed" db:"proposed"`
	VotesAmount string `json:"votes_amount" db:"votes_amount"`
}
