package db

import (
	"time"

	"github.com/arthben/go-payments-platform/services/payment/internal/domain"
)

type Accounts struct {
	ID        string    `db:"id"`
	OwnerType string    `db:"owner_type"`
	OwnerID   string    `db:"owner_id"`
	IsDeleted bool      `db:"is_deleted"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func toAccountDomain(acc Accounts) domain.Accounts {
	return domain.Accounts{
		ID:        domain.UUID(acc.ID),
		OwnerType: domain.AccountOwnerType(acc.OwnerType),
		OwnerID:   acc.OwnerID,
		IsDeleted: acc.IsDeleted,
		CreatedAt: acc.CreatedAt,
		UpdatedAt: acc.UpdatedAt,
	}
}
