package banking

import (
	"time"
)

type Account struct {
	ID        string    `db:"id"`
	BankID    string    `db:"bank_id"`
	Name      string    `db:"name"`
	Balance   float64   `db:"balance"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
