package server

import (
	"context"
	"net/http"

	"github.com/artemmarkaryan/fisha-facade/pkg/marchy"
	"github.com/artemmarkaryan/fisha-facade/pkg/network"
	"github.com/artemmarkaryan/fisha-facade/pkg/pb/gen/api"
)

func (s Server) login(ctx context.Context) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		userIdReq, err := marchy.Obj[*api.UserIdRequest](ctx, r.Body)
		if err != nil {
			network.WriteError(w, "internal: "+err.Error(), 500)
			return
		}
		if userIdReq.GetUserId() == 0 {
			network.WriteError(w, "bad user_id", 400)
			return
		}

		isNew, err := s.userSvc.Login(ctx, userIdReq.GetUserId())
		if err != nil {
			network.WriteError(w, "cant login: "+err.Error(), 500)
			return
		}

		network.Write(w, api.IsNewMessage{New: isNew})
	}
}
