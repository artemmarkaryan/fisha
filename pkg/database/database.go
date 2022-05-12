package database

import (
	"context"
	"fmt"

	"github.com/artemmarkaryan/fisha-facade/pkg/logy"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const databaseKey = "database"

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func (c Config) psql() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.DBName)
}

type closeDB func() error
type DBProvider func() (db *sqlx.DB, closer closeDB, err error)

func check(ctx context.Context, cfg Config) error {
	db, err := sqlx.Connect("postgres", cfg.psql())
	if err != nil {
		return err
	}

	defer func() { _ = db.Close() }()

	if err = db.Ping(); err != nil {
		return err
	}

	logy.Log(ctx).Infoln("connected to database")

	return nil
}

func Init(ctx context.Context, cfg Config) (context.Context, error) {
	err := check(ctx, cfg)
	if err != nil {
		return ctx, err
	}

	var g DBProvider = func() (db *sqlx.DB, closer closeDB, err error) {
		db, err = sqlx.Open("postgres", cfg.psql())
		if err != nil {
			return
		}

		closer = func() error { return db.Close() }

		return
	}

	return context.WithValue(ctx, databaseKey, g), nil
}

func Get(ctx context.Context) DBProvider {
	v := ctx.Value(databaseKey)
	p, ok := v.(DBProvider)
	if !ok {
		panic(fmt.Errorf(databaseKey+" has wrong type: %[1]v %[1]T", v))
	}

	return p
}
