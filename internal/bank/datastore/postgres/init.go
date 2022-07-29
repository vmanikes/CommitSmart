package postgres

import (
	"BankingService/internal/bank/datastore"
	"database/sql"
	_ "github.com/lib/pq"
)

type dbClient struct {
	db *sql.DB
}

func New(db *sql.DB) datastore.BankDataStore {
	return &dbClient{db: db}
}
