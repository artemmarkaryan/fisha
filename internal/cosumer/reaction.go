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
	for t := range objs {
		if err = svc.Calculate(ctx, t); err != nil {
			logy.Log(ctx).Errorf("cant calculare reaction: %q", err)
		}
	}
}
