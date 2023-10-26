package services

import (
	"villa_go/config"

	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type MidtransService interface {
	GenerateNewSnapTransaction(echo.Context)
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

func (p *MidtransServiceImpl) GenerateNewSnapTransaction(ctx echo.Context) {

}
