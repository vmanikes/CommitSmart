package v1

import (
	"context"
	"github.com/flannel-dev-lab/cyclops/v2/logger"
	"github.com/shopspring/decimal"
)

type DepositRequest struct {
	Amount        decimal.Decimal `json:"amount" validate:"required,gt=0"`
	AccountNumber int             `json:"account_number" validate:"required"`
}

func (d *DepositRequest) IsValid(ctx context.Context) bool {
	err := validate.Struct(d)
	if err != nil {
		logger.Error(ctx, "deposit request is invalid", nil)
		return false
	}

	return true
}
