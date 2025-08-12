package banking

import (
	"time"
)

type Account struct {
	ID        string    `json:"id"`
	BankID    string    `json:"bank_id"`
	Name      string    `json:"name"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
