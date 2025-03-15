package usecase

import (
	"bazar/internal/domain"
	"bazar/internal/domain/requests"
	"context"
)

type NFTEventHandler struct {
	nftService domain.NFTService
}

func (n *NFTEventHandler) OnTokenListedForSale(ctx context.Context, event domain.TokenListedForSaleEvent) error {
	panic("unimplemented")
}

// TODO доделать реализацию обновления id токена. Генерировать и искать по tokenURI
func (n *NFTEventHandler) OnTokenMinted(ctx context.Context, event domain.TokenMintedEvent) error {
	req := &requests.UpdateTokenIdReq{
		TokenURI: event.TokenId.String(),
		TokenId:  event.TokenURI,
	}

	
}

func (n *NFTEventHandler) OnTokenSold(ctx context.Context, event domain.TokenSold) error {
	panic("unimplemented")
}

func NewNFTEventHandler(nftService domain.NFTService) domain.EventHandler {
	return &NFTEventHandler{nftService: nftService}
}
