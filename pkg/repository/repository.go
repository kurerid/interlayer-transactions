package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"main.go/pkg/transactor"
)

type (
	Account interface {
		Exist(email string) (bool, error)
		Register(email string, password string, ctx context.Context) error
	}

	Transactor interface {
		WithinTransaction(ctx context.Context, tFunc transactor.Func) (interface{}, error)
	}

	Repository struct {
		Account
		Transactor
	}
)

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Account:    NewAccountPostgres(db),
		Transactor: NewTransactorPostgres(db),
	}
}
