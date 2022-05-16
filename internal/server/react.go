package server

import (
	"context"
	"net/http"

	"github.com/artemmarkaryan/fisha-facade/internal/config"
	"github.com/artemmarkaryan/fisha-facade/internal/service/reaction"
	"github.com/artemmarkaryan/fisha-facade/pkg/logy"
	"github.com/artemmarkaryan/fisha-facade/pkg/marchy"
	"github.com/artemmarkaryan/fisha-facade/pkg/network"
	"github.com/artemmarkaryan/fisha-facade/pkg/pb/gen/api"
	"github.com/artemmarkaryan/fisha-facade/pkg/rabbit"
)

func (s Server) react(ctx context.Context) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		reqObj, err := marchy.Obj[*api.ReactRequest](ctx, r.Body)
		if err != nil {
			network.InternalError(w)
		}

		reactionType := reqObj.GetReaction().String()
		if !reaction.ValidateReactionType(reactionType) {
			network.WriteBadRequestError(w, "unknown reaction type: "+reactionType)
		}

		if err = rabbit.Produce(ctx, config.ReactionQueueName, marchy.Force(reaction.Reaction{
			UserID:     reqObj.GetUserId(),
			ActivityID: reqObj.GetActivityId(),
			Reaction:   reactionType,
		})); err != nil {
			logy.Log(ctx).Errorf("producing reaction error: %v", err)
			network.InternalError(w)
		}
	}
}
