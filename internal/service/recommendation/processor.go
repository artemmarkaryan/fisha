package recommendation

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/artemmarkaryan/fisha-facade/internal/service/activity"
	"github.com/artemmarkaryan/fisha-facade/internal/service/user"
	"github.com/artemmarkaryan/fisha-facade/pkg/logy"
)

var NoUserLocation = errors.New("no user location")

type processorCfg struct {
	findLimit          uint64
	saveLimit          int
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
	logy.Time(ctx, time.Now(), "recommendations processor")

	u, err := new(user.Service).Get(ctx, userID)
	if err != nil {
		return fmt.Errorf("cant get user: %w", err)
	}

	distance := p.initialDistance
	attempts := 0

	currentR12n, err := new(Service).GetExistingActivities(ctx, userID)
	if err != nil {
		return err
	}

	for {
		var r12ns = make([]R12n, 0, int(p.findLimit))
		if !u.ValidLocation() {
			return NoUserLocation
		}

		var activities []activity.Activity
		activities, err = new(activity.Service).GetNear(ctx, *u.Lon, *u.Lat, distance, p.findLimit, currentR12n)
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

		if err = new(Service).CalculateAndSave(ctx, r12ns, p.saveLimit); err != nil {
			return fmt.Errorf("cant calculate recommendations: %w", err)
		}

		break
	}

	return nil
}
