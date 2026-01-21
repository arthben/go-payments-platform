package domain

import "time"

type Balances struct {
	AccountID UUID
	Currency  string
	Amount    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
