package consumer

import (
	"context"

	"github.com/artemmarkaryan/fisha-facade/internal/config"
	"github.com/artemmarkaryan/fisha-facade/internal/service/reaction"
	"github.com/artemmarkaryan/fisha-facade/pkg/logy"
	"github.com/artemmarkaryan/fisha-facade/pkg/rabbit"
)

func HandleReaction(ctx context.Context, stop chan struct{}) {
	objs, err := rabbit.Consume[reaction.Reaction](ctx, config.ReactionQueueName, stop)
	if err != nil {
		panic(config.ReactionQueueName + " cant consume: " + err.Error())
	}

	svc := new(reaction.Service)

	logy.Log(ctx).Infoln("running reaction consumer ...")
	for r := range objs {
		if err = svc.Calculate(ctx, r); err != nil {
			logy.Log(ctx).Errorf("cant calculate reaction: %q", err)
		} else {
			logy.Log(ctx).Debugf("reaction effect calculated: user: %v, activity: %v, reaction: %v", r.UserID, r.ActivityID, r.Reaction)
		}
	}

	logy.Log(ctx).Infoln("running reaction consumer finished")
}
