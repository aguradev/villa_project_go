package services

import (
	"errors"
	"time"
	"villa_go/models/entities"
	"villa_go/payloads/request"
	"villa_go/payloads/resources"
	"villa_go/repositories"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type ReservationService interface {
	GetListReservation() ([]resources.ReservationResource, error)
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

func (r *ReservationServiceImpl) GetListReservation() ([]resources.ReservationResource, error) {

	ListsReservations, ListResErr := r.ReservationRepo.GetListReservation()

	if ListResErr != nil {
		return nil, ListResErr
	}

	return ListsReservations, nil

}

func (r *ReservationServiceImpl) CreateNewReservation(ctx echo.Context, reservationRequest request.ReservationRequest) (*resources.ReservationResource, error) {

	var Reservation entities.Reservation

	GetUserAccess, UserException := r.CredentialRepo.UserLoginProfile(ctx)

	if UserException != nil {
		return nil, UserException
	}

	VillaId, IdVillaException := uuid.FromString(reservationRequest.Villa_id)

	if IdVillaException != nil {
		return nil, errors.New("wrong uuid format")
	}

	CheckInParsing, TimeInErr := time.ParseInLocation("2023-01-02", reservationRequest.Check_in_date, time.Local)
	CheckOutParsing, TimeOutErr := time.ParseInLocation("2023-01-02", reservationRequest.Check_out_date, time.Local)

	if TimeInErr != nil {
		return nil, errors.New("Wrong check in date format")
	}

	if TimeOutErr != nil {
		return nil, errors.New("Wrong check in date format")
	}

	GetDataVilla, IsExist := r.VillaRepo.CheckVillaIsExists(VillaId)

	if IsExist != nil {
		return nil, errors.New("Villa does not exists")
	}

	Reservation.GetReservationRequest(reservationRequest, *GetDataVilla.Price_per_night, GetUserAccess.Id, GetDataVilla.Id, &CheckInParsing, &CheckOutParsing)
	CreateReservation, CreateException := r.ReservationRepo.CreateNewReservation(Reservation)

	if CreateException != nil {
		return nil, CreateException
	}

	GenerateTransactionURL, TransactionErr := r.MidtransService.GenerateSnapURL(ctx, *GetDataVilla, *GetUserAccess, *CreateReservation)

	if TransactionErr != nil {
		return nil, TransactionErr
	}

	SetUpdateRequestReservation := request.ReservationRequest{}
	SetUpdateRequestReservation.ReservationDetail = &request.ReservationDetailRequest{
		SnapURL: GenerateTransactionURL.RedirectURL,
	}

	GetReservationData, QueryException := r.ReservationRepo.UpdateSnapUrlReservation(*CreateReservation.Id, SetUpdateRequestReservation)

	if QueryException != nil {
		return nil, errors.New("Reservation data not found")
	}

	return GetReservationData, nil

}
