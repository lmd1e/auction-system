package tests

import (
	"auction-system/domain/entities"
	"auction-system/infrastructure/data_access"
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreateLot(t *testing.T) {
	db, err := sql.Open("postgres", "user=auction_user password=auction_pass dbname=auction_db sslmode=disable")
	assert.NoError(t, err)
	defer db.Close()

	repo := data_access.NewPostgresLotRepository(db)

	lot := &entities.Lot{
		Name:        "Test Lot",
		Description: "Test Description",
		StartPrice:  50.0,
		SellerID:    1,
	}

	err = repo.CreateLot(lot)
	assert.NoError(t, err)
	assert.NotZero(t, lot.ID)
}

func TestGetLotByID(t *testing.T) {
	db, err := sql.Open("postgres", "user=auction_user password=auction_pass dbname=auction_db sslmode=disable")
	assert.NoError(t, err)
	defer db.Close()

	repo := data_access.NewPostgresLotRepository(db)

	lot, err := repo.GetLotByID(1)
	assert.NoError(t, err)
	assert.Equal(t, "Test Lot 1", lot.Name)
	assert.Equal(t, "Description 1", lot.Description)
	assert.Equal(t, 50.0, lot.StartPrice)
	assert.Equal(t, 1, lot.SellerID)
}
