package domain

import "time"

type BalanceReservationStatus string

const (
	BalanceReservationStatusReserved  BalanceReservationStatus = "RESERVED"
	BalanceReservationStatusConfirmed BalanceReservationStatus = "CONFIRMED"
	BalanceReservationStatusReleased  BalanceReservationStatus = "RELEASED"
)

type BalanceReservations struct {
	ID        UUID
	AccountID UUID
	OrderID   string
	Currency  string
	Amount    int
	Status    BalanceReservationStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}
