package main

import (
	"context"
	"log"
	"os"

	"github.com/artemmarkaryan/fisha-facade/internal/config"
	"github.com/artemmarkaryan/fisha-facade/internal/server"
	"github.com/artemmarkaryan/fisha-facade/pkg/database"
	"github.com/artemmarkaryan/fisha-facade/pkg/logy"
	"github.com/artemmarkaryan/fisha-facade/pkg/rabbit"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var ctx = context.Background()
	var err error

	ctx = initLogger(ctx)

	if ctx, err = initDatabase(ctx); err != nil {
		log.Fatalln("unable to connect to database: ", err)
	}

	if ctx, err = initRabbit(ctx); err != nil {
		log.Fatalln("unable to connect to RabbitMQ: ", err)
	}

	s := new(server.Server)

	if err = s.Serve(ctx); err != nil {
		logy.Log(ctx).Errorf("failed to serve: %w", err)
	}
}

func initLogger(ctx context.Context) context.Context {
	ctx2 := logy.New(ctx)
	logy.Log(ctx2).Infoln("logger configured")
	return ctx2
}

func initDatabase(ctx context.Context) (ctx2 context.Context, err error) {
	ctx2, err = database.Init(ctx, database.Config{
		Host:     os.Getenv(config.DatabaseHostKey),
		Port:     os.Getenv(config.DatabasePortKey),
		User:     os.Getenv(config.DatabaseUserKey),
		Password: os.Getenv(config.DatabasePasswordKey),
		DBName:   os.Getenv(config.DatabaseDBNameKey),
	})

	logy.Log(ctx).Infoln("connected to database")

	return
}

func initRabbit(ctx context.Context) (ctx2 context.Context, err error) {
	ctx2, err = rabbit.Init(ctx, rabbit.Config{
		Host:     os.Getenv(config.RabbitHostKey),
		Port:     os.Getenv(config.RabbitPortKey),
		User:     os.Getenv(config.RabbitUserKey),
		Password: os.Getenv(config.RabbitPasswordKey),
	})

	logy.Log(ctx).Infoln("connected to rabbitmq")

	return
}
