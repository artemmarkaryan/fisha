package recommendation

import (
	"context"
	"errors"
	"fmt"

	"github.com/artemmarkaryan/fisha-facade/internal/service/activity"
	"github.com/artemmarkaryan/fisha-facade/internal/service/user"
)

var NoUserLocation = errors.New("no user location")

type processorCfg struct {
	limit              uint64
	initialDistance    float64 // meters
	distanceMultiplier float64
	maxAttempts        int
}

type processor struct {
	processorCfg
}

func NewProcessor(processorCfg processorCfg) *processor {
	return &processor{processorCfg: processorCfg}
}

func (p processor) Process(ctx context.Context, userID int64) error {
	u, err := new(user.Service).Get(ctx, userID)
	if err != nil {
		return fmt.Errorf("cant get user: %w", err)
	}

	distance := p.initialDistance
	attempts := 0

	for {
		var r12ns = make([]R12n, 0, int(p.limit))
		if !u.ValidLocation() {
			return NoUserLocation
		}

		var activities []activity.Activity
		activities, err = new(activity.Service).GetNear(ctx, *u.Lon, *u.Lat, distance, p.limit)
		if err != nil {
			return fmt.Errorf("cant find relevant activities: %w", err)
		}

		if len(activities) == 0 {
			if attempts == p.maxAttempts {
				return errors.New("no more activities")
			}

			attempts++
			distance *= p.distanceMultiplier

			continue
		}

		for _, a := range activities {
			r12ns = append(r12ns, R12n{UserId: u.Id, ActivityId: a.Id})
		}

		if err = new(Service).Calculate(ctx, r12ns); err != nil {
			return fmt.Errorf("cant calculate recommendations: %w", err)

		}

		break
	}

	return nil
}
