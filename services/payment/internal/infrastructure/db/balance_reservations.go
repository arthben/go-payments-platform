package db

import "time"

type BalanceReservations struct {
	ID        string    `db:"id"`
	AccountID string    `db:"account_id"`
	OrderID   string    `db:"order_id"`
	Currency  string    `db:"currency"`
	Amount    int       `db:"amount"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
