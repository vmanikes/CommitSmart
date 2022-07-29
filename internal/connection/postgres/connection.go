package postgres

import (
	"BankingService/internal/config"
	"context"
	"database/sql"
	"fmt"
	"github.com/flannel-dev-lab/cyclops/v2/logger"
	_ "github.com/lib/pq"
)

func GetConnection(ctx context.Context, cfg *config.Config) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.Username, cfg.Database.Password, cfg.Database.DBName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		logger.Error(ctx, "unable to open db connection", err)
		return nil, err
	}

	return db, nil
}
