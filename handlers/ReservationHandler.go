package handlers

import (
	"fmt"
	"net/http"
	"villa_go/exceptions"
	"villa_go/payloads/request"
	"villa_go/payloads/response"
	"villa_go/services"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type ReservationHandler interface {
	GetAllReservationHandler(echo.Context) error
	GetReservationByIdHandler(echo.Context) error
	CreateReservationHandler(echo.Context) error
	NotificationReservationHandler(echo.Context) error
	GetReservationDataByUserLogin(echo.Context) error
}

type ReservationHandlerImpl struct {
	ReservationService services.ReservationService
	MidtransService    services.MidtransService
}

func NewReservationHandler(reservation services.ReservationService, midtrans services.MidtransService) ReservationHandler {
	return &ReservationHandlerImpl{
		ReservationService: reservation,
		MidtransService:    midtrans,
	}
}

func (r *ReservationHandlerImpl) GetAllReservationHandler(ctx echo.Context) error {

	GetListReservations, ErrMessage := r.ReservationService.GetListReservation()

	if ErrMessage != nil {
		return exceptions.NotFoundException(ctx, ErrMessage.Error())
	}

	return response.HandleSuccess(ctx, GetListReservations, "Success get list reservation", http.StatusOK)

}

func (r *ReservationHandlerImpl) CreateReservationHandler(ctx echo.Context) error {

	var Request request.ReservationRequest

	BindingRequest := ctx.Bind(&Request)

	if BindingRequest != nil {
		return exceptions.BadRequestException(ctx, BindingRequest.Error())
	}

	ReservationResponse, ReservationValidation, ReservationException := r.ReservationService.CreateNewReservation(ctx, Request)

	if ReservationValidation != nil {
		return exceptions.ValidationException(ctx, "One or more validation errors occurred", ReservationValidation)
	}

	if ReservationException != nil {
		return exceptions.NotFoundException(ctx, ReservationException.Error())
	}

	return response.HandleSuccess(ctx, ReservationResponse, "Reservation Created, Finish Your Payment Transaction", http.StatusCreated)
}

func (r *ReservationHandlerImpl) NotificationReservationHandler(ctx echo.Context) error {

	var NotificationPayment map[string]interface{}

	BindingNotification := ctx.Bind(&NotificationPayment)

	if BindingNotification != nil {
		return exceptions.AppException(ctx, BindingNotification.Error())
	}

	NotificationBool, NotificationMessage, ErrResponse := r.MidtransService.NotificationPayment(NotificationPayment)
	if !NotificationBool {

		fmt.Println("Error : ", ErrResponse)

		return ctx.JSON(http.StatusInternalServerError, ErrResponse)
	}

	fmt.Println("Notification :", NotificationMessage)

	return ctx.String(http.StatusCreated, NotificationMessage)
}

func (r *ReservationHandlerImpl) GetReservationByIdHandler(ctx echo.Context) error {

	GetId := ctx.Param("villa_id")

	ParseToUUID, errPasing := uuid.FromString(GetId)

	if errPasing != nil {
		return exceptions.BadRequestException(ctx, "Invalid format uuid")
	}

	Results, ErrMessage := r.ReservationService.GetReservationById(ParseToUUID)

	if ErrMessage != nil {
		return exceptions.AppException(ctx, ErrMessage.Error())
	}

	return response.HandleSuccess(ctx, Results, "Success Get Reservation Detail", http.StatusOK)

}

func (r *ReservationHandlerImpl) GetReservationDataByUserLogin(ctx echo.Context) error {

	ReservationResponse, httpStatus, Err := r.ReservationService.GetReservationListUser(ctx)

	if Err != nil {
		if httpStatus == http.StatusNotFound {
			return exceptions.NotFoundException(ctx, Err.Error())
		}
		if httpStatus == http.StatusNoContent {
			return exceptions.NotFoundException(ctx, Err.Error())
		}
	}

	return response.HandleSuccess(ctx, ReservationResponse, "Success get reservation data user", httpStatus)

}
