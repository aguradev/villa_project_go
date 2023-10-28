package repositories

import (
	"errors"
	"villa_go/models/entities"
	"villa_go/payloads/request"
	"villa_go/payloads/resources"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ReservationRepository interface {
	GetListReservation() ([]resources.ReservationResource, error)
	CreateNewReservation(entities.Reservation) (*entities.Reservation, error)
	GetReservationById(uuid.UUID) (*resources.ReservationResource, error)
	UpdateSnapUrlReservation(uuid.UUID, request.ReservationRequest) (*resources.ReservationResource, error)
	UpdateStatusReservation(uuid.UUID, string) (bool, error)
}

type ReservationRepositoryImpl struct {
	db *gorm.DB
}

func NewReservationRepositoryImpl(Db *gorm.DB) ReservationRepository {
	return &ReservationRepositoryImpl{
		db: Db,
	}
}

func (r *ReservationRepositoryImpl) GetListReservation() ([]resources.ReservationResource, error) {

	var Reservation []entities.Reservation

	GetReservationErr := r.db.Preload("User").Preload("Reservation_detail.Villa").Find(&Reservation)

	if GetReservationErr.RowsAffected == 0 {
		return nil, errors.New("Reservation record empty")
	}

	if GetReservationErr.Error != nil {
		return nil, GetReservationErr.Error
	}

	ReservationResponse := resources.GetListReservationResponse(Reservation)

	return ReservationResponse, nil

}

func (r *ReservationRepositoryImpl) CreateNewReservation(entitiy entities.Reservation) (*entities.Reservation, error) {

	BeginTransaction := r.db.Begin()

	if CreateException := BeginTransaction.Create(&entitiy); CreateException.Error != nil {
		BeginTransaction.Rollback()
		return nil, errors.New("Error when create reservation")
	}

	BeginTransaction.Commit()

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

func (r *ReservationRepositoryImpl) UpdateSnapUrlReservation(id uuid.UUID, reservationdetailRequest request.ReservationRequest) (*resources.ReservationResource, error) {

	var Reservation entities.Reservation

	if CheckReservationExists := r.db.Preload("Reservation_detail").First(&Reservation, "id = ?", id); CheckReservationExists.Error != nil {
		return nil, errors.New("Reservation not found")
	}

	Reservation.Reservation_detail.SnapURL = reservationdetailRequest.ReservationDetail.SnapURL
	UpdateQueryException := r.db.Model(&Reservation.Reservation_detail).Updates(&Reservation.Reservation_detail)

	if UpdateQueryException.Error != nil {
		return nil, errors.New("Error when updating reservation")
	}

	GetResponseReservation, ErrResponse := r.GetReservationById(id)

	if ErrResponse != nil {
		return nil, errors.New("Failed to get reservation response")
	}

	return GetResponseReservation, nil

}

func (r *ReservationRepositoryImpl) UpdateStatusReservation(id uuid.UUID, status string) (bool, error) {

	var Reservation entities.Reservation

	if CheckReservation := r.db.First(&Reservation, "id = ?", id); CheckReservation.Error != nil {
		return false, errors.New("Reservation not found")
	}

	UpdateStatusError := r.db.Model(&Reservation).Update("status", status)

	if UpdateStatusError.Error != nil {
		return false, errors.New("Error when update status reservation")
	}

	return true, nil

}
