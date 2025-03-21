package usecase

import (
	"bazar/internal/domain"
	"bazar/internal/domain/requests"
	"log"
)

type NFTEventHandler struct {
	nftService domain.NFTService
}

func (n *NFTEventHandler) OnTokenListedForSale(event *domain.TokenListedForSaleEvent) error {
	req := &requests.UpdateTokenPriceReq{
		TokenId: event.TokenId, Price: event.Price.String(),
	}

	_, err := n.nftService.SetTokenPrice(req)
	if err != nil {
		return err
	}

	return nil
}

func (n *NFTEventHandler) OnTokenMinted(event *domain.TokenMintedEvent) error {
	req := &requests.UpdateTokenIdReq{
		TokenURI: event.TokenURI,
		TokenId:  event.TokenId.String(),
	}

	_, err := n.nftService.SetTokenId(req)
	if err != nil {
		log.Printf("Ошибка обновления NFT в БД: %v", err)
		return err
	}

	return nil
}

func (n *NFTEventHandler) OnTokenSold(event domain.TokenSold) error {
	panic("unimplemented")
}

func NewNFTEventHandler(nftService domain.NFTService) domain.EventHandler {
	return &NFTEventHandler{nftService: nftService}
}
