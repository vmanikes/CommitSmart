package withdraw

import "github.com/shopspring/decimal"

type Withdraw struct {
	Amount        decimal.Decimal
	AccountNumber int
	UserID        string
}

type Transfer struct {
	Amount                   decimal.Decimal
	CurrentAccountNumber     int
	DestinationAccountNumber int
}
