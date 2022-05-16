package reaction

import (
	"context"
	"fmt"
	"math"
	"time"

	ai "github.com/artemmarkaryan/fisha-facade/internal/service/activity_interest"
	ui "github.com/artemmarkaryan/fisha-facade/internal/service/user-interest"
)

const coefficient float64 = 0.05

type Service struct{}

func (Service) Calculate(ctx context.Context, reaction Reaction) (err error) {
	if !ValidateReactionType(reaction.Reaction) {
		return fmt.Errorf("unknwown reaction type: %s", reaction.Reaction)
	}

	userInterests, err := new(ui.Service).ByUserIds(ctx, []int64{reaction.UserID})
	if err != nil {
		return fmt.Errorf("cant get user interests: %w", err)
	}

	activityInterests, err := new(ai.Service).ByActivityIds(ctx, []int64{reaction.ActivityID})
	if err != nil {
		return fmt.Errorf("cant get activity interests: %w", err)
	}

	interestsSet := calcInterestsSet(userInterests, activityInterests)

	userInterestsByID := calcUserInterestsByID(userInterests)

	activityInterestsByID := caclActivityInterestsByID(activityInterests)

	var newUserInterests = make([]ui.UserInterest, 0, len(interestsSet))
	var newActivityInterests = make([]ai.ActivityInterest, 0, len(interestsSet))

	tNow := time.Now()
	for interestID := range interestsSet {
		var userInterestRank float64 = 0
		if r, ok := userInterestsByID[interestID]; ok {
			userInterestRank = r.Rank
		}

		var activityInterestRank float64 = 0
		if r, ok := activityInterestsByID[interestID]; ok {
			activityInterestRank = r.Rank
		}

		c := coefficient
		if reaction.Reaction == ReactionTypeDislike {
			c = -c // dislike will shift the weigths in opposite direction
		}

		newUserInterestRank, newActivityInterestRank := calcNewRanks(userInterestRank, activityInterestRank, c)
		newUserInterests = append(newUserInterests, ui.UserInterest{
			UserId:     reaction.UserID,
			InterestId: interestID,
			Rank:       newUserInterestRank,
			UpdatedAt:  &tNow,
		})

		newActivityInterests = append(newActivityInterests, ai.ActivityInterest{
			ActivityId: reaction.ActivityID,
			InterestId: interestID,
			Rank:       newActivityInterestRank,
			UpdatedAt:  &tNow,
		})
	}

	if err = new(ui.Service).Upsert(ctx, newUserInterests); err != nil {
		return fmt.Errorf("cant upsert user interests: %w", err)
	}

	if err = new(ai.Service).Upsert(ctx, newActivityInterests); err != nil {
		return fmt.Errorf("cant upsert activity interests: %w", err)
	}

	return nil
}

func calcNewRanks(userRank, activityRank float64, k float64) (newUserRank, newActivityRank float64) {
	newUserRank = userRank - k*(userRank-activityRank)
	newUserRank = math.Max(-1, newUserRank)
	newUserRank = math.Min(1, newUserRank)

	newActivityRank = activityRank - k*(activityRank-userRank)
	newActivityRank = math.Max(-1, newActivityRank)
	newActivityRank = math.Min(1, newActivityRank)
	return
}

func caclActivityInterestsByID(activityInterests []ai.ActivityInterest) map[int64]ai.ActivityInterest {
	activityInterestByID := make(map[int64]ai.ActivityInterest, len(activityInterests))
	for _, interest := range activityInterests {
		activityInterestByID[interest.InterestId] = interest
	}
	return activityInterestByID
}

func calcUserInterestsByID(userInerests []ui.UserInterest) map[int64]ui.UserInterest {
	userInterestByID := make(map[int64]ui.UserInterest, len(userInerests))
	for _, interest := range userInerests {
		userInterestByID[interest.InterestId] = interest
	}
	return userInterestByID
}

func calcInterestsSet(userInerests []ui.UserInterest, activityInterests []ai.ActivityInterest) map[int64]struct{} {
	interestsSet := make(map[int64]struct{}, len(userInerests)+len(activityInterests))
	for _, interest := range userInerests {
		interestsSet[interest.InterestId] = struct{}{}
	}
	for _, interest := range activityInterests {
		interestsSet[interest.InterestId] = struct{}{}
	}
	return interestsSet
}
