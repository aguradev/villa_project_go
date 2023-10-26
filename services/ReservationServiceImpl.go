package services

import (
	"errors"
	"fmt"
	"villa_go/models/entities"
	"villa_go/payloads/request"
	"villa_go/repositories"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type ReservationService interface {
	CreateNewReservation(echo.Context, request.ReservationRequest) error
}

type ReservationServiceImpl struct {
	ReservationRepo repositories.ReservationRepository
	MidtransService MidtransService
	VillaRepo       repositories.VillaRepository
	CredentialRepo  repositories.CredentialRepository
}

func NewReservationServiceImplement(reservation repositories.ReservationRepository, midtrans MidtransService, villa repositories.VillaRepository, credential repositories.CredentialRepository) ReservationService {
	return &ReservationServiceImpl{
		ReservationRepo: reservation,
		MidtransService: midtrans,
		VillaRepo:       villa,
		CredentialRepo:  credential,
	}
}

func (r *ReservationServiceImpl) CreateNewReservation(ctx echo.Context, request request.ReservationRequest) error {

	var Reservation entities.Reservation

	GetUserAccess, UserException := r.CredentialRepo.UserLoginProfile(ctx)

	if UserException != nil {
		return UserException
	}

	VillaId, IdVillaException := uuid.FromString(request.Villa_id)

	if IdVillaException != nil {
		return errors.New("wrong uuid format")
	}

	GetDataVilla, IsExist := r.VillaRepo.CheckVillaIsExists(VillaId)

	if IsExist != nil {
		return errors.New("Villa does not exists")
	}

	Reservation.GetReservationRequest(request, *GetDataVilla.Price_per_night, GetUserAccess.Id, GetDataVilla.Id)

	fmt.Println(Reservation)

	return nil

}
