package postgres

import (
	"BankingService/internal/entities/deposit"
	transactionEntity "BankingService/internal/entities/transaction"
	"BankingService/internal/entities/withdraw"
	"BankingService/internal/errors"
	"context"
	"database/sql"
	"github.com/flannel-dev-lab/cyclops/v2/logger"
	"github.com/shopspring/decimal"
)

func (d *dbClient) Deposit(ctx context.Context, deposit deposit.Deposit) *errors.Error {
	tx, err := d.db.Begin()
	if err != nil {
		logger.Error(ctx, "unable to start transaction", err)
		return errors.ErrUnableToBeginTransaction
	}

	var balance decimal.Decimal

	updateBalanceQuery := `
	UPDATE account_user
		SET balance = balance + $1
		WHERE account_number = $2
		RETURNING balance;
	`

	err = tx.QueryRowContext(ctx, updateBalanceQuery, deposit.Amount, deposit.AccountNumber).Scan(&balance)
	if err != nil {
		tx.Rollback()
		logger.Error(ctx, "unable to update balance", err)
		return errors.ErrUnableToUpdateBalance
	}

	transactionRecordQuery := `
	INSERT INTO account_transaction (user_id, balance, amount, transaction_type)
    VALUES (
            (
            SELECT id 
            FROM account_user 
            WHERE account_number = $1
            ), $2, $3, $4)
	`

	_, err = d.db.ExecContext(ctx, transactionRecordQuery,
		deposit.AccountNumber,
		balance,
		deposit.Amount,
		"deposit")
	if err != nil {
		tx.Rollback()
		logger.Error(ctx, "unable to create transaction", err)
		return errors.ErrUnableToCreateBankTransaction
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error(ctx, "unable to commit transaction", err)
		return errors.ErrUnableToCommitTransaction
	}

	return nil
}

func (d *dbClient) Withdraw(ctx context.Context, withdraw withdraw.Withdraw) *errors.Error {
	tx, err := d.db.Begin()
	if err != nil {
		logger.Error(ctx, "unable to start transaction", err)
		return errors.ErrUnableToBeginTransaction
	}

	var (
		balance decimal.Decimal
	)

	updateBalanceQuery := `
	UPDATE account_user
		SET balance = balance - $1
		WHERE account_number = $2 AND id = $3 AND balance >= $1
		RETURNING balance;
	`

	err = tx.QueryRowContext(ctx, updateBalanceQuery,
		withdraw.Amount, withdraw.AccountNumber, withdraw.UserID).Scan(&balance)
	if err != nil {
		tx.Rollback()
		logger.Error(ctx, "unable to update balance", err)
		return errors.ErrUnableToUpdateBalance
	}

	transactionRecordQuery := `
	INSERT INTO account_transaction (user_id, balance, amount, transaction_type)
    VALUES (
            $1, $2, $3, $4)
	`

	_, err = d.db.ExecContext(ctx, transactionRecordQuery,
		withdraw.UserID,
		balance,
		withdraw.Amount,
		"withdraw")
	if err != nil {
		tx.Rollback()
		logger.Error(ctx, "unable to create transaction", err)
		return errors.ErrUnableToCreateBankTransaction
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error(ctx, "unable to commit transaction", err)
		return errors.ErrUnableToCommitTransaction
	}

	return nil
}

func (d *dbClient) Transfer(ctx context.Context, transfer withdraw.Transfer) *errors.Error {
	tx, err := d.db.Begin()
	if err != nil {
		logger.Error(ctx, "unable to start transaction", err)
		return errors.ErrUnableToBeginTransaction
	}

	var (
		senderBalance   decimal.Decimal
		receiverBalance decimal.Decimal
	)

	updateBalanceQuery := `
	UPDATE account_user
		SET balance = balance - $1
		WHERE account_number = $2 AND balance >= $1
		RETURNING balance;
	`

	err = tx.QueryRowContext(ctx, updateBalanceQuery, transfer.Amount, transfer.CurrentAccountNumber).Scan(&senderBalance)
	if err != nil {
		tx.Rollback()
		logger.Error(ctx, "unable to update balance", err)
		return errors.ErrUnableToUpdateBalance
	}

	updateBalanceQuery = `
	UPDATE account_user
		SET balance = balance + $1
		WHERE account_number = $2
		RETURNING balance;
	`

	err = tx.QueryRowContext(ctx, updateBalanceQuery, transfer.Amount, transfer.DestinationAccountNumber).Scan(&senderBalance)
	if err != nil {
		tx.Rollback()
		logger.Error(ctx, "unable to update balance", err)
		return errors.ErrUnableToUpdateBalance
	}

	transactionRecordQuery := `
	INSERT INTO account_transaction
            (user_id,
             balance,
             amount,
             transaction_type)
	VALUES ((SELECT id
              FROM   account_user
              WHERE  account_number = $1),
             $2,
             $3,
             'withdraw'),
            ((SELECT id
              FROM   account_user
              WHERE  account_number = $4),
             $5,
             $3,
             'deposit') 
	`

	_, err = d.db.ExecContext(ctx, transactionRecordQuery,
		transfer.CurrentAccountNumber,
		senderBalance,
		transfer.Amount,
		transfer.DestinationAccountNumber,
		receiverBalance)
	if err != nil {
		tx.Rollback()
		logger.Error(ctx, "unable to create transaction", err)
		return errors.ErrUnableToCreateBankTransaction
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error(ctx, "unable to commit transaction", err)
		return errors.ErrUnableToCommitTransaction
	}

	return nil
}

func (d *dbClient) GetTransactions(ctx context.Context, userID, filterType, filterValue string) ([]transactionEntity.Transaction, *errors.Error) {
	var (
		query string
		rows  *sql.Rows
		err   error
	)

	if filterType == "date" {
		query = `
			SELECT * FROM account_transaction WHERE user_id = $1 AND created_on > $2
		`

		rows, err = d.db.QueryContext(ctx, query, userID, filterValue)
		if err != nil {
			logger.Error(ctx, "unable to fetch transactions", err)
			return nil, errors.ErrUnableToFetchTransactions
		}
	} else if filterType == "" {
		query = `
			SELECT * FROM account_transaction WHERE user_id = $1
		`

		rows, err = d.db.QueryContext(ctx, query, userID)
		if err != nil {
			logger.Error(ctx, "unable to fetch transactions", err)
			return nil, errors.ErrUnableToFetchTransactions
		}
	} else {
		query = `
			SELECT * FROM account_transaction WHERE user_id = $1 AND transaction_type = $2;
		`

		rows, err = d.db.QueryContext(ctx, query, userID, filterType)
		if err != nil {
			logger.Error(ctx, "unable to fetch transactions", err)
			return nil, errors.ErrUnableToFetchTransactions
		}
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			logger.Error(ctx, "unable to close rows", err)
		}
	}()

	var transactions []transaction

	for rows.Next() {
		var transaction transaction

		err = rows.Scan(
			&transaction.ID,
			&transaction.UserID,
			&transaction.Balance,
			&transaction.Amount,
			&transaction.CreatedOn,
			&transaction.Type)
		if err != nil {
			logger.Error(ctx, "unable to scan rows", err)
			return nil, errors.ErrUnableToScanRows
		}

		transactions = append(transactions, transaction)
	}

	transactionEntities := make([]transactionEntity.Transaction, 0, len(transactions))

	for _, transaction := range transactions {
		transactionEntities = append(transactionEntities, transaction.serializeTransactionsToEntity())
	}

	return transactionEntities, nil
}
