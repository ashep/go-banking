package banking

import (
	"time"
)

type Transaction struct {
	ID             string    `db:"id"`
	BankID         string    `db:"bank_id"`
	AccountID      string    `db:"account_id"`
	Amount         float64   `db:"amount"`
	Balance        float64   `db:"balance"`
	Description    string    `db:"description"`
	CurrencyCode   int       `db:"currency_code"`
	MCC            int       `db:"mcc"`
	FeeAmount      float64   `db:"fee_amount"`
	CashbackAmount float64   `db:"cashback_amount"`
	CreatedAt      time.Time `db:"created_at"`
}
