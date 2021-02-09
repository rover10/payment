package database

import (
	"github.com/rover10/payment/src/config"
	"github.com/rover10/payment/src/model"
)

type DBClient interface {
	DBConnect(config *config.Config) error
	CreatePayment(model.Payment) (model.Payment, error)
}
