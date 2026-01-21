package db

import "time"

type Balances struct {
	AccountID string    `db:"account_id"`
	Currency  string    `db:"currency"`
	Amount    int       `db:"amount"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
