package dto

type UpdateTokenAddressReq struct {
	Id              uint   `json:"id"`
	ContractAddress string `json:"contract_address"`
}
