package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type CreatorPostgres struct {
	db *sqlx.DB
}

func NewCreatorPostgres(db *sqlx.DB) *CreatorPostgres {
	return &CreatorPostgres{db: db}
}

// Create default person.
func (r *CreatorPostgres) Create(ctx context.Context) (int, error) {
	tx := extractTx(ctx)

	var id int
	err := tx.Get(&id, `INSERT INTO "Person" ("name", age) VALUES ($1, $2) RETURNING "id"`, "wakadoo", 18)
	if err != nil {
		return -1, err
	}

	return id, nil
}
