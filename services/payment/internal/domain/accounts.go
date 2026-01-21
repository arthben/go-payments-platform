package domain

import "time"

type AccountOwnerType string

const (
	AccountOwnerTypeUser     AccountOwnerType = "USER"
	AccountOwnerTypeMerchant AccountOwnerType = "MERCHANT"
	AccountOwnerTypeSystem   AccountOwnerType = "SYSTEM"
)

type Accounts struct {
	ID        UUID
	OwnerType AccountOwnerType
	OwnerID   string
	IsDeleted bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
