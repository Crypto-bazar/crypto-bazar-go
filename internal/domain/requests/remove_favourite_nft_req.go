package requests

type RemoveFavouriteNFT struct {
	EthAddress string `json:"eth_address"`
	NftId  string `json:"nft_id"`
}
