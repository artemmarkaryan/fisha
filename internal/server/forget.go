package server

import (
	"context"
	"net/http"

	"github.com/artemmarkaryan/fisha-facade/pkg/logy"
	"github.com/artemmarkaryan/fisha-facade/pkg/marchy"
	"github.com/artemmarkaryan/fisha-facade/pkg/network"
	"github.com/artemmarkaryan/fisha-facade/pkg/pb/gen/api"
)

func (s Server) forget(ctx context.Context) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		userIdReq, err := marchy.Obj[*api.UserIdRequest](ctx, r.Body)
		if err != nil {
			network.InternalError(w)
			return
		}
		if userIdReq.GetUserId() == 0 {
			network.WriteBadRequestError(w, "bad user_id")
			return
		}

		if err = s.userSvc.Forget(ctx, userIdReq.GetUserId()); err != nil {
			logy.Log(ctx).Errorf("cant forget user: %v", err)
			network.InternalError(w)
			return
		}
	}
}
