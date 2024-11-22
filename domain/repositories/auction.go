package repositories

import (
	"auction-system/domain/entities"
	"time"
)

type AuctionRepository interface {
	CreateAuction(auction *entities.Auction) error
	GetAuctionByLotID(lotID int) (*entities.Auction, error)
	GetAuctionsEndingBefore(endTime time.Time) ([]*entities.Auction, error)
}
