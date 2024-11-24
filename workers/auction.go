package workers

import (
	"auction-system/domain/repositories"
	"context"
	"log"
	"time"
)

type AuctionWorker struct {
	auctionRepo repositories.AuctionRepository
	bidRepo     repositories.BidRepository
	userRepo    repositories.UserRepository
}

func NewAuctionWorker(auctionRepo repositories.AuctionRepository, bidRepo repositories.BidRepository, userRepo repositories.UserRepository) *AuctionWorker {
	return &AuctionWorker{
		auctionRepo: auctionRepo,
		bidRepo:     bidRepo,
		userRepo:    userRepo,
	}
}

func (w *AuctionWorker) Start(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			w.processEndedAuctions(ctx)
		}
	}
}

func (w *AuctionWorker) processEndedAuctions(ctx context.Context) {
	now := time.Now()
	auctions, err := w.auctionRepo.GetAuctionsEndingBefore(now)
	if err != nil {
		log.Printf("failed to get ended auctions: %v", err)
		return
	}

	for _, auction := range auctions {
		bids, err := w.bidRepo.GetBidsByLotID(auction.LotID)
		if err != nil {
			log.Printf("failed to get bids for auction %d: %v", auction.ID, err)
			continue
		}

		var winnerID int
		var maxBid float64
		for _, bid := range bids {
			if bid.Amount > maxBid {
				maxBid = bid.Amount
				winnerID = bid.BidderID
			}
		}

		if winnerID > 0 {
			_, err := w.userRepo.GetUserByID(winnerID)
			if err != nil {
				log.Printf("failed to get winner user %d: %v", winnerID, err)
				continue
			}

			// Process transaction and notify participants
			log.Printf("Auction %d ended. Winner: %d, Amount: %f", auction.ID, winnerID, maxBid)
		} else {
			log.Printf("Auction %d ended with no winner", auction.ID)
		}
	}
}
