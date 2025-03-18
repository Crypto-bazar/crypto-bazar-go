package requests

type CreateNFTRequest struct {
	TokenID      string `form:"token_id"`
	Name         string `form:"name"`
	Description  string `form:"description"`
	OwnerAddress string `form:"owner_address"`
}
