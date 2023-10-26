package handlers

import (
	"villa_go/exceptions"
	"villa_go/payloads/request"
	"villa_go/services"

	"github.com/labstack/echo/v4"
)

type ReservationHandler interface {
	CreateReservationHandler(ctx echo.Context) error
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

	ReservationException := r.ReservationService.CreateNewReservation(ctx, Request)

	if ReservationException != nil {
		return exceptions.AppException(ctx, ReservationException.Error())
	}

	return nil
}
