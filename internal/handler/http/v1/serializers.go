package v1

import (
	"BankingService/internal/entities/deposit"
	"BankingService/internal/entities/transaction"
	"BankingService/internal/entities/withdraw"
	v1 "BankingService/models/http/v1"
	"time"
)

func serializeDepositToEntity(d v1.DepositRequest) deposit.Deposit {
	return deposit.Deposit{
		Amount:        d.Amount,
		AccountNumber: d.AccountNumber,
	}
}

func serializeWithdrawToEntity(w v1.WithdrawRequest) withdraw.Withdraw {
	return withdraw.Withdraw{
		Amount:        w.Amount,
		AccountNumber: w.AccountNumber,
		UserID:        w.UserID,
	}
}

func serializeTransferToEntity(t v1.TransferRequest) withdraw.Transfer {
	return withdraw.Transfer{
		Amount:                   t.Amount,
		CurrentAccountNumber:     t.CurrentAccountNumber,
		DestinationAccountNumber: t.DestinationAccountNumber,
	}
}

func serializeTransactionsToModel(transactions []transaction.Transaction) []v1.Transaction {
	modelTransactions := make([]v1.Transaction, 0, len(transactions))

	for _, t := range transactions {
		modelTransactions = append(modelTransactions, v1.Transaction{
			ID:        t.ID,
			UserID:    t.UserID,
			Type:      t.Type,
			Balance:   t.Balance,
			Amount:    t.Amount,
			CreatedOn: t.CreatedOn.Format(time.RFC3339),
		})
	}

	return modelTransactions
}
