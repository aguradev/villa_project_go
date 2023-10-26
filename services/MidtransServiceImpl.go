package services

import (
	"errors"
	"villa_go/config"
	"villa_go/models/entities"

	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type MidtransService interface {
	GenerateSnapURL(echo.Context, entities.Villa, entities.Users, decimal.Decimal) (*snap.Response, error)
}

type MidtransServiceImpl struct {
	Client snap.Client
}

func NewMidtransServiceImpl(midtransConfig *config.PaymentGatewayConfig) MidtransService {
	var setClient snap.Client

	enviroment := midtrans.Sandbox

	setClient.New(midtransConfig.MidtransClientKey, enviroment)

	return &MidtransServiceImpl{
		Client: setClient,
	}
}

func (p *MidtransServiceImpl) GenerateSnapURL(ctx echo.Context, villa entities.Villa, User entities.Users, Amount decimal.Decimal) (*snap.Response, error) {

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  uuid.NewV4().String(),
			GrossAmt: Amount.IntPart(),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: User.First_name,
			LName: User.Last_name,
			Email: User.Email,
		},
	}

	SnapRes, errSnap := p.Client.CreateTransaction(req)

	if errSnap != nil {
		return nil, errors.New("Error when create transaction")
	}

	return SnapRes, nil
}
