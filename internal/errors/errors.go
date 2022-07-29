package errors

import (
	"net/http"
	"strconv"
)

// Error is a struct that implements the standard error but has additional functionality
type Error struct {
	// HttpCode is used to contain status codes for HTTP based services
	HttpCode int
	// Code is application specific internal status code
	Code Code
	// Message contains more verbose information about the error that can be passed to user
	Message string
}

var (
	ErrUnableToBeginTransaction      = new(http.StatusInternalServerError, UnableToBeginTransaction, "internal server error")
	ErrUnableToCommitTransaction     = new(http.StatusInternalServerError, UnableToCommitTransaction, "unable to persist record")
	ErrUnableToUpdateBalance         = new(http.StatusInternalServerError, UnableToUpdateBalance, "unable to update balance")
	ErrUnableToCreateBankTransaction = new(http.StatusInternalServerError, UnableToCreateBankTransaction, "unable to create transaction")
	ErrUnableToFetchTransactions     = new(http.StatusInternalServerError, UnableToFetchTransactions, "unable to fetch transactions")
	ErrUnableToScanRows              = new(http.StatusInternalServerError, UnableToScanRows, "unable to fetch records")
)

func new(httpCode int, code Code, message string) *Error {
	return &Error{
		HttpCode: httpCode,
		Code:     code,
		Message:  message,
	}
}

// Error method returns a stringified version of the error
func (e *Error) Error() string {
	return strconv.Itoa(int(e.Code)) + ": " + e.Message
}

// Code is application specific internal status code
type Code uint64

const (
	BaseError Code = iota + 1000
	UnableToBeginTransaction
	UnableToCommitTransaction
	UnableToUpdateBalance
	UnableToCreateBankTransaction
	UnableToFetchTransactions
	UnableToScanRows
)
