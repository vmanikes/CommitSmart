package postgres

import (
	transactionEntity "BankingService/internal/entities/transaction"
	"github.com/shopspring/decimal"
	"time"
)

type transaction struct {
	ID        string          `json:"id"`
	UserID    string          `json:"user_id"`
	Balance   decimal.Decimal `json:"balance"`
	Amount    decimal.Decimal `json:"amount"`
	Type      string          `json:"transaction_type"`
	CreatedOn time.Time       `json:"created_on"`
}

func (t *transaction) serializeTransactionsToEntity() transactionEntity.Transaction {
	return transactionEntity.Transaction{
		ID:        t.ID,
		UserID:    t.UserID,
		Type:      t.Type,
		Balance:   t.Balance,
		Amount:    t.Amount,
		CreatedOn: t.CreatedOn,
	}
}
