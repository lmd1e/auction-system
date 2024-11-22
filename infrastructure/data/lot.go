package data_access

import (
	"auction-system/domain/entities"
	"auction-system/domain/repositories"
	"database/sql"
)

type PostgresLotRepository struct {
	db *sql.DB
}

func NewPostgresLotRepository(db *sql.DB) repositories.LotRepository {
	return &PostgresLotRepository{db: db}
}

func (r *PostgresLotRepository) CreateLot(lot *entities.Lot) error {
	query := `INSERT INTO lots (name, description, start_price, seller_id) VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.db.QueryRow(query, lot.Name, lot.Description, lot.StartPrice, lot.SellerID).Scan(&lot.ID)
	return err
}

func (r *PostgresLotRepository) GetLotByID(id int) (*entities.Lot, error) {
	query := `SELECT id, name, description, start_price, seller_id FROM lots WHERE id = $1`
	lot := &entities.Lot{}
	err := r.db.QueryRow(query, id).Scan(&lot.ID, &lot.Name, &lot.Description, &lot.StartPrice, &lot.SellerID)
	if err != nil {
		return nil, err
	}
	return lot, nil
}
