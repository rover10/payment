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

func (client *Client) PaymentHistory(account_id int, offset int64, limit int64) ([]model.PaymentHistory, error) {
	var rows *sql.Rows
	var err error
	// Todo - Paging
	// Fetch in reverse order to get latest payment first
	//
	fmt.Printf("\noffset %d", offset)
	rows, err = client.db.Query(`
		SELECT 
		ta.id, ta.utr, ta.amount, ta.from_account_id, ta.to_account_id, ta.payment_time, ta.status 
			FROM transaction ta 
			INNER JOIN
				users_account ua ON (from_account_id = ua.id  OR to_account_id = ua.id) AND user_id = $1
			WHERE 
				id < $2
			ORDER BY id DESC 
			LIMIT $3
		`,
		account_id, offset, limit,
	)
	fmt.Println("limit")
	fmt.Println(limit)

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
