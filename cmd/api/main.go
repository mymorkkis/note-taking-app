package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/caarlos0/env/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mymorkkis/notes-app/internal/dbal"
)

type config struct {
	APIPort          int    `env:"API_PORT,notEmpty"`
	APIEnv           string `env:"API_ENV" envDefault:"development"`
	APIVersion       string `env:"API_VERSION,notEmpty"`
	PostgresDB       string `env:"POSTGRES_DB,notEmpty"`
	PostgresUser     string `env:"POSTGRES_USER,notEmpty"`
	PostgresPassword string `env:"POSTGRES_PASSWORD,notEmpty"`
	PostgresPort     int    `env:"POSTGRES_PORT,notEmpty"`
	PoolMaxConns     int    `env:"POOL_MAX_CONNS" envDefault:"50"`
}

func main() {
	var logLevel = new(slog.LevelVar)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	pool, err := openConnectionPool(cfg)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer pool.Close()

	if strings.ToLower(cfg.APIEnv) == "development" {
		logLevel.Set(slog.LevelDebug)
	}

	app := &application{
		systemInfo: systemInfo{
			version:     cfg.APIVersion,
			environment: cfg.APIEnv,
		},
		logger:  logger,
		queries: dbal.New(pool),
	}

	app.logger.Info(fmt.Sprintf("Running app on port: %d", cfg.APIPort))

	http.ListenAndServe(fmt.Sprintf(":%d", cfg.APIPort), app.routes())
}

func openConnectionPool(config config) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@notes_db:%s/%s?sslmode=disable&pool_max_conns=%s",
		config.PostgresUser,
		config.PostgresPassword,
		strconv.Itoa(config.PostgresPort),
		config.PostgresDB,
		strconv.Itoa(config.PoolMaxConns),
	)

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return pool, nil
}
