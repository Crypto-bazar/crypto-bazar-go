package responses

type SignatureResponse struct {
	IsValid          bool   `json:"isValid"`
	RecoveredAddress string `json:"recoveredAddress"`
}
