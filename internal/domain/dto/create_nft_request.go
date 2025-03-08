package dto

type CreateNFTRequest struct {
	TokenID      string `form:"token_id"`
	Name         string `form:"name"`
	Description  string `form:"description"`
	Price        string `form:"price"`
	OwnerAddress string `form:"owner_address"`
}
