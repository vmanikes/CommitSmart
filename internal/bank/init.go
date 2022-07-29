package bank

import "BankingService/internal/bank/datastore"

type bank struct {
	datastore.BankDataStore
}

func New(db datastore.BankDataStore) Bank {
	return &bank{
		db,
	}
}
