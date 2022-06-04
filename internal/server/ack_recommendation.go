package server

import (
	"context"
	"net/http"

	"github.com/artemmarkaryan/fisha-facade/internal/service/recommendation"
	"github.com/artemmarkaryan/fisha-facade/pkg/logy"
	"github.com/artemmarkaryan/fisha-facade/pkg/marchy"
	"github.com/artemmarkaryan/fisha-facade/pkg/network"
	"github.com/artemmarkaryan/fisha-facade/pkg/pb/gen/api"
)

func (s Server) ackRecommendation(ctx context.Context) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		obj, err := marchy.Obj[*api.AckRecommendationMessage](ctx, r.Body)
		if err != nil {
			network.InternalError(w)
			return
		}

		if err = new(recommendation.Service).Ack(ctx, obj.GetUserId(), obj.GetActivityId()); err != nil {
			logy.Log(ctx).Errorf("cant ack ecommendation: %v", err)
			network.InternalError(w)
			return
		}

		network.Write(w, api.EmptyMessage{})
	}
}
