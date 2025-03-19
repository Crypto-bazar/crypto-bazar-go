package usecase

import (
	"bazar/internal/domain"
	"bazar/internal/domain/requests"
	"context"
	"log"
)

type NFTEventHandler struct {
	nftService domain.NFTService
}

func (n *NFTEventHandler) OnTokenListedForSale(ctx context.Context, event *domain.TokenListedForSaleEvent) error {
	req := &requests.UpdateTokenPriceReq{
		TokenId: event.TokenId, Price: event.Price.String(),
	}

	_, err := n.nftService.SetTokenPrice(req)
	if err != nil {
		return err
	}

	return nil
}

func (n *NFTEventHandler) OnTokenMinted(ctx context.Context, event *domain.TokenMintedEvent) error {
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

func (n *NFTEventHandler) OnTokenSold(ctx context.Context, event domain.TokenSold) error {
	panic("unimplemented")
}

func NewNFTEventHandler(nftService domain.NFTService) domain.EventHandler {
	return &NFTEventHandler{nftService: nftService}
}
