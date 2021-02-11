package payservice

import (
	"log"

	"github.com/rover10/payment/lib/database"
	"github.com/rover10/payment/lib/model"
)

type PaymentServiceInterface interface {
	PaymentHistory(userId int, offSet int64, limit int64) ([]model.PaymentHistory, error)
}

type PaymentService struct{}

func (p *PaymentService) PaymentHistory(userId int, offset int64, limit int64) ([]model.PaymentHistory, error) {
	log.Printf("PaymentHistoryService: userId: %v, offset: %d, limit: %d", userId, offset, limit)
	return database.DB.PaymentHistory(userId, offset, limit)
}
