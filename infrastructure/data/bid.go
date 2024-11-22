package data_access

import (
	"auction-system/domain/entities"
	"auction-system/domain/repositories"
	"database/sql"
)

type PostgresBidRepository struct {
	db *sql.DB
}

func NewPostgresBidRepository(db *sql.DB) repositories.BidRepository {
	return &PostgresBidRepository{db: db}
}

func (r *PostgresBidRepository) CreateBid(bid *entities.Bid) error {
	query := `INSERT INTO bids (amount, lot_id, bidder_id) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(query, bid.Amount, bid.LotID, bid.BidderID).Scan(&bid.ID)
	return err
}

func (r *PostgresBidRepository) GetBidsByLotID(lotID int) ([]*entities.Bid, error) {
	query := `SELECT id, amount, lot_id, bidder_id FROM bids WHERE lot_id = $1`
	rows, err := r.db.Query(query, lotID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bids := []*entities.Bid{}
	for rows.Next() {
		bid := &entities.Bid{}
		if err := rows.Scan(&bid.ID, &bid.Amount, &bid.LotID, &bid.BidderID); err != nil {
			return nil, err
		}
		bids = append(bids, bid)
	}
	return bids, nil
}
