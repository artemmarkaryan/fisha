package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/artemmarkaryan/fisha/facade/internal/config"
	"github.com/artemmarkaryan/fisha/facade/internal/cron/recommendation"
	"github.com/artemmarkaryan/fisha/facade/pkg/database"
	"github.com/artemmarkaryan/fisha/facade/pkg/logy"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var ctx = context.Background()
	ctx = initLogger(ctx)
	ctx, err := initDatabase(ctx)
	if err != nil {
		log.Fatalln("unable to connect to database: ", err)
	}

	defer func(t time.Time) { logy.Log(ctx).Debugf("rec-s cron; elapsed: %v", time.Since(t)) }(time.Now())
	if err = new(recommendation.Cron).Process(ctx); err != nil {
		logy.Log(ctx).Errorf("error at rec-s cron: %q", err)
	}
}

func initLogger(ctx context.Context) context.Context {
	return logy.New(ctx)
}

func initDatabase(ctx context.Context) (context.Context, error) {
	return database.Init(ctx, database.Config{
		Host:     os.Getenv(config.DatabaseHostKey),
		Port:     os.Getenv(config.DatabasePortKey),
		User:     os.Getenv(config.DatabaseUserKey),
		Password: os.Getenv(config.DatabasePasswordKey),
		DBName:   os.Getenv(config.DatabaseDBNameKey),
	})
}
