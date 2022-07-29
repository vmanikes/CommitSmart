package transaction

import (
	"github.com/shopspring/decimal"
	"time"
)

type Transaction struct {
	ID        string
	UserID    string
	Type      string
	Balance   decimal.Decimal
	Amount    decimal.Decimal
	CreatedOn time.Time
}
