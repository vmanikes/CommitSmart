package bank

import (
	"BankingService/internal/entities/deposit"
	"BankingService/internal/entities/transaction"
	"BankingService/internal/entities/withdraw"
	"BankingService/internal/errors"
	"context"
)

func (b *bank) Deposit(ctx context.Context, deposit deposit.Deposit) *errors.Error {
	return b.BankDataStore.Deposit(ctx, deposit)
}

func (b *bank) Withdraw(ctx context.Context, withdraw withdraw.Withdraw) *errors.Error {
	return b.BankDataStore.Withdraw(ctx, withdraw)
}

func (b *bank) Transfer(ctx context.Context, transfer withdraw.Transfer) *errors.Error {
	return b.BankDataStore.Transfer(ctx, transfer)
}

func (b *bank) GetTransactions(ctx context.Context, userID, filterType, filterValue string) ([]transaction.Transaction, *errors.Error) {
	return b.BankDataStore.GetTransactions(ctx, userID, filterType, filterValue)
}
