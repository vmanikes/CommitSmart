package deposit

import "github.com/shopspring/decimal"

type Deposit struct {
	Amount        decimal.Decimal
	AccountNumber int
}
