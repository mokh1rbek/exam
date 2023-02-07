package storage_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"exam/config"
	"exam/storage/postgres"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	categoryRepo *postgres.CategoriesRepo
)

func TestMain(m *testing.M) {

	cfg := config.Load()

	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		panic(err)
	}

	config.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		panic(err)
	}

	categoryRepo = postgres.NewCategoriesRepo(pool)

	os.Exit(m.Run())
}
