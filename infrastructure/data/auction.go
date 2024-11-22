package data_access

import (
	"auction-system/domain/entities"
	"auction-system/domain/repositories"
	"database/sql"
	"time"
)

type PostgresAuctionRepository struct {
	db *sql.DB
}

func NewPostgresAuctionRepository(db *sql.DB) repositories.AuctionRepository {
	return &PostgresAuctionRepository{db: db}
}

func (r *PostgresAuctionRepository) CreateAuction(auction *entities.Auction) error {
	query := `INSERT INTO auctions (lot_id, start_time, end_time) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(query, auction.LotID, auction.StartTime, auction.EndTime).Scan(&auction.ID)
	return err
}

func (r *PostgresAuctionRepository) GetAuctionByLotID(lotID int) (*entities.Auction, error) {
	query := `SELECT id, lot_id, start_time, end_time FROM auctions WHERE lot_id = $1`
	auction := &entities.Auction{}
	err := r.db.QueryRow(query, lotID).Scan(&auction.ID, &auction.LotID, &auction.StartTime, &auction.EndTime)
	if err != nil {
		return nil, err
	}
	return auction, nil
}

func (r *PostgresAuctionRepository) GetAuctionsEndingBefore(endTime time.Time) ([]*entities.Auction, error) {
	query := `SELECT id, lot_id, start_time, end_time FROM auctions WHERE end_time <= $1`
	rows, err := r.db.Query(query, endTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	auctions := []*entities.Auction{}
	for rows.Next() {
		auction := &entities.Auction{}
		if err := rows.Scan(&auction.ID, &auction.LotID, &auction.StartTime, &auction.EndTime); err != nil {
			return nil, err
		}
		auctions = append(auctions, auction)
	}
	return auctions, nil
}
