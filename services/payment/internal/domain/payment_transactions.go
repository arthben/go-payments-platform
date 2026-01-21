package domain

import "time"

type PaymentTransactionStatus string

const (
	PaymentTransactionStatusSuccess PaymentTransactionStatus = "SUCCESS"
	PaymentTransactionStatusFailed  PaymentTransactionStatus = "FAILED"
)

type PaymentTransactions struct {
	ID            UUID
	OrderID       string
	AccountID     UUID
	Currency      string
	Amount        int
	Status        PaymentTransactionStatus
	FailureReason string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
