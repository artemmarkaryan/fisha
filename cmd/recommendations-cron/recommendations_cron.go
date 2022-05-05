package main

import (
	"context"
	"log"
	"os"

	"github.com/artemmarkaryan/fisha/facade/internal/config"
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

	new(recommendation.Cron).Process(ctx)
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
