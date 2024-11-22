package repositories

import "auction-system/domain/entities"

type UserRepository interface {
	CreateUser(user *entities.User) error
	GetUserByID(id int) (*entities.User, error)
	UpdateUser(user *entities.User) error
}
