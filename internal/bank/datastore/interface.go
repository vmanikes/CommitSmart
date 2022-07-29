package datastore

import (
	"BankingService/internal/entities/deposit"
	"BankingService/internal/entities/transaction"
	"BankingService/internal/entities/withdraw"
	"BankingService/internal/errors"
	"context"
)

type BankDataStore interface {
	Deposit(ctx context.Context, deposit deposit.Deposit) *errors.Error
	Withdraw(ctx context.Context, withdraw withdraw.Withdraw) *errors.Error
	Transfer(ctx context.Context, transfer withdraw.Transfer) *errors.Error
	GetTransactions(ctx context.Context, userID, filterType, filterValue string) ([]transaction.Transaction, *errors.Error)
}
