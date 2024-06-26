package main

import (
	"log"
	"net/http"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"main.go/pkg/handler"
	"main.go/pkg/repository"
	"main.go/pkg/service"
	"main.go/pkg/usecase"
)

func main() {
	db, err := NewPostgresDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	repo := repository.NewRepository(db)
	services := service.NewServices(repo)
	usecases := usecase.NewUsecase(services)
	handlers := handler.NewHandler(usecases)

	if err := http.ListenAndServe(":8080", handlers.InitRoutes()); err != nil {
		log.Fatal(err.Error())
	}
}

func NewPostgresDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect(
		"pgx", "host=localhost port=5432 user=postgres dbname=testdb password=1234 sslmode=disable")
	if err != nil {
		return nil, err
	}

	return db, err
}
