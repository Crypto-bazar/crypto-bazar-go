package requests

type UpdateTokenPriceReq struct {
	TokenId uint64 `json:"token_id"`
	Price   string `json:"price"`
}
