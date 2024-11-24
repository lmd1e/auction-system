package tests

import (
	"auction-system/domain/entities"
	"auction-system/infrastructure/data_access"
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db, err := sql.Open("postgres", "user=auction_user password=auction_pass dbname=auction_db sslmode=disable")
	assert.NoError(t, err)
	defer db.Close()

	repo := data_access.NewPostgresUserRepository(db)

	user := &entities.User{
		Name:    "Test User",
		Balance: 100.0,
	}

	err = repo.CreateUser(user)
	assert.NoError(t, err)
	assert.NotZero(t, user.ID)
}

func TestGetUserByID(t *testing.T) {
	db, err := sql.Open("postgres", "user=auction_user password=auction_pass dbname=auction_db sslmode=disable")
	assert.NoError(t, err)
	defer db.Close()

	repo := data_access.NewPostgresUserRepository(db)

	user, err := repo.GetUserByID(1)
	assert.NoError(t, err)
	assert.Equal(t, "Test User 1", user.Name)
	assert.Equal(t, 100.0, user.Balance)
}
