package usecase

import (
	"context"
	"errors"
)

func (u *Usecase) RegisterAccount(email string, password string) (interface{}, error) {
	return u.Service.Transactor.WithinTransaction(context.Background(), func(ctx context.Context) (interface{}, error) {
		exist, err := u.Service.Account.Exist(email)
		if err != nil {
			return nil, err
		} else if exist {
			return nil, errors.New("account with this email already exist")
		}

		err = u.Service.Account.Register(email, password, ctx)
		if err != nil {
			return nil, err
		}

		err = u.Service.Account.SendConfirmation(ctx, email)
		if err != nil {
			return nil, err
		}

		return nil, nil
	})
}
