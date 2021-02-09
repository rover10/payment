package database

import (
	"github.com/rover10/payment/lib/config"
	"github.com/rover10/payment/lib/model"
)

type DBClient interface {
	DBConnect(config *config.Config) error
	CreatePayment(model.Payment) (model.Payment, error)
}
