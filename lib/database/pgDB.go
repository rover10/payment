package database

import (
	"database/sql"
	"log"
	"time"

	"github.com/rover10/payment/lib/config"
	"github.com/rover10/payment/lib/model"
)

type Client struct {
	db *sql.DB
}

func (client *Client) DBConnect(config *config.Config) error {
	dbinfo := "user=potgres port=5431 password=root dbname=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatalf("Error connecting database", err)
		return err
	}
	db.SetConnMaxLifetime(60 * time.Minute)
	client.db = db
	return nil
}

func (client *Client) CreatePayment(model.Payment) (model.Payment, error) {
	return model.Payment{}, nil
}
