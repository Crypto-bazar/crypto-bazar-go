package requests

type RemoveFavouriteNFT struct {
    // Ethereum адрес пользователя
    // required: true
    // example: 0x71C7656EC7ab88b098defB751B7401B5f6d8976F
    EthAddress string `json:"eth_address"`
    
    // ID NFT для удаления из избранного
    // required: true
    // example: 123
    NftId string `json:"nft_id"`
}