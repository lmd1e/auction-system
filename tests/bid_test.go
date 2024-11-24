package tests

import (
	"auction-system/domain/entities"
	"auction-system/infrastructure/data_access"
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreateBid(t *testing.T) {
	db, err := sql.Open("postgres", "user=auction_user password=auction_pass dbname=auction_db sslmode=disable")
	assert.NoError(t, err)
	defer db.Close()

	repo := data_access.NewPostgresBidRepository(db)

	bid := &entities.Bid{
		Amount:   60.0,
		LotID:    1,
		BidderID: 2,
	}

	err = repo.CreateBid(bid)
	assert.NoError(t, err)
	assert.NotZero(t, bid.ID)
}

func TestGetBidsByLotID(t *testing.T) {
	db, err := sql.Open("postgres", "user=auction_user password=auction_pass dbname=auction_db sslmode=disable")
	assert.NoError(t, err)
	defer db.Close()

	repo := data_access.NewPostgresBidRepository(db)

	bids, err := repo.GetBidsByLotID(1)
	assert.NoError(t, err)
	assert.Len(t, bids, 2)
	assert.Equal(t, 60.0, bids[0].Amount)
	assert.Equal(t, 70.0, bids[1].Amount)
}
