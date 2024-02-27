package service

import (
	"context"

	"main.go/models"
	"main.go/pkg/repository"
)

type PersonService struct {
	repo repository.Person
}

func NewCreatorService(repo repository.Person) *PersonService {
	return &PersonService{repo: repo}
}

func (s *PersonService) Create(ctx context.Context) (int, error) {
	return s.repo.Create(ctx)
}

func (s *PersonService) Edit(ctx context.Context, id int, name string, age int) (*models.CreatePersonOutput, error) {
	return s.repo.Edit(ctx, id, name, age)
}
