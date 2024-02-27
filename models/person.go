package models

type (
	Person struct {
		Name string `json:"name" db:"name"`
		Age  int    `json:"age" db:"age"`
	}

	CreatePersonInput struct {
		Person
	}

	CreatePersonOutput struct {
		Id int `json:""json:"id" db:"id"`
		Person
	}
)
