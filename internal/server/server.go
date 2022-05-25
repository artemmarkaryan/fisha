package server

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/artemmarkaryan/fisha-facade/internal/config"
	"github.com/artemmarkaryan/fisha-facade/internal/service/interest"
	"github.com/artemmarkaryan/fisha-facade/internal/service/user"
	ui "github.com/artemmarkaryan/fisha-facade/internal/service/user-interest"
	"github.com/artemmarkaryan/fisha-facade/pkg/logy"
	"github.com/gernest/alien"
)

const module = "server"

type Server struct {
	interestSvc     interest.Service
	userSvc         user.Service
	userInterestSvc ui.Service
}

type handler func(w http.ResponseWriter, r *http.Request)

func (s Server) Serve(ctx context.Context) (err error) {
	m := alien.New()

	if m, err = s.registerHandlers(ctx, m); err != nil {
		return
	}

	logy.Log(ctx).Infoln("Running server at " + os.Getenv(config.ServerPort) + "...")
	if err = http.ListenAndServe(":"+os.Getenv(config.ServerPort), m); err != nil {
		return fmt.Errorf("running server error: %w", err)
	}

	return nil
}

func (s Server) registerHandlers(ctx context.Context, m *alien.Mux) (nm *alien.Mux, err error) {
	for _, err = range []error{
		m.Get("/interests", s.interests(ctx)),
		m.Post("/login", s.login(ctx)),
		m.Post("/react", s.react(ctx)),
		m.Post("/forget", s.forget(ctx)),
		m.Post("/addInterest", s.addInterest(ctx)),
	} {
		if err != nil {
			return nil, err
		}
	}

	return m, nil
}
