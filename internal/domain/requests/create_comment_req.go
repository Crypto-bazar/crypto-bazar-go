package requests

type CreateCommentReq struct {
	TokenId      uint   `json:"token_id"`
	OwnerComment uint   `json:"owner_comment"`
	Content      string `json:"content"`
}
