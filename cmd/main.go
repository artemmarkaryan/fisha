package main

import (
	"context"

	"github.com/artemmarkaryan/fisha/facade/internal/server"
	"github.com/artemmarkaryan/fisha/facade/pkg/logy"
)

func main() {
	var ctx context.Context
	ctx = initLogger(context.Background())

	if err := server.Serve(ctx); err != nil {
		logy.Log(ctx).Errorf("failed to serve: %w", err)
	}
}

func initLogger(ctx context.Context) context.Context {
	return logy.New(ctx)
}
