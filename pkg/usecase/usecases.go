package usecase

import "main.go/pkg/service"

type Usecase struct {
	Service *service.Services
}

func NewUsecase(service *service.Services) *Usecase {
	return &Usecase{Service: service}
}
