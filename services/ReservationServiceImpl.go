package services

import (
	"errors"
	"villa_go/models/entities"
	"villa_go/payloads/request"
	"villa_go/payloads/resources"
	"villa_go/repositories"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type ReservationService interface {
	CreateNewReservation(echo.Context, request.ReservationRequest) (*resources.ReservationResource, error)
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

func (r *ReservationServiceImpl) CreateNewReservation(ctx echo.Context, request request.ReservationRequest) (*resources.ReservationResource, error) {

	var Reservation entities.Reservation

	GetUserAccess, UserException := r.CredentialRepo.UserLoginProfile(ctx)

	if UserException != nil {
		return nil, UserException
	}

	VillaId, IdVillaException := uuid.FromString(request.Villa_id)

	if IdVillaException != nil {
		return nil, errors.New("wrong uuid format")
	}

	GetDataVilla, IsExist := r.VillaRepo.CheckVillaIsExists(VillaId)

	if IsExist != nil {
		return nil, errors.New("Villa does not exists")
	}

	Reservation.GetReservationRequest(request, *GetDataVilla.Price_per_night, GetUserAccess.Id, GetDataVilla.Id)
	GenerateTransactionURL, TransactionErr := r.MidtransService.GenerateSnapURL(ctx, *GetDataVilla, *GetUserAccess, Reservation.Reservation_detail.Total)

	if TransactionErr != nil {
		return nil, TransactionErr
	}

	Reservation.Reservation_detail.SnapURL = GenerateTransactionURL.RedirectURL
	CreateReservation, CreateException := r.ReservationRepo.CreateNewReservation(Reservation)

	if CreateException != nil {
		return nil, CreateException
	}

	GetReservationData, QueryException := r.ReservationRepo.GetReservationById(*CreateReservation.Id)

	if QueryException != nil {
		return nil, errors.New("Reservation data not found")
	}

	return GetReservationData, nil

}
