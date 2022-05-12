package server

import (
	"context"
	"net/http"

	"github.com/artemmarkaryan/fisha/facade/internal/service/interest"
	"github.com/artemmarkaryan/fisha/facade/pkg/logy"
	"github.com/gernest/alien"
)

const module = "server"

type Server struct {
	interest interest.Service
}

type handler func(w http.ResponseWriter, r *http.Request)

func (s Server) Serve(ctx context.Context) (err error) {
	m := alien.New()

	if m, err = s.registerHandlers(ctx, m); err != nil {
		return
	}

	logy.Log(ctx).Infoln("Running server...")
	if err = http.ListenAndServe(":8090", m); err != nil {
		return
	}

	return nil
}

func (s Server) registerHandlers(ctx context.Context, m *alien.Mux) (nm *alien.Mux, err error) {
	for _, err = range []error{
		m.Get("/interests", s.interests(ctx)),
		m.Post("/login", s.login(ctx)),
	} {
		if err != nil {
			return nil, err
		}
	}

	return m, nil
}
