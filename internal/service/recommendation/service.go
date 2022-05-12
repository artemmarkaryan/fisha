package recommendation

import (
	"context"

	ai "github.com/artemmarkaryan/fisha/facade/internal/service/activity_interest"
	ui "github.com/artemmarkaryan/fisha/facade/internal/service/user-interest"
)

type Service struct{}

func (Service) Calculate(ctx context.Context, recs []R12n) error {
	var uIds = make(map[int64]struct{}, len(recs))
	var aIds = make(map[int64]struct{}, len(recs))
	for _, rec := range recs {
		uIds[rec.UserId] = struct{}{}
		aIds[rec.ActivityId] = struct{}{}
	}

	uis, err := new(ui.Service).ByUserIds(ctx, keys(uIds))
	if err != nil {
		return err
	}

	ais, err := new(ai.Service).ByActivityIds(ctx, keys(aIds))
	if err != nil {
		return err
	}

	var interestByUser = make(map[int64][]ui.UserInterest)
	for _, uu := range uis {
		interestByUser[uu.UserId] = append(interestByUser[uu.UserId], uu)
	}

	var interestByActivity = make(map[int64][]ai.ActivityInterest)
	for _, aa := range ais {
		interestByActivity[aa.ActivityId] = append(interestByActivity[aa.ActivityId], aa)
	}
}

func keys[K comparable](m map[K]struct{}) (r []K) {
	for k := range m {
		r = append(r, k)
	}

	return
}
