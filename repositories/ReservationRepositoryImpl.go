package repositories

import (
	"errors"
	"villa_go/models/entities"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ReservationRepository interface {
	CreateNewReservation(entities.Reservation) (*entities.Reservation, error)
	GetReservationById(uuid.UUID)
}

type ReservationRepositoryImpl struct {
	db *gorm.DB
}

func NewReservationRepositoryImpl(Db *gorm.DB) ReservationRepository {
	return &ReservationRepositoryImpl{
		db: Db,
	}
}

func (r *ReservationRepositoryImpl) CreateNewReservation(entitiy entities.Reservation) (*entities.Reservation, error) {

	if CreateException := r.db.Create(&entitiy); CreateException != nil {
		return nil, errors.New("Error when create reservation")
	}

	return &entitiy, nil

}

func (r *ReservationRepositoryImpl) GetReservationById(id uuid.UUID) {

}
