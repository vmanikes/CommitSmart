package v1

import "BankingService/internal/bank"

type Handler struct {
	bank.Bank
}

func New(b bank.Bank) *Handler {
	return &Handler{b}
}
