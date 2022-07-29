package main

import (
	"BankingService/internal/bank"
	pgDataStore "BankingService/internal/bank/datastore/postgres"
	"BankingService/internal/config"
	"BankingService/internal/connection/postgres"
	"BankingService/internal/handler/http/v1"
	"BankingService/internal/server/http"
	"context"
	"github.com/flannel-dev-lab/cyclops/v2"
	"os"
)

func main() {
	ctx := context.Background()

	cfg := config.GetConfig(ctx)

	postgresConn, err := postgres.GetConnection(ctx, cfg)
	if err != nil {
		os.Exit(1)
	}

	bankService := bank.New(pgDataStore.New(postgresConn))

	handler := v1.New(bankService)

	router := http.GetRoutes(handler)

	cyclops.StartServer(":8080", router)
}
