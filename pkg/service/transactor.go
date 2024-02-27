package service

import (
	"context"

	"main.go/pkg/repository"
	"main.go/pkg/transactor"
)

type TransactorService struct {
	repo repository.Transactor
}

func NewTransactorService(repo repository.Transactor) *TransactorService {
	return &TransactorService{repo: repo}
}

func (t *TransactorService) WithinTransaction(ctx context.Context, tFunc transactor.Func) (interface{}, error) {
	return t.repo.WithinTransaction(ctx, tFunc)
}
