package payservice

import (
	"github.com/rover10/payment/lib/database"
	"github.com/rover10/payment/lib/model"
)

type PaymentServiceInterface interface {
	PaymentHistory(userId int, offSet int64, limit int64) ([]model.PaymentHistory, error)
}

type PaymentService struct{}

func (p *PaymentService) PaymentHistory(userId int, offSet int64, limit int64) ([]model.PaymentHistory, error) {
	return database.DB.PaymentHistory(userId, offSet, limit)
}
