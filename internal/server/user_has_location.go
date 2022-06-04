package server

import (
	"context"
	"net/http"

	"github.com/artemmarkaryan/fisha-facade/internal/service/user"
	"github.com/artemmarkaryan/fisha-facade/pkg/marchy"
	"github.com/artemmarkaryan/fisha-facade/pkg/network"
	"github.com/artemmarkaryan/fisha-facade/pkg/pb/gen/api"
)

func (s Server) userHasLocation(ctx context.Context) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		reqObj, err := marchy.Obj[*api.IdMessage](ctx, r.Body)
		if err != nil {
			network.InternalError(w)
			return
		}

		userID := reqObj.GetId()
		u, err := new(user.Service).Get(ctx, userID)
		if err != nil {
			network.InternalError(w)
			return
		}

		network.Write(w, api.BooleanMessage{Result: u.ValidLocation()})
	}
}
