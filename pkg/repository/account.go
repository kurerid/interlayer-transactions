package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

type AccountPostgres struct {
	db *sqlx.DB
}

func NewAccountPostgres(db *sqlx.DB) *AccountPostgres {
	return &AccountPostgres{db: db}
}

func (r *AccountPostgres) Exist(email string) (bool, error) {
	var nullableEmail sql.NullString
	err := r.db.Get(&nullableEmail, `SELECT "email" FROM "Account" WHERE "email" = $1`, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func (r *AccountPostgres) Register(email string, password string, ctx context.Context) error {
	tx := extractTx(ctx)
	_, err := tx.Exec(`INSERT INTO "Account"("email","password") VALUES ($1, $2)`, email, password)
	if err != nil {
		return err
	}
	return nil
}
