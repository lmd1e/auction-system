package tests

import (
	"auction-system/domain/entities"
	"auction-system/infrastructure/data_access"
	"database/sql"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreateAuction(t *testing.T) {
	db, err := sql.Open("postgres", "user=auction_user password=auction_pass dbname=auction_db sslmode=disable")
	assert.NoError(t, err)
	defer db.Close()

	repo := data_access.NewPostgresAuctionRepository(db)

	auction := &entities.Auction{
		LotID:     1,
		StartTime: time.Now(),
		EndTime:   time.Now().Add(time.Hour),
	}

	err = repo.CreateAuction(auction)
	assert.NoError(t, err)
	assert.NotZero(t, auction.ID)
}

func TestGetAuctionByLotID(t *testing.T) {
	db, err := sql.Open("postgres", "user=auction_user password=auction_pass dbname=auction_db sslmode=disable")
	assert.NoError(t, err)
	defer db.Close()

	repo := data_access.NewPostgresAuctionRepository(db)

	auction, err := repo.GetAuctionByLotID(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, auction.LotID)
	assert.NotZero(t, auction.StartTime)
	assert.NotZero(t, auction.EndTime)
}
