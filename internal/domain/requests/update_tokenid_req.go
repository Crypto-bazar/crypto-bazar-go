package requests

import "math/big"

type UpdateTokenIdReq struct {
	TokenURI string `json:"token_URI"`
	TokenId *big.Int `json:"token_id"`
}
