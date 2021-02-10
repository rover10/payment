package database

import (
	"github.com/rover10/payment/lib/config"
	"github.com/rover10/payment/lib/model"
)

var DB DBClient

type DBClient interface {
	DBConnect(config *config.Config) error
	CreatePayment(model.Payment) (model.Payment, error)
	PaymentHistory(account_id int, offset int64, limit int64) ([]model.PaymentHistory, error)
}
