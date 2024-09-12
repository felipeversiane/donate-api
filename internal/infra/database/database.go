package database

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	Database DatabaseInterface
	once     sync.Once
)

type database struct {
	db *pgxpool.Pool
}

type DatabaseInterface interface {
	GetDB() *pgxpool.Pool
	Ping(ctx context.Context) error
	Close()
}

func NewDatabase(db *pgxpool.Pool) DatabaseInterface {
	return &database{db: db}
}

func NewDatabaseConnection(ctx context.Context, dsn string) error {
	var err error
	once.Do(func() {
		poolConfig, parseErr := pgxpool.ParseConfig(dsn)
		if parseErr != nil {
			err = parseErr
			return
		}

		pool, poolErr := pgxpool.NewWithConfig(ctx, poolConfig)
		if poolErr != nil {
			err = poolErr
			return
		}

		Database = NewDatabase(pool)

		if pingErr := Database.Ping(ctx); pingErr != nil {
			err = pingErr
		}
	})

	return err
}

func (pg *database) Ping(ctx context.Context) error {
	return pg.db.Ping(ctx)
}

func (pg *database) Close() {
	pg.db.Close()
}

func (pg *database) GetDB() *pgxpool.Pool {
	return pg.db
}
