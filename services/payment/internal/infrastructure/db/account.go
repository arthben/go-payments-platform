package db

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/arthben/go-payments-platform/services/payment/internal/domain"
)

type AccountDAO struct {
	pool *pgxpool.Pool
}

func NewAccountDAO(dbPool *pgxpool.Pool) *AccountDAO {
	return &AccountDAO{pool: dbPool}
}

func (dbase *AccountDAO) CreateAccount(ctx context.Context, acc domain.Accounts) error {
	sql := "INSERT INTO account (id, owner_type, owner_id) VALUES ($1, $2, $3)"
	_, err := dbase.pool.Exec(ctx, sql, acc.ID, acc.OwnerType, acc.OwnerID)
	return toInternalError(err)
}

func (dbase *AccountDAO) GetAccountByID(ctx context.Context, id domain.UUID) (domain.Accounts, error) {
	var account Accounts
	sql := `SELECT id, owner_type, owner_id, created_at, updated_at 
			FROM accounts
			WHERE id=$1 AND is_deleted=false`
	err := pgxscan.Select(ctx, dbase.pool, &account, sql, string(id))
	if err != nil {
		return domain.Accounts{}, toInternalError(err)
	}

	return toAccountDomain(account), nil
}

func (dbase *AccountDAO) DeleteAccountByID(ctx context.Context, id domain.UUID) error {
	sql := "UPDATE accounts SET is_deleted=true WHERE id=$1"
	_, err := dbase.pool.Exec(ctx, sql, string(id))
	return toInternalError(err)
}
