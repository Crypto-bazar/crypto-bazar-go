package requests

type CreateCommentReq struct {
	TokenId      uint   `json:"token_id"`
	OwnerAddress string   `json:"owner_address"`
	Content      string `json:"content"`
}
