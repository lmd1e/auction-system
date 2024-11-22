package repositories

import "auction-system/domain/entities"

type LotRepository interface {
	CreateLot(lot *entities.Lot) error
	GetLotByID(id int) (*entities.Lot, error)
}
