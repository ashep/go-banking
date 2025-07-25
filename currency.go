package banking

import (
	"errors"
	"time"

	"github.com/shopspring/decimal"
)

var ErrCurrencyNotFound = errors.New("currency not found")

var currencyIcons = map[string]string{
	"CHF": "🇨🇭",
	"EUR": "🇪🇺",
	"UAH": "🇺🇦",
	"USD": "🇺🇸",
}

var currencyNumbers = map[string]int{
	"CHF": 756,
	"EUR": 978,
	"UAH": 980,
	"USD": 840,
}

var currencyCodes = map[int]string{
	756: "CHF",
	978: "EUR",
	980: "UAH",
	840: "USD",
}

var currencyDigits = map[int]int{
	756: 2,
	978: 2,
	980: 2,
	840: 2,
}

type Currency struct {
	Code   string `json:"code"`   // ISO 4217 3-letter code
	Num    int    `json:"num"`    // ISO 4217 number code
	Digits int    `json:"digits"` // Number of decimal places
	Icon   string `json:"icon"`   // Emoji icon representation
}

type CurrencyRate struct {
	Base   Currency        `json:"base"`   // The currency being converted from
	Target Currency        `json:"target"` // The currency being converted to
	Rate   decimal.Decimal `json:"rate"`   // How much of currency `Base` you need to get 1 unit of currency `Target`
	Date   time.Time       `json:"date"`   // The date and time when the rate was last updated
}

func NewCurrencyByCode(code string) (Currency, error) {
	num, exists := currencyNumbers[code]
	if !exists {
		return Currency{}, ErrCurrencyNotFound
	}

	return Currency{
		Code:   code,
		Num:    num,
		Digits: currencyDigits[num],
		Icon:   currencyIcons[code],
	}, nil
}

func NewCurrencyByNum(num int) (Currency, error) {
	code, exists := currencyCodes[num]
	if !exists {
		return Currency{}, ErrCurrencyNotFound
	}

	return Currency{
		Code:   code,
		Num:    num,
		Digits: currencyDigits[num],
		Icon:   currencyIcons[code],
	}, nil
}
