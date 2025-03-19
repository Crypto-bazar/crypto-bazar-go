package requests

type UpdateTokenIdReq struct {
	TokenURI string `json:"token_URI"`
	TokenId  string `json:"token_id"`
}
