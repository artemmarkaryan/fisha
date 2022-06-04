package server

import (
	"context"
	"net/http"

	"github.com/artemmarkaryan/fisha-facade/pkg/logy"
	"github.com/artemmarkaryan/fisha-facade/pkg/marchy"
	"github.com/artemmarkaryan/fisha-facade/pkg/network"
	"github.com/artemmarkaryan/fisha-facade/pkg/pb/gen/api"
)

func (s Server) recommend(ctx context.Context) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		reqObj, err := marchy.Obj[*api.IdMessage](ctx, r.Body)
		if err != nil {
			network.InternalError(w)
			return
		}

		userID := reqObj.GetId()
		a, err := s.r12nSvc.Get(ctx, userID)
		if err != nil {
			logy.Log(ctx).Errorf("getting recommendation err: %v", err)
			network.InternalError(w)
			return
		}

		network.Write(w, api.ActivityMessage{
			Id:      a.Id,
			Name:    a.Name,
			Address: a.Address,
			Meta:    a.Meta,
			Lon:     a.Lon,
			Lat:     a.Lat,
		})

		return
	}
}
