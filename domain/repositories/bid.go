package repositories

import "auction-system/domain/entities"

type BidRepository interface {
	CreateBid(bid *entities.Bid) error
	GetBidsByLotID(lotID int) ([]*entities.Bid, error)
}
