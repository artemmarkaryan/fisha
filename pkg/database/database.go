package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const databaseKey = "database"

type Config struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func (c Config) psql() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.host, c.port, c.user, c.password, c.dbname)
}

type closeDB func() error
type DBProvider func() (db *sql.DB, closer closeDB, err error)

func check(cfg Config) error {
	db, err := sql.Open("postgres", cfg.psql())
	if err != nil {
		return err
	}

	defer func() { _ = db.Close() }()

	if err = db.Ping(); err != nil {
		return err
	}

	return nil
}

func Init(ctx context.Context, cfg Config) (context.Context, error) {
	err := check(cfg)
	if err != nil {
		return ctx, err
	}

	var g DBProvider = func() (db *sql.DB, closer closeDB, err error) {
		db, err = sql.Open("postgres", cfg.psql())
		if err != nil {
			return
		}

		closer = func() error { return db.Close() }

		return
	}

	return context.WithValue(ctx, databaseKey, g), nil
}

func Get(ctx context.Context) (DBProvider, error) {
	v := ctx.Value(databaseKey)
	p, ok := v.(DBProvider)
	if !ok {
		return nil, fmt.Errorf(databaseKey+" has wrong type: %[1]v %[1]T", v)
	}

	return p, nil
}
