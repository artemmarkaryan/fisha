package server

import (
	"context"
	"net/http"

	"github.com/artemmarkaryan/fisha/facade/pkg/logy"
	"github.com/artemmarkaryan/fisha/facade/pkg/marchy"
	"github.com/artemmarkaryan/fisha/facade/pkg/pb/gen/api"
	"github.com/gernest/alien"
)

const module = "server"

func Serve(ctx context.Context) (err error) {
	m := alien.New()

	if m, err = addHandlers(ctx, m); err != nil {
		return
	}

	logy.Log(ctx).Infoln("Running server...")
	if err = http.ListenAndServe(":8090", m); err != nil {
		return
	}

	return nil
}

func addHandlers(ctx context.Context, m *alien.Mux) (nm *alien.Mux, err error) {

	for _, err = range []error{
		m.Post("/", func(w http.ResponseWriter, r *http.Request) {
			_, _ = marchy.Obj[*api.EmptyRequest](ctx, r)
		}),

		m.Post("/string", func(w http.ResponseWriter, r *http.Request) {
			_, _ = marchy.Obj[*api.StringRequest](ctx, r)
		}),
	} {
	}

	return m, nil
}
