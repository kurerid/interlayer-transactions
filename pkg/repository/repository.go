package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"main.go/models"
	"main.go/pkg/transactor"
)

type (
	Creator interface {
		Create(ctx context.Context) (int, error)
	}

	Transactor interface {
		WithinTransaction(ctx context.Context, tFunc transactor.Func) (interface{}, error)
	}

	Editor interface {
		EditPerson(ctx context.Context, id int, name string, age int) (*models.CreatePersonOutput, error)
	}

	Repository struct {
		Creator
		Transactor
		Editor
	}
)

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Creator:    NewCreatorPostgres(db),
		Transactor: NewTransactorPostgres(db),
		Editor:     NewEditorPostgres(db),
	}
}
