package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"main.go/models"
)

type EditorPostgres struct {
	db *sqlx.DB
}

func NewEditorPostgres(db *sqlx.DB) *EditorPostgres {
	return &EditorPostgres{db: db}
}

func (e *EditorPostgres) EditPerson(ctx context.Context, id int, name string, age int) (*models.CreatePersonOutput, error) {
	tx := extractTx(ctx)

	var output models.CreatePersonOutput
	err := tx.Get(&output, `UPDATE "Person" SET name = $1, age = $2 WHERE id =$3 RETURNING "id", "name", "age"`, name, age, id)
	if err != nil {
		return nil, err
	}

	return &output, nil
}
