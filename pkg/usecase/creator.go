package usecase

import (
	"context"
)

func (s *Usecase) CreatePerson(name string, age int) (interface{}, error) {
	return s.Service.Transactor.WithinTransaction(context.Background(), func(ctx context.Context) (interface{}, error) {
		id, err := s.Service.Creator.Create(ctx)
		if err != nil {
			return nil, err
		}

		output, err := s.Service.Editor.EditPerson(ctx, id, name, age)
		if err != nil {
			return nil, err
		}

		return &output, nil
	})
}
