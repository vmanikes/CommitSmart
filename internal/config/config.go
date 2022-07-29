package config

import (
	"context"
	"github.com/caarlos0/env/v6"
	"github.com/flannel-dev-lab/cyclops/v2/logger"
	"os"
)

type Config struct {
	Database
}

type Database struct {
	Host     string `env:"DBHOST" envDefault:"localhost"`
	Username string `env:"DBUSER" envDefault:"root"`
	Password string `env:"DBPASSWORD" envDefault:"root"`
	DBName   string `env:"DBNAME" envDefault:"postgres"`
	Port     uint16 `env:"DBPORT" envDefault:"5432"`
}

func GetConfig(ctx context.Context) *Config {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		logger.Error(ctx, "unable to build config", nil)
		os.Exit(1)
	}

	return &cfg
}
