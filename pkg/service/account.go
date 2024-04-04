package service

import (
	"context"
	"errors"
	"fmt"
	"main.go/pkg/repository"
	"time"
)

type AccountService struct {
	repo repository.Account
}

func NewAccountService(repo repository.Account) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) Exist(email string) (bool, error) {
	return s.repo.Exist(email)
}

func (s *AccountService) Register(email string, password string, ctx context.Context) error {
	return s.repo.Register(email, password, ctx)
}

func (s *AccountService) SendConfirmation(ctx context.Context, email string) error {
	fmt.Println("trying to send confirmation email")
	time.Sleep(time.Second * 3)
	return errors.New("email service connection refused")
}
