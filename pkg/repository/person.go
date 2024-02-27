package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"main.go/models"
)

type PersonPostgres struct {
	db *sqlx.DB
}

func NewCreatorPostgres(db *sqlx.DB) *PersonPostgres {
	return &PersonPostgres{db: db}
}

// Create default person.
func (r *PersonPostgres) Create(ctx context.Context) (int, error) {
	tx := extractTx(ctx)

	var id int
	err := tx.Get(&id, `INSERT INTO "Person" ("name", age) VALUES ($1, $2) RETURNING "id"`, "wakadoo", 18)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (r *PersonPostgres) Edit(ctx context.Context, id int, name string, age int) (*models.CreatePersonOutput, error) {
	tx := extractTx(ctx)

	var output models.CreatePersonOutput
	err := tx.Get(&output, `UPDATE "Person" SET name = $1, age = $2 WHERE id =$3 RETURNING "id", "name", "age"`, name, age, id)
	if err != nil {
		return nil, err
	}

	return &output, nil
}
