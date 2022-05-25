package server

import (
	"context"
	"net/http"

	"github.com/artemmarkaryan/fisha-facade/pkg/marchy"
	"github.com/artemmarkaryan/fisha-facade/pkg/network"
	"github.com/artemmarkaryan/fisha-facade/pkg/pb/gen/api"
)

func (s Server) interestById(ctx context.Context) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		obj, err := marchy.Obj[*api.IdMessage](ctx, r.Body)
		if err != nil {
			network.InternalError(w)
		}

		interest, err := s.interestSvc.Get(ctx, obj.GetId())
		if err != nil {
			network.WriteError(w, err.Error(), 500)
		}

		network.Write(w, marchy.Force(api.StringMessage{S: interest.Name}))
	}
}
