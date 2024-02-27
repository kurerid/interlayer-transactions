package service

import (
	"context"

	"main.go/models"
	"main.go/pkg/repository"
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

	Services struct {
		Person
		Transactor
	}
)

func NewServices(repo *repository.Repository) *Services {
	return &Services{
		Person:     NewCreatorService(repo.Person),
		Transactor: NewTransactorService(repo.Transactor),
	}
}
