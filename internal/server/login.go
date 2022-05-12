package server

import (
	"context"
	"net/http"

	"github.com/artemmarkaryan/fisha/facade/pkg/marchy"
	"github.com/artemmarkaryan/fisha/facade/pkg/pb/gen/api"
)

func (s Server) login(ctx context.Context) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, _ := marchy.Obj[*api.UserIdRequest](ctx, r)

	}
}
