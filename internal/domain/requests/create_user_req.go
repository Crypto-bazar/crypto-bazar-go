package requests

type CreateUserRequest struct {
	EthAddress string `json:"eth_address"`
	Signature  string `json:"signature"`
	Message    string `json:"message"`
}
