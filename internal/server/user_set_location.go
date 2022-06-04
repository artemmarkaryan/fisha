package server

import (
	"context"
	"net/http"

	"github.com/artemmarkaryan/fisha-facade/internal/service/user"
	"github.com/artemmarkaryan/fisha-facade/pkg/marchy"
	"github.com/artemmarkaryan/fisha-facade/pkg/network"
	"github.com/artemmarkaryan/fisha-facade/pkg/pb/gen/api"
)

func (s Server) userSetLocation(ctx context.Context) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		reqObj, err := marchy.Obj[*api.SetLocationMessage](ctx, r.Body)
		if err != nil {
			network.InternalError(w)
			return
		}

		err = new(user.Service).SetLocation(ctx, reqObj.GetUserId(), reqObj.GetLon(), reqObj.GetLat())
		if err != nil {
			network.InternalError(w)
			return
		}

		network.Write(w, api.EmptyMessage{})
	}
}
