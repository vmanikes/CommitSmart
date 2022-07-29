package v1

import (
	"context"
	"github.com/flannel-dev-lab/cyclops/v2/logger"
	"github.com/shopspring/decimal"
)

type WithdrawRequest struct {
	Amount        decimal.Decimal `json:"amount" validate:"required,gt=0"`
	AccountNumber int             `json:"account_number" validate:"required"`
	UserID        string          `json:"user_id" validate:"required"`
}

func (w *WithdrawRequest) IsValid(ctx context.Context) bool {
	err := validate.Struct(w)
	if err != nil {
		logger.Error(ctx, "withdraw request is invalid", nil)
		return false
	}

	return true
}

type TransferRequest struct {
	Amount                   decimal.Decimal `json:"amount" validate:"required,gt=0"`
	CurrentAccountNumber     int             `json:"current_account_number" validate:"required"`
	DestinationAccountNumber int             `json:"destination_account_number" validate:"required"`
}

func (t *TransferRequest) IsValid(ctx context.Context) bool {
	err := validate.Struct(t)
	if err != nil {
		logger.Error(ctx, "transfer request is invalid", nil)
		return false
	}

	return true
}
