package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"main.go/models"
	"main.go/pkg/transactor"
)

type (
	Person interface {
		Create(ctx context.Context) (int, error)
		Edit(ctx context.Context, id int, name string, age int) (*models.CreatePersonOutput, error)
	}

	Transactor interface {
		WithinTransaction(ctx context.Context, tFunc transactor.Func) (interface{}, error)
	}

	Editor interface {
	}

	Repository struct {
		Person
		Transactor
	}
)

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Person:     NewCreatorPostgres(db),
		Transactor: NewTransactorPostgres(db),
	}
}
