package requests

type UpdateTokenStatusReq struct {
	Status  bool   `db:"status" json:"status"`
	TokenID uint64 `db:"token_id" json:"token_id"`
}
