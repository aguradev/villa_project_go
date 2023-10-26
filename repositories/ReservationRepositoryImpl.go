package repositories

import (
	"errors"
	"villa_go/models/entities"
	"villa_go/payloads/resources"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ReservationRepository interface {
	CreateNewReservation(entities.Reservation) (*entities.Reservation, error)
	GetReservationById(uuid.UUID) (*resources.ReservationResource, error)
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

	if CreateException := r.db.Create(&entitiy); CreateException.Error != nil {
		return nil, errors.New("Error when create reservation")
	}

	return &entitiy, nil

}

func (r *ReservationRepositoryImpl) GetReservationById(id uuid.UUID) (*resources.ReservationResource, error) {

	var Reservation entities.Reservation
	var MappingReservation resources.ReservationResource

	if FindReservationDataException := r.db.Joins("User").Joins("Reservation_detail").Joins("Reservation_detail.Villa").Joins("Reservation_detail.Villa.Location").First(&Reservation, "reservations.id = ?", id); FindReservationDataException.Error != nil {
		return nil, errors.New("Reservation data not found")
	}

	MappingReservation.GetDetailReservationResponse(Reservation)

	return &MappingReservation, nil

}
