package services

import (
	"errors"
	"villa_go/exceptions"
	"villa_go/models/entities"
	"villa_go/payloads/request"
	"villa_go/payloads/resources"
	"villa_go/repositories"
	"villa_go/utils"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type ReservationService interface {
	GetListReservation() ([]resources.ReservationResource, error)
	CreateNewReservation(echo.Context, request.ReservationRequest) (*resources.ReservationResource, []exceptions.ValidationMessage, error)
}

type ReservationServiceImpl struct {
	ReservationRepo repositories.ReservationRepository
	MidtransService MidtransService
	VillaRepo       repositories.VillaRepository
	CredentialRepo  repositories.CredentialRepository
	Validate        *validator.Validate
	Trans           ut.Translator
}

func NewReservationServiceImplement(reservation repositories.ReservationRepository, midtrans MidtransService, villa repositories.VillaRepository, credential repositories.CredentialRepository, validate *validator.Validate, trans ut.Translator) ReservationService {
	return &ReservationServiceImpl{
		ReservationRepo: reservation,
		MidtransService: midtrans,
		VillaRepo:       villa,
		CredentialRepo:  credential,
		Validate:        validate,
		Trans:           trans,
	}
}

func (r *ReservationServiceImpl) GetListReservation() ([]resources.ReservationResource, error) {

	ListsReservations, ListResErr := r.ReservationRepo.GetListReservation()

	if ListResErr != nil {
		return nil, ListResErr
	}

	return ListsReservations, nil

}

func (r *ReservationServiceImpl) CreateNewReservation(ctx echo.Context, reservationRequest request.ReservationRequest) (*resources.ReservationResource, []exceptions.ValidationMessage, error) {

	var Reservation entities.Reservation

	ValidationErr := r.Validate.Struct(reservationRequest)

	if ValidationErr != nil {
		return nil, utils.ValidationError(ctx, r.Trans, ValidationErr), nil
	}

	GetUserAccess, UserException := r.CredentialRepo.UserLoginProfile(ctx)

	if UserException != nil {
		return nil, nil, UserException
	}

	VillaId, IdVillaException := uuid.FromString(reservationRequest.Villa_id)

	if IdVillaException != nil {
		return nil, nil, errors.New("wrong uuid format")
	}

	CheckInParsing, TimeInErr := utils.ConvertDate(reservationRequest.Check_in_date)
	CheckOutParsing, TimeOutErr := utils.ConvertDate(reservationRequest.Check_out_date)

	if TimeInErr != nil {
		return nil, nil, errors.New("Wrong check in date format")
	}

	if TimeOutErr != nil {
		return nil, nil, errors.New("Wrong check in date format")
	}

	GetDurationDays := utils.GetSubDate(*CheckInParsing, *CheckOutParsing)

	GetDataVilla, IsExist := r.VillaRepo.CheckVillaIsExists(VillaId)

	if IsExist != nil {
		return nil, nil, errors.New("Villa does not exists")
	}

	Reservation.GetReservationRequest(reservationRequest, *GetDataVilla.Price_per_night, GetUserAccess.Id, GetDataVilla.Id, CheckInParsing, CheckOutParsing, GetDurationDays)
	CreateReservation, CreateException := r.ReservationRepo.CreateNewReservation(Reservation)

	if CreateException != nil {
		return nil, nil, CreateException
	}

	GenerateTransactionURL, TransactionErr := r.MidtransService.GenerateSnapURL(ctx, *GetDataVilla, *GetUserAccess, *CreateReservation)

	if TransactionErr != nil {
		return nil, nil, TransactionErr
	}

	SetUpdateRequestReservation := request.ReservationRequest{}
	SetUpdateRequestReservation.ReservationDetail = &request.ReservationDetailRequest{
		SnapURL: GenerateTransactionURL.RedirectURL,
	}

	GetReservationData, QueryException := r.ReservationRepo.UpdateSnapUrlReservation(*CreateReservation.Id, SetUpdateRequestReservation)

	if QueryException != nil {
		return nil, nil, errors.New("Reservation data not found")
	}

	return GetReservationData, nil, nil

}
