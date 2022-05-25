package server

import (
	"context"
	"net/http"

	ui "github.com/artemmarkaryan/fisha-facade/internal/service/user-interest"
	"github.com/artemmarkaryan/fisha-facade/pkg/logy"
	"github.com/artemmarkaryan/fisha-facade/pkg/marchy"
	"github.com/artemmarkaryan/fisha-facade/pkg/network"
	"github.com/artemmarkaryan/fisha-facade/pkg/pb/gen/api"
)

func (s Server) addInterest(ctx context.Context) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		o, err := marchy.Obj[*api.AddInterestRequest](ctx, r.Body)
		if err != nil {
			network.InternalError(w)
			return
		}

		inserted, err := s.userInterestSvc.Insert(ctx, ui.UserInterest{
			UserId:     o.GetUserId(),
			InterestId: o.GetInterestId(),
		})
		if err != nil {
			logy.Log(ctx).Errorf("error inserting user inerest: %w", err)
			network.InternalError(w)
			return
		}

		network.Write(w, api.IsNewMessage{New: inserted})
	}
}
