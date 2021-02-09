package database

import (
	"github.com/rover10/payment/config"
	"github.com/rover10/payment/model"
)

type DBClient interface {
	DBConnect(config *config.Config) error
	CreatePayment(model.Payment) (model.Payment, error)
}
