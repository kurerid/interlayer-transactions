package service

import (
	"context"

	"main.go/models"
	"main.go/pkg/repository"
	"main.go/pkg/transactor"
)

type (
	Creator interface {
		Create(ctx context.Context) (int, error)
	}

	Transactor interface {
		WithinTransaction(ctx context.Context, tFunc transactor.Func) (interface{}, error)
	}

	Editor interface {
		EditPerson(ctx context.Context, id int, name string, age int) (*models.CreatePersonOutput, error)
	}

	Services struct {
		Creator
		Transactor
		Editor
	}
)

func NewServices(repo *repository.Repository) *Services {
	return &Services{
		Creator:    NewCreatorService(repo.Creator),
		Transactor: NewTransactorService(repo.Transactor),
		Editor:     NewEditorService(repo.Editor),
	}
}
