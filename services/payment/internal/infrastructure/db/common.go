package db

import (
	"database/sql"
	"errors"

	"github.com/arthben/go-payments-platform/services/payment/internal/domain"
)

func toInternalError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		return domain.ErrNotFound
	}

	// var pgErr pgdriver.Error
	// if errors.As(err, &pgErr) {
	// 	if pgErr.Field('C') == "23505" {
	// 		return errors.Join(err, domain.ErrAlreadyExists)
	// 	}
	// }

	return err
}
