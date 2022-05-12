package server

import (
	"context"
	"net/http"

	"github.com/artemmarkaryan/fisha-facade/pkg/marchy"
	"github.com/artemmarkaryan/fisha-facade/pkg/network"
	"github.com/artemmarkaryan/fisha-facade/pkg/pb/gen/api"
)

func (s Server) interests(ctx context.Context) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = marchy.Obj[*api.EmptyRequest](ctx, r)

		i, err := s.interestSvc.Names(ctx)
		if err != nil {
			network.WriteError(w, err.Error(), 500)
		}

		network.Write(w, i)
	}
}
