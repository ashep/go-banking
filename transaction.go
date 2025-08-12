package banking

import (
	"time"
)

type Transaction struct {
	ID             string    `json:"id"`
	BankID         string    `json:"bank_id"`
	AccountID      string    `json:"account_id"`
	Amount         float64   `json:"amount"`
	Balance        float64   `json:"balance"`
	Description    string    `json:"description"`
	CurrencyCode   int       `json:"currency_code"`
	MCC            int       `json:"mcc"`
	FeeAmount      float64   `json:"fee_amount"`
	CashbackAmount float64   `json:"cashback_amount"`
	CreatedAt      time.Time `json:"created_at"`
}
