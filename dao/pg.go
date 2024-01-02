package dao

import (
	"context"
	"database/sql"
	"log-management/config"
	"log-management/domain"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

// InitPG initialises a bun.DB instance
func InitPG() *bun.DB {
	if config.DB_URL == "" {
		panic("POSTGRESQL_URL environment variable is required to launch")
	}

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(config.DB_URL)))

	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		// disable the hook
		bundebug.WithEnabled(false),

		// BUNDEBUG=1 logs failed queries
		// BUNDEBUG=2 logs all queries
		bundebug.FromEnv("BUNDEBUG"),
	))

	return db
}

//lint:ignore U1000 we don't really use this anywhere, just to bootstrap initial migration
func CreateTables(ctx context.Context, db *bun.DB) error {
	db.NewCreateTable().Model((*domain.Microservice)(nil)).
		IfNotExists().
		Temp().
		Exec(ctx)

	db.NewCreateTable().
		Model((*domain.Log)(nil)).
		IfNotExists().
		Temp().
		Exec(ctx)

	return nil
}

type pgDAO struct {
	DB bun.IDB
}

func NewPGDAO(db bun.IDB) *pgDAO {
	return &pgDAO{DB: db}
}
