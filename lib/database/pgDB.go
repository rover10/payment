package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/rover10/payment/lib/config"
	"github.com/rover10/payment/lib/model"
)

type Client struct {
	db *sql.DB
}

func (client *Client) DBConnect(config *config.Config) error {
	dbinfo := "user=postgres port=5432 password=root dbname=tes2ac host=localhost sslmode=disable"
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

func (client *Client) PaymentHistory(account_id int) ([]model.PaymentHistory, error) {
	var rows *sql.Rows
	var err error
	// Todo - Paging
	rows, err = client.db.Query(`
		SELECT id,utr,amount,from_account_id, to_account_id, payment_time, status   FROM transaction 
			WHERE 
				from_account_id 
					IN (select id from users_account where user_id = $1)  
			OR 
				to_account_id 
					IN (select id from users_account where user_id = $1)`, account_id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	allPays := make([]model.PaymentHistory, 0)
	for rows.Next() {
		pay := model.PaymentHistory{}
		err := rows.Scan(
			&pay.ID,
			&pay.UTR,
			&pay.Amount,
			&pay.FromAccount,
			&pay.ToAccount,
			&pay.PaymentTime,
			&pay.Status,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		allPays = append(allPays, pay)
	}
	return allPays, nil
}
