package recommendation

import (
	"context"

	"github.com/artemmarkaryan/fisha-facade/internal/service/activity"
	"github.com/artemmarkaryan/fisha-facade/internal/service/recommendation"
	"github.com/artemmarkaryan/fisha-facade/internal/service/user"
)

const (
	userBatch            uint64 = 100
	nearActivitiesLimit  uint64 = 20
	userActivityDistance int    = 50 * 1000 // 50 km
)

type Cron struct {
	user     user.Service
	activity activity.Service
	r12n     recommendation.Service
}

func (c Cron) Process(ctx context.Context) error {
	var lastUserId int64

	for {
		users, err := c.user.GetBatch(ctx, lastUserId, userBatch)
		if err != nil {
			return err
		}

		var userIds = make([]int64, 0, len(users))
		var r12ns = make([]recommendation.R12n, 0, len(users)*int(nearActivitiesLimit))
		for _, u := range users {
			if u.LastLocationLon == nil || *u.LastLocationLon == 0 || u.LastLocationLat == nil || *u.LastLocationLat == 0 {
				continue
			}
			userIds = append(userIds, u.Id)

			var activities []activity.Activity
			activities, err = c.activity.GetNear(ctx, *u.LastLocationLon, *u.LastLocationLat, userActivityDistance, nearActivitiesLimit)
			if err != nil {
				return err
			}

			for _, a := range activities {
				r12ns = append(r12ns, recommendation.R12n{UserId: u.Id, ActivityId: a.Id})
			}
		}

		if err = c.r12n.Calculate(ctx, r12ns); err != nil {
			return err
		}

		if len(users) < int(userBatch) {
			break
		}
	}

	return nil
}
