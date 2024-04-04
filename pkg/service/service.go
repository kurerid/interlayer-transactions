package service

import (
	"context"

	"main.go/pkg/repository"
	"main.go/pkg/transactor"
)

type (
	Account interface {
		Exist(email string) (bool, error)
		Register(email string, password string, ctx context.Context) error
		SendConfirmation(ctx context.Context, email string) error
	}

	Transactor interface {
		WithinTransaction(ctx context.Context, tFunc transactor.Func) (interface{}, error)
	}

	Services struct {
		Account
		Transactor
	}
)

func NewServices(repo *repository.Repository) *Services {
	return &Services{
		Account:    NewAccountService(repo.Account),
		Transactor: NewTransactorService(repo.Transactor),
	}
}
