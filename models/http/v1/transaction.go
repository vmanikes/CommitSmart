package v1

import (
	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID        string          `json:"id"`
	UserID    string          `json:"user_id"`
	Type      string          `json:"type"`
	Balance   decimal.Decimal `json:"balance"`
	Amount    decimal.Decimal `json:"amount"`
	CreatedOn string          `json:"created_on"`
}
