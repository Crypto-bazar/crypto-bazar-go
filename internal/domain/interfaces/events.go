package interfaces

import "bazar/internal/domain"

type EventHandler interface {
	OnTokenMinted(event *domain.TokenMintedEvent) error
	OnTokenListedForSale(event *domain.TokenListedForSaleEvent) error
	OnTokenSold(event *domain.TokenSold) error
}
