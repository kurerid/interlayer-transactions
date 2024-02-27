package service

import (
	"context"

	"main.go/models"
	"main.go/pkg/repository"
)

type EditorService struct {
	repo repository.Editor
}

func NewEditorService(repo repository.Editor) *EditorService {
	return &EditorService{repo: repo}
}

func (e *EditorService) EditPerson(ctx context.Context, id int, name string, age int) (*models.CreatePersonOutput, error) {
	return e.repo.EditPerson(ctx, id, name, age)
}
