package data

import (
	"auction-system/domain/entities"
	"auction-system/domain/repositories"
	"database/sql"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) repositories.UserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) CreateUser(user *entities.User) error {
	query := `INSERT INTO users (name, balance) VALUES ($1, $2) RETURNING id`
	err := r.db.QueryRow(query, user.Name, user.Balance).Scan(&user.ID)
	return err
}

func (r *PostgresUserRepository) GetUserByID(id int) (*entities.User, error) {
	query := `SELECT id, name, balance FROM users WHERE id = $1`
	user := &entities.User{}
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Balance)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *PostgresUserRepository) UpdateUser(user *entities.User) error {
	query := `UPDATE users SET name = $1, balance = $2 WHERE id = $3`
	_, err := r.db.Exec(query, user.Name, user.Balance, user.ID)
	return err
}
