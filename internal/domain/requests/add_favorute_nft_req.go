package requests

type AddFavouriteNFT struct {
	EthAddress string `json:"eth_address"`
	NftId  string `json:"nft_id"`
}
