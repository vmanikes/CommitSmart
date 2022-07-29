package v1

import (
	v1 "BankingService/models/http/v1"
	"encoding/json"
	"github.com/flannel-dev-lab/cyclops/v2"
	"github.com/flannel-dev-lab/cyclops/v2/response"
	"net/http"
	"strings"
	"time"
)

func (h *Handler) Deposit(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var depositRequestBody v1.DepositRequest

	if err := json.NewDecoder(r.Body).Decode(&depositRequestBody); err != nil {
		response.ErrorResponse(http.StatusBadRequest, "unparseable request body", w)
		return
	}

	if !depositRequestBody.IsValid(ctx) {
		response.ErrorResponse(http.StatusBadRequest, "invalid request body", w)
		return
	}

	err := h.Bank.Deposit(ctx, serializeDepositToEntity(depositRequestBody))
	if err != nil {
		response.ErrorResponse(err.HttpCode, err.Message, w)
		return
	}

	response.SuccessResponse(http.StatusOK, w, nil)
	return
}

func (h *Handler) Withdraw(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var withdrawRequestBody v1.WithdrawRequest

	if err := json.NewDecoder(r.Body).Decode(&withdrawRequestBody); err != nil {
		response.ErrorResponse(http.StatusBadRequest, "unparseable request body", w)
		return
	}

	if !withdrawRequestBody.IsValid(ctx) {
		response.ErrorResponse(http.StatusBadRequest, "invalid request body", w)
		return
	}

	err := h.Bank.Withdraw(ctx, serializeWithdrawToEntity(withdrawRequestBody))
	if err != nil {
		response.ErrorResponse(err.HttpCode, err.Message, w)
		return
	}

	response.SuccessResponse(http.StatusOK, w, nil)
	return
}

func (h *Handler) Transfer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var transferRequestBody v1.TransferRequest

	if err := json.NewDecoder(r.Body).Decode(&transferRequestBody); err != nil {
		response.ErrorResponse(http.StatusBadRequest, "unparseable request body", w)
		return
	}

	if !transferRequestBody.IsValid(ctx) {
		response.ErrorResponse(http.StatusBadRequest, "invalid request body", w)
		return
	}

	err := h.Bank.Transfer(ctx, serializeTransferToEntity(transferRequestBody))
	if err != nil {
		response.ErrorResponse(err.HttpCode, err.Message, w)
		return
	}

	response.SuccessResponse(http.StatusOK, w, nil)
	return
}

func (h *Handler) Transactions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	filterType := r.URL.Query().Get("filter_type")
	filterValue := r.URL.Query().Get("filter_value")

	if len(filterType) > 0 {
		if !strings.Contains(filterType, "deposit") && !strings.Contains(filterType, "withdraw") && !strings.Contains(filterType, "date") {
			response.ErrorResponse(http.StatusBadRequest, "invalid query params", w)
			return
		}
	}

	if filterType == "date" {
		_, err := time.Parse("2006-01-02", filterValue)
		if err != nil {
			response.ErrorResponse(http.StatusBadRequest, "invalid date format", w)
			return
		}
	}

	transactions, err := h.Bank.GetTransactions(ctx, cyclops.Param(r, "user_id"), filterType, filterValue)
	if err != nil {
		response.ErrorResponse(err.HttpCode, err.Message, w)
		return
	}

	response.SuccessResponse(http.StatusOK, w, serializeTransactionsToModel(transactions))
	return
}
