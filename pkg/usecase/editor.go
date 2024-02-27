package usecase

import "context"

func (s *Usecase) EditPerson(ctx context.Context, id int, name string, age int) (interface{}, error) {
	return s.Service.Editor.EditPerson(ctx, id, name, age)
}
