package services

import (
	"errors"
	"fmt"
	"villa_go/config"
	"villa_go/models/entities"
	"villa_go/payloads/request"
	"villa_go/repositories"
	"villa_go/utils"

	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
	uuid "github.com/satori/go.uuid"
)

type MidtransService interface {
	GenerateSnapURL(echo.Context, entities.Villa, entities.Users, entities.Reservation) (*snap.Response, error)
	NotificationPayment(map[string]interface{}) (bool, string, error)
}

type MidtransServiceImpl struct {
	Client          snap.Client
	ClientOpenCore  coreapi.Client
	ReservationRepo repositories.ReservationRepository
}

func NewMidtransServiceImpl(midtransConfig *config.PaymentGatewayConfig, reservation repositories.ReservationRepository) MidtransService {
	var setClient snap.Client
	var setCoreApi coreapi.Client

	enviroment := midtrans.Sandbox

	setClient.New(midtransConfig.MidtransClientKey, enviroment)
	setCoreApi.New(midtransConfig.MidtransClientKey, enviroment)

	return &MidtransServiceImpl{
		Client:          setClient,
		ClientOpenCore:  setCoreApi,
		ReservationRepo: reservation,
	}
}

func (p *MidtransServiceImpl) GenerateSnapURL(ctx echo.Context, villa entities.Villa, User entities.Users, reservation entities.Reservation) (*snap.Response, error) {
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  reservation.Id.String(),
			GrossAmt: reservation.Reservation_detail.Total.IntPart(),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: User.First_name,
			LName: User.Last_name,
			Email: User.Email,
		},
		Items: &[]midtrans.ItemDetails{
			{
				ID:    villa.Id.String(),
				Price: reservation.Reservation_detail.Total.IntPart(),
				Name:  villa.Name,
				Qty:   1,
			},
		},
	}

	SnapRes, errSnap := p.Client.CreateTransaction(req)

	if errSnap != nil {
		return nil, errors.New("Error when create transaction")
	}

	return SnapRes, nil
}

func (p *MidtransServiceImpl) NotificationPayment(transaction map[string]interface{}) (bool, string, error) {

	orderId, exists := transaction["order_id"].(string)

	if !exists {
		return false, "", errors.New("order id not found")
	}

	SetToUuid, ErrParse := uuid.FromString(orderId)

	if ErrParse != nil {
		return false, "", errors.New("invalid format uuid")
	}

	TransactionStatusReps, TransactionExp := p.ClientOpenCore.CheckTransaction(orderId)
	GetReservationData, ErrGetReservation := p.ReservationRepo.GetReservationById(SetToUuid)

	if ErrGetReservation != nil {
		return false, "", errors.New("failed to find reservation")
	}

	if TransactionExp != nil {
		return false, "", errors.New("Transaction not found")
	} else {
		if TransactionStatusReps != nil {

			TransactionStatus := TransactionStatusReps.TransactionStatus
			RequestTransactionEmail := request.ReservationEmailRequest{
				Name_customer:  GetReservationData.User.First_name + " " + GetReservationData.User.Last_name,
				Villa_name:     GetReservationData.Reservation_detail.Villa.Name,
				Check_in_date:  GetReservationData.Reservation_detail.Check_in_date.String(),
				Check_out_date: GetReservationData.Reservation_detail.Check_out_date.String(),
				Guest_count:    uint(GetReservationData.Reservation_detail.Guest_count),
				Total:          GetReservationData.Reservation_detail.Total,
			}

			fmt.Println(RequestTransactionEmail)

			if TransactionStatus == "capture" {
				if TransactionStatusReps.FraudStatus == "challange" {
					StatusUpdated, ErrMessage := p.ReservationRepo.UpdateStatusReservation(SetToUuid, TransactionStatus)

					if !StatusUpdated {
						return false, "", ErrMessage
					}

					ErrSendingEmail := utils.SendingEmail(GetReservationData.User.Email, "Notification Payment", RequestTransactionEmail)

					if ErrSendingEmail != nil {
						return false, ErrSendingEmail.Error(), nil
					}

					return true, TransactionStatusReps.StatusMessage, nil
				} else if TransactionStatusReps.FraudStatus == "accept" {
					StatusUpdated, ErrMessage := p.ReservationRepo.UpdateStatusReservation(SetToUuid, TransactionStatus)

					if !StatusUpdated {
						return false, "", ErrMessage
					}

					ErrSendingEmail := utils.SendingEmail(GetReservationData.User.Email, "Notification Payment", RequestTransactionEmail)

					if ErrSendingEmail != nil {
						return false, ErrSendingEmail.Error(), nil
					}

					return true, TransactionStatusReps.StatusMessage, nil
				}
			} else if TransactionStatus == "settlement" {

				StatusUpdated, ErrMessage := p.ReservationRepo.UpdateStatusReservation(SetToUuid, TransactionStatus)

				if !StatusUpdated {
					return false, "", ErrMessage
				}

				ErrSendingEmail := utils.SendingEmail(GetReservationData.User.Email, "Notification Payment", RequestTransactionEmail)

				if ErrSendingEmail != nil {
					return false, ErrSendingEmail.Error(), nil
				}

				return true, "Reservation transaction settlement", nil

			} else if TransactionStatus == "deny" {
				StatusUpdated, ErrMessage := p.ReservationRepo.UpdateStatusReservation(SetToUuid, TransactionStatus)

				if !StatusUpdated {
					return false, "", ErrMessage
				}

				return true, TransactionStatusReps.StatusMessage, nil
			} else if TransactionStatus == "cancel" || TransactionStatus == "expire" {
				StatusUpdated, ErrMessage := p.ReservationRepo.UpdateStatusReservation(SetToUuid, TransactionStatus)

				if !StatusUpdated {
					return false, "", ErrMessage
				}

				return true, "Reservation transaction is cancel or expired", nil
			} else if TransactionStatus == "pending" {
				return true, TransactionStatusReps.StatusMessage, nil
			}
		}
	}

	return false, "", nil
}
