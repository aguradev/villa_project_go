package handlers

import (
	"fmt"
	"net/http"
	"villa_go/exceptions"
	"villa_go/payloads/request"
	"villa_go/payloads/response"
	"villa_go/services"

	"github.com/labstack/echo/v4"
)

type ReservationHandler interface {
	CreateReservationHandler(ctx echo.Context) error
	NotificationReservationHandler(ctx echo.Context) error
}

type ReservationHandlerImpl struct {
	ReservationService services.ReservationService
}

func NewReservationHandler(reservation services.ReservationService) ReservationHandler {
	return &ReservationHandlerImpl{
		ReservationService: reservation,
	}
}

func (r *ReservationHandlerImpl) CreateReservationHandler(ctx echo.Context) error {

	var Request request.ReservationRequest

	BindingRequest := ctx.Bind(&Request)

	if BindingRequest != nil {
		return exceptions.BadRequestException(ctx, BindingRequest.Error())
	}

	ReservationResponse, ReservationException := r.ReservationService.CreateNewReservation(ctx, Request)

	if ReservationException != nil {
		return exceptions.AppException(ctx, ReservationException.Error())
	}

	return response.HandleSuccess(ctx, ReservationResponse, "Reservation Created, Finish Your Payment Transaction", http.StatusCreated)
}

func (r *ReservationHandlerImpl) NotificationReservationHandler(ctx echo.Context) error {

	var NotificationPayment map[string]interface{}

	BindingNotification := ctx.Bind(&NotificationPayment)

	if BindingNotification != nil {
		return exceptions.AppException(ctx, BindingNotification.Error())
	}

	fmt.Println(NotificationPayment)

	return nil
}
