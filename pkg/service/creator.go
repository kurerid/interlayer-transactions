package service

import (
	"context"

	"main.go/pkg/repository"
)

type CreatorService struct {
	repo repository.Creator
}

func NewCreatorService(repo repository.Creator) *CreatorService {
	return &CreatorService{repo: repo}
}

func (d *CreatorService) Create(ctx context.Context) (int, error) {
	return d.repo.Create(ctx)
}
