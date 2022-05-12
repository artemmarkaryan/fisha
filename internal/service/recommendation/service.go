package recommendation

import (
	"context"
	"math"

	ai "github.com/artemmarkaryan/fisha-facade/internal/service/activity_interest"
	ui "github.com/artemmarkaryan/fisha-facade/internal/service/user-interest"
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

	for i, rec := range recs {
		recs[i].Rank = calculateRank(interestByUser[rec.UserId], interestByActivity[rec.ActivityId])
	}

	return new(repo).Upsert(ctx, recs)
}

func calculateInterestsCorrelation(userRank, activityRank float64) float64 {
	diff := userRank - activityRank
	return math.Pow(diff, 2)
}

func calculateRank(userInterests []ui.UserInterest, activityInterests []ai.ActivityInterest) float64 {
	var (
		interestIdSet        = make(map[int64]struct{})
		userInterestById     = make(map[int64]ui.UserInterest)
		activityInterestById = make(map[int64]ai.ActivityInterest)
	)

	for _, interest := range userInterests {
		userInterestById[interest.InterestId] = interest
		interestIdSet[interest.InterestId] = struct{}{}
	}

	for _, interest := range activityInterests {
		activityInterestById[interest.InterestId] = interest
		interestIdSet[interest.InterestId] = struct{}{}
	}

	var sum float64 = 0
	for interestId := range interestIdSet {
		var userRank float64
		var activityRank float64

		if i, ok := userInterestById[interestId]; ok {
			userRank = i.Rank
		}

		if i, ok := activityInterestById[interestId]; ok {
			activityRank = i.Rank
		}

		sum += calculateInterestsCorrelation(userRank, activityRank)
	}

	return sum
}

func keys[K comparable](m map[K]struct{}) (r []K) {
	for k := range m {
		r = append(r, k)
	}

	return
}
