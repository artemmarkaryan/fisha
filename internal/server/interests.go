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
		_, _ = marchy.Obj[*api.EmptyMessage](ctx, r.Body)

		interests, err := s.interestSvc.List(ctx)
		if err != nil {
			network.WriteError(w, err.Error(), 500)
		}

		response := api.InterestsResponse{}
		for _, interest := range interests {
			response.Interest = append(response.Interest, &api.InterestsResponse_Interest{
				Id:   interest.Id,
				Name: interest.Name,
			})
		}

		network.Write(w, &response)
	}
}
