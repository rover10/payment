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
	// Fetch latest payments first
	rows, err := client.db.Query(`
		SELECT 
			ta.id, ta.utr, ta.amount, ta.payment_time, ta.status, 
			CASE WHEN ta.from_account_id = ua.id THEN 'sent' ELSE 'received' END AS payment,
			(SELECT b.name as sender_user_bank from users_account ua2 
				INNER JOIN transaction t ON ua2.id = t.from_account_id 
				INNER JOIN bank b ON b.id = ua2.bank_id WHERE ta.id = t.id),
			(SELECT left(ua2.account_number, -4) || '****' as sender_acc_no from users_account ua2 
				INNER JOIN transaction t ON ua2.id = t.from_account_id 
				INNER JOIN bank b ON b.id = ua2.bank_id WHERE ta.id = t.id),
			(SELECT b.icon_url as sender_user_bank_icon from users_account ua2 
				INNER JOIN transaction t ON ua2.id = t.from_account_id 
				INNER JOIN bank b ON b.id = ua2.bank_id WHERE ta.id = t.id),
			(SELECT b.name as receiver_user_bank from users_account ua2 
				INNER JOIN transaction t ON ua2.id = t.to_account_id 
				INNER JOIN bank b ON b.id = ua2.bank_id WHERE ta.id = t.id),
			(SELECT left(ua2.account_number, -4) || '****' as receiver_acc_no from users_account ua2 
				INNER JOIN transaction t ON ua2.id = t.to_account_id 
				INNER JOIN bank b ON b.id = ua2.bank_id WHERE ta.id = t.id),
			(SELECT b.icon_url as receiver_user_bank_icon from users_account ua2 
				INNER JOIN transaction t ON ua2.id = t.to_account_id 
				INNER JOIN bank b ON b.id = ua2.bank_id WHERE ta.id = t.id)
		FROM 
			transaction ta 
		INNER JOIN
			users_account ua ON (ta.from_account_id = ua.id  OR ta.to_account_id = ua.id) AND user_id = $1
		WHERE 
			ta.id < $2  
		ORDER BY 
			ta.id DESC 
		LIMIT $3
		`,
		account_id, offset, limit,
	)
	if err != nil {
		log.Printf("PaymentHistory: error executing query: %v, account_id: %d", err, account_id)
		return nil, err
	}

	// Handle nullable fields in subquery
	fromBankName := sql.NullString{}
	fromBankAcc := sql.NullString{}
	fromBankIcon := sql.NullString{}
	toBankName := sql.NullString{}
	toBankAcc := sql.NullString{}
	toBankIcon := sql.NullString{}

	// Same user's account transaction filter for receive
	receivePayFilter := make(map[string]int, 0)
	indexCount := 0
	allPays := make([]model.PaymentHistory, 0)
	for rows.Next() {
		pay := model.PaymentHistory{}
		err := rows.Scan(
			&pay.ID,
			&pay.UTR,
			&pay.Amount,
			&pay.PaymentTime,
			&pay.Status,
			&pay.Payment,
			&fromBankName,
			&fromBankAcc,
			&fromBankIcon,
			&toBankName,
			&toBankAcc,
			&toBankIcon,
		)
		// Assign values
		pay.FromBank.Name = fromBankName.String
		pay.FromBank.Account = fromBankAcc.String
		pay.FromBank.BankIcon = fromBankIcon.String
		pay.ToBank.Name = toBankName.String
		pay.ToBank.Account = toBankAcc.String
		pay.ToBank.BankIcon = toBankIcon.String

		if err != nil {
			log.Printf("PaymentHistory: error reading row: %v, err: %v ", rows, err)
			return nil, err
		}

		// Filter out equivalent 'received' transaction if transfer made to another linked account of the same user
		if index, ok := receivePayFilter[pay.UTR]; ok {
			if pay.Payment == "received" {
				// No need to include this transaction as equivalent 'sent' transaction already added
				continue
			}
			if pay.Payment == "sent" {
				// Overwrite equivalent 'received' transaction with 'sent'
				allPays[index] = pay
				continue
			}
		}

		allPays = append(allPays, pay)
		receivePayFilter[pay.UTR] = indexCount
		indexCount++
	}
	return allPays, nil
}
