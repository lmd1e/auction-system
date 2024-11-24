package grpc

import (
	"auction-system/api/grpc/proto/auction"
	"auction-system/domain/entities"
	"auction-system/domain/repositories"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuctionService struct {
	userRepo    repositories.UserRepository
	lotRepo     repositories.LotRepository
	bidRepo     repositories.BidRepository
	auctionRepo repositories.AuctionRepository
}

func NewAuctionService(userRepo repositories.UserRepository, lotRepo repositories.LotRepository, bidRepo repositories.BidRepository, auctionRepo repositories.AuctionRepository) *AuctionService {
	return &AuctionService{
		userRepo:    userRepo,
		lotRepo:     lotRepo,
		bidRepo:     bidRepo,
		auctionRepo: auctionRepo,
	}
}

func (s *AuctionService) CreateUser(ctx context.Context, req *auction.CreateUserRequest) (*auction.CreateUserResponse, error) {
	user := &entities.User{
		Name:    req.Name,
		Balance: req.Balance,
	}
	if err := s.userRepo.CreateUser(user); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}
	return &auction.CreateUserResponse{Id: int32(user.ID)}, nil
}

func (s *AuctionService) CreateLot(ctx context.Context, req *auction.CreateLotRequest) (*auction.CreateLotResponse, error) {
	lot := &entities.Lot{
		Name:        req.Name,
		Description: req.Description,
		StartPrice:  req.StartPrice,
		SellerID:    int(req.SellerId),
	}
	if err := s.lotRepo.CreateLot(lot); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create lot: %v", err)
	}
	return &auction.CreateLotResponse{Id: int32(lot.ID)}, nil
}

func (s *AuctionService) PlaceBid(ctx context.Context, req *auction.PlaceBidRequest) (*auction.PlaceBidResponse, error) {
	bid := &entities.Bid{
		Amount:   req.Amount,
		LotID:    int(req.LotId),
		BidderID: int(req.BidderId),
	}
	if err := s.bidRepo.CreateBid(bid); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to place bid: %v", err)
	}
	return &auction.PlaceBidResponse{Id: int32(bid.ID)}, nil
}

func (s *AuctionService) CloseAuction(ctx context.Context, req *auction.CloseAuctionRequest) (*auction.CloseAuctionResponse, error) {
	lotID := int(req.LotId)
	auction, err := s.auctionRepo.GetAuctionByLotID(lotID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "auction not found: %v", err)
	}

	bids, err := s.bidRepo.GetBidsByLotID(lotID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get bids: %v", err)
	}

	var winnerID int
	var maxBid float64
	for _, bid := range bids {
		if bid.Amount > maxBid {
			maxBid = bid.Amount
			winnerID = bid.BidderID
		}
	}

	return &auction.CloseAuctionResponse{WinnerId: int32(winnerID)}, nil
}
