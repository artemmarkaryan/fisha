package recommendation

import (
	"context"

	"github.com/artemmarkaryan/fisha/facade/internal/service/activity"
	"github.com/artemmarkaryan/fisha/facade/internal/service/recommendation"
	"github.com/artemmarkaryan/fisha/facade/internal/service/user"
)

const userBatch = 100
const userActivityDistance = 50 * 1000 // 50 km
const nearActivitiesLimit = 20

type Cron struct {
	user     user.Service
	activity activity.Service
	r12n     recommendation.Service
}

func (c Cron) Process(ctx context.Context) {
	var lastUserId int64

	for {
		users, err := c.user.GetBatch(ctx, lastUserId, userBatch)
		if err != nil {
			return
		}

		var userIds = make([]int64, 0, len(users))
		var r12ns = make([]recommendation.R12n, 0, len(users))
		for _, u := range users {
			userIds = append(userIds, u.Id)

			as, err := c.activity.GetNear(ctx, u.LastLocationLon, u.LastLocationLat, userActivityDistance)
			if err != nil {
				return
			}

			for _, a := range as {
				r12ns = append(r12ns, recommendation.R12n{UserId: u.Id, ActivityId: a.Id})
			}
		}

		if err = c.r12n.Calculate(ctx, r12ns); err != nil {
			return
		}

		if len(users) < userBatch {
			break
		}
	}
}
