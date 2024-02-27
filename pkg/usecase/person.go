package usecase

import (
	"context"
)

func (s *Usecase) CreatePerson(name string, age int) (interface{}, error) {
	return s.Service.Transactor.WithinTransaction(context.Background(), func(ctx context.Context) (interface{}, error) {
		id, err := s.Service.Person.Create(ctx)
		if err != nil {
			return nil, err
		}

		output, err := s.Service.Person.Edit(ctx, id, name, age)
		if err != nil {
			return nil, err
		}

		return &output, nil
	})
}
