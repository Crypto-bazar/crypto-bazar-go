package dto

type SignatureResponse struct {
	IsValid          bool   `json:"isValid"`
	RecoveredAddress string `json:"recoveredAddress"`
}
