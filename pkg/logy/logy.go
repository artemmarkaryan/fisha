package logy

import (
	"context"
	"io"
	"os"

	log "github.com/amoghe/distillog"
)

const key = "logger"

func New(ctx context.Context) context.Context {
	logFile, err := os.OpenFile("./tmp/log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)

	l := log.NewStreamLogger("", struct {
		io.Writer
		io.Closer
	}{
		Writer: mw,
	})

	return context.WithValue(ctx, key, l)
}

func Log(ctx context.Context) (l log.Logger) {
	l, _ = ctx.Value(key).(log.Logger)
	return
}
