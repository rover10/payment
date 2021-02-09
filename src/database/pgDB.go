package database

import (
	"database/sql"

	"github.com/rover10/payment/src/config"
	"github.com/rover10/payment/src/model"
)

type Client struct {
	db *sql.DB
}

func (client *Client) DBConnect(config *config.Config) error {
	return nil
}

func CreatePayment(model.Payment) (model.Payment, error) {
	return model.Payment{}, nil
}
