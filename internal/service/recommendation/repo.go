package recommendation

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/artemmarkaryan/fisha-facade/pkg/database"
	"github.com/artemmarkaryan/fisha-facade/pkg/logy"
)

type repo struct{}

func (repo) upsert(ctx context.Context, recs []R12n) error {
	db, c, err := database.Get(ctx)()
	if err != nil {
		return err
	}

	defer c()

	for _, rec := range recs {
		b := sq.
			Insert("recommendations").
			Values(rec.UserId, rec.ActivityId, rec.Rank, rec.Shown).
			Suffix("on conflict (user_id, activity_id) do update set rank = ?, shown = false", rec.Rank).
			PlaceholderFormat(sq.Dollar)
		q, a, err := b.ToSql()
		if err != nil {
			return err
		}

		if _, err = db.ExecContext(ctx, q, a...); err != nil {
			return err
		}
	}

	return nil
}

func (repo) getRecommendedActivity(ctx context.Context, user int64) (activity int64, err error) {
	db, c, err := database.Get(ctx)()
	if err != nil {
		return
	}

	defer func() { _ = c() }()

	q, a, err := sq.
		Select("activity_id").
		From("recommendations").
		Where(sq.Eq{"user_id": user, "shown": false}).
		OrderBy("rank").
		Limit(1).
		PlaceholderFormat(sq.Dollar).ToSql()

	return activity, db.GetContext(ctx, &activity, q, a...)
}

func (repo) ack(ctx context.Context, user, activity int64) error {
	db, c, err := database.Get(ctx)()
	if err != nil {
		return err
	}

	defer func() { _ = c() }()

	q, a, err := sq.
		Update("recommendations").
		Where(sq.Eq{"user_id": user, "activity_id": activity}).
		Set("shown", true).
		PlaceholderFormat(sq.Dollar).ToSql()

	_, err = db.ExecContext(ctx, q, a...)

	if err != nil {
		logy.Log(ctx).Debugf("recommendation acked. user_id: %v; activity_id: %v", user, activity)
	}

	return err
}

func (repo) getExistingActivities(ctx context.Context, user int64) (as []int64, err error) {
	db, c, err := database.Get(ctx)()
	if err != nil {
		return
	}

	defer func() { _ = c() }()

	q, a, err := sq.
		Select("activity_id").
		From("recommendations").
		Where(sq.Eq{"user_id": user}).
		PlaceholderFormat(sq.Dollar).ToSql()

	logy.Log(ctx).Debugln(q)

	err = db.SelectContext(ctx, &as, q, a...)

	return
}
