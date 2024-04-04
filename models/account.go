package models

type (
	Account struct {
		Email    string `json:"email" db:"email"`
		Password string `json:"password" db:"password"`
	}

	RegisterAccountInput struct {
		Account
	}
)
