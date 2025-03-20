package requests

type GetUserNFTReq struct {
	UserAddress string `jsob:"eth_address" db:"eth_address"`
}
