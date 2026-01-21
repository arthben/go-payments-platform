package db

import "time"

type PaymentTransactions struct {
	ID            string    `db:"id"`
	OrderID       string    `db:"order_id"`
	AccountID     string    `db:"account_id"`
	Currency      string    `db:"currency"`
	Amount        int       `db:"amount"`
	Status        string    `db:"status"`
	FailureReason string    `db:"failure_reason"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}
